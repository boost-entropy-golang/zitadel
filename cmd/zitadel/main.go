package main

import (
	"context"
	"flag"
	metrics "github.com/caos/zitadel/internal/telemetry/metrics/config"

	"github.com/caos/logging"
	admin_es "github.com/caos/zitadel/internal/admin/repository/eventsourcing"
	"github.com/caos/zitadel/internal/api"
	internal_authz "github.com/caos/zitadel/internal/api/authz"
	"github.com/caos/zitadel/internal/api/grpc/admin"
	"github.com/caos/zitadel/internal/api/grpc/auth"
	"github.com/caos/zitadel/internal/api/grpc/management"
	"github.com/caos/zitadel/internal/api/oidc"
	auth_es "github.com/caos/zitadel/internal/auth/repository/eventsourcing"
	"github.com/caos/zitadel/internal/authz"
	authz_repo "github.com/caos/zitadel/internal/authz/repository/eventsourcing"
	"github.com/caos/zitadel/internal/config"
	sd "github.com/caos/zitadel/internal/config/systemdefaults"
	es_int "github.com/caos/zitadel/internal/eventstore"
	mgmt_es "github.com/caos/zitadel/internal/management/repository/eventsourcing"
	"github.com/caos/zitadel/internal/notification"
	"github.com/caos/zitadel/internal/setup"
	tracing "github.com/caos/zitadel/internal/telemetry/tracing/config"
	"github.com/caos/zitadel/internal/ui"
	"github.com/caos/zitadel/internal/ui/console"
	"github.com/caos/zitadel/internal/ui/login"
)

type Config struct {
	Log            logging.Config
	Tracing        tracing.TracingConfig
	Metrics        metrics.MetricsConfig
	InternalAuthZ  internal_authz.Config
	SystemDefaults sd.SystemDefaults

	AuthZ authz.Config
	Auth  auth_es.Config
	Admin admin_es.Config
	Mgmt  mgmt_es.Config

	API api.Config
	UI  ui.Config

	Notification notification.Config
}

type setupConfig struct {
	Log logging.Config

	Eventstore     es_int.Config
	SystemDefaults sd.SystemDefaults
	SetUp          setup.IAMSetUp
}

var (
	configPaths         = config.NewArrayFlags("authz.yaml", "startup.yaml", "system-defaults.yaml")
	setupPaths          = config.NewArrayFlags("system-defaults.yaml", "setup.yaml")
	adminEnabled        = flag.Bool("admin", true, "enable admin api")
	managementEnabled   = flag.Bool("management", true, "enable management api")
	authEnabled         = flag.Bool("auth", true, "enable auth api")
	oidcEnabled         = flag.Bool("oidc", true, "enable oidc api")
	loginEnabled        = flag.Bool("login", true, "enable login ui")
	consoleEnabled      = flag.Bool("console", true, "enable console ui")
	notificationEnabled = flag.Bool("notification", true, "enable notification handler")
	localDevMode        = flag.Bool("localDevMode", false, "enable local development specific configs")
)

const (
	cmdStart = "start"
	cmdSetup = "setup"
)

func main() {
	flag.Var(configPaths, "config-files", "paths to the config files")
	flag.Var(configPaths, "setup-files", "paths to the setup files")
	flag.Parse()
	arg := flag.Arg(0)
	switch arg {
	case cmdStart:
		startZitadel(configPaths.Values())
	case cmdSetup:
		startSetup(setupPaths.Values(), *localDevMode)
	default:
		logging.Log("MAIN-afEQ2").Fatal("please provide an valid argument [start, setup]")
	}
}

func startZitadel(configPaths []string) {
	conf := new(Config)
	err := config.Read(conf, configPaths...)
	logging.Log("MAIN-FaF2r").OnError(err).Fatal("cannot read config")

	ctx := context.Background()
	authZRepo, err := authz.Start(ctx, conf.AuthZ, conf.InternalAuthZ, conf.SystemDefaults)
	logging.Log("MAIN-s9KOw").OnError(err).Fatal("error starting authz repo")
	var authRepo *auth_es.EsRepository
	if *authEnabled || *oidcEnabled || *loginEnabled {
		authRepo, err = auth_es.Start(conf.Auth, conf.InternalAuthZ, conf.SystemDefaults, authZRepo)
		logging.Log("MAIN-9oRw6").OnError(err).Fatal("error starting auth repo")
	}

	startAPI(ctx, conf, authZRepo, authRepo)
	startUI(ctx, conf, authRepo)

	if *notificationEnabled {
		notification.Start(ctx, conf.Notification, conf.SystemDefaults)
	}

	<-ctx.Done()
	logging.Log("MAIN-s8d2h").Info("stopping zitadel")
}

func startUI(ctx context.Context, conf *Config, authRepo *auth_es.EsRepository) {
	uis := ui.Create(conf.UI)
	if *loginEnabled {
		login, prefix := login.Start(conf.UI.Login, authRepo, conf.SystemDefaults, *localDevMode)
		uis.RegisterHandler(prefix, login.Handler())
	}
	if *consoleEnabled {
		consoleHandler, prefix, err := console.Start(conf.UI.Console)
		logging.Log("API-AGD1f").OnError(err).Fatal("error starting console")
		uis.RegisterHandler(prefix, consoleHandler)
	}
	uis.Start(ctx)
}

func startAPI(ctx context.Context, conf *Config, authZRepo *authz_repo.EsRepository, authRepo *auth_es.EsRepository) {
	roles := make([]string, len(conf.InternalAuthZ.RolePermissionMappings))
	for i, role := range conf.InternalAuthZ.RolePermissionMappings {
		roles[i] = role.Role
	}
	adminRepo, err := admin_es.Start(ctx, conf.Admin, conf.SystemDefaults, roles)
	logging.Log("API-D42tq").OnError(err).Fatal("error starting auth repo")

	apis := api.Create(conf.API, conf.InternalAuthZ, authZRepo, authRepo, adminRepo, conf.SystemDefaults)

	if *adminEnabled {
		apis.RegisterServer(ctx, admin.CreateServer(adminRepo))
	}
	if *managementEnabled {
		managementRepo, err := mgmt_es.Start(conf.Mgmt, conf.SystemDefaults, roles)
		logging.Log("API-Gd2qq").OnError(err).Fatal("error starting management repo")
		apis.RegisterServer(ctx, management.CreateServer(managementRepo, conf.SystemDefaults))
	}
	if *authEnabled {
		apis.RegisterServer(ctx, auth.CreateServer(authRepo))
	}
	if *oidcEnabled {
		op := oidc.NewProvider(ctx, conf.API.OIDC, authRepo, *localDevMode)
		apis.RegisterHandler("/oauth/v2", op.HttpHandler())
	}
	apis.Start(ctx)
}

func startSetup(configPaths []string, localDevMode bool) {
	conf := new(setupConfig)
	err := config.Read(conf, configPaths...)
	logging.Log("MAIN-FaF2r").OnError(err).Fatal("cannot read config")

	ctx := context.Background()

	setup, err := setup.StartSetup(conf.Eventstore, conf.SystemDefaults)
	logging.Log("SERVE-fD252").OnError(err).Panic("failed to start setup")
	err = setup.Execute(ctx, conf.SetUp)
	logging.Log("SERVE-djs3R").OnError(err).Panic("failed to execute setup")
}

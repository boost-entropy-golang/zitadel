package eventsourcing

import (
	"context"

	"github.com/caos/zitadel/internal/api/authz"
	"github.com/caos/zitadel/internal/auth/repository/eventsourcing/eventstore"
	"github.com/caos/zitadel/internal/auth/repository/eventsourcing/handler"
	"github.com/caos/zitadel/internal/auth/repository/eventsourcing/spooler"
	auth_view "github.com/caos/zitadel/internal/auth/repository/eventsourcing/view"
	"github.com/caos/zitadel/internal/auth_request/repository/cache"
	authz_repo "github.com/caos/zitadel/internal/authz/repository/eventsourcing"
	sd "github.com/caos/zitadel/internal/config/systemdefaults"
	"github.com/caos/zitadel/internal/config/types"
	"github.com/caos/zitadel/internal/crypto"
	es_int "github.com/caos/zitadel/internal/eventstore"
	es_spol "github.com/caos/zitadel/internal/eventstore/spooler"
	es_iam "github.com/caos/zitadel/internal/iam/repository/eventsourcing"
	"github.com/caos/zitadel/internal/id"
	es_key "github.com/caos/zitadel/internal/key/repository/eventsourcing"
	es_org "github.com/caos/zitadel/internal/org/repository/eventsourcing"
	es_policy "github.com/caos/zitadel/internal/policy/repository/eventsourcing"
	es_proj "github.com/caos/zitadel/internal/project/repository/eventsourcing"
	es_user "github.com/caos/zitadel/internal/user/repository/eventsourcing"
)

type Config struct {
	SearchLimit uint64
	Domain      string
	Eventstore  es_int.Config
	AuthRequest cache.Config
	View        types.SQL
	Spooler     spooler.SpoolerConfig
	KeyConfig   es_key.KeyConfig
}

type EsRepository struct {
	spooler *es_spol.Spooler
	eventstore.UserRepo
	eventstore.AuthRequestRepo
	eventstore.TokenRepo
	eventstore.KeyRepository
	eventstore.ApplicationRepo
	eventstore.UserSessionRepo
	eventstore.UserGrantRepo
	eventstore.OrgRepository
	eventstore.IAMRepository
	eventstore.PolicyRepo
}

func Start(conf Config, authZ authz.Config, systemDefaults sd.SystemDefaults, authZRepo *authz_repo.EsRepository) (*EsRepository, error) {
	es, err := es_int.Start(conf.Eventstore)
	if err != nil {
		return nil, err
	}

	sqlClient, err := conf.View.Start()
	if err != nil {
		return nil, err
	}

	keyAlgorithm, err := crypto.NewAESCrypto(conf.KeyConfig.EncryptionConfig)
	if err != nil {
		return nil, err
	}
	idGenerator := id.SonyFlakeGenerator

	view, err := auth_view.StartView(sqlClient, keyAlgorithm, idGenerator)
	if err != nil {
		return nil, err
	}
	policy, err := es_policy.StartPolicy(
		es_policy.PolicyConfig{
			Eventstore: es,
			Cache:      conf.Eventstore.Cache,
		},
		systemDefaults,
	)
	if err != nil {
		return nil, err
	}
	user, err := es_user.StartUser(
		es_user.UserConfig{
			Eventstore: es,
			Cache:      conf.Eventstore.Cache,
		},
		systemDefaults,
	)
	if err != nil {
		return nil, err
	}
	authReq, err := cache.Start(conf.AuthRequest)
	if err != nil {
		return nil, err
	}

	key, err := es_key.StartKey(es, conf.KeyConfig, keyAlgorithm, idGenerator)
	if err != nil {
		return nil, err
	}
	iam, err := es_iam.StartIAM(
		es_iam.IAMConfig{
			Eventstore: es,
			Cache:      conf.Eventstore.Cache,
		},
		systemDefaults,
	)
	if err != nil {
		return nil, err
	}

	project, err := es_proj.StartProject(
		es_proj.ProjectConfig{
			Cache:      conf.Eventstore.Cache,
			Eventstore: es,
		},
		systemDefaults,
	)
	if err != nil {
		return nil, err
	}

	org := es_org.StartOrg(es_org.OrgConfig{Eventstore: es, IAMDomain: conf.Domain}, systemDefaults)

	repos := handler.EventstoreRepos{UserEvents: user, ProjectEvents: project, OrgEvents: org, IamEvents: iam}
	spool := spooler.StartSpooler(conf.Spooler, es, view, sqlClient, repos, systemDefaults)

	return &EsRepository{
		spool,
		eventstore.UserRepo{
			Eventstore:   es,
			UserEvents:   user,
			OrgEvents:    org,
			PolicyEvents: policy,
			View:         view,
		},
		eventstore.AuthRequestRepo{
			UserEvents:               user,
			AuthRequests:             authReq,
			View:                     view,
			UserSessionViewProvider:  view,
			UserViewProvider:         view,
			UserEventProvider:        user,
			OrgViewProvider:          view,
			IdGenerator:              idGenerator,
			PasswordCheckLifeTime:    systemDefaults.VerificationLifetimes.PasswordCheck.Duration,
			MfaInitSkippedLifeTime:   systemDefaults.VerificationLifetimes.MfaInitSkip.Duration,
			MfaSoftwareCheckLifeTime: systemDefaults.VerificationLifetimes.MfaSoftwareCheck.Duration,
			MfaHardwareCheckLifeTime: systemDefaults.VerificationLifetimes.MfaHardwareCheck.Duration,
		},
		eventstore.TokenRepo{View: view},
		eventstore.KeyRepository{
			KeyEvents:          key,
			View:               view,
			SigningKeyRotation: conf.KeyConfig.SigningKeyRotation.Duration,
		},
		eventstore.ApplicationRepo{
			View:          view,
			ProjectEvents: project,
		},
		eventstore.UserSessionRepo{
			View: view,
		},
		eventstore.UserGrantRepo{
			SearchLimit: conf.SearchLimit,
			View:        view,
			IamID:       systemDefaults.IamID,
			Auth:        authZ,
			AuthZRepo:   authZRepo,
		},
		eventstore.OrgRepository{
			SearchLimit:      conf.SearchLimit,
			View:             view,
			OrgEventstore:    org,
			PolicyEventstore: policy,
			UserEventstore:   user,
		},
		eventstore.IAMRepository{
			IAMEvents: iam,
			IAMID:     systemDefaults.IamID,
		},
		eventstore.PolicyRepo{
			PolicyEvents: policy,
		},
	}, nil
}

func (repo *EsRepository) Health(ctx context.Context) error {
	if err := repo.UserRepo.Health(ctx); err != nil {
		return err
	}
	return repo.AuthRequestRepo.Health(ctx)
}

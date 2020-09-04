package eventsourcing

import (
	"context"

	sd "github.com/caos/zitadel/internal/config/systemdefaults"
	"github.com/caos/zitadel/internal/config/types"
	es_int "github.com/caos/zitadel/internal/eventstore"
	es_spol "github.com/caos/zitadel/internal/eventstore/spooler"
	es_iam "github.com/caos/zitadel/internal/iam/repository/eventsourcing"
	"github.com/caos/zitadel/internal/management/repository/eventsourcing/eventstore"
	"github.com/caos/zitadel/internal/management/repository/eventsourcing/handler"
	"github.com/caos/zitadel/internal/management/repository/eventsourcing/spooler"
	mgmt_view "github.com/caos/zitadel/internal/management/repository/eventsourcing/view"
	es_org "github.com/caos/zitadel/internal/org/repository/eventsourcing"
	es_pol "github.com/caos/zitadel/internal/policy/repository/eventsourcing"
	es_proj "github.com/caos/zitadel/internal/project/repository/eventsourcing"
	es_usr "github.com/caos/zitadel/internal/user/repository/eventsourcing"
	es_grant "github.com/caos/zitadel/internal/usergrant/repository/eventsourcing"
)

type Config struct {
	SearchLimit uint64
	Domain      string
	Eventstore  es_int.Config
	View        types.SQL
	Spooler     spooler.SpoolerConfig
}

type EsRepository struct {
	spooler *es_spol.Spooler
	eventstore.OrgRepository
	eventstore.ProjectRepo
	eventstore.UserRepo
	eventstore.UserGrantRepo
	eventstore.PolicyRepo
	eventstore.IAMRepository
}

func Start(conf Config, systemDefaults sd.SystemDefaults, roles []string) (*EsRepository, error) {
	es, err := es_int.Start(conf.Eventstore)
	if err != nil {
		return nil, err
	}

	sqlClient, err := conf.View.Start()
	if err != nil {
		return nil, err
	}
	view, err := mgmt_view.StartView(sqlClient)
	if err != nil {
		return nil, err
	}

	project, err := es_proj.StartProject(es_proj.ProjectConfig{
		Eventstore: es,
		Cache:      conf.Eventstore.Cache,
	}, systemDefaults)
	if err != nil {
		return nil, err
	}
	policy, err := es_pol.StartPolicy(es_pol.PolicyConfig{
		Eventstore: es,
		Cache:      conf.Eventstore.Cache,
	}, systemDefaults)
	if err != nil {
		return nil, err
	}
	user, err := es_usr.StartUser(es_usr.UserConfig{
		Eventstore: es,
		Cache:      conf.Eventstore.Cache,
	}, systemDefaults)
	if err != nil {
		return nil, err
	}
	usergrant, err := es_grant.StartUserGrant(es_grant.UserGrantConfig{
		Eventstore: es,
		Cache:      conf.Eventstore.Cache,
	})
	if err != nil {
		return nil, err
	}
	org := es_org.StartOrg(es_org.OrgConfig{Eventstore: es, IAMDomain: conf.Domain}, systemDefaults)

	iam, err := es_iam.StartIAM(es_iam.IAMConfig{
		Eventstore: es,
		Cache:      conf.Eventstore.Cache,
	}, systemDefaults)
	if err != nil {
		return nil, err
	}
	eventstoreRepos := handler.EventstoreRepos{ProjectEvents: project, UserEvents: user, OrgEvents: org, IamEvents: iam}
	spool := spooler.StartSpooler(conf.Spooler, es, view, sqlClient, eventstoreRepos, systemDefaults)

	return &EsRepository{
		spooler:       spool,
		OrgRepository: eventstore.OrgRepository{conf.SearchLimit, org, user, view, roles, systemDefaults},
		ProjectRepo:   eventstore.ProjectRepo{es, conf.SearchLimit, project, usergrant, user, iam, view, roles, systemDefaults.IamID},
		UserRepo:      eventstore.UserRepo{conf.SearchLimit, user, policy, org, view, systemDefaults},
		UserGrantRepo: eventstore.UserGrantRepo{conf.SearchLimit, usergrant, view},
		PolicyRepo:    eventstore.PolicyRepo{policy},
		IAMRepository: eventstore.IAMRepository{iam},
	}, nil
}

func (repo *EsRepository) Health() error {
	return repo.ProjectEvents.Health(context.Background())
}

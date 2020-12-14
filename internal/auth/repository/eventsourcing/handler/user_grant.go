package handler

import (
	"context"
	"strings"

	"github.com/caos/logging"
	"github.com/caos/zitadel/internal/errors"
	caos_errs "github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/eventstore/models"
	es_models "github.com/caos/zitadel/internal/eventstore/models"
	"github.com/caos/zitadel/internal/eventstore/spooler"
	iam_model "github.com/caos/zitadel/internal/iam/model"
	iam_events "github.com/caos/zitadel/internal/iam/repository/eventsourcing"
	iam_es_model "github.com/caos/zitadel/internal/iam/repository/eventsourcing/model"
	org_model "github.com/caos/zitadel/internal/org/model"
	org_events "github.com/caos/zitadel/internal/org/repository/eventsourcing"
	org_es_model "github.com/caos/zitadel/internal/org/repository/eventsourcing/model"
	proj_model "github.com/caos/zitadel/internal/project/model"
	proj_event "github.com/caos/zitadel/internal/project/repository/eventsourcing"
	proj_es_model "github.com/caos/zitadel/internal/project/repository/eventsourcing/model"
	usr_model "github.com/caos/zitadel/internal/user/model"
	usr_events "github.com/caos/zitadel/internal/user/repository/eventsourcing"
	usr_es_model "github.com/caos/zitadel/internal/user/repository/eventsourcing/model"
	grant_es_model "github.com/caos/zitadel/internal/usergrant/repository/eventsourcing/model"
	view_model "github.com/caos/zitadel/internal/usergrant/repository/view/model"
)

type UserGrant struct {
	handler
	eventstore    eventstore.Eventstore
	projectEvents *proj_event.ProjectEventstore
	userEvents    *usr_events.UserEventstore
	orgEvents     *org_events.OrgEventstore
	iamEvents     *iam_events.IAMEventstore
	iamID         string
	iamProjectID  string
}

const (
	userGrantTable = "auth.user_grants"
)

func (u *UserGrant) ViewModel() string {
	return userGrantTable
}

func (u *UserGrant) EventQuery() (*models.SearchQuery, error) {
	if u.iamProjectID == "" {
		err := u.setIamProjectID()
		if err != nil {
			return nil, err
		}
	}
	sequence, err := u.view.GetLatestUserGrantSequence()
	if err != nil {
		return nil, err
	}
	return es_models.NewSearchQuery().
		AggregateTypeFilter(grant_es_model.UserGrantAggregate, iam_es_model.IAMAggregate, org_es_model.OrgAggregate, usr_es_model.UserAggregate, proj_es_model.ProjectAggregate).
		LatestSequenceFilter(sequence.CurrentSequence), nil
}

func (u *UserGrant) Reduce(event *models.Event) (err error) {
	switch event.AggregateType {
	case grant_es_model.UserGrantAggregate:
		err = u.processUserGrant(event)
	case usr_es_model.UserAggregate:
		err = u.processUser(event)
	case proj_es_model.ProjectAggregate:
		err = u.processProject(event)
	case iam_es_model.IAMAggregate:
		err = u.processIAMMember(event, "IAM", false)
	case org_es_model.OrgAggregate:
		return u.processOrg(event)
	}
	return err
}

func (u *UserGrant) processUserGrant(event *models.Event) (err error) {
	grant := new(view_model.UserGrantView)
	switch event.Type {
	case grant_es_model.UserGrantAdded:
		err = grant.AppendEvent(event)
		if err != nil {
			return err
		}
		err = u.fillData(grant, event.ResourceOwner)
	case grant_es_model.UserGrantChanged,
		grant_es_model.UserGrantCascadeChanged,
		grant_es_model.UserGrantDeactivated,
		grant_es_model.UserGrantReactivated:
		grant, err = u.view.UserGrantByID(event.AggregateID)
		if err != nil {
			return err
		}
		err = grant.AppendEvent(event)
	case grant_es_model.UserGrantRemoved, grant_es_model.UserGrantCascadeRemoved:
		return u.view.DeleteUserGrant(event.AggregateID, event.Sequence, event.CreationDate)
	default:
		return u.view.ProcessedUserGrantSequence(event.Sequence, event.CreationDate)
	}
	if err != nil {
		return err
	}
	return u.view.PutUserGrant(grant, grant.Sequence, event.CreationDate)
}

func (u *UserGrant) processUser(event *models.Event) (err error) {
	switch event.Type {
	case usr_es_model.UserProfileChanged,
		usr_es_model.UserEmailChanged,
		usr_es_model.HumanProfileChanged,
		usr_es_model.HumanEmailChanged,
		usr_es_model.MachineChanged:
		grants, err := u.view.UserGrantsByUserID(event.AggregateID)
		if err != nil {
			return err
		}
		if len(grants) == 0 {
			return u.view.ProcessedUserGrantSequence(event.Sequence, event.CreationDate)
		}
		user, err := u.userEvents.UserByID(context.Background(), event.AggregateID)
		if err != nil {
			return err
		}
		for _, grant := range grants {
			u.fillUserData(grant, user)
		}
		return u.view.PutUserGrants(grants, event.Sequence, event.CreationDate)
	default:
		return u.view.ProcessedUserGrantSequence(event.Sequence, event.CreationDate)
	}
}

func (u *UserGrant) processProject(event *models.Event) (err error) {
	switch event.Type {
	case proj_es_model.ProjectChanged:
		grants, err := u.view.UserGrantsByProjectID(event.AggregateID)
		if err != nil {
			return err
		}
		project, err := u.projectEvents.ProjectByID(context.Background(), event.AggregateID)
		if err != nil {
			return err
		}
		for _, grant := range grants {
			u.fillProjectData(grant, project)
		}
		return u.view.PutUserGrants(grants, event.Sequence, event.CreationDate)
	case proj_es_model.ProjectMemberAdded, proj_es_model.ProjectMemberChanged, proj_es_model.ProjectMemberRemoved:
		member := new(proj_es_model.ProjectMember)
		member.SetData(event)
		return u.processMember(event, "PROJECT", event.AggregateID, member.UserID, member.Roles)
	case proj_es_model.ProjectGrantMemberAdded, proj_es_model.ProjectGrantMemberChanged, proj_es_model.ProjectGrantMemberRemoved:
		member := new(proj_es_model.ProjectGrantMember)
		member.SetData(event)
		return u.processMember(event, "PROJECT_GRANT", member.GrantID, member.UserID, member.Roles)
	default:
		return u.view.ProcessedUserGrantSequence(event.Sequence, event.CreationDate)
	}
}

func (u *UserGrant) processOrg(event *models.Event) (err error) {
	switch event.Type {
	case org_es_model.OrgMemberAdded, org_es_model.OrgMemberChanged, org_es_model.OrgMemberRemoved:
		member := new(org_es_model.OrgMember)
		member.SetData(event)
		return u.processMember(event, "ORG", "", member.UserID, member.Roles)
	default:
		return u.view.ProcessedUserGrantSequence(event.Sequence, event.CreationDate)
	}
}

func (u *UserGrant) processIAMMember(event *models.Event, rolePrefix string, suffix bool) error {
	member := new(iam_es_model.IAMMember)

	switch event.Type {
	case iam_es_model.IAMMemberAdded, iam_es_model.IAMMemberChanged:
		member.SetData(event)

		grant, err := u.view.UserGrantByIDs(u.iamID, u.iamProjectID, member.UserID)
		if err != nil && !errors.IsNotFound(err) {
			return err
		}
		if errors.IsNotFound(err) {
			grant = &view_model.UserGrantView{
				ID:            u.iamProjectID + member.UserID,
				ResourceOwner: u.iamID,
				OrgName:       u.iamID,
				ProjectID:     u.iamProjectID,
				UserID:        member.UserID,
				RoleKeys:      member.Roles,
				CreationDate:  event.CreationDate,
			}
			if suffix {
				grant.RoleKeys = suffixRoles(event.AggregateID, grant.RoleKeys)
			}
		} else {
			newRoles := member.Roles
			if grant.RoleKeys != nil {
				grant.RoleKeys = mergeExistingRoles(rolePrefix, "", grant.RoleKeys, newRoles)
			} else {
				grant.RoleKeys = newRoles
			}
		}
		grant.Sequence = event.Sequence
		grant.ChangeDate = event.CreationDate
		return u.view.PutUserGrant(grant, grant.Sequence, event.CreationDate)
	case iam_es_model.IAMMemberRemoved:
		member.SetData(event)
		grant, err := u.view.UserGrantByIDs(u.iamID, u.iamProjectID, member.UserID)
		if err != nil {
			return err
		}
		return u.view.DeleteUserGrant(grant.ID, event.Sequence, event.CreationDate)
	default:
		return u.view.ProcessedUserGrantSequence(event.Sequence, event.CreationDate)
	}
}

func (u *UserGrant) processMember(event *models.Event, rolePrefix, roleSuffix string, userID string, roleKeys []string) error {
	switch event.Type {
	case org_es_model.OrgMemberAdded, proj_es_model.ProjectMemberAdded, proj_es_model.ProjectGrantMemberAdded,
		org_es_model.OrgMemberChanged, proj_es_model.ProjectMemberChanged, proj_es_model.ProjectGrantMemberChanged:

		grant, err := u.view.UserGrantByIDs(event.ResourceOwner, u.iamProjectID, userID)
		if err != nil && !errors.IsNotFound(err) {
			return err
		}
		if roleSuffix != "" {
			roleKeys = suffixRoles(roleSuffix, roleKeys)
		}
		if errors.IsNotFound(err) {
			grant = &view_model.UserGrantView{
				ID:            u.iamProjectID + event.ResourceOwner + userID,
				ResourceOwner: event.ResourceOwner,
				ProjectID:     u.iamProjectID,
				UserID:        userID,
				RoleKeys:      roleKeys,
				CreationDate:  event.CreationDate,
			}
			u.fillData(grant, event.ResourceOwner)
		} else {
			newRoles := roleKeys
			if grant.RoleKeys != nil {
				grant.RoleKeys = mergeExistingRoles(rolePrefix, roleSuffix, grant.RoleKeys, newRoles)
			} else {
				grant.RoleKeys = newRoles
			}
		}
		grant.Sequence = event.Sequence
		grant.ChangeDate = event.CreationDate
		return u.view.PutUserGrant(grant, event.Sequence, event.CreationDate)
	case org_es_model.OrgMemberRemoved,
		proj_es_model.ProjectMemberRemoved,
		proj_es_model.ProjectGrantMemberRemoved:

		grant, err := u.view.UserGrantByIDs(event.ResourceOwner, u.iamProjectID, userID)
		if err != nil && !errors.IsNotFound(err) {
			return err
		}
		if errors.IsNotFound(err) {
			return u.view.ProcessedUserGrantSequence(event.Sequence, event.CreationDate)
		}
		if roleSuffix != "" {
			roleKeys = suffixRoles(roleSuffix, roleKeys)
		}
		if grant.RoleKeys == nil {
			return u.view.ProcessedUserGrantSequence(event.Sequence, event.CreationDate)
		}
		grant.RoleKeys = mergeExistingRoles(rolePrefix, roleSuffix, grant.RoleKeys, nil)
		return u.view.PutUserGrant(grant, event.Sequence, event.CreationDate)
	default:
		return u.view.ProcessedUserGrantSequence(event.Sequence, event.CreationDate)
	}
}

func suffixRoles(suffix string, roles []string) []string {
	suffixedRoles := make([]string, len(roles))
	for i := 0; i < len(roles); i++ {
		suffixedRoles[i] = roles[i] + ":" + suffix
	}
	return suffixedRoles
}

func mergeExistingRoles(rolePrefix, suffix string, existingRoles, newRoles []string) []string {
	mergedRoles := make([]string, 0)
	for _, existingRole := range existingRoles {
		if !strings.HasPrefix(existingRole, rolePrefix) {
			mergedRoles = append(mergedRoles, existingRole)
			continue
		}
		if suffix != "" && !strings.HasSuffix(existingRole, suffix) {
			mergedRoles = append(mergedRoles, existingRole)
		}
	}
	return append(mergedRoles, newRoles...)
}

func (u *UserGrant) setIamProjectID() error {
	if u.iamProjectID != "" {
		return nil
	}
	iam, err := u.iamEvents.IAMByID(context.Background(), u.iamID)
	if err != nil {
		return err
	}

	if iam.SetUpDone < iam_model.StepCount-1 {
		return caos_errs.ThrowPreconditionFailed(nil, "HANDL-s5DTs", "Setup not done")
	}
	u.iamProjectID = iam.IAMProjectID
	return nil
}

func (u *UserGrant) fillData(grant *view_model.UserGrantView, resourceOwner string) (err error) {
	user, err := u.userEvents.UserByID(context.Background(), grant.UserID)
	if err != nil {
		return err
	}
	u.fillUserData(grant, user)
	project, err := u.projectEvents.ProjectByID(context.Background(), grant.ProjectID)
	if err != nil {
		return err
	}
	u.fillProjectData(grant, project)

	org, err := u.orgEvents.OrgByID(context.TODO(), org_model.NewOrg(resourceOwner))
	if err != nil {
		return err
	}
	u.fillOrgData(grant, org)
	return nil
}

func (u *UserGrant) fillUserData(grant *view_model.UserGrantView, user *usr_model.User) {
	grant.UserName = user.UserName
	if user.Human != nil {
		grant.FirstName = user.FirstName
		grant.LastName = user.LastName
		grant.DisplayName = user.FirstName + " " + user.LastName
		grant.Email = user.EmailAddress
	}
	if user.Machine != nil {
		grant.DisplayName = user.Machine.Name
	}
}

func (u *UserGrant) fillProjectData(grant *view_model.UserGrantView, project *proj_model.Project) {
	grant.ProjectName = project.Name
	grant.ProjectOwner = project.ResourceOwner
}

func (u *UserGrant) fillOrgData(grant *view_model.UserGrantView, org *org_model.Org) {
	grant.OrgName = org.Name
	for _, domain := range org.Domains {
		if domain.Primary {
			grant.OrgPrimaryDomain = domain.Domain
			break
		}
	}
}

func (u *UserGrant) OnError(event *models.Event, err error) error {
	logging.LogWithFields("SPOOL-UZmc7", "id", event.AggregateID).WithError(err).Warn("something went wrong in user grant handler")
	return spooler.HandleError(event, err, u.view.GetLatestUserGrantFailedEvent, u.view.ProcessedUserGrantFailedEvent, u.view.ProcessedUserGrantSequence, u.errorCountUntilSkip)
}

func (u *UserGrant) OnSuccess() error {
	return spooler.HandleSuccess(u.view.UpdateUserGrantSpoolerRunTimestamp)
}

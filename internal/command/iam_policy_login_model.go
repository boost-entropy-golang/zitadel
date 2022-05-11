package command

import (
	"context"

	"github.com/zitadel/zitadel/internal/eventstore"

	"github.com/zitadel/zitadel/internal/domain"
	"github.com/zitadel/zitadel/internal/repository/iam"
	"github.com/zitadel/zitadel/internal/repository/policy"
)

type IAMLoginPolicyWriteModel struct {
	LoginPolicyWriteModel
}

func NewIAMLoginPolicyWriteModel() *IAMLoginPolicyWriteModel {
	return &IAMLoginPolicyWriteModel{
		LoginPolicyWriteModel{
			WriteModel: eventstore.WriteModel{
				AggregateID:   domain.IAMID,
				ResourceOwner: domain.IAMID,
			},
		},
	}
}

func (wm *IAMLoginPolicyWriteModel) AppendEvents(events ...eventstore.Event) {
	for _, event := range events {
		switch e := event.(type) {
		case *iam.LoginPolicyAddedEvent:
			wm.LoginPolicyWriteModel.AppendEvents(&e.LoginPolicyAddedEvent)
		case *iam.LoginPolicyChangedEvent:
			wm.LoginPolicyWriteModel.AppendEvents(&e.LoginPolicyChangedEvent)
		}
	}
}

func (wm *IAMLoginPolicyWriteModel) IsValid() bool {
	return wm.AggregateID != ""
}

func (wm *IAMLoginPolicyWriteModel) Reduce() error {
	return wm.LoginPolicyWriteModel.Reduce()
}

func (wm *IAMLoginPolicyWriteModel) Query() *eventstore.SearchQueryBuilder {
	return eventstore.NewSearchQueryBuilder(eventstore.ColumnsEvent).
		ResourceOwner(wm.ResourceOwner).
		AddQuery().
		AggregateTypes(iam.AggregateType).
		AggregateIDs(wm.LoginPolicyWriteModel.AggregateID).
		EventTypes(
			iam.LoginPolicyAddedEventType,
			iam.LoginPolicyChangedEventType).
		Builder()
}

func (wm *IAMLoginPolicyWriteModel) NewChangedEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	allowUsernamePassword,
	allowRegister,
	allowExternalIDP,
	forceMFA,
	hidePasswordReset,
	ignoreUnknownUsernames bool,
	passwordlessType domain.PasswordlessType,
	defaultRedirectURI string,
) (*iam.LoginPolicyChangedEvent, bool) {

	changes := make([]policy.LoginPolicyChanges, 0)
	if wm.AllowUserNamePassword != allowUsernamePassword {
		changes = append(changes, policy.ChangeAllowUserNamePassword(allowUsernamePassword))
	}
	if wm.AllowRegister != allowRegister {
		changes = append(changes, policy.ChangeAllowRegister(allowRegister))
	}
	if wm.AllowExternalIDP != allowExternalIDP {
		changes = append(changes, policy.ChangeAllowExternalIDP(allowExternalIDP))
	}
	if wm.ForceMFA != forceMFA {
		changes = append(changes, policy.ChangeForceMFA(forceMFA))
	}
	if passwordlessType.Valid() && wm.PasswordlessType != passwordlessType {
		changes = append(changes, policy.ChangePasswordlessType(passwordlessType))
	}
	if wm.HidePasswordReset != hidePasswordReset {
		changes = append(changes, policy.ChangeHidePasswordReset(hidePasswordReset))
	}
	if wm.IgnoreUnknownUsernames != ignoreUnknownUsernames {
		changes = append(changes, policy.ChangeIgnoreUnknownUsernames(ignoreUnknownUsernames))
	}
	if wm.DefaultRedirectURI != defaultRedirectURI {
		changes = append(changes, policy.ChangeDefaultRedirectURI(defaultRedirectURI))
	}
	if len(changes) == 0 {
		return nil, false
	}
	changedEvent, err := iam.NewLoginPolicyChangedEvent(ctx, aggregate, changes)
	if err != nil {
		return nil, false
	}
	return changedEvent, true
}

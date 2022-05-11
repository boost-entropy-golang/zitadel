package command

import (
	"context"

	"github.com/zitadel/zitadel/internal/domain"
	caos_errs "github.com/zitadel/zitadel/internal/errors"
	"github.com/zitadel/zitadel/internal/eventstore"
	"github.com/zitadel/zitadel/internal/repository/org"
)

func (c *Commands) SetOrgFeatures(ctx context.Context, resourceOwner string, features *domain.Features) (*domain.ObjectDetails, error) {
	if resourceOwner == "" {
		return nil, caos_errs.ThrowInvalidArgument(nil, "Features-G5tg", "Errors.ResourceOwnerMissing")
	}
	err := c.checkOrgExists(ctx, resourceOwner)
	if err != nil {
		return nil, err
	}
	existingFeatures := NewOrgFeaturesWriteModel(resourceOwner)
	err = c.eventstore.FilterToQueryReducer(ctx, existingFeatures)
	if err != nil {
		return nil, err
	}
	setEvent, hasChanged := existingFeatures.NewSetEvent(
		ctx,
		OrgAggregateFromWriteModel(&existingFeatures.FeaturesWriteModel.WriteModel),
		features.TierName,
		features.TierDescription,
		features.State,
		features.StateDescription,
		features.AuditLogRetention,
		features.LoginPolicyFactors,
		features.LoginPolicyIDP,
		features.LoginPolicyPasswordless,
		features.LoginPolicyRegistration,
		features.LoginPolicyUsernameLogin,
		features.LoginPolicyPasswordReset,
		features.PasswordComplexityPolicy,
		features.LabelPolicyPrivateLabel,
		features.LabelPolicyWatermark,
		features.CustomDomain,
		features.PrivacyPolicy,
		features.MetadataUser,
		features.CustomTextMessage,
		features.CustomTextLogin,
		features.LockoutPolicy,
		features.ActionsAllowed,
		features.MaxActions,
	)
	if !hasChanged {
		return nil, caos_errs.ThrowPreconditionFailed(nil, "Features-GE4h2", "Errors.Features.NotChanged")
	}

	events, err := c.ensureOrgSettingsToFeatures(ctx, resourceOwner, features)
	if err != nil {
		return nil, err
	}
	events = append(events, setEvent)

	pushedEvents, err := c.eventstore.Push(ctx, events...)
	if err != nil {
		return nil, err
	}
	err = AppendAndReduce(existingFeatures, pushedEvents...)
	if err != nil {
		return nil, err
	}
	return writeModelToObjectDetails(&existingFeatures.WriteModel), nil
}

func (c *Commands) RemoveOrgFeatures(ctx context.Context, orgID string) (*domain.ObjectDetails, error) {
	if orgID == "" {
		return nil, caos_errs.ThrowInvalidArgument(nil, "Features-G5tg", "Errors.ResourceOwnerMissing")
	}
	existingFeatures := NewOrgFeaturesWriteModel(orgID)
	err := c.eventstore.FilterToQueryReducer(ctx, existingFeatures)
	if err != nil {
		return nil, err
	}
	if existingFeatures.State == domain.FeaturesStateUnspecified || existingFeatures.State == domain.FeaturesStateRemoved {
		return nil, caos_errs.ThrowNotFound(nil, "Features-Bg32G", "Errors.Features.NotFound")
	}
	removedEvent := org.NewFeaturesRemovedEvent(ctx, OrgAggregateFromWriteModel(&existingFeatures.FeaturesWriteModel.WriteModel))

	features, err := c.getDefaultFeatures(ctx)
	if err != nil {
		return nil, err
	}
	events, err := c.ensureOrgSettingsToFeatures(ctx, orgID, features)
	if err != nil {
		return nil, err
	}

	events = append(events, removedEvent)
	pushedEvents, err := c.eventstore.Push(ctx, events...)
	if err != nil {
		return nil, err
	}
	err = AppendAndReduce(existingFeatures, pushedEvents...)
	if err != nil {
		return nil, err
	}
	return writeModelToObjectDetails(&existingFeatures.WriteModel), nil
}

func (c *Commands) ensureOrgSettingsToFeatures(ctx context.Context, orgID string, features *domain.Features) ([]eventstore.Command, error) {
	events, err := c.setAllowedLoginPolicy(ctx, orgID, features)
	if err != nil {
		return nil, err
	}
	if !features.PasswordComplexityPolicy {
		removePasswordComplexityEvent, err := c.removePasswordComplexityPolicyIfExists(ctx, orgID)
		if err != nil {
			return nil, err
		}
		if removePasswordComplexityEvent != nil {
			events = append(events, removePasswordComplexityEvent)
		}
	}
	labelPolicyEvents, err := c.setAllowedLabelPolicy(ctx, orgID, features)
	if err != nil {
		return nil, err
	}
	events = append(events, labelPolicyEvents...)

	if !features.CustomDomain {
		removeCustomDomainsEvents, err := c.removeCustomDomains(ctx, orgID)
		if err != nil {
			return nil, err
		}
		if removeCustomDomainsEvents != nil {
			events = append(events, removeCustomDomainsEvents...)
		}
	}
	if !features.CustomTextMessage {
		removeCustomMessageTextEvents, err := c.removeOrgMessageTextsIfExists(ctx, orgID)
		if err != nil {
			return nil, err
		}
		if removeCustomMessageTextEvents != nil {
			events = append(events, removeCustomMessageTextEvents...)
		}
	}
	if !features.CustomTextLogin {
		removeCustomLoginTextEvents, err := c.removeOrgLoginTextsIfExists(ctx, orgID)
		if err != nil {
			return nil, err
		}
		if removeCustomLoginTextEvents != nil {
			events = append(events, removeCustomLoginTextEvents...)
		}
	}
	if !features.PrivacyPolicy {
		removePrivacyPolicyEvent, err := c.removePrivacyPolicyIfExists(ctx, orgID)
		if err != nil {
			return nil, err
		}
		if removePrivacyPolicyEvent != nil {
			events = append(events, removePrivacyPolicyEvent)
		}
	}
	if !features.LockoutPolicy {
		removeLockoutPolicyEvent, err := c.removeLockoutPolicyIfExists(ctx, orgID)
		if err != nil {
			return nil, err
		}
		if removeLockoutPolicyEvent != nil {
			events = append(events, removeLockoutPolicyEvent)
		}
	}
	if !features.MetadataUser {
		removeOrgUserMetadatas, err := c.removeUserMetadataFromOrg(ctx, orgID)
		if err != nil {
			return nil, err
		}
		if len(removeOrgUserMetadatas) > 0 {
			events = append(events, removeOrgUserMetadatas...)
		}
	}
	if features.ActionsAllowed == domain.ActionsNotAllowed {
		removeOrgActions, err := c.removeActionsFromOrg(ctx, orgID)
		if err != nil {
			return nil, err
		}
		if len(removeOrgActions) > 0 {
			events = append(events, removeOrgActions...)
		}
	}
	if features.ActionsAllowed == domain.ActionsMaxAllowed {
		deactivateActions, err := c.deactivateNotAllowedActionsFromOrg(ctx, orgID, features.MaxActions)
		if err != nil {
			return nil, err
		}
		if len(deactivateActions) > 0 {
			events = append(events, deactivateActions...)
		}
	}
	return events, nil
}

func (c *Commands) setAllowedLoginPolicy(ctx context.Context, orgID string, features *domain.Features) ([]eventstore.Command, error) {
	events := make([]eventstore.Command, 0)
	existingPolicy, err := c.orgLoginPolicyWriteModelByID(ctx, orgID)
	if err != nil {
		return nil, err
	}
	if existingPolicy.State == domain.PolicyStateUnspecified || existingPolicy.State == domain.PolicyStateRemoved {
		return nil, nil
	}
	defaultPolicy, err := c.getDefaultLoginPolicy(ctx)
	if err != nil {
		return nil, err
	}
	policy := *existingPolicy
	if !features.LoginPolicyFactors {
		if defaultPolicy.ForceMFA != existingPolicy.ForceMFA {
			policy.ForceMFA = defaultPolicy.ForceMFA
		}
		authFactorsEvents, err := c.setDefaultAuthFactorsInCustomLoginPolicy(ctx, orgID)
		if err != nil {
			return nil, err
		}
		events = append(events, authFactorsEvents...)
	}
	if !features.LoginPolicyIDP {
		if defaultPolicy.AllowExternalIDP != existingPolicy.AllowExternalIDP {
			policy.AllowExternalIDP = defaultPolicy.AllowExternalIDP
		}
		//TODO: handle idps
	}
	if !features.LoginPolicyRegistration && defaultPolicy.AllowRegister != existingPolicy.AllowRegister {
		policy.AllowRegister = defaultPolicy.AllowRegister
	}
	if !features.LoginPolicyPasswordless && defaultPolicy.PasswordlessType != existingPolicy.PasswordlessType {
		policy.PasswordlessType = defaultPolicy.PasswordlessType
	}
	if !features.LoginPolicyUsernameLogin && defaultPolicy.AllowUsernamePassword != existingPolicy.AllowUserNamePassword {
		policy.AllowUserNamePassword = defaultPolicy.AllowUsernamePassword
	}
	if !features.LoginPolicyPasswordReset && defaultPolicy.HidePasswordReset != existingPolicy.HidePasswordReset {
		policy.HidePasswordReset = defaultPolicy.HidePasswordReset
	}
	changedEvent, hasChanged := existingPolicy.NewChangedEvent(
		ctx,
		OrgAggregateFromWriteModel(&existingPolicy.WriteModel),
		policy.AllowUserNamePassword,
		policy.AllowRegister,
		policy.AllowExternalIDP,
		policy.ForceMFA,
		policy.HidePasswordReset,
		policy.IgnoreUnknownUsernames,
		policy.PasswordlessType,
		policy.DefaultRedirectURI,
	)
	if hasChanged {
		events = append(events, changedEvent)
	}
	return events, nil
}

func (c *Commands) setDefaultAuthFactorsInCustomLoginPolicy(ctx context.Context, orgID string) ([]eventstore.Command, error) {
	orgAuthFactors, err := c.orgLoginPolicyAuthFactorsWriteModel(ctx, orgID)
	if err != nil {
		return nil, err
	}
	events := make([]eventstore.Command, 0)
	for _, factor := range domain.SecondFactorTypes() {
		state := orgAuthFactors.SecondFactors[factor]
		if state == nil || state.IAM == state.Org {
			continue
		}
		secondFactorWriteModel := orgAuthFactors.ToSecondFactorWriteModel(factor)
		if state.IAM == domain.FactorStateActive {
			event, err := c.addSecondFactorToLoginPolicy(ctx, secondFactorWriteModel, factor)
			if err != nil {
				return nil, err
			}
			if event != nil {
				events = append(events, event)
			}
			continue
		}
		event, err := c.removeSecondFactorFromLoginPolicy(ctx, secondFactorWriteModel, factor)
		if err != nil {
			return nil, err
		}
		if event != nil {
			events = append(events, event)
		}
	}

	for _, factor := range domain.MultiFactorTypes() {
		state := orgAuthFactors.MultiFactors[factor]
		if state == nil || state.IAM == state.Org {
			continue
		}
		multiFactorWriteModel := orgAuthFactors.ToMultiFactorWriteModel(factor)
		if state.IAM == domain.FactorStateActive {
			event, err := c.addMultiFactorToLoginPolicy(ctx, multiFactorWriteModel, factor)
			if err != nil {
				return nil, err
			}
			if event != nil {
				events = append(events, event)
			}
			continue
		}
		event, err := c.removeMultiFactorFromLoginPolicy(ctx, multiFactorWriteModel, factor)
		if err != nil {
			return nil, err
		}
		if event != nil {
			events = append(events, event)
		}
	}
	return events, nil
}

func (c *Commands) setAllowedLabelPolicy(ctx context.Context, orgID string, features *domain.Features) ([]eventstore.Command, error) {
	events := make([]eventstore.Command, 0)
	existingPolicy, err := c.orgLabelPolicyWriteModelByID(ctx, orgID)
	if err != nil {
		return nil, err
	}
	if existingPolicy.State == domain.PolicyStateUnspecified || existingPolicy.State == domain.PolicyStateRemoved {
		return nil, nil
	}
	if !features.LabelPolicyPrivateLabel && !features.LabelPolicyWatermark {
		removeEvent, err := c.removeLabelPolicy(ctx, existingPolicy)
		if err != nil {
			return nil, err
		}
		return append(events, removeEvent), nil
	}
	defaultPolicy, err := c.getDefaultLabelPolicy(ctx)
	if err != nil {
		return nil, err
	}
	policy := *existingPolicy
	if !features.LabelPolicyWatermark && defaultPolicy.DisableWatermark != existingPolicy.DisableWatermark {
		policy.DisableWatermark = defaultPolicy.DisableWatermark
	}
	if !features.LabelPolicyPrivateLabel {
		if defaultPolicy.HideLoginNameSuffix != existingPolicy.HideLoginNameSuffix {
			policy.HideLoginNameSuffix = defaultPolicy.HideLoginNameSuffix
		}
		policy.PrimaryColor = ""
		policy.BackgroundColor = ""
		policy.WarnColor = ""
		policy.FontColor = ""
		policy.PrimaryColorDark = ""
		policy.BackgroundColorDark = ""
		policy.WarnColorDark = ""
		policy.FontColorDark = ""

		assetsEvent, err := c.removeLabelPolicyAssets(ctx, existingPolicy)
		if err != nil {
			return nil, err
		}
		events = append(events, assetsEvent)
	}
	changedEvent, hasChangedEvent := existingPolicy.NewChangedEvent(ctx, OrgAggregateFromWriteModel(&existingPolicy.WriteModel),
		policy.PrimaryColor, policy.BackgroundColor, policy.WarnColor, policy.FontColor,
		policy.PrimaryColorDark, policy.BackgroundColorDark, policy.WarnColorDark, policy.FontColorDark,
		policy.HideLoginNameSuffix, policy.ErrorMsgPopup, policy.HideLoginNameSuffix)
	if hasChangedEvent {
		events = append(events, changedEvent)
	}
	if len(events) > 0 {
		events = append(events, org.NewLabelPolicyActivatedEvent(ctx, OrgAggregateFromWriteModel(&existingPolicy.WriteModel)))
	}
	return events, nil
}

func (c *Commands) getOrgFeaturesOrDefault(ctx context.Context, orgID string) (*domain.Features, error) {
	existingFeatures := NewOrgFeaturesWriteModel(orgID)
	err := c.eventstore.FilterToQueryReducer(ctx, existingFeatures)
	if err != nil {
		return nil, err
	}
	if existingFeatures.State != domain.FeaturesStateUnspecified && existingFeatures.State != domain.FeaturesStateRemoved {
		return writeModelToFeatures(&existingFeatures.FeaturesWriteModel), nil
	}

	existingIAMFeatures := NewIAMFeaturesWriteModel()
	err = c.eventstore.FilterToQueryReducer(ctx, existingIAMFeatures)
	if err != nil {
		return nil, err
	}
	return writeModelToFeatures(&existingIAMFeatures.FeaturesWriteModel), nil
}

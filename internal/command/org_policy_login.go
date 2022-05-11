package command

import (
	"context"
	"reflect"

	"github.com/zitadel/logging"

	"github.com/zitadel/zitadel/internal/domain"
	caos_errs "github.com/zitadel/zitadel/internal/errors"
	"github.com/zitadel/zitadel/internal/eventstore"
	"github.com/zitadel/zitadel/internal/repository/org"
	"github.com/zitadel/zitadel/internal/telemetry/tracing"
)

func (c *Commands) AddLoginPolicy(ctx context.Context, resourceOwner string, policy *domain.LoginPolicy) (*domain.LoginPolicy, error) {
	if resourceOwner == "" {
		return nil, caos_errs.ThrowInvalidArgument(nil, "Org-Fn8ds", "Errors.ResourceOwnerMissing")
	}
	if ok := domain.ValidateDefaultRedirectURI(policy.DefaultRedirectURI); !ok {
		return nil, caos_errs.ThrowInvalidArgument(nil, "Org-WSfdq", "Errors.Org.LoginPolicy.RedirectURIInvalid")
	}
	addedPolicy := NewOrgLoginPolicyWriteModel(resourceOwner)
	err := c.eventstore.FilterToQueryReducer(ctx, addedPolicy)
	if err != nil {
		return nil, err
	}
	if addedPolicy.State == domain.PolicyStateActive {
		return nil, caos_errs.ThrowAlreadyExists(nil, "Org-Dgfb2", "Errors.Org.LoginPolicy.AlreadyExists")
	}

	err = c.checkLoginPolicyAllowed(ctx, resourceOwner, policy)
	if err != nil {
		return nil, err
	}

	orgAgg := OrgAggregateFromWriteModel(&addedPolicy.WriteModel)
	pushedEvents, err := c.eventstore.Push(
		ctx,
		org.NewLoginPolicyAddedEvent(
			ctx,
			orgAgg,
			policy.AllowUsernamePassword,
			policy.AllowRegister,
			policy.AllowExternalIDP,
			policy.ForceMFA,
			policy.HidePasswordReset,
			policy.IgnoreUnknownUsernames,
			policy.PasswordlessType,
			policy.DefaultRedirectURI))
	if err != nil {
		return nil, err
	}
	err = AppendAndReduce(addedPolicy, pushedEvents...)
	if err != nil {
		return nil, err
	}
	return writeModelToLoginPolicy(&addedPolicy.LoginPolicyWriteModel), nil
}

func (c *Commands) orgLoginPolicyWriteModelByID(ctx context.Context, orgID string) (*OrgLoginPolicyWriteModel, error) {
	policyWriteModel := NewOrgLoginPolicyWriteModel(orgID)
	err := c.eventstore.FilterToQueryReducer(ctx, policyWriteModel)
	if err != nil {
		return nil, err
	}
	return policyWriteModel, nil
}

func (c *Commands) getOrgLoginPolicy(ctx context.Context, orgID string) (*domain.LoginPolicy, error) {
	policy, err := c.orgLoginPolicyWriteModelByID(ctx, orgID)
	if err != nil {
		return nil, err
	}
	if policy.State == domain.PolicyStateActive {
		return writeModelToLoginPolicy(&policy.LoginPolicyWriteModel), nil
	}
	return c.getDefaultLoginPolicy(ctx)
}

func (c *Commands) ChangeLoginPolicy(ctx context.Context, resourceOwner string, policy *domain.LoginPolicy) (*domain.LoginPolicy, error) {
	if resourceOwner == "" {
		return nil, caos_errs.ThrowInvalidArgument(nil, "Org-Mf9sf", "Errors.ResourceOwnerMissing")
	}
	if ok := domain.ValidateDefaultRedirectURI(policy.DefaultRedirectURI); !ok {
		return nil, caos_errs.ThrowInvalidArgument(nil, "Org-Sfd21", "Errors.Org.LoginPolicy.RedirectURIInvalid")
	}
	existingPolicy := NewOrgLoginPolicyWriteModel(resourceOwner)
	err := c.eventstore.FilterToQueryReducer(ctx, existingPolicy)
	if err != nil {
		return nil, err
	}
	if existingPolicy.State == domain.PolicyStateUnspecified || existingPolicy.State == domain.PolicyStateRemoved {
		return nil, caos_errs.ThrowNotFound(nil, "Org-M0sif", "Errors.Org.LoginPolicy.NotFound")
	}

	err = c.checkLoginPolicyAllowed(ctx, resourceOwner, policy)
	if err != nil {
		return nil, err
	}

	orgAgg := OrgAggregateFromWriteModel(&existingPolicy.LoginPolicyWriteModel.WriteModel)
	changedEvent, hasChanged := existingPolicy.NewChangedEvent(
		ctx,
		orgAgg,
		policy.AllowUsernamePassword,
		policy.AllowRegister,
		policy.AllowExternalIDP,
		policy.ForceMFA,
		policy.HidePasswordReset,
		policy.IgnoreUnknownUsernames,
		policy.PasswordlessType,
		policy.DefaultRedirectURI)

	if !hasChanged {
		return nil, caos_errs.ThrowPreconditionFailed(nil, "Org-5M9vdd", "Errors.Org.LoginPolicy.NotChanged")
	}

	pushedEvents, err := c.eventstore.Push(ctx, changedEvent)
	if err != nil {
		return nil, err
	}
	err = AppendAndReduce(existingPolicy, pushedEvents...)
	if err != nil {
		return nil, err
	}
	return writeModelToLoginPolicy(&existingPolicy.LoginPolicyWriteModel), nil
}

func (c *Commands) checkLoginPolicyAllowed(ctx context.Context, resourceOwner string, policy *domain.LoginPolicy) error {
	defaultPolicy, err := c.getDefaultLoginPolicy(ctx)
	if err != nil {
		return err
	}
	requiredFeatures := make([]string, 0)
	if defaultPolicy.ForceMFA != policy.ForceMFA || !reflect.DeepEqual(defaultPolicy.MultiFactors, policy.MultiFactors) || !reflect.DeepEqual(defaultPolicy.SecondFactors, policy.SecondFactors) {
		requiredFeatures = append(requiredFeatures, domain.FeatureLoginPolicyFactors)
	}
	if defaultPolicy.AllowExternalIDP != policy.AllowExternalIDP || !reflect.DeepEqual(defaultPolicy.IDPProviders, policy.IDPProviders) {
		requiredFeatures = append(requiredFeatures, domain.FeatureLoginPolicyIDP)
	}
	if defaultPolicy.AllowRegister != policy.AllowRegister {
		requiredFeatures = append(requiredFeatures, domain.FeatureLoginPolicyRegistration)
	}
	if defaultPolicy.PasswordlessType != policy.PasswordlessType {
		requiredFeatures = append(requiredFeatures, domain.FeatureLoginPolicyPasswordless)
	}
	if defaultPolicy.AllowUsernamePassword != policy.AllowUsernamePassword {
		requiredFeatures = append(requiredFeatures, domain.FeatureLoginPolicyUsernameLogin)
	}
	if defaultPolicy.HidePasswordReset != policy.HidePasswordReset {
		requiredFeatures = append(requiredFeatures, domain.FeatureLoginPolicyPasswordReset)
	}
	return c.tokenVerifier.CheckOrgFeatures(ctx, resourceOwner, requiredFeatures...)
}

func (c *Commands) RemoveLoginPolicy(ctx context.Context, orgID string) (*domain.ObjectDetails, error) {
	if orgID == "" {
		return nil, caos_errs.ThrowInvalidArgument(nil, "Org-55Mg9", "Errors.ResourceOwnerMissing")
	}
	existingPolicy := NewOrgLoginPolicyWriteModel(orgID)
	err := c.eventstore.FilterToQueryReducer(ctx, existingPolicy)
	if err != nil {
		return nil, err
	}
	if existingPolicy.State == domain.PolicyStateUnspecified || existingPolicy.State == domain.PolicyStateRemoved {
		return nil, caos_errs.ThrowNotFound(nil, "Org-GHB37", "Errors.Org.LoginPolicy.NotFound")
	}
	orgAgg := OrgAggregateFromWriteModel(&existingPolicy.LoginPolicyWriteModel.WriteModel)
	pushedEvents, err := c.eventstore.Push(ctx, org.NewLoginPolicyRemovedEvent(ctx, orgAgg))
	if err != nil {
		return nil, err
	}
	err = AppendAndReduce(existingPolicy, pushedEvents...)
	if err != nil {
		return nil, err
	}
	return writeModelToObjectDetails(&existingPolicy.LoginPolicyWriteModel.WriteModel), nil
}

func (c *Commands) AddIDPProviderToLoginPolicy(ctx context.Context, resourceOwner string, idpProvider *domain.IDPProvider) (*domain.IDPProvider, error) {
	if resourceOwner == "" {
		return nil, caos_errs.ThrowInvalidArgument(nil, "Org-M0fs9", "Errors.ResourceOwnerMissing")
	}
	if !idpProvider.IsValid() {
		return nil, caos_errs.ThrowInvalidArgument(nil, "Org-9nf88", "Errors.Org.LoginPolicy.IDP.")
	}
	existingPolicy, err := c.orgLoginPolicyWriteModelByID(ctx, resourceOwner)
	if err != nil {
		return nil, err
	}
	if existingPolicy.State == domain.PolicyStateUnspecified || existingPolicy.State == domain.PolicyStateRemoved {
		return nil, caos_errs.ThrowNotFound(nil, "Org-Ffgw2", "Errors.Org.LoginPolicy.NotFound")
	}

	if idpProvider.Type == domain.IdentityProviderTypeOrg {
		_, err = c.getOrgIDPConfigByID(ctx, idpProvider.IDPConfigID, resourceOwner)
	} else {
		_, err = c.getIAMIDPConfigByID(ctx, idpProvider.IDPConfigID)
	}
	if err != nil {
		return nil, caos_errs.ThrowPreconditionFailed(err, "Org-3N9fs", "Errors.IDPConfig.NotExisting")
	}
	idpModel := NewOrgIdentityProviderWriteModel(resourceOwner, idpProvider.IDPConfigID)
	err = c.eventstore.FilterToQueryReducer(ctx, idpModel)
	if err != nil {
		return nil, err
	}
	if idpModel.State == domain.IdentityProviderStateActive {
		return nil, caos_errs.ThrowAlreadyExists(nil, "Org-2B0ps", "Errors.Org.LoginPolicy.IDP.AlreadyExists")
	}

	orgAgg := OrgAggregateFromWriteModel(&idpModel.WriteModel)
	pushedEvents, err := c.eventstore.Push(ctx, org.NewIdentityProviderAddedEvent(ctx, orgAgg, idpProvider.IDPConfigID, idpProvider.Type))
	if err != nil {
		return nil, err
	}
	err = AppendAndReduce(idpModel, pushedEvents...)
	if err != nil {
		return nil, err
	}
	return writeModelToIDPProvider(&idpModel.IdentityProviderWriteModel), nil
}

func (c *Commands) RemoveIDPProviderFromLoginPolicy(ctx context.Context, resourceOwner string, idpProvider *domain.IDPProvider, cascadeExternalIDPs ...*domain.UserIDPLink) (*domain.ObjectDetails, error) {
	if resourceOwner == "" {
		return nil, caos_errs.ThrowInvalidArgument(nil, "Org-M0fs9", "Errors.ResourceOwnerMissing")
	}
	if !idpProvider.IsValid() {
		return nil, caos_errs.ThrowInvalidArgument(nil, "Org-66m9s", "Errors.Org.LoginPolicy.IDP.Invalid")
	}
	existingPolicy, err := c.orgLoginPolicyWriteModelByID(ctx, resourceOwner)
	if err != nil {
		return nil, err
	}
	if existingPolicy.State == domain.PolicyStateUnspecified || existingPolicy.State == domain.PolicyStateRemoved {
		return nil, caos_errs.ThrowNotFound(nil, "Org-GVDfe", "Errors.Org.LoginPolicy.NotFound")
	}

	idpModel := NewOrgIdentityProviderWriteModel(resourceOwner, idpProvider.IDPConfigID)
	err = c.eventstore.FilterToQueryReducer(ctx, idpModel)
	if err != nil {
		return nil, err
	}
	if idpModel.State == domain.IdentityProviderStateUnspecified || idpModel.State == domain.IdentityProviderStateRemoved {
		return nil, caos_errs.ThrowNotFound(nil, "Org-39fjs", "Errors.Org.LoginPolicy.IDP.NotExisting")
	}

	orgAgg := OrgAggregateFromWriteModel(&idpModel.IdentityProviderWriteModel.WriteModel)
	events := c.removeIDPProviderFromLoginPolicy(ctx, orgAgg, idpProvider.IDPConfigID, false, cascadeExternalIDPs...)

	pushedEvents, err := c.eventstore.Push(ctx, events...)
	if err != nil {
		return nil, err
	}
	err = AppendAndReduce(idpModel, pushedEvents...)
	if err != nil {
		return nil, err
	}
	return writeModelToObjectDetails(&idpModel.WriteModel), nil
}

func (c *Commands) removeIDPProviderFromLoginPolicy(ctx context.Context, orgAgg *eventstore.Aggregate, idpConfigID string, cascade bool, cascadeExternalIDPs ...*domain.UserIDPLink) []eventstore.Command {
	var events []eventstore.Command
	if cascade {
		events = append(events, org.NewIdentityProviderCascadeRemovedEvent(ctx, orgAgg, idpConfigID))
	} else {
		events = append(events, org.NewIdentityProviderRemovedEvent(ctx, orgAgg, idpConfigID))
	}

	for _, idp := range cascadeExternalIDPs {
		event, _, err := c.removeUserIDPLink(ctx, idp, true)
		if err != nil {
			logging.LogWithFields("COMMAND-n8RRf", "userid", idp.AggregateID, "idpconfigid", idp.IDPConfigID).WithError(err).Warn("could not cascade remove external idp")
			continue
		}
		events = append(events, event)
	}
	return events
}

func (c *Commands) AddSecondFactorToLoginPolicy(ctx context.Context, secondFactor domain.SecondFactorType, orgID string) (domain.SecondFactorType, *domain.ObjectDetails, error) {
	if orgID == "" {
		return domain.SecondFactorTypeUnspecified, nil, caos_errs.ThrowInvalidArgument(nil, "Org-M0fs9", "Errors.ResourceOwnerMissing")
	}
	if !secondFactor.Valid() {
		return domain.SecondFactorTypeUnspecified, nil, caos_errs.ThrowInvalidArgument(nil, "Org-5m9fs", "Errors.Org.LoginPolicy.MFA.Unspecified")
	}
	secondFactorModel := NewOrgSecondFactorWriteModel(orgID, secondFactor)
	addedEvent, err := c.addSecondFactorToLoginPolicy(ctx, secondFactorModel, secondFactor)
	if err != nil {
		return domain.SecondFactorTypeUnspecified, nil, err
	}

	pushedEvents, err := c.eventstore.Push(ctx, addedEvent)
	if err != nil {
		return domain.SecondFactorTypeUnspecified, nil, err
	}

	err = AppendAndReduce(secondFactorModel, pushedEvents...)
	if err != nil {
		return domain.SecondFactorTypeUnspecified, nil, err
	}
	return secondFactorModel.MFAType, writeModelToObjectDetails(&secondFactorModel.WriteModel), nil
}

func (c *Commands) addSecondFactorToLoginPolicy(ctx context.Context, secondFactorModel *OrgSecondFactorWriteModel, secondFactor domain.SecondFactorType) (*org.LoginPolicySecondFactorAddedEvent, error) {
	err := c.eventstore.FilterToQueryReducer(ctx, secondFactorModel)
	if err != nil {
		return nil, err
	}

	if secondFactorModel.State == domain.FactorStateActive {
		return nil, caos_errs.ThrowAlreadyExists(nil, "Org-2B0ps", "Errors.Org.LoginPolicy.MFA.AlreadyExists")
	}

	orgAgg := OrgAggregateFromWriteModel(&secondFactorModel.SecondFactorWriteModel.WriteModel)
	return org.NewLoginPolicySecondFactorAddedEvent(ctx, orgAgg, secondFactor), nil
}

func (c *Commands) RemoveSecondFactorFromLoginPolicy(ctx context.Context, secondFactor domain.SecondFactorType, orgID string) (*domain.ObjectDetails, error) {
	if orgID == "" {
		return nil, caos_errs.ThrowInvalidArgument(nil, "Org-fM0gs", "Errors.ResourceOwnerMissing")
	}
	if !secondFactor.Valid() {
		return nil, caos_errs.ThrowInvalidArgument(nil, "Org-55n8s", "Errors.Org.LoginPolicy.MFA.Unspecified")
	}
	secondFactorModel := NewOrgSecondFactorWriteModel(orgID, secondFactor)
	removedEvent, err := c.removeSecondFactorFromLoginPolicy(ctx, secondFactorModel, secondFactor)
	if err != nil {
		return nil, err
	}

	pushedEvents, err := c.eventstore.Push(ctx, removedEvent)
	if err != nil {
		return nil, err
	}
	err = AppendAndReduce(secondFactorModel, pushedEvents...)
	if err != nil {
		return nil, err
	}
	return writeModelToObjectDetails(&secondFactorModel.WriteModel), nil
}

func (c *Commands) removeSecondFactorFromLoginPolicy(ctx context.Context, secondFactorModel *OrgSecondFactorWriteModel, secondFactor domain.SecondFactorType) (*org.LoginPolicySecondFactorRemovedEvent, error) {
	err := c.eventstore.FilterToQueryReducer(ctx, secondFactorModel)
	if err != nil {
		return nil, err
	}
	if secondFactorModel.State == domain.FactorStateUnspecified || secondFactorModel.State == domain.FactorStateRemoved {
		return nil, caos_errs.ThrowNotFound(nil, "Org-3M9od", "Errors.Org.LoginPolicy.MFA.NotExisting")
	}
	orgAgg := OrgAggregateFromWriteModel(&secondFactorModel.SecondFactorWriteModel.WriteModel)
	return org.NewLoginPolicySecondFactorRemovedEvent(ctx, orgAgg, secondFactor), nil
}

func (c *Commands) AddMultiFactorToLoginPolicy(ctx context.Context, multiFactor domain.MultiFactorType, orgID string) (domain.MultiFactorType, *domain.ObjectDetails, error) {
	if orgID == "" {
		return domain.MultiFactorTypeUnspecified, nil, caos_errs.ThrowInvalidArgument(nil, "Org-M0fsf", "Errors.ResourceOwnerMissing")
	}
	if !multiFactor.Valid() {
		return domain.MultiFactorTypeUnspecified, nil, caos_errs.ThrowInvalidArgument(nil, "Org-5m9fs", "Errors.Org.LoginPolicy.MFA.Unspecified")
	}
	multiFactorModel := NewOrgMultiFactorWriteModel(orgID, multiFactor)
	addedEvent, err := c.addMultiFactorToLoginPolicy(ctx, multiFactorModel, multiFactor)
	if err != nil {
		return domain.MultiFactorTypeUnspecified, nil, err
	}

	pushedEvents, err := c.eventstore.Push(ctx, addedEvent)
	if err != nil {
		return domain.MultiFactorTypeUnspecified, nil, err
	}
	err = AppendAndReduce(multiFactorModel, pushedEvents...)
	if err != nil {
		return domain.MultiFactorTypeUnspecified, nil, err
	}
	return multiFactorModel.MultiFactorWriteModel.MFAType, writeModelToObjectDetails(&multiFactorModel.WriteModel), nil
}

func (c *Commands) addMultiFactorToLoginPolicy(ctx context.Context, multiFactorModel *OrgMultiFactorWriteModel, multiFactor domain.MultiFactorType) (*org.LoginPolicyMultiFactorAddedEvent, error) {
	err := c.eventstore.FilterToQueryReducer(ctx, multiFactorModel)
	if err != nil {
		return nil, err
	}
	if multiFactorModel.State == domain.FactorStateActive {
		return nil, caos_errs.ThrowAlreadyExists(nil, "Org-3M9od", "Errors.Org.LoginPolicy.MFA.AlreadyExists")
	}

	orgAgg := OrgAggregateFromWriteModel(&multiFactorModel.WriteModel)
	return org.NewLoginPolicyMultiFactorAddedEvent(ctx, orgAgg, multiFactor), nil
}

func (c *Commands) RemoveMultiFactorFromLoginPolicy(ctx context.Context, multiFactor domain.MultiFactorType, orgID string) (*domain.ObjectDetails, error) {
	if orgID == "" {
		return nil, caos_errs.ThrowInvalidArgument(nil, "Org-M0fsf", "Errors.ResourceOwnerMissing")
	}
	if !multiFactor.Valid() {
		return nil, caos_errs.ThrowInvalidArgument(nil, "Org-5m9fs", "Errors.Org.LoginPolicy.MFA.Unspecified")
	}
	multiFactorModel := NewOrgMultiFactorWriteModel(orgID, multiFactor)
	removedEvent, err := c.removeMultiFactorFromLoginPolicy(ctx, multiFactorModel, multiFactor)
	if err != nil {
		return nil, err
	}

	pushedEvents, err := c.eventstore.Push(ctx, removedEvent)
	if err != nil {
		return nil, err
	}
	err = AppendAndReduce(multiFactorModel, pushedEvents...)
	if err != nil {
		return nil, err
	}
	return writeModelToObjectDetails(&multiFactorModel.WriteModel), nil
}

func (c *Commands) removeMultiFactorFromLoginPolicy(ctx context.Context, multiFactorModel *OrgMultiFactorWriteModel, multiFactor domain.MultiFactorType) (*org.LoginPolicyMultiFactorRemovedEvent, error) {
	err := c.eventstore.FilterToQueryReducer(ctx, multiFactorModel)
	if err != nil {
		return nil, err
	}
	if multiFactorModel.State == domain.FactorStateUnspecified || multiFactorModel.State == domain.FactorStateRemoved {
		return nil, caos_errs.ThrowNotFound(nil, "Org-3M9df", "Errors.Org.LoginPolicy.MFA.NotExisting")
	}
	orgAgg := OrgAggregateFromWriteModel(&multiFactorModel.MultiFactorWriteModel.WriteModel)

	return org.NewLoginPolicyMultiFactorRemovedEvent(ctx, orgAgg, multiFactor), nil
}

func (c *Commands) orgLoginPolicyAuthFactorsWriteModel(ctx context.Context, orgID string) (_ *OrgAuthFactorsAllowedWriteModel, err error) {
	ctx, span := tracing.NewSpan(ctx)
	defer func() { span.EndWithError(err) }()

	writeModel := NewOrgAuthFactorsAllowedWriteModel(orgID)
	err = c.eventstore.FilterToQueryReducer(ctx, writeModel)
	if err != nil {
		return nil, err
	}
	return writeModel, nil
}

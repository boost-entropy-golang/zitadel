package iam

import (
	"github.com/caos/zitadel/internal/eventstore"
)

func RegisterEventMappers(es *eventstore.Eventstore) {
	es.RegisterFilterEventMapper(SetupStartedEventType, SetupStepMapper).
		RegisterFilterEventMapper(SetupDoneEventType, SetupStepMapper).
		RegisterFilterEventMapper(GlobalOrgSetEventType, GlobalOrgSetMapper).
		RegisterFilterEventMapper(ProjectSetEventType, ProjectSetMapper).
		RegisterFilterEventMapper(UniqueConstraintsMigratedEventType, MigrateUniqueConstraintEventMapper).
		RegisterFilterEventMapper(LabelPolicyAddedEventType, LabelPolicyAddedEventMapper).
		RegisterFilterEventMapper(LabelPolicyChangedEventType, LabelPolicyChangedEventMapper).
		RegisterFilterEventMapper(LabelPolicyActivatedEventType, LabelPolicyActivatedEventMapper).
		RegisterFilterEventMapper(LoginPolicyAddedEventType, LoginPolicyAddedEventMapper).
		RegisterFilterEventMapper(LoginPolicyChangedEventType, LoginPolicyChangedEventMapper).
		RegisterFilterEventMapper(OrgIAMPolicyAddedEventType, OrgIAMPolicyAddedEventMapper).
		RegisterFilterEventMapper(OrgIAMPolicyChangedEventType, OrgIAMPolicyChangedEventMapper).
		RegisterFilterEventMapper(PasswordAgePolicyAddedEventType, PasswordAgePolicyAddedEventMapper).
		RegisterFilterEventMapper(PasswordAgePolicyChangedEventType, PasswordAgePolicyChangedEventMapper).
		RegisterFilterEventMapper(PasswordComplexityPolicyAddedEventType, PasswordComplexityPolicyAddedEventMapper).
		RegisterFilterEventMapper(PasswordComplexityPolicyChangedEventType, PasswordComplexityPolicyChangedEventMapper).
		RegisterFilterEventMapper(PasswordLockoutPolicyAddedEventType, PasswordLockoutPolicyAddedEventMapper).
		RegisterFilterEventMapper(PasswordLockoutPolicyChangedEventType, PasswordLockoutPolicyChangedEventMapper).
		RegisterFilterEventMapper(MemberAddedEventType, MemberAddedEventMapper).
		RegisterFilterEventMapper(MemberChangedEventType, MemberChangedEventMapper).
		RegisterFilterEventMapper(MemberRemovedEventType, MemberRemovedEventMapper).
		RegisterFilterEventMapper(IDPConfigAddedEventType, IDPConfigAddedEventMapper).
		RegisterFilterEventMapper(IDPConfigChangedEventType, IDPConfigChangedEventMapper).
		RegisterFilterEventMapper(IDPConfigRemovedEventType, IDPConfigRemovedEventMapper).
		RegisterFilterEventMapper(IDPConfigDeactivatedEventType, IDPConfigDeactivatedEventMapper).
		RegisterFilterEventMapper(IDPConfigReactivatedEventType, IDPConfigReactivatedEventMapper).
		RegisterFilterEventMapper(IDPOIDCConfigAddedEventType, IDPOIDCConfigAddedEventMapper).
		RegisterFilterEventMapper(IDPOIDCConfigChangedEventType, IDPOIDCConfigChangedEventMapper).
		RegisterFilterEventMapper(LoginPolicyIDPProviderAddedEventType, IdentityProviderAddedEventMapper).
		RegisterFilterEventMapper(LoginPolicyIDPProviderRemovedEventType, IdentityProviderRemovedEventMapper).
		RegisterFilterEventMapper(LoginPolicyIDPProviderCascadeRemovedEventType, IdentityProviderCascadeRemovedEventMapper).
		RegisterFilterEventMapper(LoginPolicySecondFactorAddedEventType, SecondFactorAddedEventMapper).
		RegisterFilterEventMapper(LoginPolicySecondFactorRemovedEventType, SecondFactorRemovedEventMapper).
		RegisterFilterEventMapper(LoginPolicyMultiFactorAddedEventType, MultiFactorAddedEventMapper).
		RegisterFilterEventMapper(LoginPolicyMultiFactorRemovedEventType, MultiFactorRemovedEventMapper).
		RegisterFilterEventMapper(MailTemplateAddedEventType, MailTemplateAddedEventMapper).
		RegisterFilterEventMapper(MailTemplateChangedEventType, MailTemplateChangedEventMapper).
		RegisterFilterEventMapper(MailTextAddedEventType, MailTextAddedEventMapper).
		RegisterFilterEventMapper(MailTextChangedEventType, MailTextChangedEventMapper).
		RegisterFilterEventMapper(FeaturesSetEventType, FeaturesSetEventMapper)
}

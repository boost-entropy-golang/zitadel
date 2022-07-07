---
title: zitadel/admin.proto
---
> This document reflects the state from API 1.0 (available from 20.04.2021)


## AdminService {#zitadeladminv1adminservice}


### Healthz

> **rpc** Healthz([HealthzRequest](#healthzrequest))
[HealthzResponse](#healthzresponse)

Indicates if ZITADEL is running.
It respondes as soon as ZITADEL started



    GET: /healthz


### GetSupportedLanguages

> **rpc** GetSupportedLanguages([GetSupportedLanguagesRequest](#getsupportedlanguagesrequest))
[GetSupportedLanguagesResponse](#getsupportedlanguagesresponse)

Returns the default languages



    GET: /languages


### GetOrgByID

> **rpc** GetOrgByID([GetOrgByIDRequest](#getorgbyidrequest))
[GetOrgByIDResponse](#getorgbyidresponse)

Returns an organisation by id



    GET: /orgs/{id}


### IsOrgUnique

> **rpc** IsOrgUnique([IsOrgUniqueRequest](#isorguniquerequest))
[IsOrgUniqueResponse](#isorguniqueresponse)

Checks whether an organisation exists by the given parameters



    GET: /orgs/_is_unique


### ListOrgs

> **rpc** ListOrgs([ListOrgsRequest](#listorgsrequest))
[ListOrgsResponse](#listorgsresponse)

Returns all organisations matching the request
all queries need to match (AND)



    POST: /orgs/_search


### SetUpOrg

> **rpc** SetUpOrg([SetUpOrgRequest](#setuporgrequest))
[SetUpOrgResponse](#setuporgresponse)

Creates a new org and user 
and adds the user to the orgs members as ORG_OWNER



    POST: /orgs/_setup


### GetIDPByID

> **rpc** GetIDPByID([GetIDPByIDRequest](#getidpbyidrequest))
[GetIDPByIDResponse](#getidpbyidresponse)

Returns a identity provider configuration of the IAM



    GET: /idps/{id}


### ListIDPs

> **rpc** ListIDPs([ListIDPsRequest](#listidpsrequest))
[ListIDPsResponse](#listidpsresponse)

Returns all identity provider configurations of the IAM



    POST: /idps/_search


### AddOIDCIDP

> **rpc** AddOIDCIDP([AddOIDCIDPRequest](#addoidcidprequest))
[AddOIDCIDPResponse](#addoidcidpresponse)

Adds a new oidc identity provider configuration the IAM



    POST: /idps/oidc


### AddJWTIDP

> **rpc** AddJWTIDP([AddJWTIDPRequest](#addjwtidprequest))
[AddJWTIDPResponse](#addjwtidpresponse)

Adds a new jwt identity provider configuration the IAM



    POST: /idps/jwt


### UpdateIDP

> **rpc** UpdateIDP([UpdateIDPRequest](#updateidprequest))
[UpdateIDPResponse](#updateidpresponse)

Updates the specified idp
all fields are updated. If no value is provided the field will be empty afterwards.



    PUT: /idps/{idp_id}


### DeactivateIDP

> **rpc** DeactivateIDP([DeactivateIDPRequest](#deactivateidprequest))
[DeactivateIDPResponse](#deactivateidpresponse)

Sets the state of the idp to IDP_STATE_INACTIVE
the state MUST be IDP_STATE_ACTIVE for this call



    POST: /idps/{idp_id}/_deactivate


### ReactivateIDP

> **rpc** ReactivateIDP([ReactivateIDPRequest](#reactivateidprequest))
[ReactivateIDPResponse](#reactivateidpresponse)

Sets the state of the idp to IDP_STATE_ACTIVE
the state MUST be IDP_STATE_INACTIVE for this call



    POST: /idps/{idp_id}/_reactivate


### RemoveIDP

> **rpc** RemoveIDP([RemoveIDPRequest](#removeidprequest))
[RemoveIDPResponse](#removeidpresponse)

RemoveIDP deletes the IDP permanetly



    DELETE: /idps/{idp_id}


### UpdateIDPOIDCConfig

> **rpc** UpdateIDPOIDCConfig([UpdateIDPOIDCConfigRequest](#updateidpoidcconfigrequest))
[UpdateIDPOIDCConfigResponse](#updateidpoidcconfigresponse)

Updates the oidc configuration of the specified idp
all fields are updated. If no value is provided the field will be empty afterwards.



    PUT: /idps/{idp_id}/oidc_config


### UpdateIDPJWTConfig

> **rpc** UpdateIDPJWTConfig([UpdateIDPJWTConfigRequest](#updateidpjwtconfigrequest))
[UpdateIDPJWTConfigResponse](#updateidpjwtconfigresponse)

Updates the jwt configuration of the specified idp
all fields are updated. If no value is provided the field will be empty afterwards.



    PUT: /idps/{idp_id}/jwt_config


### GetDefaultFeatures

> **rpc** GetDefaultFeatures([GetDefaultFeaturesRequest](#getdefaultfeaturesrequest))
[GetDefaultFeaturesResponse](#getdefaultfeaturesresponse)





    GET: /features


### SetDefaultFeatures

> **rpc** SetDefaultFeatures([SetDefaultFeaturesRequest](#setdefaultfeaturesrequest))
[SetDefaultFeaturesResponse](#setdefaultfeaturesresponse)





    PUT: /features


### GetOrgFeatures

> **rpc** GetOrgFeatures([GetOrgFeaturesRequest](#getorgfeaturesrequest))
[GetOrgFeaturesResponse](#getorgfeaturesresponse)





    GET: /orgs/{org_id}/features


### SetOrgFeatures

> **rpc** SetOrgFeatures([SetOrgFeaturesRequest](#setorgfeaturesrequest))
[SetOrgFeaturesResponse](#setorgfeaturesresponse)





    PUT: /orgs/{org_id}/features


### ResetOrgFeatures

> **rpc** ResetOrgFeatures([ResetOrgFeaturesRequest](#resetorgfeaturesrequest))
[ResetOrgFeaturesResponse](#resetorgfeaturesresponse)





    DELETE: /orgs/{org_id}/features


### GetOrgIAMPolicy

> **rpc** GetOrgIAMPolicy([GetOrgIAMPolicyRequest](#getorgiampolicyrequest))
[GetOrgIAMPolicyResponse](#getorgiampolicyresponse)

Returns the IAM policy defined by the administrators of ZITADEL



    GET: /policies/orgiam


### UpdateOrgIAMPolicy

> **rpc** UpdateOrgIAMPolicy([UpdateOrgIAMPolicyRequest](#updateorgiampolicyrequest))
[UpdateOrgIAMPolicyResponse](#updateorgiampolicyresponse)

Updates the default IAM policy.
it impacts all organisations without a customised policy



    PUT: /policies/orgiam


### GetCustomOrgIAMPolicy

> **rpc** GetCustomOrgIAMPolicy([GetCustomOrgIAMPolicyRequest](#getcustomorgiampolicyrequest))
[GetCustomOrgIAMPolicyResponse](#getcustomorgiampolicyresponse)

Returns the customised policy or the default if not customised



    GET: /orgs/{org_id}/policies/orgiam


### AddCustomOrgIAMPolicy

> **rpc** AddCustomOrgIAMPolicy([AddCustomOrgIAMPolicyRequest](#addcustomorgiampolicyrequest))
[AddCustomOrgIAMPolicyResponse](#addcustomorgiampolicyresponse)

Defines a custom ORGIAM policy as specified



    POST: /orgs/{org_id}/policies/orgiam


### UpdateCustomOrgIAMPolicy

> **rpc** UpdateCustomOrgIAMPolicy([UpdateCustomOrgIAMPolicyRequest](#updatecustomorgiampolicyrequest))
[UpdateCustomOrgIAMPolicyResponse](#updatecustomorgiampolicyresponse)

Updates a custom ORGIAM policy as specified



    PUT: /orgs/{org_id}/policies/orgiam


### ResetCustomOrgIAMPolicyToDefault

> **rpc** ResetCustomOrgIAMPolicyToDefault([ResetCustomOrgIAMPolicyToDefaultRequest](#resetcustomorgiampolicytodefaultrequest))
[ResetCustomOrgIAMPolicyToDefaultResponse](#resetcustomorgiampolicytodefaultresponse)

Resets the org iam policy of the organisation to default
ZITADEL will fallback to the default policy defined by the ZITADEL administrators



    DELETE: /orgs/{org_id}/policies/orgiam


### GetLabelPolicy

> **rpc** GetLabelPolicy([GetLabelPolicyRequest](#getlabelpolicyrequest))
[GetLabelPolicyResponse](#getlabelpolicyresponse)

Returns the label policy defined by the administrators of ZITADEL



    GET: /policies/label


### GetPreviewLabelPolicy

> **rpc** GetPreviewLabelPolicy([GetPreviewLabelPolicyRequest](#getpreviewlabelpolicyrequest))
[GetPreviewLabelPolicyResponse](#getpreviewlabelpolicyresponse)

Returns the preview label policy defined by the administrators of ZITADEL



    GET: /policies/label/_preview


### UpdateLabelPolicy

> **rpc** UpdateLabelPolicy([UpdateLabelPolicyRequest](#updatelabelpolicyrequest))
[UpdateLabelPolicyResponse](#updatelabelpolicyresponse)

Updates the default label policy of ZITADEL
it impacts all organisations without a customised policy



    PUT: /policies/label


### ActivateLabelPolicy

> **rpc** ActivateLabelPolicy([ActivateLabelPolicyRequest](#activatelabelpolicyrequest))
[ActivateLabelPolicyResponse](#activatelabelpolicyresponse)

Activates all changes of the label policy



    POST: /policies/label/_activate


### RemoveLabelPolicyLogo

> **rpc** RemoveLabelPolicyLogo([RemoveLabelPolicyLogoRequest](#removelabelpolicylogorequest))
[RemoveLabelPolicyLogoResponse](#removelabelpolicylogoresponse)

Removes the logo of the label policy



    DELETE: /policies/label/logo


### RemoveLabelPolicyLogoDark

> **rpc** RemoveLabelPolicyLogoDark([RemoveLabelPolicyLogoDarkRequest](#removelabelpolicylogodarkrequest))
[RemoveLabelPolicyLogoDarkResponse](#removelabelpolicylogodarkresponse)

Removes the logo dark of the label policy



    DELETE: /policies/label/logo_dark


### RemoveLabelPolicyIcon

> **rpc** RemoveLabelPolicyIcon([RemoveLabelPolicyIconRequest](#removelabelpolicyiconrequest))
[RemoveLabelPolicyIconResponse](#removelabelpolicyiconresponse)

Removes the icon of the label policy



    DELETE: /policies/label/icon


### RemoveLabelPolicyIconDark

> **rpc** RemoveLabelPolicyIconDark([RemoveLabelPolicyIconDarkRequest](#removelabelpolicyicondarkrequest))
[RemoveLabelPolicyIconDarkResponse](#removelabelpolicyicondarkresponse)

Removes the logo dark of the label policy



    DELETE: /policies/label/icon_dark


### RemoveLabelPolicyFont

> **rpc** RemoveLabelPolicyFont([RemoveLabelPolicyFontRequest](#removelabelpolicyfontrequest))
[RemoveLabelPolicyFontResponse](#removelabelpolicyfontresponse)

Removes the font of the label policy



    DELETE: /policies/label/font


### GetLoginPolicy

> **rpc** GetLoginPolicy([GetLoginPolicyRequest](#getloginpolicyrequest))
[GetLoginPolicyResponse](#getloginpolicyresponse)

Returns the login policy defined by the administrators of ZITADEL



    GET: /policies/login


### UpdateLoginPolicy

> **rpc** UpdateLoginPolicy([UpdateLoginPolicyRequest](#updateloginpolicyrequest))
[UpdateLoginPolicyResponse](#updateloginpolicyresponse)

Updates the default login policy of ZITADEL
it impacts all organisations without a customised policy



    PUT: /policies/login


### ListLoginPolicyIDPs

> **rpc** ListLoginPolicyIDPs([ListLoginPolicyIDPsRequest](#listloginpolicyidpsrequest))
[ListLoginPolicyIDPsResponse](#listloginpolicyidpsresponse)

Returns the idps linked to the default login policy,
defined by the administrators of ZITADEL



    POST: /policies/login/idps/_search


### AddIDPToLoginPolicy

> **rpc** AddIDPToLoginPolicy([AddIDPToLoginPolicyRequest](#addidptologinpolicyrequest))
[AddIDPToLoginPolicyResponse](#addidptologinpolicyresponse)

Adds the povided idp to the default login policy.
It impacts all organisations without a customised policy



    POST: /policies/login/idps


### RemoveIDPFromLoginPolicy

> **rpc** RemoveIDPFromLoginPolicy([RemoveIDPFromLoginPolicyRequest](#removeidpfromloginpolicyrequest))
[RemoveIDPFromLoginPolicyResponse](#removeidpfromloginpolicyresponse)

Removes the povided idp from the default login policy.
It impacts all organisations without a customised policy



    DELETE: /policies/login/idps/{idp_id}


### ListLoginPolicySecondFactors

> **rpc** ListLoginPolicySecondFactors([ListLoginPolicySecondFactorsRequest](#listloginpolicysecondfactorsrequest))
[ListLoginPolicySecondFactorsResponse](#listloginpolicysecondfactorsresponse)

Returns the available second factors defined by the administrators of ZITADEL



    POST: /policies/login/second_factors/_search


### AddSecondFactorToLoginPolicy

> **rpc** AddSecondFactorToLoginPolicy([AddSecondFactorToLoginPolicyRequest](#addsecondfactortologinpolicyrequest))
[AddSecondFactorToLoginPolicyResponse](#addsecondfactortologinpolicyresponse)

Adds a second factor to the default login policy.
It impacts all organisations without a customised policy



    POST: /policies/login/second_factors


### RemoveSecondFactorFromLoginPolicy

> **rpc** RemoveSecondFactorFromLoginPolicy([RemoveSecondFactorFromLoginPolicyRequest](#removesecondfactorfromloginpolicyrequest))
[RemoveSecondFactorFromLoginPolicyResponse](#removesecondfactorfromloginpolicyresponse)

Removes a second factor from the default login policy.
It impacts all organisations without a customised policy



    DELETE: /policies/login/second_factors/{type}


### ListLoginPolicyMultiFactors

> **rpc** ListLoginPolicyMultiFactors([ListLoginPolicyMultiFactorsRequest](#listloginpolicymultifactorsrequest))
[ListLoginPolicyMultiFactorsResponse](#listloginpolicymultifactorsresponse)

Returns the available multi factors defined by the administrators of ZITADEL



    POST: /policies/login/multi_factors/_search


### AddMultiFactorToLoginPolicy

> **rpc** AddMultiFactorToLoginPolicy([AddMultiFactorToLoginPolicyRequest](#addmultifactortologinpolicyrequest))
[AddMultiFactorToLoginPolicyResponse](#addmultifactortologinpolicyresponse)

Adds a multi factor to the default login policy.
It impacts all organisations without a customised policy



    POST: /policies/login/multi_factors


### RemoveMultiFactorFromLoginPolicy

> **rpc** RemoveMultiFactorFromLoginPolicy([RemoveMultiFactorFromLoginPolicyRequest](#removemultifactorfromloginpolicyrequest))
[RemoveMultiFactorFromLoginPolicyResponse](#removemultifactorfromloginpolicyresponse)

Removes a multi factor from the default login policy.
It impacts all organisations without a customised policy



    DELETE: /policies/login/multi_factors/{type}


### GetPasswordComplexityPolicy

> **rpc** GetPasswordComplexityPolicy([GetPasswordComplexityPolicyRequest](#getpasswordcomplexitypolicyrequest))
[GetPasswordComplexityPolicyResponse](#getpasswordcomplexitypolicyresponse)

Returns the password complexity policy defined by the administrators of ZITADEL



    GET: /policies/password/complexity


### UpdatePasswordComplexityPolicy

> **rpc** UpdatePasswordComplexityPolicy([UpdatePasswordComplexityPolicyRequest](#updatepasswordcomplexitypolicyrequest))
[UpdatePasswordComplexityPolicyResponse](#updatepasswordcomplexitypolicyresponse)

Updates the default password complexity policy of ZITADEL
it impacts all organisations without a customised policy



    PUT: /policies/password/complexity


### GetPasswordAgePolicy

> **rpc** GetPasswordAgePolicy([GetPasswordAgePolicyRequest](#getpasswordagepolicyrequest))
[GetPasswordAgePolicyResponse](#getpasswordagepolicyresponse)

Returns the password age policy defined by the administrators of ZITADEL



    GET: /policies/password/age


### UpdatePasswordAgePolicy

> **rpc** UpdatePasswordAgePolicy([UpdatePasswordAgePolicyRequest](#updatepasswordagepolicyrequest))
[UpdatePasswordAgePolicyResponse](#updatepasswordagepolicyresponse)

Updates the default password age policy of ZITADEL
it impacts all organisations without a customised policy



    PUT: /policies/password/age


### GetLockoutPolicy

> **rpc** GetLockoutPolicy([GetLockoutPolicyRequest](#getlockoutpolicyrequest))
[GetLockoutPolicyResponse](#getlockoutpolicyresponse)

Returns the lockout policy defined by the administrators of ZITADEL



    GET: /policies/lockout


### UpdateLockoutPolicy

> **rpc** UpdateLockoutPolicy([UpdateLockoutPolicyRequest](#updatelockoutpolicyrequest))
[UpdateLockoutPolicyResponse](#updatelockoutpolicyresponse)

Updates the default lockout policy of ZITADEL
it impacts all organisations without a customised policy



    PUT: /policies/password/lockout


### GetPrivacyPolicy

> **rpc** GetPrivacyPolicy([GetPrivacyPolicyRequest](#getprivacypolicyrequest))
[GetPrivacyPolicyResponse](#getprivacypolicyresponse)

Returns the privacy policy defined by the administrators of ZITADEL



    GET: /policies/privacy


### UpdatePrivacyPolicy

> **rpc** UpdatePrivacyPolicy([UpdatePrivacyPolicyRequest](#updateprivacypolicyrequest))
[UpdatePrivacyPolicyResponse](#updateprivacypolicyresponse)

Updates the default privacy policy of ZITADEL
it impacts all organisations without a customised policy
Variable {{.Lang}} can be set to have different links based on the language



    PUT: /policies/privacy


### GetDefaultInitMessageText

> **rpc** GetDefaultInitMessageText([GetDefaultInitMessageTextRequest](#getdefaultinitmessagetextrequest))
[GetDefaultInitMessageTextResponse](#getdefaultinitmessagetextresponse)

Returns the default text for initial message (translation file)



    GET: /text/default/message/init/{language}


### GetCustomInitMessageText

> **rpc** GetCustomInitMessageText([GetCustomInitMessageTextRequest](#getcustominitmessagetextrequest))
[GetCustomInitMessageTextResponse](#getcustominitmessagetextresponse)

Returns the custom text for initial message (overwritten in eventstore)



    GET: /text/message/init/{language}


### SetDefaultInitMessageText

> **rpc** SetDefaultInitMessageText([SetDefaultInitMessageTextRequest](#setdefaultinitmessagetextrequest))
[SetDefaultInitMessageTextResponse](#setdefaultinitmessagetextresponse)

Sets the default custom text for initial message
it impacts all organisations without customized initial message text
The Following Variables can be used:
{{.Code}} {{.UserName}} {{.FirstName}} {{.LastName}} {{.NickName}} {{.DisplayName}} {{.LastEmail}} {{.VerifiedEmail}} {{.LastPhone}} {{.VerifiedPhone}} {{.PreferredLoginName}} {{.LoginNames}} {{.ChangeDate}}



    PUT: /text/message/init/{language}


### ResetCustomInitMessageTextToDefault

> **rpc** ResetCustomInitMessageTextToDefault([ResetCustomInitMessageTextToDefaultRequest](#resetcustominitmessagetexttodefaultrequest))
[ResetCustomInitMessageTextToDefaultResponse](#resetcustominitmessagetexttodefaultresponse)

Removes the custom init message text of the system
The default text from the translation file will trigger after



    DELETE: /text/message/init/{language}


### GetDefaultPasswordResetMessageText

> **rpc** GetDefaultPasswordResetMessageText([GetDefaultPasswordResetMessageTextRequest](#getdefaultpasswordresetmessagetextrequest))
[GetDefaultPasswordResetMessageTextResponse](#getdefaultpasswordresetmessagetextresponse)

Returns the default text for password reset message (translation file)



    GET: /text/deafult/message/passwordreset/{language}


### GetCustomPasswordResetMessageText

> **rpc** GetCustomPasswordResetMessageText([GetCustomPasswordResetMessageTextRequest](#getcustompasswordresetmessagetextrequest))
[GetCustomPasswordResetMessageTextResponse](#getcustompasswordresetmessagetextresponse)

Returns the custom text for password reset message (overwritten in eventstore)



    GET: /text/message/passwordreset/{language}


### SetDefaultPasswordResetMessageText

> **rpc** SetDefaultPasswordResetMessageText([SetDefaultPasswordResetMessageTextRequest](#setdefaultpasswordresetmessagetextrequest))
[SetDefaultPasswordResetMessageTextResponse](#setdefaultpasswordresetmessagetextresponse)

Sets the default custom text for password reset message
it impacts all organisations without customized password reset message text
The Following Variables can be used:
{{.Code}} {{.UserName}} {{.FirstName}} {{.LastName}} {{.NickName}} {{.DisplayName}} {{.LastEmail}} {{.VerifiedEmail}} {{.LastPhone}} {{.VerifiedPhone}} {{.PreferredLoginName}} {{.LoginNames}} {{.ChangeDate}}



    PUT: /text/message/passwordreset/{language}


### ResetCustomPasswordResetMessageTextToDefault

> **rpc** ResetCustomPasswordResetMessageTextToDefault([ResetCustomPasswordResetMessageTextToDefaultRequest](#resetcustompasswordresetmessagetexttodefaultrequest))
[ResetCustomPasswordResetMessageTextToDefaultResponse](#resetcustompasswordresetmessagetexttodefaultresponse)

Removes the custom password reset message text of the system
The default text from the translation file will trigger after



    DELETE: /text/message/passwordreset/{language}


### GetDefaultVerifyEmailMessageText

> **rpc** GetDefaultVerifyEmailMessageText([GetDefaultVerifyEmailMessageTextRequest](#getdefaultverifyemailmessagetextrequest))
[GetDefaultVerifyEmailMessageTextResponse](#getdefaultverifyemailmessagetextresponse)

Returns the default text for verify email message (translation files)



    GET: /text/default/message/verifyemail/{language}


### GetCustomVerifyEmailMessageText

> **rpc** GetCustomVerifyEmailMessageText([GetCustomVerifyEmailMessageTextRequest](#getcustomverifyemailmessagetextrequest))
[GetCustomVerifyEmailMessageTextResponse](#getcustomverifyemailmessagetextresponse)

Returns the custom text for verify email message (overwritten in eventstore)



    GET: /text/message/verifyemail/{language}


### SetDefaultVerifyEmailMessageText

> **rpc** SetDefaultVerifyEmailMessageText([SetDefaultVerifyEmailMessageTextRequest](#setdefaultverifyemailmessagetextrequest))
[SetDefaultVerifyEmailMessageTextResponse](#setdefaultverifyemailmessagetextresponse)

Sets the default custom text for verify email message
it impacts all organisations without customized verify email message text
The Following Variables can be used:
{{.Code}} {{.UserName}} {{.FirstName}} {{.LastName}} {{.NickName}} {{.DisplayName}} {{.LastEmail}} {{.VerifiedEmail}} {{.LastPhone}} {{.VerifiedPhone}} {{.PreferredLoginName}} {{.LoginNames}} {{.ChangeDate}}



    PUT: /text/message/verifyemail/{language}


### ResetCustomVerifyEmailMessageTextToDefault

> **rpc** ResetCustomVerifyEmailMessageTextToDefault([ResetCustomVerifyEmailMessageTextToDefaultRequest](#resetcustomverifyemailmessagetexttodefaultrequest))
[ResetCustomVerifyEmailMessageTextToDefaultResponse](#resetcustomverifyemailmessagetexttodefaultresponse)

Removes the custom verify email message text of the system
The default text from the translation file will trigger after



    DELETE: /text/message/verifyemail/{language}


### GetDefaultVerifyPhoneMessageText

> **rpc** GetDefaultVerifyPhoneMessageText([GetDefaultVerifyPhoneMessageTextRequest](#getdefaultverifyphonemessagetextrequest))
[GetDefaultVerifyPhoneMessageTextResponse](#getdefaultverifyphonemessagetextresponse)

Returns the default text for verify phone message (translation file)



    GET: /text/default/message/verifyphone/{language}


### GetCustomVerifyPhoneMessageText

> **rpc** GetCustomVerifyPhoneMessageText([GetCustomVerifyPhoneMessageTextRequest](#getcustomverifyphonemessagetextrequest))
[GetCustomVerifyPhoneMessageTextResponse](#getcustomverifyphonemessagetextresponse)

Returns the custom text for verify phone message



    GET: /text/message/verifyphone/{language}


### SetDefaultVerifyPhoneMessageText

> **rpc** SetDefaultVerifyPhoneMessageText([SetDefaultVerifyPhoneMessageTextRequest](#setdefaultverifyphonemessagetextrequest))
[SetDefaultVerifyPhoneMessageTextResponse](#setdefaultverifyphonemessagetextresponse)

Sets the default custom text for verify phone message
it impacts all organisations without customized verify phone message text
The Following Variables can be used:
{{.Code}} {{.UserName}} {{.FirstName}} {{.LastName}} {{.NickName}} {{.DisplayName}} {{.LastEmail}} {{.VerifiedEmail}} {{.LastPhone}} {{.VerifiedPhone}} {{.PreferredLoginName}} {{.LoginNames}} {{.ChangeDate}}



    PUT: /text/message/verifyphone/{language}


### ResetCustomVerifyPhoneMessageTextToDefault

> **rpc** ResetCustomVerifyPhoneMessageTextToDefault([ResetCustomVerifyPhoneMessageTextToDefaultRequest](#resetcustomverifyphonemessagetexttodefaultrequest))
[ResetCustomVerifyPhoneMessageTextToDefaultResponse](#resetcustomverifyphonemessagetexttodefaultresponse)

Removes the custom verify phone text of the system
The default text from the translation file will trigger after



    DELETE: /text/message/verifyphone/{language}


### GetDefaultDomainClaimedMessageText

> **rpc** GetDefaultDomainClaimedMessageText([GetDefaultDomainClaimedMessageTextRequest](#getdefaultdomainclaimedmessagetextrequest))
[GetDefaultDomainClaimedMessageTextResponse](#getdefaultdomainclaimedmessagetextresponse)

Returns the default text for domain claimed message (translation file)



    GET: /text/default/message/domainclaimed/{language}


### GetCustomDomainClaimedMessageText

> **rpc** GetCustomDomainClaimedMessageText([GetCustomDomainClaimedMessageTextRequest](#getcustomdomainclaimedmessagetextrequest))
[GetCustomDomainClaimedMessageTextResponse](#getcustomdomainclaimedmessagetextresponse)

Returns the custom text for domain claimed message (overwritten in eventstore)



    GET: /text/message/domainclaimed/{language}


### SetDefaultDomainClaimedMessageText

> **rpc** SetDefaultDomainClaimedMessageText([SetDefaultDomainClaimedMessageTextRequest](#setdefaultdomainclaimedmessagetextrequest))
[SetDefaultDomainClaimedMessageTextResponse](#setdefaultdomainclaimedmessagetextresponse)

Sets the default custom text for domain claimed phone message
it impacts all organisations without customized domain claimed message text
The Following Variables can be used:
{{.Domain}} {{.TempUsername}} {{.UserName}} {{.FirstName}} {{.LastName}} {{.NickName}} {{.DisplayName}} {{.LastEmail}} {{.VerifiedEmail}} {{.LastPhone}} {{.VerifiedPhone}} {{.PreferredLoginName}} {{.LoginNames}} {{.ChangeDate}}



    PUT: /text/message/domainclaimed/{language}


### ResetCustomDomainClaimedMessageTextToDefault

> **rpc** ResetCustomDomainClaimedMessageTextToDefault([ResetCustomDomainClaimedMessageTextToDefaultRequest](#resetcustomdomainclaimedmessagetexttodefaultrequest))
[ResetCustomDomainClaimedMessageTextToDefaultResponse](#resetcustomdomainclaimedmessagetexttodefaultresponse)

Removes the custom domain claimed message text of the system
The default text from the translation file will trigger after



    DELETE: /text/message/domainclaimed/{language}


### GetDefaultPasswordlessRegistrationMessageText

> **rpc** GetDefaultPasswordlessRegistrationMessageText([GetDefaultPasswordlessRegistrationMessageTextRequest](#getdefaultpasswordlessregistrationmessagetextrequest))
[GetDefaultPasswordlessRegistrationMessageTextResponse](#getdefaultpasswordlessregistrationmessagetextresponse)

Returns the default text for passwordless registration message (translation file)



    GET: /text/default/message/passwordless_registration/{language}


### GetCustomPasswordlessRegistrationMessageText

> **rpc** GetCustomPasswordlessRegistrationMessageText([GetCustomPasswordlessRegistrationMessageTextRequest](#getcustompasswordlessregistrationmessagetextrequest))
[GetCustomPasswordlessRegistrationMessageTextResponse](#getcustompasswordlessregistrationmessagetextresponse)

Returns the custom text for passwordless registration message (overwritten in eventstore)



    GET: /text/message/passwordless_registration/{language}


### SetDefaultPasswordlessRegistrationMessageText

> **rpc** SetDefaultPasswordlessRegistrationMessageText([SetDefaultPasswordlessRegistrationMessageTextRequest](#setdefaultpasswordlessregistrationmessagetextrequest))
[SetDefaultPasswordlessRegistrationMessageTextResponse](#setdefaultpasswordlessregistrationmessagetextresponse)

Sets the default custom text for passwordless registration message
it impacts all organisations without customized passwordless registration message text
The Following Variables can be used:
{{.UserName}} {{.FirstName}} {{.LastName}} {{.NickName}} {{.DisplayName}} {{.LastEmail}} {{.VerifiedEmail}} {{.LastPhone}} {{.VerifiedPhone}} {{.PreferredLoginName}} {{.LoginNames}} {{.ChangeDate}}



    PUT: /text/message/passwordless_registration/{language}


### ResetCustomPasswordlessRegistrationMessageTextToDefault

> **rpc** ResetCustomPasswordlessRegistrationMessageTextToDefault([ResetCustomPasswordlessRegistrationMessageTextToDefaultRequest](#resetcustompasswordlessregistrationmessagetexttodefaultrequest))
[ResetCustomPasswordlessRegistrationMessageTextToDefaultResponse](#resetcustompasswordlessregistrationmessagetexttodefaultresponse)

Removes the custom passwordless link message text of the system
The default text from the translation file will trigger after



    DELETE: /text/message/passwordless_registration/{language}


### GetDefaultLoginTexts

> **rpc** GetDefaultLoginTexts([GetDefaultLoginTextsRequest](#getdefaultlogintextsrequest))
[GetDefaultLoginTextsResponse](#getdefaultlogintextsresponse)

Returns the default custom texts for login ui (translation file)



    GET: /text/default/login/{language}


### GetCustomLoginTexts

> **rpc** GetCustomLoginTexts([GetCustomLoginTextsRequest](#getcustomlogintextsrequest))
[GetCustomLoginTextsResponse](#getcustomlogintextsresponse)

Returns the custom texts for login ui



    GET: /text/login/{language}


### SetCustomLoginText

> **rpc** SetCustomLoginText([SetCustomLoginTextsRequest](#setcustomlogintextsrequest))
[SetCustomLoginTextsResponse](#setcustomlogintextsresponse)

Sets the custom text for login ui
it impacts all organisations without customized login ui texts



    PUT: /text/login/{language}


### ResetCustomLoginTextToDefault

> **rpc** ResetCustomLoginTextToDefault([ResetCustomLoginTextsToDefaultRequest](#resetcustomlogintextstodefaultrequest))
[ResetCustomLoginTextsToDefaultResponse](#resetcustomlogintextstodefaultresponse)

Removes the custom texts for login ui
it impacts all organisations without customized login ui texts
The default text form translation file will trigger after



    DELETE: /text/login/{language}


### ListIAMMemberRoles

> **rpc** ListIAMMemberRoles([ListIAMMemberRolesRequest](#listiammemberrolesrequest))
[ListIAMMemberRolesResponse](#listiammemberrolesresponse)

Returns the IAM roles visible for the requested user



    POST: /members/roles/_search


### ListIAMMembers

> **rpc** ListIAMMembers([ListIAMMembersRequest](#listiammembersrequest))
[ListIAMMembersResponse](#listiammembersresponse)

Returns all members matching the request
all queries need to match (ANDed)



    POST: /members/_search


### AddIAMMember

> **rpc** AddIAMMember([AddIAMMemberRequest](#addiammemberrequest))
[AddIAMMemberResponse](#addiammemberresponse)

Adds a user to the membership list of ZITADEL with the given roles
undefined roles will be dropped



    POST: /members


### UpdateIAMMember

> **rpc** UpdateIAMMember([UpdateIAMMemberRequest](#updateiammemberrequest))
[UpdateIAMMemberResponse](#updateiammemberresponse)

Sets the given roles on a member.
The member has only roles provided by this call



    PUT: /members/{user_id}


### RemoveIAMMember

> **rpc** RemoveIAMMember([RemoveIAMMemberRequest](#removeiammemberrequest))
[RemoveIAMMemberResponse](#removeiammemberresponse)

Removes the user from the membership list of ZITADEL



    DELETE: /members/{user_id}


### ListViews

> **rpc** ListViews([ListViewsRequest](#listviewsrequest))
[ListViewsResponse](#listviewsresponse)

Returns all stored read models of ZITADEL
views are used for search optimisation and optimise request latencies
they represent the delta of the event happend on the objects



    POST: /views/_search


### ClearView

> **rpc** ClearView([ClearViewRequest](#clearviewrequest))
[ClearViewResponse](#clearviewresponse)

Truncates the delta of the change stream
be carefull with this function because ZITADEL has to 
recompute the deltas after they got cleared. 
Search requests will return wrong results until all deltas are recomputed



    POST: /views/{database}/{view_name}


### ListFailedEvents

> **rpc** ListFailedEvents([ListFailedEventsRequest](#listfailedeventsrequest))
[ListFailedEventsResponse](#listfailedeventsresponse)

Returns event descriptions which cannot be processed.
It's possible that some events need some retries. 
For example if the SMTP-API wasn't able to send an email at the first time



    POST: /failedevents/_search


### RemoveFailedEvent

> **rpc** RemoveFailedEvent([RemoveFailedEventRequest](#removefailedeventrequest))
[RemoveFailedEventResponse](#removefailedeventresponse)

Deletes the event from failed events view.
the event is not removed from the change stream
This call is usefull if the system was able to process the event later. 
e.g. if the second try of sending an email was successful. the first try produced a
failed event. You can find out if it worked on the `failure_count`



    DELETE: /failedevents/{database}/{view_name}/{failed_sequence}


### ExportData

> **rpc** ExportData([ExportDataRequest](#exportdatarequest))
[ExportDataResponse](#exportdataresponse)

Export data into instance and creates different objects



    POST: /export







## Messages


### ActivateLabelPolicyRequest
This is an empty request




### ActivateLabelPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### AddCustomOrgIAMPolicyRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| org_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| user_login_must_be_domain |  bool | the username has to end with the domain of it's organisation (uniqueness is organisation based) |  |




### AddCustomOrgIAMPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### AddIAMMemberRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| user_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| roles | repeated string | if no roles provided the user won't have any rights |  |




### AddIAMMemberResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### AddIDPToLoginPolicyRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| idp_id |  string | Id of the predefined idp configuration | string.min_len: 1<br /> string.max_len: 200<br />  |




### AddIDPToLoginPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### AddJWTIDPRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| styling_type |  zitadel.idp.v1.IDPStylingType | - | enum.defined_only: true<br />  |
| jwt_endpoint |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| issuer |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| keys_endpoint |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| header_name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| auto_register |  bool | - |  |




### AddJWTIDPResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |
| idp_id |  string | - |  |




### AddMultiFactorToLoginPolicyRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| type |  zitadel.policy.v1.MultiFactorType | - | enum.defined_only: true<br /> enum.not_in: [0]<br />  |




### AddMultiFactorToLoginPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### AddOIDCIDPRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| styling_type |  zitadel.idp.v1.IDPStylingType | - | enum.defined_only: true<br />  |
| client_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| client_secret |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| issuer |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| scopes | repeated string | - |  |
| display_name_mapping |  zitadel.idp.v1.OIDCMappingField | - | enum.defined_only: true<br />  |
| username_mapping |  zitadel.idp.v1.OIDCMappingField | - | enum.defined_only: true<br />  |
| auto_register |  bool | - |  |




### AddOIDCIDPResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |
| idp_id |  string | - |  |




### AddSecondFactorToLoginPolicyRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| type |  zitadel.policy.v1.SecondFactorType | - | enum.defined_only: true<br /> enum.not_in: [0]<br />  |




### AddSecondFactorToLoginPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### ClearViewRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| database |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| view_name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### ClearViewResponse
This is an empty response




### DataAPIApplication



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| app_id |  string | - |  |
| app |  zitadel.management.v1.AddAPIAppRequest | - |  |




### DataAction



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| action_id |  string | - |  |
| action |  zitadel.management.v1.CreateActionRequest | - |  |




### DataHumanUser



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| user_id |  string | - |  |
| user |  ExportHumanUser | - |  |




### DataJWTIDP



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| idp_id |  string | - |  |
| idp |  zitadel.management.v1.AddOrgJWTIDPRequest | - |  |




### DataMachineUser



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| user_id |  string | - |  |
| user |  zitadel.management.v1.AddMachineUserRequest | - |  |




### DataOIDCApplication



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| app_id |  string | - |  |
| app |  zitadel.management.v1.AddOIDCAppRequest | - |  |




### DataOIDCIDP



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| idp_id |  string | - |  |
| idp |  zitadel.management.v1.AddOrgOIDCIDPRequest | - |  |




### DataOrg



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| org_id |  string | - |  |
| org |  zitadel.management.v1.AddOrgRequest | - |  |
| iam_policy |  AddCustomOrgIAMPolicyRequest | - |  |
| label_policy |  zitadel.management.v1.AddCustomLabelPolicyRequest | - |  |
| lockout_policy |  zitadel.management.v1.AddCustomLockoutPolicyRequest | - |  |
| login_policy |  zitadel.management.v1.AddCustomLoginPolicyRequest | - |  |
| password_complexity_policy |  zitadel.management.v1.AddCustomPasswordComplexityPolicyRequest | - |  |
| privacy_policy |  zitadel.management.v1.AddCustomPrivacyPolicyRequest | - |  |
| projects | repeated DataProject | - |  |
| project_roles | repeated zitadel.management.v1.AddProjectRoleRequest | - |  |
| api_apps | repeated DataAPIApplication | - |  |
| oidc_apps | repeated DataOIDCApplication | - |  |
| human_users | repeated DataHumanUser | - |  |
| machine_users | repeated DataMachineUser | - |  |
| trigger_actions | repeated zitadel.management.v1.SetTriggerActionsRequest | - |  |
| actions | repeated DataAction | - |  |
| project_grants | repeated DataProjectGrant | - |  |
| user_grants | repeated zitadel.management.v1.AddUserGrantRequest | - |  |
| org_members | repeated zitadel.management.v1.AddOrgMemberRequest | - |  |
| project_members | repeated zitadel.management.v1.AddProjectMemberRequest | - |  |
| project_grant_members | repeated zitadel.management.v1.AddProjectGrantMemberRequest | - |  |
| user_metadata | repeated zitadel.management.v1.SetUserMetadataRequest | - |  |
| login_texts | repeated zitadel.management.v1.SetCustomLoginTextsRequest | - |  |
| init_messages | repeated zitadel.management.v1.SetCustomInitMessageTextRequest | - |  |
| password_reset_messages | repeated zitadel.management.v1.SetCustomPasswordResetMessageTextRequest | - |  |
| verify_email_messages | repeated zitadel.management.v1.SetCustomVerifyEmailMessageTextRequest | - |  |
| verify_phone_messages | repeated zitadel.management.v1.SetCustomVerifyPhoneMessageTextRequest | - |  |
| domain_claimed_messages | repeated zitadel.management.v1.SetCustomDomainClaimedMessageTextRequest | - |  |
| passwordless_registration_messages | repeated zitadel.management.v1.SetCustomPasswordlessRegistrationMessageTextRequest | - |  |
| oidc_idps | repeated DataOIDCIDP | - |  |
| jwt_idps | repeated DataJWTIDP | - |  |
| second_factors | repeated zitadel.management.v1.AddSecondFactorToLoginPolicyRequest | - |  |
| multi_factors | repeated zitadel.management.v1.AddMultiFactorToLoginPolicyRequest | - |  |
| idps | repeated zitadel.management.v1.AddIDPToLoginPolicyRequest | - |  |
| user_links | repeated zitadel.idp.v1.IDPUserLink | - |  |
| domains | repeated zitadel.org.v1.Domain | - |  |




### DataProject



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| project_id |  string | - |  |
| project |  zitadel.management.v1.AddProjectRequest | - |  |




### DataProjectGrant



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| grant_id |  string | - |  |
| project_grant |  zitadel.management.v1.AddProjectGrantRequest | - |  |




### DeactivateIDPRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| idp_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### DeactivateIDPResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### ExportDataRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| org_ids | repeated string | - |  |
| excluded_org_ids | repeated string | - |  |
| with_passwords |  bool | - |  |
| with_otp |  bool | - |  |




### ExportDataResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| orgs | repeated DataOrg | - |  |




### ExportHumanUser



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| user_name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| profile |  ExportHumanUser.Profile | - | message.required: true<br />  |
| email |  ExportHumanUser.Email | - | message.required: true<br />  |
| phone |  ExportHumanUser.Phone | - |  |
| password |  string | - |  |
| hashed_password |  ExportHumanUser.HashedPassword | - |  |
| password_change_required |  bool | - |  |
| request_passwordless_registration |  bool | - |  |
| otp_code |  string | - |  |




### ExportHumanUser.Email



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| email |  string | TODO: check if no value is allowed | string.email: true<br />  |
| is_email_verified |  bool | - |  |




### ExportHumanUser.HashedPassword



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| value |  string | - |  |
| algorithm |  string | - |  |




### ExportHumanUser.Phone



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| phone |  string | has to be a global number | string.min_len: 1<br /> string.max_len: 50<br /> string.prefix: +<br />  |
| is_phone_verified |  bool | - |  |




### ExportHumanUser.Profile



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| first_name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| last_name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| nick_name |  string | - | string.max_len: 200<br />  |
| display_name |  string | - | string.max_len: 200<br />  |
| preferred_language |  string | - | string.max_len: 10<br />  |
| gender |  zitadel.user.v1.Gender | - |  |




### FailedEvent



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| database |  string | - |  |
| view_name |  string | - |  |
| failed_sequence |  uint64 | - |  |
| failure_count |  uint64 | - |  |
| error_message |  string | - |  |




### GetCustomDomainClaimedMessageTextRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### GetCustomDomainClaimedMessageTextResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| custom_text |  zitadel.text.v1.MessageCustomText | - |  |




### GetCustomInitMessageTextRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### GetCustomInitMessageTextResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| custom_text |  zitadel.text.v1.MessageCustomText | - |  |




### GetCustomLoginTextsRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### GetCustomLoginTextsResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| custom_text |  zitadel.text.v1.LoginCustomText | - |  |




### GetCustomOrgIAMPolicyRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| org_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### GetCustomOrgIAMPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| policy |  zitadel.policy.v1.OrgIAMPolicy | - |  |
| is_default |  bool | deprecated: is_default is also defined in zitadel.policy.v1.OrgIAMPolicy |  |




### GetCustomPasswordResetMessageTextRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### GetCustomPasswordResetMessageTextResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| custom_text |  zitadel.text.v1.MessageCustomText | - |  |




### GetCustomPasswordlessRegistrationMessageTextRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### GetCustomPasswordlessRegistrationMessageTextResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| custom_text |  zitadel.text.v1.MessageCustomText | - |  |




### GetCustomVerifyEmailMessageTextRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### GetCustomVerifyEmailMessageTextResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| custom_text |  zitadel.text.v1.MessageCustomText | - |  |




### GetCustomVerifyPhoneMessageTextRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### GetCustomVerifyPhoneMessageTextResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| custom_text |  zitadel.text.v1.MessageCustomText | - |  |




### GetDefaultDomainClaimedMessageTextRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### GetDefaultDomainClaimedMessageTextResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| custom_text |  zitadel.text.v1.MessageCustomText | - |  |




### GetDefaultFeaturesRequest





### GetDefaultFeaturesResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| features |  zitadel.features.v1.Features | - |  |




### GetDefaultInitMessageTextRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### GetDefaultInitMessageTextResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| custom_text |  zitadel.text.v1.MessageCustomText | - |  |




### GetDefaultLoginTextsRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### GetDefaultLoginTextsResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| custom_text |  zitadel.text.v1.LoginCustomText | - |  |




### GetDefaultPasswordResetMessageTextRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### GetDefaultPasswordResetMessageTextResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| custom_text |  zitadel.text.v1.MessageCustomText | - |  |




### GetDefaultPasswordlessRegistrationMessageTextRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### GetDefaultPasswordlessRegistrationMessageTextResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| custom_text |  zitadel.text.v1.MessageCustomText | - |  |




### GetDefaultVerifyEmailMessageTextRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### GetDefaultVerifyEmailMessageTextResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| custom_text |  zitadel.text.v1.MessageCustomText | - |  |




### GetDefaultVerifyPhoneMessageTextRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### GetDefaultVerifyPhoneMessageTextResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| custom_text |  zitadel.text.v1.MessageCustomText | - |  |




### GetIDPByIDRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### GetIDPByIDResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| idp |  zitadel.idp.v1.IDP | - |  |




### GetLabelPolicyRequest
This is an empty request




### GetLabelPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| policy |  zitadel.policy.v1.LabelPolicy | - |  |




### GetLockoutPolicyRequest
This is an empty request




### GetLockoutPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| policy |  zitadel.policy.v1.LockoutPolicy | - |  |




### GetLoginPolicyRequest
This is an empty request




### GetLoginPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| policy |  zitadel.policy.v1.LoginPolicy | - |  |




### GetOrgByIDRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### GetOrgByIDResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| org |  zitadel.org.v1.Org | - |  |




### GetOrgFeaturesRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| org_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### GetOrgFeaturesResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| features |  zitadel.features.v1.Features | - |  |




### GetOrgIAMPolicyRequest





### GetOrgIAMPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| policy |  zitadel.policy.v1.OrgIAMPolicy | - |  |




### GetPasswordAgePolicyRequest
This is an empty request




### GetPasswordAgePolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| policy |  zitadel.policy.v1.PasswordAgePolicy | - |  |




### GetPasswordComplexityPolicyRequest





### GetPasswordComplexityPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| policy |  zitadel.policy.v1.PasswordComplexityPolicy | - |  |




### GetPreviewLabelPolicyRequest
This is an empty request




### GetPreviewLabelPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| policy |  zitadel.policy.v1.LabelPolicy | - |  |




### GetPrivacyPolicyRequest
This is an empty request




### GetPrivacyPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| policy |  zitadel.policy.v1.PrivacyPolicy | - |  |




### GetSupportedLanguagesRequest
This is an empty request




### GetSupportedLanguagesResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| languages | repeated string | - |  |




### HealthzRequest
This is an empty request




### HealthzResponse
This is an empty response




### IDPQuery



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) query.idp_id_query |  zitadel.idp.v1.IDPIDQuery | - |  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) query.idp_name_query |  zitadel.idp.v1.IDPNameQuery | - |  |




### IsOrgUniqueRequest
if name or domain is already in use, org is not unique


| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| domain |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### IsOrgUniqueResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| is_unique |  bool | - |  |




### ListFailedEventsRequest
This is an empty request




### ListFailedEventsResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| result | repeated FailedEvent | TODO: list details |  |




### ListIAMMemberRolesRequest
This is an empty request




### ListIAMMemberRolesResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ListDetails | - |  |
| roles | repeated string | - |  |




### ListIAMMembersRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| query |  zitadel.v1.ListQuery | list limitations and ordering |  |
| queries | repeated zitadel.member.v1.SearchQuery | criterias the client is looking for |  |




### ListIAMMembersResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ListDetails | - |  |
| result | repeated zitadel.member.v1.Member | - |  |




### ListIDPsRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| query |  zitadel.v1.ListQuery | list limitations and ordering |  |
| sorting_column |  zitadel.idp.v1.IDPFieldName | the field the result is sorted |  |
| queries | repeated IDPQuery | criterias the client is looking for |  |




### ListIDPsResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ListDetails | - |  |
| sorting_column |  zitadel.idp.v1.IDPFieldName | - |  |
| result | repeated zitadel.idp.v1.IDP | - |  |




### ListLoginPolicyIDPsRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| query |  zitadel.v1.ListQuery | list limitations and ordering |  |




### ListLoginPolicyIDPsResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ListDetails | - |  |
| result | repeated zitadel.idp.v1.IDPLoginPolicyLink | - |  |




### ListLoginPolicyMultiFactorsRequest
This is an empty request




### ListLoginPolicyMultiFactorsResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ListDetails | - |  |
| result | repeated zitadel.policy.v1.MultiFactorType | - |  |




### ListLoginPolicySecondFactorsRequest
This is an empty request




### ListLoginPolicySecondFactorsResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ListDetails | - |  |
| result | repeated zitadel.policy.v1.SecondFactorType | - |  |




### ListOrgsRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| query |  zitadel.v1.ListQuery | list limitations and ordering |  |
| sorting_column |  zitadel.org.v1.OrgFieldName | the field the result is sorted |  |
| queries | repeated zitadel.org.v1.OrgQuery | criterias the client is looking for |  |




### ListOrgsResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ListDetails | - |  |
| sorting_column |  zitadel.org.v1.OrgFieldName | - |  |
| result | repeated zitadel.org.v1.Org | - |  |




### ListViewsRequest
This is an empty request




### ListViewsResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| result | repeated View | TODO: list details |  |




### ReactivateIDPRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| idp_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### ReactivateIDPResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### RemoveFailedEventRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| database |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| view_name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| failed_sequence |  uint64 | - |  |




### RemoveFailedEventResponse
This is an empty response




### RemoveIAMMemberRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| user_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### RemoveIAMMemberResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### RemoveIDPFromLoginPolicyRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| idp_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### RemoveIDPFromLoginPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### RemoveIDPRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| idp_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### RemoveIDPResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### RemoveLabelPolicyFontRequest
This is an empty request




### RemoveLabelPolicyFontResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### RemoveLabelPolicyIconDarkRequest
This is an empty request




### RemoveLabelPolicyIconDarkResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### RemoveLabelPolicyIconRequest
This is an empty request




### RemoveLabelPolicyIconResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### RemoveLabelPolicyLogoDarkRequest
This is an empty request




### RemoveLabelPolicyLogoDarkResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### RemoveLabelPolicyLogoRequest
This is an empty request




### RemoveLabelPolicyLogoResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### RemoveMultiFactorFromLoginPolicyRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| type |  zitadel.policy.v1.MultiFactorType | - | enum.defined_only: true<br /> enum.not_in: [0]<br />  |




### RemoveMultiFactorFromLoginPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### RemoveSecondFactorFromLoginPolicyRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| type |  zitadel.policy.v1.SecondFactorType | - | enum.defined_only: true<br /> enum.not_in: [0]<br />  |




### RemoveSecondFactorFromLoginPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### ResetCustomDomainClaimedMessageTextToDefaultRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### ResetCustomDomainClaimedMessageTextToDefaultResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### ResetCustomInitMessageTextToDefaultRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### ResetCustomInitMessageTextToDefaultResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### ResetCustomLoginTextsToDefaultRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### ResetCustomLoginTextsToDefaultResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### ResetCustomOrgIAMPolicyToDefaultRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| org_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### ResetCustomOrgIAMPolicyToDefaultResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### ResetCustomPasswordResetMessageTextToDefaultRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### ResetCustomPasswordResetMessageTextToDefaultResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### ResetCustomPasswordlessRegistrationMessageTextToDefaultRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### ResetCustomPasswordlessRegistrationMessageTextToDefaultResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### ResetCustomVerifyEmailMessageTextToDefaultRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### ResetCustomVerifyEmailMessageTextToDefaultResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### ResetCustomVerifyPhoneMessageTextToDefaultRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### ResetCustomVerifyPhoneMessageTextToDefaultResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### ResetOrgFeaturesRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| org_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### ResetOrgFeaturesResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### SetCustomLoginTextsRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| select_account_text |  zitadel.text.v1.SelectAccountScreenText | - |  |
| login_text |  zitadel.text.v1.LoginScreenText | - |  |
| password_text |  zitadel.text.v1.PasswordScreenText | - |  |
| username_change_text |  zitadel.text.v1.UsernameChangeScreenText | - |  |
| username_change_done_text |  zitadel.text.v1.UsernameChangeDoneScreenText | - |  |
| init_password_text |  zitadel.text.v1.InitPasswordScreenText | - |  |
| init_password_done_text |  zitadel.text.v1.InitPasswordDoneScreenText | - |  |
| email_verification_text |  zitadel.text.v1.EmailVerificationScreenText | - |  |
| email_verification_done_text |  zitadel.text.v1.EmailVerificationDoneScreenText | - |  |
| initialize_user_text |  zitadel.text.v1.InitializeUserScreenText | - |  |
| initialize_done_text |  zitadel.text.v1.InitializeUserDoneScreenText | - |  |
| init_mfa_prompt_text |  zitadel.text.v1.InitMFAPromptScreenText | - |  |
| init_mfa_otp_text |  zitadel.text.v1.InitMFAOTPScreenText | - |  |
| init_mfa_u2f_text |  zitadel.text.v1.InitMFAU2FScreenText | - |  |
| init_mfa_done_text |  zitadel.text.v1.InitMFADoneScreenText | - |  |
| mfa_providers_text |  zitadel.text.v1.MFAProvidersText | - |  |
| verify_mfa_otp_text |  zitadel.text.v1.VerifyMFAOTPScreenText | - |  |
| verify_mfa_u2f_text |  zitadel.text.v1.VerifyMFAU2FScreenText | - |  |
| passwordless_text |  zitadel.text.v1.PasswordlessScreenText | - |  |
| password_change_text |  zitadel.text.v1.PasswordChangeScreenText | - |  |
| password_change_done_text |  zitadel.text.v1.PasswordChangeDoneScreenText | - |  |
| password_reset_done_text |  zitadel.text.v1.PasswordResetDoneScreenText | - |  |
| registration_option_text |  zitadel.text.v1.RegistrationOptionScreenText | - |  |
| registration_user_text |  zitadel.text.v1.RegistrationUserScreenText | - |  |
| registration_org_text |  zitadel.text.v1.RegistrationOrgScreenText | - |  |
| linking_user_done_text |  zitadel.text.v1.LinkingUserDoneScreenText | - |  |
| external_user_not_found_text |  zitadel.text.v1.ExternalUserNotFoundScreenText | - |  |
| success_login_text |  zitadel.text.v1.SuccessLoginScreenText | - |  |
| logout_text |  zitadel.text.v1.LogoutDoneScreenText | - |  |
| footer_text |  zitadel.text.v1.FooterText | - |  |
| passwordless_prompt_text |  zitadel.text.v1.PasswordlessPromptScreenText | - |  |
| passwordless_registration_text |  zitadel.text.v1.PasswordlessRegistrationScreenText | - |  |
| passwordless_registration_done_text |  zitadel.text.v1.PasswordlessRegistrationDoneScreenText | - |  |
| external_registration_user_overview_text |  zitadel.text.v1.ExternalRegistrationUserOverviewScreenText | - |  |




### SetCustomLoginTextsResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### SetDefaultDomainClaimedMessageTextRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| title |  string | - | string.max_len: 200<br />  |
| pre_header |  string | - | string.max_len: 200<br />  |
| subject |  string | - | string.max_len: 200<br />  |
| greeting |  string | - | string.max_len: 200<br />  |
| text |  string | - | string.max_len: 800<br />  |
| button_text |  string | - | string.max_len: 200<br />  |
| footer_text |  string | - | string.max_len: 200<br />  |




### SetDefaultDomainClaimedMessageTextResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### SetDefaultFeaturesRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| tier_name |  string | - | string.max_len: 200<br />  |
| description |  string | - | string.max_len: 200<br />  |
| audit_log_retention |  google.protobuf.Duration | - | duration.gte.seconds: 0<br /> duration.gte.nanos: 0<br />  |
| login_policy_username_login |  bool | - |  |
| login_policy_registration |  bool | - |  |
| login_policy_idp |  bool | - |  |
| login_policy_factors |  bool | - |  |
| login_policy_passwordless |  bool | - |  |
| password_complexity_policy |  bool | - |  |
| label_policy |  bool | - |  |
| custom_domain |  bool | - |  |
| login_policy_password_reset |  bool | - |  |
| label_policy_private_label |  bool | - |  |
| label_policy_watermark |  bool | - |  |
| custom_text |  bool | - |  |
| privacy_policy |  bool | - |  |
| metadata_user |  bool | - |  |
| custom_text_message |  bool | - |  |
| custom_text_login |  bool | - |  |
| lockout_policy |  bool | - |  |
| actions |  bool | - |  |
| actions_allowed |  zitadel.features.v1.ActionsAllowed | - |  |
| max_actions |  int32 | - |  |




### SetDefaultFeaturesResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### SetDefaultInitMessageTextRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| title |  string | - | string.max_len: 200<br />  |
| pre_header |  string | - | string.max_len: 200<br />  |
| subject |  string | - | string.max_len: 200<br />  |
| greeting |  string | - | string.max_len: 200<br />  |
| text |  string | - | string.max_len: 1000<br />  |
| button_text |  string | - | string.max_len: 200<br />  |
| footer_text |  string | - | string.max_len: 200<br />  |




### SetDefaultInitMessageTextResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### SetDefaultPasswordResetMessageTextRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| title |  string | - | string.max_len: 200<br />  |
| pre_header |  string | - | string.max_len: 200<br />  |
| subject |  string | - | string.max_len: 200<br />  |
| greeting |  string | - | string.max_len: 200<br />  |
| text |  string | - | string.max_len: 800<br />  |
| button_text |  string | - | string.max_len: 200<br />  |
| footer_text |  string | - | string.max_len: 200<br />  |




### SetDefaultPasswordResetMessageTextResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### SetDefaultPasswordlessRegistrationMessageTextRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| title |  string | - | string.max_len: 200<br />  |
| pre_header |  string | - | string.max_len: 200<br />  |
| subject |  string | - | string.max_len: 200<br />  |
| greeting |  string | - | string.max_len: 200<br />  |
| text |  string | - | string.max_len: 800<br />  |
| button_text |  string | - | string.max_len: 200<br />  |
| footer_text |  string | - | string.max_len: 200<br />  |




### SetDefaultPasswordlessRegistrationMessageTextResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### SetDefaultVerifyEmailMessageTextRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| title |  string | - | string.max_len: 200<br />  |
| pre_header |  string | - | string.max_len: 200<br />  |
| subject |  string | - | string.max_len: 200<br />  |
| greeting |  string | - | string.max_len: 200<br />  |
| text |  string | - | string.max_len: 800<br />  |
| button_text |  string | - | string.max_len: 200<br />  |
| footer_text |  string | - | string.max_len: 200<br />  |




### SetDefaultVerifyEmailMessageTextResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### SetDefaultVerifyPhoneMessageTextRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| language |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| title |  string | - | string.max_len: 200<br />  |
| pre_header |  string | - | string.max_len: 200<br />  |
| subject |  string | - | string.max_len: 200<br />  |
| greeting |  string | - | string.max_len: 200<br />  |
| text |  string | - | string.max_len: 800<br />  |
| button_text |  string | - | string.max_len: 200<br />  |
| footer_text |  string | - | string.max_len: 200<br />  |




### SetDefaultVerifyPhoneMessageTextResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### SetOrgFeaturesRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| org_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| tier_name |  string | - | string.max_len: 200<br />  |
| description |  string | - | string.max_len: 200<br />  |
| state |  zitadel.features.v1.FeaturesState | - |  |
| state_description |  string | - | string.max_len: 200<br />  |
| audit_log_retention |  google.protobuf.Duration | - | duration.gte.seconds: 0<br /> duration.gte.nanos: 0<br />  |
| login_policy_username_login |  bool | - |  |
| login_policy_registration |  bool | - |  |
| login_policy_idp |  bool | - |  |
| login_policy_factors |  bool | - |  |
| login_policy_passwordless |  bool | - |  |
| password_complexity_policy |  bool | - |  |
| label_policy |  bool | - |  |
| custom_domain |  bool | - |  |
| login_policy_password_reset |  bool | - |  |
| label_policy_private_label |  bool | - |  |
| label_policy_watermark |  bool | - |  |
| custom_text |  bool | - |  |
| privacy_policy |  bool | - |  |
| metadata_user |  bool | - |  |
| custom_text_message |  bool | - |  |
| custom_text_login |  bool | - |  |
| lockout_policy |  bool | - |  |
| actions |  bool | - |  |
| actions_allowed |  zitadel.features.v1.ActionsAllowed | - |  |
| max_actions |  int32 | - |  |




### SetOrgFeaturesResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### SetUpOrgRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| org |  SetUpOrgRequest.Org | - | message.required: true<br />  |
| [**oneof**](https://developers.google.com/protocol-buffers/docs/proto3#oneof) user.human |  SetUpOrgRequest.Human | oneof field for the user managing the organisation |  |




### SetUpOrgRequest.Human



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| user_name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| profile |  SetUpOrgRequest.Human.Profile | - | message.required: true<br />  |
| email |  SetUpOrgRequest.Human.Email | - | message.required: true<br />  |
| phone |  SetUpOrgRequest.Human.Phone | - |  |
| password |  string | - |  |




### SetUpOrgRequest.Human.Email



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| email |  string | TODO: check if no value is allowed | string.email: true<br />  |
| is_email_verified |  bool | - |  |




### SetUpOrgRequest.Human.Phone



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| phone |  string | has to be a global number | string.min_len: 1<br /> string.max_len: 50<br /> string.prefix: +<br />  |
| is_phone_verified |  bool | - |  |




### SetUpOrgRequest.Human.Profile



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| first_name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| last_name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| nick_name |  string | - | string.max_len: 200<br />  |
| display_name |  string | - | string.max_len: 200<br />  |
| preferred_language |  string | - | string.max_len: 10<br />  |
| gender |  zitadel.user.v1.Gender | - |  |




### SetUpOrgRequest.Org



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| domain |  string | - | string.max_len: 200<br />  |




### SetUpOrgResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |
| org_id |  string | - |  |
| user_id |  string | - |  |




### UpdateCustomOrgIAMPolicyRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| org_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| user_login_must_be_domain |  bool | - |  |




### UpdateCustomOrgIAMPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### UpdateIAMMemberRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| user_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| roles | repeated string | if no roles provided the user won't have any rights |  |




### UpdateIAMMemberResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### UpdateIDPJWTConfigRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| idp_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| jwt_endpoint |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| issuer |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| keys_endpoint |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| header_name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |




### UpdateIDPJWTConfigResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### UpdateIDPOIDCConfigRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| idp_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| issuer |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| client_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| client_secret |  string | - | string.max_len: 200<br />  |
| scopes | repeated string | - |  |
| display_name_mapping |  zitadel.idp.v1.OIDCMappingField | - | enum.defined_only: true<br />  |
| username_mapping |  zitadel.idp.v1.OIDCMappingField | - | enum.defined_only: true<br />  |




### UpdateIDPOIDCConfigResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### UpdateIDPRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| idp_id |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| name |  string | - | string.min_len: 1<br /> string.max_len: 200<br />  |
| styling_type |  zitadel.idp.v1.IDPStylingType | - | enum.defined_only: true<br />  |
| auto_register |  bool | - |  |




### UpdateIDPResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### UpdateLabelPolicyRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| primary_color |  string | - | string.max_len: 50<br />  |
| hide_login_name_suffix |  bool | - |  |
| warn_color |  string | - | string.max_len: 50<br />  |
| background_color |  string | - | string.max_len: 50<br />  |
| font_color |  string | - | string.max_len: 50<br />  |
| primary_color_dark |  string | - | string.max_len: 50<br />  |
| background_color_dark |  string | - | string.max_len: 50<br />  |
| warn_color_dark |  string | - | string.max_len: 50<br />  |
| font_color_dark |  string | - | string.max_len: 50<br />  |
| disable_watermark |  bool | - |  |




### UpdateLabelPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### UpdateLockoutPolicyRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| max_password_attempts |  uint32 | failed attempts until a user gets locked |  |




### UpdateLockoutPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### UpdateLoginPolicyRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| allow_username_password |  bool | - |  |
| allow_register |  bool | - |  |
| allow_external_idp |  bool | - |  |
| force_mfa |  bool | - |  |
| passwordless_type |  zitadel.policy.v1.PasswordlessType | - | enum.defined_only: true<br />  |
| hide_password_reset |  bool | - |  |
| ignore_unknown_usernames |  bool | - |  |
| default_redirect_uri |  string | - |  |




### UpdateLoginPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### UpdateOrgIAMPolicyRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| user_login_must_be_domain |  bool | - |  |




### UpdateOrgIAMPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### UpdatePasswordAgePolicyRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| max_age_days |  uint32 | - |  |
| expire_warn_days |  uint32 | - |  |




### UpdatePasswordAgePolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### UpdatePasswordComplexityPolicyRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| min_length |  uint32 | - |  |
| has_uppercase |  bool | - |  |
| has_lowercase |  bool | - |  |
| has_number |  bool | - |  |
| has_symbol |  bool | - |  |




### UpdatePasswordComplexityPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### UpdatePrivacyPolicyRequest



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| tos_link |  string | - |  |
| privacy_link |  string | - |  |
| help_link |  string | - |  |




### UpdatePrivacyPolicyResponse



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| details |  zitadel.v1.ObjectDetails | - |  |




### View



| Field | Type | Description | Validation |
| ----- | ---- | ----------- | ----------- |
| database |  string | - |  |
| view_name |  string | - |  |
| processed_sequence |  uint64 | - |  |
| event_timestamp |  google.protobuf.Timestamp | The timestamp the event occured |  |
| last_successful_spooler_run |  google.protobuf.Timestamp | - |  |







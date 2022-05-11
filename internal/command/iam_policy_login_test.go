package command

import (
	"context"
	"testing"

	"github.com/zitadel/zitadel/internal/domain"
	caos_errs "github.com/zitadel/zitadel/internal/errors"
	"github.com/zitadel/zitadel/internal/eventstore"
	"github.com/zitadel/zitadel/internal/eventstore/repository"
	"github.com/zitadel/zitadel/internal/eventstore/v1/models"
	"github.com/zitadel/zitadel/internal/repository/iam"
	"github.com/zitadel/zitadel/internal/repository/policy"
	"github.com/zitadel/zitadel/internal/repository/user"

	"github.com/stretchr/testify/assert"
)

func TestCommandSide_AddDefaultLoginPolicy(t *testing.T) {
	type fields struct {
		eventstore *eventstore.Eventstore
	}
	type args struct {
		ctx    context.Context
		policy *domain.LoginPolicy
	}
	type res struct {
		want *domain.LoginPolicy
		err  func(error) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		res    res
	}{
		{
			name: "loginpolicy already existing, already exists error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							iam.NewLoginPolicyAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								true,
								true,
								false,
								false,
								false,
								false,
								domain.PasswordlessTypeAllowed,
								"",
							),
						),
					),
				),
			},
			args: args{
				ctx: context.Background(),
				policy: &domain.LoginPolicy{
					AllowRegister:         true,
					AllowUsernamePassword: true,
					PasswordlessType:      domain.PasswordlessTypeAllowed,
				},
			},
			res: res{
				err: caos_errs.IsErrorAlreadyExists,
			},
		},
		{
			name: "add policy,ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								iam.NewLoginPolicyAddedEvent(context.Background(),
									&iam.NewAggregate().Aggregate,
									true,
									true,
									true,
									true,
									true,
									true,
									domain.PasswordlessTypeAllowed,
									"https://example.com/redirect",
								),
							),
						},
					),
				),
			},
			args: args{
				ctx: context.Background(),
				policy: &domain.LoginPolicy{
					AllowRegister:          true,
					AllowUsernamePassword:  true,
					AllowExternalIDP:       true,
					ForceMFA:               true,
					HidePasswordReset:      true,
					IgnoreUnknownUsernames: true,
					PasswordlessType:       domain.PasswordlessTypeAllowed,
					DefaultRedirectURI:     "https://example.com/redirect",
				},
			},
			res: res{
				want: &domain.LoginPolicy{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "IAM",
						ResourceOwner: "IAM",
					},
					AllowRegister:          true,
					AllowUsernamePassword:  true,
					AllowExternalIDP:       true,
					ForceMFA:               true,
					HidePasswordReset:      true,
					IgnoreUnknownUsernames: true,
					PasswordlessType:       domain.PasswordlessTypeAllowed,
					DefaultRedirectURI:     "https://example.com/redirect",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Commands{
				eventstore: tt.fields.eventstore,
			}
			got, err := r.AddDefaultLoginPolicy(tt.args.ctx, tt.args.policy)
			if tt.res.err == nil {
				assert.NoError(t, err)
			}
			if tt.res.err != nil && !tt.res.err(err) {
				t.Errorf("got wrong err: %v ", err)
			}
			if tt.res.err == nil {
				assert.Equal(t, tt.res.want, got)
			}
		})
	}
}

func TestCommandSide_ChangeDefaultLoginPolicy(t *testing.T) {
	type fields struct {
		eventstore *eventstore.Eventstore
	}
	type args struct {
		ctx    context.Context
		policy *domain.LoginPolicy
	}
	type res struct {
		want *domain.LoginPolicy
		err  func(error) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		res    res
	}{
		{
			name: "loginpolicy not existing, not found error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(),
				),
			},
			args: args{
				ctx: context.Background(),
				policy: &domain.LoginPolicy{
					AllowRegister:    true,
					AllowExternalIDP: true,
				},
			},
			res: res{
				err: caos_errs.IsNotFound,
			},
		},
		{
			name: "no changes, precondition error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							iam.NewLoginPolicyAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								true,
								true,
								true,
								true,
								true,
								true,
								domain.PasswordlessTypeAllowed,
								"https://example.com/redirect",
							),
						),
					),
				),
			},
			args: args{
				ctx: context.Background(),
				policy: &domain.LoginPolicy{
					AllowRegister:          true,
					AllowUsernamePassword:  true,
					AllowExternalIDP:       true,
					ForceMFA:               true,
					HidePasswordReset:      true,
					IgnoreUnknownUsernames: true,
					PasswordlessType:       domain.PasswordlessTypeAllowed,
					DefaultRedirectURI:     "https://example.com/redirect",
				},
			},
			res: res{
				err: caos_errs.IsPreconditionFailed,
			},
		},
		{
			name: "change, ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							iam.NewLoginPolicyAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								true,
								true,
								true,
								true,
								true,
								true,
								domain.PasswordlessTypeAllowed,
								"https://example.com/redirect",
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								newDefaultLoginPolicyChangedEvent(
									context.Background(),
									false,
									false,
									false,
									false,
									false,
									false,
									domain.PasswordlessTypeNotAllowed,
									"",
								),
							),
						},
					),
				),
			},
			args: args{
				ctx: context.Background(),
				policy: &domain.LoginPolicy{
					AllowRegister:          false,
					AllowUsernamePassword:  false,
					AllowExternalIDP:       false,
					ForceMFA:               false,
					HidePasswordReset:      false,
					IgnoreUnknownUsernames: false,
					PasswordlessType:       domain.PasswordlessTypeNotAllowed,
					DefaultRedirectURI:     "",
				},
			},
			res: res{
				want: &domain.LoginPolicy{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "IAM",
						ResourceOwner: "IAM",
					},
					AllowRegister:          false,
					AllowUsernamePassword:  false,
					AllowExternalIDP:       false,
					ForceMFA:               false,
					HidePasswordReset:      false,
					IgnoreUnknownUsernames: false,
					PasswordlessType:       domain.PasswordlessTypeNotAllowed,
					DefaultRedirectURI:     "",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Commands{
				eventstore: tt.fields.eventstore,
			}
			got, err := r.ChangeDefaultLoginPolicy(tt.args.ctx, tt.args.policy)
			if tt.res.err == nil {
				assert.NoError(t, err)
			}
			if tt.res.err != nil && !tt.res.err(err) {
				t.Errorf("got wrong err: %v ", err)
			}
			if tt.res.err == nil {
				assert.Equal(t, tt.res.want, got)
			}
		})
	}
}

func TestCommandSide_AddIDPProviderDefaultLoginPolicy(t *testing.T) {
	type fields struct {
		eventstore *eventstore.Eventstore
	}
	type args struct {
		ctx      context.Context
		provider *domain.IDPProvider
	}
	type res struct {
		want *domain.IDPProvider
		err  func(error) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		res    res
	}{
		{
			name: "provider invalid, invalid argument error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
				),
			},
			args: args{
				ctx:      context.Background(),
				provider: &domain.IDPProvider{},
			},
			res: res{
				err: caos_errs.IsErrorInvalidArgument,
			},
		},
		{
			name: "policy not existing, not found error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(),
				),
			},
			args: args{
				ctx: context.Background(),
				provider: &domain.IDPProvider{
					IDPConfigID: "config1",
				},
			},
			res: res{
				err: caos_errs.IsNotFound,
			},
		},
		{
			name: "config not existing, precondition error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							iam.NewLoginPolicyAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								true,
								true,
								true,
								true,
								true,
								true,
								domain.PasswordlessTypeAllowed,
								"",
							),
						),
					),
					expectFilter(),
				),
			},
			args: args{
				ctx: context.Background(),
				provider: &domain.IDPProvider{
					IDPConfigID: "config1",
				},
			},
			res: res{
				err: caos_errs.IsPreconditionFailed,
			},
		},
		{
			name: "provider already exists, already exists error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							iam.NewLoginPolicyAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								true,
								true,
								true,
								true,
								true,
								true,
								domain.PasswordlessTypeAllowed,
								"",
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							iam.NewIDPConfigAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								"config1",
								"name",
								domain.IDPConfigTypeOIDC,
								domain.IDPConfigStylingTypeUnspecified,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							iam.NewIdentityProviderAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								"config1",
							),
						),
					),
				),
			},
			args: args{
				ctx: context.Background(),
				provider: &domain.IDPProvider{
					IDPConfigID: "config1",
				},
			},
			res: res{
				err: caos_errs.IsErrorAlreadyExists,
			},
		},
		{
			name: "add provider, ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							iam.NewLoginPolicyAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								true,
								true,
								true,
								true,
								true,
								true,
								domain.PasswordlessTypeAllowed,
								"",
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							iam.NewIDPConfigAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								"config1",
								"name",
								domain.IDPConfigTypeOIDC,
								domain.IDPConfigStylingTypeUnspecified,
								true,
							),
						),
					),
					expectFilter(),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								iam.NewIdentityProviderAddedEvent(context.Background(),
									&iam.NewAggregate().Aggregate,
									"config1"),
							),
						},
					),
				),
			},
			args: args{
				ctx: context.Background(),
				provider: &domain.IDPProvider{
					IDPConfigID: "config1",
				},
			},
			res: res{
				want: &domain.IDPProvider{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "IAM",
						ResourceOwner: "IAM",
					},
					IDPConfigID: "config1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Commands{
				eventstore: tt.fields.eventstore,
			}
			got, err := r.AddIDPProviderToDefaultLoginPolicy(tt.args.ctx, tt.args.provider)
			if tt.res.err == nil {
				assert.NoError(t, err)
			}
			if tt.res.err != nil && !tt.res.err(err) {
				t.Errorf("got wrong err: %v ", err)
			}
			if tt.res.err == nil {
				assert.Equal(t, tt.res.want, got)
			}
		})
	}
}

func TestCommandSide_RemoveIDPProviderDefaultLoginPolicy(t *testing.T) {
	type fields struct {
		eventstore *eventstore.Eventstore
	}
	type args struct {
		ctx                 context.Context
		provider            *domain.IDPProvider
		cascadeExternalIDPs []*domain.UserIDPLink
	}
	type res struct {
		want *domain.ObjectDetails
		err  func(error) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		res    res
	}{
		{
			name: "provider invalid, invalid argument error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
				),
			},
			args: args{
				ctx:      context.Background(),
				provider: &domain.IDPProvider{},
			},
			res: res{
				err: caos_errs.IsErrorInvalidArgument,
			},
		},
		{
			name: "login policy not existing, not found error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(),
				),
			},
			args: args{
				ctx: context.Background(),
				provider: &domain.IDPProvider{
					IDPConfigID: "config1",
				},
			},
			res: res{
				err: caos_errs.IsNotFound,
			},
		},
		{
			name: "provider not existing, not found error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							iam.NewLoginPolicyAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								true,
								true,
								true,
								true,
								true,
								true,
								domain.PasswordlessTypeAllowed,
								"",
							),
						),
					),
					expectFilter(),
				),
			},
			args: args{
				ctx: context.Background(),
				provider: &domain.IDPProvider{
					IDPConfigID: "config1",
				},
			},
			res: res{
				err: caos_errs.IsNotFound,
			},
		},
		{
			name: "provider removed, not found error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							iam.NewLoginPolicyAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								true,
								true,
								true,
								true,
								true,
								true,
								domain.PasswordlessTypeAllowed,
								"",
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							iam.NewIdentityProviderAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								"config1",
							),
						),
						eventFromEventPusher(
							iam.NewIdentityProviderRemovedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								"config1",
							),
						),
					),
				),
			},
			args: args{
				ctx: context.Background(),
				provider: &domain.IDPProvider{
					IDPConfigID: "config1",
				},
			},
			res: res{
				err: caos_errs.IsNotFound,
			},
		},
		{
			name: "remove provider, ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							iam.NewLoginPolicyAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								true,
								true,
								true,
								true,
								true,
								true,
								domain.PasswordlessTypeAllowed,
								"",
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							iam.NewIdentityProviderAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								"config1",
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								iam.NewIdentityProviderRemovedEvent(context.Background(),
									&iam.NewAggregate().Aggregate,
									"config1"),
							),
						},
					),
				),
			},
			args: args{
				ctx: context.Background(),
				provider: &domain.IDPProvider{
					IDPConfigID: "config1",
				},
			},
			res: res{
				want: &domain.ObjectDetails{
					ResourceOwner: "IAM",
				},
			},
		},
		{
			name: "remove provider external idp not found, ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							iam.NewLoginPolicyAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								true,
								true,
								true,
								true,
								true,
								true,
								domain.PasswordlessTypeAllowed,
								"",
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							iam.NewIdentityProviderAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								"config1",
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								iam.NewIdentityProviderRemovedEvent(context.Background(),
									&iam.NewAggregate().Aggregate,
									"config1"),
							),
						},
					),
				),
			},
			args: args{
				ctx: context.Background(),
				provider: &domain.IDPProvider{
					IDPConfigID: "config1",
				},
				cascadeExternalIDPs: []*domain.UserIDPLink{
					{
						ObjectRoot: models.ObjectRoot{
							AggregateID: "user1",
						},
						IDPConfigID: "config1",
					},
				},
			},
			res: res{
				want: &domain.ObjectDetails{
					ResourceOwner: "IAM",
				},
			},
		},
		{
			name: "remove provider with external idps, ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							iam.NewLoginPolicyAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								true,
								true,
								true,
								true,
								true,
								true,
								domain.PasswordlessTypeAllowed,
								"",
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							iam.NewIdentityProviderAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								"config1",
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							user.NewUserIDPLinkAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								"config1", "", "externaluser1"),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								iam.NewIdentityProviderRemovedEvent(context.Background(),
									&iam.NewAggregate().Aggregate,
									"config1"),
							),
							eventFromEventPusher(
								user.NewUserIDPLinkCascadeRemovedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									"config1", "externaluser1")),
						},
						uniqueConstraintsFromEventConstraint(user.NewRemoveUserIDPLinkUniqueConstraint("config1", "externaluser1")),
					),
				),
			},
			args: args{
				ctx: context.Background(),
				provider: &domain.IDPProvider{
					IDPConfigID: "config1",
				},
				cascadeExternalIDPs: []*domain.UserIDPLink{
					{
						ObjectRoot: models.ObjectRoot{
							AggregateID: "user1",
						},
						IDPConfigID:    "config1",
						ExternalUserID: "externaluser1",
					},
				},
			},
			res: res{
				want: &domain.ObjectDetails{
					ResourceOwner: "IAM",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Commands{
				eventstore: tt.fields.eventstore,
			}
			got, err := r.RemoveIDPProviderFromDefaultLoginPolicy(tt.args.ctx, tt.args.provider, tt.args.cascadeExternalIDPs...)
			if tt.res.err == nil {
				assert.NoError(t, err)
			}
			if tt.res.err != nil && !tt.res.err(err) {
				t.Errorf("got wrong err: %v ", err)
			}
			if tt.res.err == nil {
				assert.Equal(t, tt.res.want, got)
			}
		})
	}
}

func TestCommandSide_AddSecondFactorDefaultLoginPolicy(t *testing.T) {
	type fields struct {
		eventstore *eventstore.Eventstore
	}
	type args struct {
		ctx    context.Context
		factor domain.SecondFactorType
	}
	type res struct {
		want *domain.ObjectDetails
		err  func(error) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		res    res
	}{
		{
			name: "factor invalid, invalid argument error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
				),
			},
			args: args{
				ctx:    context.Background(),
				factor: domain.SecondFactorTypeUnspecified,
			},
			res: res{
				err: caos_errs.IsErrorInvalidArgument,
			},
		},
		{
			name: "factor already exists, already exists error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							iam.NewLoginPolicySecondFactorAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								domain.SecondFactorTypeOTP,
							),
						),
					),
				),
			},
			args: args{
				ctx:    context.Background(),
				factor: domain.SecondFactorTypeOTP,
			},
			res: res{
				err: caos_errs.IsErrorAlreadyExists,
			},
		},
		{
			name: "add factor, ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								iam.NewLoginPolicySecondFactorAddedEvent(context.Background(),
									&iam.NewAggregate().Aggregate,
									domain.SecondFactorTypeOTP),
							),
						},
					),
				),
			},
			args: args{
				ctx:    context.Background(),
				factor: domain.SecondFactorTypeOTP,
			},
			res: res{
				want: &domain.ObjectDetails{
					ResourceOwner: "IAM",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Commands{
				eventstore: tt.fields.eventstore,
			}
			_, got, err := r.AddSecondFactorToDefaultLoginPolicy(tt.args.ctx, tt.args.factor)
			if tt.res.err == nil {
				assert.NoError(t, err)
			}
			if tt.res.err != nil && !tt.res.err(err) {
				t.Errorf("got wrong err: %v ", err)
			}
			if tt.res.err == nil {
				assert.Equal(t, tt.res.want, got)
			}
		})
	}
}

func TestCommandSide_RemoveSecondFactorDefaultLoginPolicy(t *testing.T) {
	type fields struct {
		eventstore *eventstore.Eventstore
	}
	type args struct {
		ctx    context.Context
		factor domain.SecondFactorType
	}
	type res struct {
		want *domain.ObjectDetails
		err  func(error) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		res    res
	}{
		{
			name: "factor invalid, invalid argument error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
				),
			},
			args: args{
				ctx:    context.Background(),
				factor: domain.SecondFactorTypeUnspecified,
			},
			res: res{
				err: caos_errs.IsErrorInvalidArgument,
			},
		},
		{
			name: "factor not existing, not found error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(),
				),
			},
			args: args{
				ctx:    context.Background(),
				factor: domain.SecondFactorTypeOTP,
			},
			res: res{
				err: caos_errs.IsNotFound,
			},
		},
		{
			name: "factor removed, not found error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							iam.NewLoginPolicySecondFactorAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								domain.SecondFactorTypeOTP,
							),
						),
						eventFromEventPusher(
							iam.NewLoginPolicySecondFactorRemovedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								domain.SecondFactorTypeOTP,
							),
						),
					),
				),
			},
			args: args{
				ctx:    context.Background(),
				factor: domain.SecondFactorTypeOTP,
			},
			res: res{
				err: caos_errs.IsNotFound,
			},
		},
		{
			name: "add factor, ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							iam.NewLoginPolicySecondFactorAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								domain.SecondFactorTypeOTP,
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								iam.NewLoginPolicySecondFactorRemovedEvent(context.Background(),
									&iam.NewAggregate().Aggregate,
									domain.SecondFactorTypeOTP),
							),
						},
					),
				),
			},
			args: args{
				ctx:    context.Background(),
				factor: domain.SecondFactorTypeOTP,
			},
			res: res{
				want: &domain.ObjectDetails{
					ResourceOwner: "IAM",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Commands{
				eventstore: tt.fields.eventstore,
			}
			got, err := r.RemoveSecondFactorFromDefaultLoginPolicy(tt.args.ctx, tt.args.factor)
			if tt.res.err == nil {
				assert.NoError(t, err)
			}
			if tt.res.err != nil && !tt.res.err(err) {
				t.Errorf("got wrong err: %v ", err)
			}
			if tt.res.err == nil {
				assert.Equal(t, tt.res.want, got)
			}
		})
	}
}

func TestCommandSide_AddMultiFactorDefaultLoginPolicy(t *testing.T) {
	type fields struct {
		eventstore *eventstore.Eventstore
	}
	type args struct {
		ctx    context.Context
		factor domain.MultiFactorType
	}
	type res struct {
		want *domain.ObjectDetails
		err  func(error) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		res    res
	}{
		{
			name: "factor invalid, invalid argument error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
				),
			},
			args: args{
				ctx:    context.Background(),
				factor: domain.MultiFactorTypeUnspecified,
			},
			res: res{
				err: caos_errs.IsErrorInvalidArgument,
			},
		},
		{
			name: "factor already exists, already exists error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							iam.NewLoginPolicyMultiFactorAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								domain.MultiFactorTypeU2FWithPIN,
							),
						),
					),
				),
			},
			args: args{
				ctx:    context.Background(),
				factor: domain.MultiFactorTypeU2FWithPIN,
			},
			res: res{
				err: caos_errs.IsErrorAlreadyExists,
			},
		},
		{
			name: "add factor, ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								iam.NewLoginPolicyMultiFactorAddedEvent(context.Background(),
									&iam.NewAggregate().Aggregate,
									domain.MultiFactorTypeU2FWithPIN),
							),
						},
					),
				),
			},
			args: args{
				ctx:    context.Background(),
				factor: domain.MultiFactorTypeU2FWithPIN,
			},
			res: res{
				want: &domain.ObjectDetails{
					ResourceOwner: "IAM",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Commands{
				eventstore: tt.fields.eventstore,
			}
			_, got, err := r.AddMultiFactorToDefaultLoginPolicy(tt.args.ctx, tt.args.factor)
			if tt.res.err == nil {
				assert.NoError(t, err)
			}
			if tt.res.err != nil && !tt.res.err(err) {
				t.Errorf("got wrong err: %v ", err)
			}
			if tt.res.err == nil {
				assert.Equal(t, tt.res.want, got)
			}
		})
	}
}

func TestCommandSide_RemoveMultiFactorDefaultLoginPolicy(t *testing.T) {
	type fields struct {
		eventstore *eventstore.Eventstore
	}
	type args struct {
		ctx    context.Context
		factor domain.MultiFactorType
	}
	type res struct {
		want *domain.ObjectDetails
		err  func(error) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		res    res
	}{
		{
			name: "factor invalid, invalid argument error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
				),
			},
			args: args{
				ctx:    context.Background(),
				factor: domain.MultiFactorTypeUnspecified,
			},
			res: res{
				err: caos_errs.IsErrorInvalidArgument,
			},
		},
		{
			name: "factor not existing, not found error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(),
				),
			},
			args: args{
				ctx:    context.Background(),
				factor: domain.MultiFactorTypeU2FWithPIN,
			},
			res: res{
				err: caos_errs.IsNotFound,
			},
		},
		{
			name: "factor removed, not found error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							iam.NewLoginPolicyMultiFactorAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								domain.MultiFactorTypeU2FWithPIN,
							),
						),
						eventFromEventPusher(
							iam.NewLoginPolicyMultiFactorRemovedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								domain.MultiFactorTypeU2FWithPIN,
							),
						),
					),
				),
			},
			args: args{
				ctx:    context.Background(),
				factor: domain.MultiFactorTypeU2FWithPIN,
			},
			res: res{
				err: caos_errs.IsNotFound,
			},
		},
		{
			name: "add factor, ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							iam.NewLoginPolicyMultiFactorAddedEvent(context.Background(),
								&iam.NewAggregate().Aggregate,
								domain.MultiFactorTypeU2FWithPIN,
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								iam.NewLoginPolicyMultiFactorRemovedEvent(context.Background(),
									&iam.NewAggregate().Aggregate,
									domain.MultiFactorTypeU2FWithPIN),
							),
						},
					),
				),
			},
			args: args{
				ctx:    context.Background(),
				factor: domain.MultiFactorTypeU2FWithPIN,
			},
			res: res{
				want: &domain.ObjectDetails{
					ResourceOwner: "IAM",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Commands{
				eventstore: tt.fields.eventstore,
			}
			got, err := r.RemoveMultiFactorFromDefaultLoginPolicy(tt.args.ctx, tt.args.factor)
			if tt.res.err == nil {
				assert.NoError(t, err)
			}
			if tt.res.err != nil && !tt.res.err(err) {
				t.Errorf("got wrong err: %v ", err)
			}
			if tt.res.err == nil {
				assert.Equal(t, tt.res.want, got)
			}
		})
	}
}

func newDefaultLoginPolicyChangedEvent(ctx context.Context, allowRegister, allowUsernamePassword, allowExternalIDP, forceMFA, hidePasswordReset, ignoreUnknownUsernames bool, passwordlessType domain.PasswordlessType, redirectURI string) *iam.LoginPolicyChangedEvent {
	event, _ := iam.NewLoginPolicyChangedEvent(ctx,
		&iam.NewAggregate().Aggregate,
		[]policy.LoginPolicyChanges{
			policy.ChangeAllowRegister(allowRegister),
			policy.ChangeAllowExternalIDP(allowExternalIDP),
			policy.ChangeForceMFA(forceMFA),
			policy.ChangeAllowUserNamePassword(allowUsernamePassword),
			policy.ChangeHidePasswordReset(hidePasswordReset),
			policy.ChangeIgnoreUnknownUsernames(ignoreUnknownUsernames),
			policy.ChangePasswordlessType(passwordlessType),
			policy.ChangeDefaultRedirectURI(redirectURI),
		},
	)
	return event
}

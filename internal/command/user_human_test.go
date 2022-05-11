package command

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"

	"github.com/zitadel/zitadel/internal/crypto"
	"github.com/zitadel/zitadel/internal/domain"
	caos_errs "github.com/zitadel/zitadel/internal/errors"
	"github.com/zitadel/zitadel/internal/eventstore"
	"github.com/zitadel/zitadel/internal/eventstore/repository"
	"github.com/zitadel/zitadel/internal/eventstore/v1/models"
	"github.com/zitadel/zitadel/internal/id"
	id_mock "github.com/zitadel/zitadel/internal/id/mock"
	"github.com/zitadel/zitadel/internal/repository/org"
	"github.com/zitadel/zitadel/internal/repository/user"
)

func TestCommandSide_AddHuman(t *testing.T) {
	type fields struct {
		eventstore      *eventstore.Eventstore
		idGenerator     id.Generator
		secretGenerator crypto.Generator
		userPasswordAlg crypto.HashAlgorithm
	}
	type args struct {
		ctx   context.Context
		orgID string
		human *domain.Human
	}
	type res struct {
		want *domain.Human
		err  func(error) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		res    res
	}{
		{
			name: "orgid missing, invalid argument error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
				),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
				},
			},
			res: res{
				err: caos_errs.IsErrorInvalidArgument,
			},
		},
		{
			name: "org policy not found, precondition error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(),
					expectFilter(),
				),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
				},
			},
			res: res{
				err: caos_errs.IsPreconditionFailed,
			},
		},
		{
			name: "password policy not found, precondition error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(),
					expectFilter(),
				),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
				},
			},
			res: res{
				err: caos_errs.IsPreconditionFailed,
			},
		},
		{
			name: "user invalid, invalid argument error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
				),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
					},
				},
			},
			res: res{
				err: caos_errs.IsErrorInvalidArgument,
			},
		},
		{
			name: "add human (with initial code), ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								user.NewHumanAddedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									"username",
									"firstname",
									"lastname",
									"",
									"firstname lastname",
									language.Und,
									domain.GenderUnspecified,
									"email@test.ch",
									true,
								),
							),
							eventFromEventPusher(
								user.NewHumanInitialCodeAddedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									&crypto.CryptoValue{
										CryptoType: crypto.TypeEncryption,
										Algorithm:  "enc",
										KeyID:      "id",
										Crypted:    []byte("a"),
									},
									time.Hour*1,
								),
							),
						},
						uniqueConstraintsFromEventConstraint(user.NewAddUsernameUniqueConstraint("username", "org1", true)),
					),
				),
				idGenerator:     id_mock.NewIDGeneratorExpectIDs(t, "user1"),
				secretGenerator: GetMockSecretGenerator(t),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
				},
			},
			res: res{
				want: &domain.Human{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "user1",
						ResourceOwner: "org1",
					},
					Username: "username",
					Profile: &domain.Profile{
						FirstName:         "firstname",
						LastName:          "lastname",
						DisplayName:       "firstname lastname",
						PreferredLanguage: language.Und,
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					State: domain.UserStateInitial,
				},
			},
		},
		{
			name: "add human (with password and initial code), ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								newAddHumanEvent("password", true, ""),
							),
							eventFromEventPusher(
								user.NewHumanInitialCodeAddedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									&crypto.CryptoValue{
										CryptoType: crypto.TypeEncryption,
										Algorithm:  "enc",
										KeyID:      "id",
										Crypted:    []byte("a"),
									},
									time.Hour*1,
								),
							),
						},
						uniqueConstraintsFromEventConstraint(user.NewAddUsernameUniqueConstraint("username", "org1", true)),
					),
				),
				idGenerator:     id_mock.NewIDGeneratorExpectIDs(t, "user1"),
				secretGenerator: GetMockSecretGenerator(t),
				userPasswordAlg: crypto.CreateMockHashAlg(gomock.NewController(t)),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Password: &domain.Password{
						SecretString: "password",
					},
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
				},
			},
			res: res{
				want: &domain.Human{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "user1",
						ResourceOwner: "org1",
					},
					Username: "username",
					Profile: &domain.Profile{
						FirstName:         "firstname",
						LastName:          "lastname",
						DisplayName:       "firstname lastname",
						PreferredLanguage: language.Und,
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					State: domain.UserStateInitial,
				},
			},
		},
		{
			name: "add human email verified, ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								newAddHumanEvent("password", true, ""),
							),
							eventFromEventPusher(
								user.NewHumanEmailVerifiedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate),
							),
						},
						uniqueConstraintsFromEventConstraint(user.NewAddUsernameUniqueConstraint("username", "org1", true)),
					),
				),
				idGenerator:     id_mock.NewIDGeneratorExpectIDs(t, "user1"),
				secretGenerator: GetMockSecretGenerator(t),
				userPasswordAlg: crypto.CreateMockHashAlg(gomock.NewController(t)),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Password: &domain.Password{
						SecretString: "password",
					},
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress:    "email@test.ch",
						IsEmailVerified: true,
					},
				},
			},
			res: res{
				want: &domain.Human{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "user1",
						ResourceOwner: "org1",
					},
					Username: "username",
					Profile: &domain.Profile{
						FirstName:         "firstname",
						LastName:          "lastname",
						DisplayName:       "firstname lastname",
						PreferredLanguage: language.Und,
					},
					Email: &domain.Email{
						EmailAddress:    "email@test.ch",
						IsEmailVerified: true,
					},
					State: domain.UserStateActive,
				},
			},
		},
		{
			name: "add human (with phone), ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								newAddHumanEvent("", false, "+41711234567"),
							),
							eventFromEventPusher(
								user.NewHumanInitialCodeAddedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									&crypto.CryptoValue{
										CryptoType: crypto.TypeEncryption,
										Algorithm:  "enc",
										KeyID:      "id",
										Crypted:    []byte("a"),
									},
									time.Hour*1,
								),
							),
							eventFromEventPusher(
								user.NewHumanPhoneCodeAddedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									&crypto.CryptoValue{
										CryptoType: crypto.TypeEncryption,
										Algorithm:  "enc",
										KeyID:      "id",
										Crypted:    []byte("a"),
									},
									time.Hour*1)),
						},
						uniqueConstraintsFromEventConstraint(user.NewAddUsernameUniqueConstraint("username", "org1", true)),
					),
				),
				idGenerator:     id_mock.NewIDGeneratorExpectIDs(t, "user1"),
				secretGenerator: GetMockSecretGenerator(t),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					Phone: &domain.Phone{
						PhoneNumber: "+41711234567",
					},
				},
			},
			res: res{
				want: &domain.Human{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "user1",
						ResourceOwner: "org1",
					},
					Username: "username",
					Profile: &domain.Profile{
						FirstName:         "firstname",
						LastName:          "lastname",
						DisplayName:       "firstname lastname",
						PreferredLanguage: language.Und,
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					Phone: &domain.Phone{
						PhoneNumber: "+41711234567",
					},
					State: domain.UserStateInitial,
				},
			},
		},
		{
			name: "add human (with verified phone), ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								newAddHumanEvent("", false, "+41711234567"),
							),
							eventFromEventPusher(
								user.NewHumanInitialCodeAddedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									&crypto.CryptoValue{
										CryptoType: crypto.TypeEncryption,
										Algorithm:  "enc",
										KeyID:      "id",
										Crypted:    []byte("a"),
									},
									time.Hour*1,
								),
							),
							eventFromEventPusher(
								user.NewHumanPhoneVerifiedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate),
							),
						},
						uniqueConstraintsFromEventConstraint(user.NewAddUsernameUniqueConstraint("username", "org1", true)),
					),
				),
				idGenerator:     id_mock.NewIDGeneratorExpectIDs(t, "user1"),
				secretGenerator: GetMockSecretGenerator(t),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					Phone: &domain.Phone{
						PhoneNumber:     "+41711234567",
						IsPhoneVerified: true,
					},
				},
			},
			res: res{
				want: &domain.Human{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "user1",
						ResourceOwner: "org1",
					},
					Username: "username",
					Profile: &domain.Profile{
						FirstName:         "firstname",
						LastName:          "lastname",
						DisplayName:       "firstname lastname",
						PreferredLanguage: language.Und,
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					Phone: &domain.Phone{
						PhoneNumber: "+41711234567",
					},
					State: domain.UserStateInitial,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Commands{
				eventstore:            tt.fields.eventstore,
				idGenerator:           tt.fields.idGenerator,
				initializeUserCode:    tt.fields.secretGenerator,
				phoneVerificationCode: tt.fields.secretGenerator,
				userPasswordAlg:       tt.fields.userPasswordAlg,
			}
			got, err := r.AddHuman(tt.args.ctx, tt.args.orgID, tt.args.human)
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

func TestCommandSide_ImportHuman(t *testing.T) {
	type fields struct {
		eventstore           *eventstore.Eventstore
		idGenerator          id.Generator
		secretGenerator      crypto.Generator
		userPasswordAlg      crypto.HashAlgorithm
		passwordlessInitCode crypto.Generator
	}
	type args struct {
		ctx          context.Context
		orgID        string
		human        *domain.Human
		passwordless bool
	}
	type res struct {
		wantHuman *domain.Human
		wantCode  *domain.PasswordlessInitCode
		err       func(error) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		res    res
	}{
		{
			name: "orgid missing, invalid argument error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
				),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
				},
			},
			res: res{
				err: caos_errs.IsErrorInvalidArgument,
			},
		},
		{
			name: "org policy not found, precondition error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(),
					expectFilter(),
				),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
				},
			},
			res: res{
				err: caos_errs.IsPreconditionFailed,
			},
		},
		{
			name: "password policy not found, precondition error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(),
					expectFilter(),
				),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
				},
			},
			res: res{
				err: caos_errs.IsPreconditionFailed,
			},
		},
		{
			name: "user invalid, invalid argument error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
				),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
					},
				},
			},
			res: res{
				err: caos_errs.IsErrorInvalidArgument,
			},
		},
		{
			name: "add human (with password and initial code), ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								newAddHumanEvent("password", true, ""),
							),
							eventFromEventPusher(
								user.NewHumanInitialCodeAddedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									&crypto.CryptoValue{
										CryptoType: crypto.TypeEncryption,
										Algorithm:  "enc",
										KeyID:      "id",
										Crypted:    []byte("a"),
									},
									time.Hour*1,
								),
							),
						},
						uniqueConstraintsFromEventConstraint(user.NewAddUsernameUniqueConstraint("username", "org1", true)),
					),
				),
				idGenerator:     id_mock.NewIDGeneratorExpectIDs(t, "user1"),
				secretGenerator: GetMockSecretGenerator(t),
				userPasswordAlg: crypto.CreateMockHashAlg(gomock.NewController(t)),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Password: &domain.Password{
						SecretString:   "password",
						ChangeRequired: true,
					},
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
				},
			},
			res: res{
				wantHuman: &domain.Human{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "user1",
						ResourceOwner: "org1",
					},
					Username: "username",
					Profile: &domain.Profile{
						FirstName:         "firstname",
						LastName:          "lastname",
						DisplayName:       "firstname lastname",
						PreferredLanguage: language.Und,
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					State: domain.UserStateInitial,
				},
			},
		},
		{
			name: "add human email verified password change not required, ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								newAddHumanEvent("password", false, ""),
							),
							eventFromEventPusher(
								user.NewHumanEmailVerifiedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate),
							),
						},
						uniqueConstraintsFromEventConstraint(user.NewAddUsernameUniqueConstraint("username", "org1", true)),
					),
				),
				idGenerator:     id_mock.NewIDGeneratorExpectIDs(t, "user1"),
				secretGenerator: GetMockSecretGenerator(t),
				userPasswordAlg: crypto.CreateMockHashAlg(gomock.NewController(t)),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Password: &domain.Password{
						SecretString:   "password",
						ChangeRequired: false,
					},
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress:    "email@test.ch",
						IsEmailVerified: true,
					},
				},
			},
			res: res{
				wantHuman: &domain.Human{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "user1",
						ResourceOwner: "org1",
					},
					Username: "username",
					Profile: &domain.Profile{
						FirstName:         "firstname",
						LastName:          "lastname",
						DisplayName:       "firstname lastname",
						PreferredLanguage: language.Und,
					},
					Email: &domain.Email{
						EmailAddress:    "email@test.ch",
						IsEmailVerified: true,
					},
					State: domain.UserStateActive,
				},
			},
		},
		{
			name: "add human email verified passwordless only, ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectFilter(),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								newAddHumanEvent("", false, ""),
							),
							eventFromEventPusher(
								user.NewHumanEmailVerifiedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate),
							),
							eventFromEventPusher(
								user.NewHumanPasswordlessInitCodeAddedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									"code1",
									&crypto.CryptoValue{
										CryptoType: crypto.TypeEncryption,
										Algorithm:  "enc",
										KeyID:      "id",
										Crypted:    []byte("a"),
									},
									time.Hour,
								),
							),
						},
						uniqueConstraintsFromEventConstraint(user.NewAddUsernameUniqueConstraint("username", "org1", true)),
					),
				),
				idGenerator:          id_mock.NewIDGeneratorExpectIDs(t, "user1", "code1"),
				secretGenerator:      GetMockSecretGenerator(t),
				userPasswordAlg:      crypto.CreateMockHashAlg(gomock.NewController(t)),
				passwordlessInitCode: GetMockSecretGenerator(t),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress:    "email@test.ch",
						IsEmailVerified: true,
					},
				},
				passwordless: true,
			},
			res: res{
				wantHuman: &domain.Human{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "user1",
						ResourceOwner: "org1",
					},
					Username: "username",
					Profile: &domain.Profile{
						FirstName:         "firstname",
						LastName:          "lastname",
						DisplayName:       "firstname lastname",
						PreferredLanguage: language.Und,
					},
					Email: &domain.Email{
						EmailAddress:    "email@test.ch",
						IsEmailVerified: true,
					},
					State: domain.UserStateActive,
				},
				wantCode: &domain.PasswordlessInitCode{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "user1",
						ResourceOwner: "org1",
					},
					Expiration: time.Hour,
					CodeID:     "code1",
					Code:       "a",
					State:      domain.PasswordlessInitCodeStateActive,
				},
			},
		},
		{
			name: "add human email verified passwordless and password change not required, ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectFilter(),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								newAddHumanEvent("password", false, ""),
							),
							eventFromEventPusher(
								user.NewHumanEmailVerifiedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate),
							),
							eventFromEventPusher(
								user.NewHumanPasswordlessInitCodeAddedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									"code1",
									&crypto.CryptoValue{
										CryptoType: crypto.TypeEncryption,
										Algorithm:  "enc",
										KeyID:      "id",
										Crypted:    []byte("a"),
									},
									time.Hour,
								),
							),
						},
						uniqueConstraintsFromEventConstraint(user.NewAddUsernameUniqueConstraint("username", "org1", true)),
					),
				),
				idGenerator:          id_mock.NewIDGeneratorExpectIDs(t, "user1", "code1"),
				secretGenerator:      GetMockSecretGenerator(t),
				userPasswordAlg:      crypto.CreateMockHashAlg(gomock.NewController(t)),
				passwordlessInitCode: GetMockSecretGenerator(t),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Password: &domain.Password{
						SecretString:   "password",
						ChangeRequired: false,
					},
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress:    "email@test.ch",
						IsEmailVerified: true,
					},
				},
				passwordless: true,
			},
			res: res{
				wantHuman: &domain.Human{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "user1",
						ResourceOwner: "org1",
					},
					Username: "username",
					Profile: &domain.Profile{
						FirstName:         "firstname",
						LastName:          "lastname",
						DisplayName:       "firstname lastname",
						PreferredLanguage: language.Und,
					},
					Email: &domain.Email{
						EmailAddress:    "email@test.ch",
						IsEmailVerified: true,
					},
					State: domain.UserStateActive,
				},
				wantCode: &domain.PasswordlessInitCode{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "user1",
						ResourceOwner: "org1",
					},
					Expiration: time.Hour,
					CodeID:     "code1",
					Code:       "a",
					State:      domain.PasswordlessInitCodeStateActive,
				},
			},
		},
		{
			name: "add human (with phone), ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								newAddHumanEvent("password", false, "+41711234567"),
							),
							eventFromEventPusher(
								user.NewHumanInitialCodeAddedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									&crypto.CryptoValue{
										CryptoType: crypto.TypeEncryption,
										Algorithm:  "enc",
										KeyID:      "id",
										Crypted:    []byte("a"),
									},
									time.Hour*1,
								),
							),
							eventFromEventPusher(
								user.NewHumanPhoneCodeAddedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									&crypto.CryptoValue{
										CryptoType: crypto.TypeEncryption,
										Algorithm:  "enc",
										KeyID:      "id",
										Crypted:    []byte("a"),
									},
									time.Hour*1)),
						},
						uniqueConstraintsFromEventConstraint(user.NewAddUsernameUniqueConstraint("username", "org1", true)),
					),
				),
				idGenerator:     id_mock.NewIDGeneratorExpectIDs(t, "user1"),
				secretGenerator: GetMockSecretGenerator(t),
				userPasswordAlg: crypto.CreateMockHashAlg(gomock.NewController(t)),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Password: &domain.Password{
						SecretString:   "password",
						ChangeRequired: false,
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					Phone: &domain.Phone{
						PhoneNumber: "+41711234567",
					},
				},
			},
			res: res{
				wantHuman: &domain.Human{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "user1",
						ResourceOwner: "org1",
					},
					Username: "username",
					Profile: &domain.Profile{
						FirstName:         "firstname",
						LastName:          "lastname",
						DisplayName:       "firstname lastname",
						PreferredLanguage: language.Und,
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					Phone: &domain.Phone{
						PhoneNumber: "+41711234567",
					},
					State: domain.UserStateInitial,
				},
			},
		},
		{
			name: "add human (with verified phone), ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								newAddHumanEvent("password", false, "+41711234567"),
							),
							eventFromEventPusher(
								user.NewHumanInitialCodeAddedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									&crypto.CryptoValue{
										CryptoType: crypto.TypeEncryption,
										Algorithm:  "enc",
										KeyID:      "id",
										Crypted:    []byte("a"),
									},
									time.Hour*1,
								),
							),
							eventFromEventPusher(
								user.NewHumanPhoneVerifiedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate),
							),
						},
						uniqueConstraintsFromEventConstraint(user.NewAddUsernameUniqueConstraint("username", "org1", true)),
					),
				),
				idGenerator:     id_mock.NewIDGeneratorExpectIDs(t, "user1"),
				secretGenerator: GetMockSecretGenerator(t),
				userPasswordAlg: crypto.CreateMockHashAlg(gomock.NewController(t)),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Password: &domain.Password{
						SecretString:   "password",
						ChangeRequired: false,
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					Phone: &domain.Phone{
						PhoneNumber:     "+41711234567",
						IsPhoneVerified: true,
					},
				},
			},
			res: res{
				wantHuman: &domain.Human{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "user1",
						ResourceOwner: "org1",
					},
					Username: "username",
					Profile: &domain.Profile{
						FirstName:         "firstname",
						LastName:          "lastname",
						DisplayName:       "firstname lastname",
						PreferredLanguage: language.Und,
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					Phone: &domain.Phone{
						PhoneNumber: "+41711234567",
					},
					State: domain.UserStateInitial,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Commands{
				eventstore:            tt.fields.eventstore,
				idGenerator:           tt.fields.idGenerator,
				initializeUserCode:    tt.fields.secretGenerator,
				phoneVerificationCode: tt.fields.secretGenerator,
				userPasswordAlg:       tt.fields.userPasswordAlg,
				passwordlessInitCode:  tt.fields.passwordlessInitCode,
			}
			gotHuman, gotCode, err := r.ImportHuman(tt.args.ctx, tt.args.orgID, tt.args.human, tt.args.passwordless)
			if tt.res.err == nil {
				assert.NoError(t, err)
			}
			if tt.res.err != nil && !tt.res.err(err) {
				t.Errorf("got wrong err: %v ", err)
			}
			if tt.res.err == nil {
				assert.Equal(t, tt.res.wantHuman, gotHuman)
				assert.Equal(t, tt.res.wantCode, gotCode)
			}
		})
	}
}

func TestCommandSide_RegisterHuman(t *testing.T) {
	type fields struct {
		eventstore      *eventstore.Eventstore
		idGenerator     id.Generator
		secretGenerator crypto.Generator
		userPasswordAlg crypto.HashAlgorithm
	}
	type args struct {
		ctx            context.Context
		orgID          string
		human          *domain.Human
		link           *domain.UserIDPLink
		orgMemberRoles []string
	}
	type res struct {
		want *domain.Human
		err  func(error) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		res    res
	}{
		{
			name: "orgid missing, invalid argument error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
				),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
				},
			},
			res: res{
				err: caos_errs.IsErrorInvalidArgument,
			},
		},
		{
			name: "org policy not found, precondition error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(),
					expectFilter(),
				),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					Password: &domain.Password{
						SecretString: "password",
					},
				},
			},
			res: res{
				err: caos_errs.IsPreconditionFailed,
			},
		},
		{
			name: "password policy not found, precondition error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(),
					expectFilter(),
				),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					Password: &domain.Password{
						SecretString: "password",
					},
				},
			},
			res: res{
				err: caos_errs.IsPreconditionFailed,
			},
		},
		{
			name: "login policy not found, precondition error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectFilter(),
					expectFilter(),
				),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
					},
					Password: &domain.Password{
						SecretString: "password",
					},
				},
			},
			res: res{
				err: caos_errs.IsPreconditionFailed,
			},
		},
		{
			name: "login policy registration not allowed, precondition error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewLoginPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
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
					),
				),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
					},
					Password: &domain.Password{
						SecretString: "password",
					},
				},
			},
			res: res{
				err: caos_errs.IsPreconditionFailed,
			},
		},
		{
			name: "user invalid, invalid argument error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewLoginPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								false,
								true,
								false,
								false,
								false,
								false,
								domain.PasswordlessTypeNotAllowed,
								"",
							),
						),
					),
				),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
					},
					Password: &domain.Password{
						SecretString: "password",
					},
				},
			},
			res: res{
				err: caos_errs.IsErrorInvalidArgument,
			},
		},
		{
			name: "email domain reserved, invalid argument error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								false,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewLoginPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								false,
								true,
								false,
								false,
								false,
								false,
								domain.PasswordlessTypeNotAllowed,
								"",
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewDomainAddedEvent(context.Background(),
								&org.NewAggregate("org2", "org2").Aggregate,
								"test.ch",
							),
						),
						eventFromEventPusher(
							org.NewDomainVerifiedEvent(context.Background(),
								&org.NewAggregate("org2", "org2").Aggregate,
								"test.ch",
							),
						),
					),
				),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Password: &domain.Password{
						SecretString: "password",
					},
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
				},
			},
			res: res{
				err: caos_errs.IsErrorInvalidArgument,
			},
		},
		{
			name: "email domain reserved, same org, ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								false,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewLoginPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								false,
								true,
								false,
								false,
								false,
								false,
								domain.PasswordlessTypeNotAllowed,
								"",
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewDomainAddedEvent(context.Background(),
								&org.NewAggregate("org2", "org2").Aggregate,
								"test.ch",
							),
						),
						eventFromEventPusher(
							org.NewDomainVerifiedEvent(context.Background(),
								&org.NewAggregate("org2", "org2").Aggregate,
								"test.ch",
							),
						),
						eventFromEventPusher(
							org.NewDomainRemovedEvent(context.Background(),
								&org.NewAggregate("org2", "org2").Aggregate,
								"test.ch",
								true,
							),
						),
						eventFromEventPusher(
							org.NewDomainAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								"test.ch",
							),
						),
						eventFromEventPusher(
							org.NewDomainVerifiedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								"test.ch",
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								newRegisterHumanEvent("email@test.ch", "password", false, ""),
							),
							eventFromEventPusher(
								user.NewHumanInitialCodeAddedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									&crypto.CryptoValue{
										CryptoType: crypto.TypeEncryption,
										Algorithm:  "enc",
										KeyID:      "id",
										Crypted:    []byte("a"),
									},
									time.Hour*1,
								),
							),
						},
						uniqueConstraintsFromEventConstraint(user.NewAddUsernameUniqueConstraint("email@test.ch", "org1", false)),
					),
				),
				idGenerator:     id_mock.NewIDGeneratorExpectIDs(t, "user1"),
				secretGenerator: GetMockSecretGenerator(t),
				userPasswordAlg: crypto.CreateMockHashAlg(gomock.NewController(t)),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Password: &domain.Password{
						SecretString: "password",
					},
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
				},
			},
			res: res{
				want: &domain.Human{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "user1",
						ResourceOwner: "org1",
					},
					Username: "email@test.ch",
					Profile: &domain.Profile{
						FirstName:         "firstname",
						LastName:          "lastname",
						DisplayName:       "firstname lastname",
						PreferredLanguage: language.Und,
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					State: domain.UserStateInitial,
				},
			},
		},
		{
			name: "add human (with password and initial code), ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewLoginPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								false,
								true,
								false,
								false,
								false,
								false,
								domain.PasswordlessTypeNotAllowed,
								"",
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								newRegisterHumanEvent("username", "password", false, ""),
							),
							eventFromEventPusher(
								user.NewHumanInitialCodeAddedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									&crypto.CryptoValue{
										CryptoType: crypto.TypeEncryption,
										Algorithm:  "enc",
										KeyID:      "id",
										Crypted:    []byte("a"),
									},
									time.Hour*1,
								),
							),
						},
						uniqueConstraintsFromEventConstraint(user.NewAddUsernameUniqueConstraint("username", "org1", true)),
					),
				),
				idGenerator:     id_mock.NewIDGeneratorExpectIDs(t, "user1"),
				secretGenerator: GetMockSecretGenerator(t),
				userPasswordAlg: crypto.CreateMockHashAlg(gomock.NewController(t)),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Password: &domain.Password{
						SecretString: "password",
					},
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
				},
			},
			res: res{
				want: &domain.Human{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "user1",
						ResourceOwner: "org1",
					},
					Username: "username",
					Profile: &domain.Profile{
						FirstName:         "firstname",
						LastName:          "lastname",
						DisplayName:       "firstname lastname",
						PreferredLanguage: language.Und,
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					State: domain.UserStateInitial,
				},
			},
		},
		{
			name: "add human email verified, ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&user.NewAggregate("org1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewLoginPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								false,
								true,
								false,
								false,
								false,
								false,
								domain.PasswordlessTypeNotAllowed,
								"",
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								newRegisterHumanEvent("username", "password", false, ""),
							),
							eventFromEventPusher(
								user.NewHumanEmailVerifiedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate),
							),
						},
						uniqueConstraintsFromEventConstraint(user.NewAddUsernameUniqueConstraint("username", "org1", true)),
					),
				),
				idGenerator:     id_mock.NewIDGeneratorExpectIDs(t, "user1"),
				secretGenerator: GetMockSecretGenerator(t),
				userPasswordAlg: crypto.CreateMockHashAlg(gomock.NewController(t)),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Password: &domain.Password{
						SecretString: "password",
					},
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress:    "email@test.ch",
						IsEmailVerified: true,
					},
				},
			},
			res: res{
				want: &domain.Human{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "user1",
						ResourceOwner: "org1",
					},
					Username: "username",
					Profile: &domain.Profile{
						FirstName:         "firstname",
						LastName:          "lastname",
						DisplayName:       "firstname lastname",
						PreferredLanguage: language.Und,
					},
					Email: &domain.Email{
						EmailAddress:    "email@test.ch",
						IsEmailVerified: true,
					},
					State: domain.UserStateActive,
				},
			},
		},
		{
			name: "add human (with phone), ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewLoginPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								false,
								true,
								false,
								false,
								false,
								false,
								domain.PasswordlessTypeNotAllowed,
								"",
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								newRegisterHumanEvent("username", "password", false, "+41711234567"),
							),
							eventFromEventPusher(
								user.NewHumanInitialCodeAddedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									&crypto.CryptoValue{
										CryptoType: crypto.TypeEncryption,
										Algorithm:  "enc",
										KeyID:      "id",
										Crypted:    []byte("a"),
									},
									time.Hour*1,
								),
							),
							eventFromEventPusher(
								user.NewHumanPhoneCodeAddedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									&crypto.CryptoValue{
										CryptoType: crypto.TypeEncryption,
										Algorithm:  "enc",
										KeyID:      "id",
										Crypted:    []byte("a"),
									},
									time.Hour*1)),
						},
						uniqueConstraintsFromEventConstraint(user.NewAddUsernameUniqueConstraint("username", "org1", true)),
					),
				),
				idGenerator:     id_mock.NewIDGeneratorExpectIDs(t, "user1"),
				secretGenerator: GetMockSecretGenerator(t),
				userPasswordAlg: crypto.CreateMockHashAlg(gomock.NewController(t)),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					Phone: &domain.Phone{
						PhoneNumber: "+41711234567",
					},
					Password: &domain.Password{
						SecretString: "password",
					},
				},
			},
			res: res{
				want: &domain.Human{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "user1",
						ResourceOwner: "org1",
					},
					Username: "username",
					Profile: &domain.Profile{
						FirstName:         "firstname",
						LastName:          "lastname",
						DisplayName:       "firstname lastname",
						PreferredLanguage: language.Und,
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					Phone: &domain.Phone{
						PhoneNumber: "+41711234567",
					},
					State: domain.UserStateInitial,
				},
			},
		},
		{
			name: "add human (with verified phone), ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							org.NewOrgIAMPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewPasswordComplexityPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								1,
								false,
								false,
								false,
								false,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							org.NewLoginPolicyAddedEvent(context.Background(),
								&org.NewAggregate("org1", "org1").Aggregate,
								false,
								true,
								false,
								false,
								false,
								false,
								domain.PasswordlessTypeNotAllowed,
								"",
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								newRegisterHumanEvent("username", "password", false, "+41711234567"),
							),
							eventFromEventPusher(
								user.NewHumanInitialCodeAddedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									&crypto.CryptoValue{
										CryptoType: crypto.TypeEncryption,
										Algorithm:  "enc",
										KeyID:      "id",
										Crypted:    []byte("a"),
									},
									time.Hour*1,
								),
							),
							eventFromEventPusher(
								user.NewHumanPhoneVerifiedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate),
							),
						},
						uniqueConstraintsFromEventConstraint(user.NewAddUsernameUniqueConstraint("username", "org1", true)),
					),
				),
				idGenerator:     id_mock.NewIDGeneratorExpectIDs(t, "user1"),
				secretGenerator: GetMockSecretGenerator(t),
				userPasswordAlg: crypto.CreateMockHashAlg(gomock.NewController(t)),
			},
			args: args{
				ctx:   context.Background(),
				orgID: "org1",
				human: &domain.Human{
					Username: "username",
					Profile: &domain.Profile{
						FirstName: "firstname",
						LastName:  "lastname",
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					Phone: &domain.Phone{
						PhoneNumber:     "+41711234567",
						IsPhoneVerified: true,
					},
					Password: &domain.Password{
						SecretString: "password",
					},
				},
			},
			res: res{
				want: &domain.Human{
					ObjectRoot: models.ObjectRoot{
						AggregateID:   "user1",
						ResourceOwner: "org1",
					},
					Username: "username",
					Profile: &domain.Profile{
						FirstName:         "firstname",
						LastName:          "lastname",
						DisplayName:       "firstname lastname",
						PreferredLanguage: language.Und,
					},
					Email: &domain.Email{
						EmailAddress: "email@test.ch",
					},
					Phone: &domain.Phone{
						PhoneNumber: "+41711234567",
					},
					State: domain.UserStateInitial,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Commands{
				eventstore:            tt.fields.eventstore,
				idGenerator:           tt.fields.idGenerator,
				initializeUserCode:    tt.fields.secretGenerator,
				phoneVerificationCode: tt.fields.secretGenerator,
				userPasswordAlg:       tt.fields.userPasswordAlg,
			}
			got, err := r.RegisterHuman(tt.args.ctx, tt.args.orgID, tt.args.human, tt.args.link, tt.args.orgMemberRoles)
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

func TestCommandSide_HumanMFASkip(t *testing.T) {
	type fields struct {
		eventstore *eventstore.Eventstore
	}
	type (
		args struct {
			ctx    context.Context
			orgID  string
			userID string
		}
	)
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
			name: "userid missing, invalid argument error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
				),
			},
			args: args{
				ctx:    context.Background(),
				orgID:  "org1",
				userID: "",
			},
			res: res{
				err: caos_errs.IsErrorInvalidArgument,
			},
		},
		{
			name: "user not existing, not found error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(),
				),
			},
			args: args{
				ctx:    context.Background(),
				orgID:  "org1",
				userID: "user1",
			},
			res: res{
				err: caos_errs.IsNotFound,
			},
		},
		{
			name: "skip mfa init, ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							user.NewHumanAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								"username",
								"firstname",
								"lastname",
								"nickname",
								"displayname",
								language.German,
								domain.GenderUnspecified,
								"email@test.ch",
								true,
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								user.NewHumanMFAInitSkippedEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
								),
							),
						},
					),
				),
			},
			args: args{
				ctx:    context.Background(),
				orgID:  "org1",
				userID: "user1",
			},
			res: res{
				want: &domain.ObjectDetails{
					ResourceOwner: "org1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Commands{
				eventstore: tt.fields.eventstore,
			}
			err := r.HumanSkipMFAInit(tt.args.ctx, tt.args.userID, tt.args.orgID)
			if tt.res.err == nil {
				assert.NoError(t, err)
			}
			if tt.res.err != nil && !tt.res.err(err) {
				t.Errorf("got wrong err: %v ", err)
			}
		})
	}
}

func TestCommandSide_HumanSignOut(t *testing.T) {
	type fields struct {
		eventstore *eventstore.Eventstore
	}
	type (
		args struct {
			ctx     context.Context
			agentID string
			userIDs []string
		}
	)
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
			name: "agentid missing, invalid argument error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
				),
			},
			args: args{
				ctx:     context.Background(),
				agentID: "",
				userIDs: []string{"user1"},
			},
			res: res{
				err: caos_errs.IsErrorInvalidArgument,
			},
		},
		{
			name: "userids missing, invalid argument error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
				),
			},
			args: args{
				ctx:     context.Background(),
				agentID: "agent1",
				userIDs: []string{},
			},
			res: res{
				err: caos_errs.IsErrorInvalidArgument,
			},
		},
		{
			name: "user not existing, not found error",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(),
				),
			},
			args: args{
				ctx:     context.Background(),
				agentID: "agent1",
				userIDs: []string{"user1"},
			},
			res: res{},
		},
		{
			name: "human sign out, ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							user.NewHumanAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								"username",
								"firstname",
								"lastname",
								"nickname",
								"displayname",
								language.German,
								domain.GenderUnspecified,
								"email@test.ch",
								true,
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								user.NewHumanSignedOutEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									"agent1",
								),
							),
						},
					),
				),
			},
			args: args{
				ctx:     context.Background(),
				agentID: "agent1",
				userIDs: []string{"user1"},
			},
			res: res{
				want: &domain.ObjectDetails{
					ResourceOwner: "org1",
				},
			},
		},
		{
			name: "human sign out multiple users, ok",
			fields: fields{
				eventstore: eventstoreExpect(
					t,
					expectFilter(
						eventFromEventPusher(
							user.NewHumanAddedEvent(context.Background(),
								&user.NewAggregate("user1", "org1").Aggregate,
								"username",
								"firstname",
								"lastname",
								"nickname",
								"displayname",
								language.German,
								domain.GenderUnspecified,
								"email@test.ch",
								true,
							),
						),
					),
					expectFilter(
						eventFromEventPusher(
							user.NewHumanAddedEvent(context.Background(),
								&user.NewAggregate("user2", "org1").Aggregate,
								"username",
								"firstname",
								"lastname",
								"nickname",
								"displayname",
								language.German,
								domain.GenderUnspecified,
								"email@test.ch",
								true,
							),
						),
					),
					expectPush(
						[]*repository.Event{
							eventFromEventPusher(
								user.NewHumanSignedOutEvent(context.Background(),
									&user.NewAggregate("user1", "org1").Aggregate,
									"agent1",
								),
							),
							eventFromEventPusher(
								user.NewHumanSignedOutEvent(context.Background(),
									&user.NewAggregate("user2", "org1").Aggregate,
									"agent1",
								),
							),
						},
					),
				),
			},
			args: args{
				ctx:     context.Background(),
				agentID: "agent1",
				userIDs: []string{"user1", "user2"},
			},
			res: res{
				want: &domain.ObjectDetails{
					ResourceOwner: "org1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Commands{
				eventstore: tt.fields.eventstore,
			}
			err := r.HumansSignOut(tt.args.ctx, tt.args.agentID, tt.args.userIDs)
			if tt.res.err == nil {
				assert.NoError(t, err)
			}
			if tt.res.err != nil && !tt.res.err(err) {
				t.Errorf("got wrong err: %v ", err)
			}
		})
	}
}

func newAddHumanEvent(password string, changeRequired bool, phone string) *user.HumanAddedEvent {
	event := user.NewHumanAddedEvent(context.Background(),
		&user.NewAggregate("user1", "org1").Aggregate,
		"username",
		"firstname",
		"lastname",
		"",
		"firstname lastname",
		language.Und,
		domain.GenderUnspecified,
		"email@test.ch",
		true,
	)
	if password != "" {
		event.AddPasswordData(&crypto.CryptoValue{
			CryptoType: crypto.TypeHash,
			Algorithm:  "hash",
			KeyID:      "",
			Crypted:    []byte(password),
		},
			changeRequired)
	}
	if phone != "" {
		event.AddPhoneData(phone)
	}
	return event
}

func newRegisterHumanEvent(username, password string, changeRequired bool, phone string) *user.HumanRegisteredEvent {
	event := user.NewHumanRegisteredEvent(context.Background(),
		&user.NewAggregate("user1", "org1").Aggregate,
		username,
		"firstname",
		"lastname",
		"",
		"firstname lastname",
		language.Und,
		domain.GenderUnspecified,
		"email@test.ch",
		true,
	)
	if password != "" {
		event.AddPasswordData(&crypto.CryptoValue{
			CryptoType: crypto.TypeHash,
			Algorithm:  "hash",
			KeyID:      "",
			Crypted:    []byte(password),
		},
			changeRequired)
	}
	if phone != "" {
		event.AddPhoneData(phone)
	}
	return event
}

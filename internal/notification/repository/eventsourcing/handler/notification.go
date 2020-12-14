package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/caos/logging"

	"github.com/caos/zitadel/internal/api/authz"
	sd "github.com/caos/zitadel/internal/config/systemdefaults"
	"github.com/caos/zitadel/internal/crypto"
	"github.com/caos/zitadel/internal/errors"
	caos_errs "github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/eventstore/models"
	"github.com/caos/zitadel/internal/eventstore/spooler"
	"github.com/caos/zitadel/internal/i18n"
	iam_model "github.com/caos/zitadel/internal/iam/model"
	iam_es_model "github.com/caos/zitadel/internal/iam/repository/view/model"
	"github.com/caos/zitadel/internal/notification/types"
	"github.com/caos/zitadel/internal/user/repository/eventsourcing"
	usr_event "github.com/caos/zitadel/internal/user/repository/eventsourcing"
	es_model "github.com/caos/zitadel/internal/user/repository/eventsourcing/model"
)

type Notification struct {
	handler
	eventstore     eventstore.Eventstore
	userEvents     *usr_event.UserEventstore
	systemDefaults sd.SystemDefaults
	AesCrypto      crypto.EncryptionAlgorithm
	i18n           *i18n.Translator
	statikDir      http.FileSystem
}

const (
	notificationTable   = "notification.notifications"
	NotifyUserID        = "NOTIFICATION"
	labelPolicyTableOrg = "management.label_policies"
	labelPolicyTableDef = "adminapi.label_policies"
)

func (n *Notification) ViewModel() string {
	return notificationTable
}

func (n *Notification) EventQuery() (*models.SearchQuery, error) {
	sequence, err := n.view.GetLatestNotificationSequence()
	if err != nil {
		return nil, err
	}
	return eventsourcing.UserQuery(sequence.CurrentSequence), nil
}

func (n *Notification) Reduce(event *models.Event) (err error) {
	switch event.Type {
	case es_model.InitializedUserCodeAdded,
		es_model.InitializedHumanCodeAdded:
		err = n.handleInitUserCode(event)
	case es_model.UserEmailCodeAdded,
		es_model.HumanEmailCodeAdded:
		err = n.handleEmailVerificationCode(event)
	case es_model.UserPhoneCodeAdded,
		es_model.HumanPhoneCodeAdded:
		err = n.handlePhoneVerificationCode(event)
	case es_model.UserPasswordCodeAdded,
		es_model.HumanPasswordCodeAdded:
		err = n.handlePasswordCode(event)
	case es_model.DomainClaimed:
		err = n.handleDomainClaimed(event)
	default:
		return n.view.ProcessedNotificationSequence(event.Sequence, event.CreationDate)
	}
	if err != nil {
		return err
	}
	return n.view.ProcessedNotificationSequence(event.Sequence, event.CreationDate)
}

func (n *Notification) handleInitUserCode(event *models.Event) (err error) {
	initCode := new(es_model.InitUserCode)
	if err := initCode.SetData(event); err != nil {
		return err
	}
	alreadyHandled, err := n.checkIfCodeAlreadyHandledOrExpired(event, initCode.Expiry,
		es_model.InitializedUserCodeAdded, es_model.InitializedUserCodeSent,
		es_model.InitializedHumanCodeAdded, es_model.InitializedHumanCodeSent)
	if err != nil || alreadyHandled {
		return err
	}

	colors, err := n.getLabelPolicy(context.Background())
	if err != nil {
		return err
	}

	user, err := n.view.NotifyUserByID(event.AggregateID)
	if err != nil {
		return err
	}
	err = types.SendUserInitCode(n.statikDir, n.i18n, user, initCode, n.systemDefaults, n.AesCrypto, colors)
	if err != nil {
		return err
	}
	return n.userEvents.InitCodeSent(getSetNotifyContextData(event.ResourceOwner), event.AggregateID)
}

func (n *Notification) handlePasswordCode(event *models.Event) (err error) {
	pwCode := new(es_model.PasswordCode)
	if err := pwCode.SetData(event); err != nil {
		return err
	}
	alreadyHandled, err := n.checkIfCodeAlreadyHandledOrExpired(event, pwCode.Expiry,
		es_model.UserPasswordCodeAdded, es_model.UserPasswordCodeSent,
		es_model.HumanPasswordCodeAdded, es_model.HumanPasswordCodeSent)
	if err != nil || alreadyHandled {
		return err
	}

	colors, err := n.getLabelPolicy(context.Background())
	if err != nil {
		return err
	}

	user, err := n.view.NotifyUserByID(event.AggregateID)
	if err != nil {
		return err
	}
	err = types.SendPasswordCode(n.statikDir, n.i18n, user, pwCode, n.systemDefaults, n.AesCrypto, colors)
	if err != nil {
		return err
	}
	return n.userEvents.PasswordCodeSent(getSetNotifyContextData(event.ResourceOwner), event.AggregateID)
}

func (n *Notification) handleEmailVerificationCode(event *models.Event) (err error) {
	emailCode := new(es_model.EmailCode)
	if err := emailCode.SetData(event); err != nil {
		return err
	}
	alreadyHandled, err := n.checkIfCodeAlreadyHandledOrExpired(event, emailCode.Expiry,
		es_model.UserEmailCodeAdded, es_model.UserEmailCodeSent,
		es_model.HumanEmailCodeAdded, es_model.HumanEmailCodeSent)
	if err != nil || alreadyHandled {
		return nil
	}

	colors, err := n.getLabelPolicy(context.Background())
	if err != nil {
		return err
	}

	user, err := n.view.NotifyUserByID(event.AggregateID)
	if err != nil {
		return err
	}
	err = types.SendEmailVerificationCode(n.statikDir, n.i18n, user, emailCode, n.systemDefaults, n.AesCrypto, colors)
	if err != nil {
		return err
	}
	return n.userEvents.EmailVerificationCodeSent(getSetNotifyContextData(event.ResourceOwner), event.AggregateID)
}

func (n *Notification) handlePhoneVerificationCode(event *models.Event) (err error) {
	phoneCode := new(es_model.PhoneCode)
	if err := phoneCode.SetData(event); err != nil {
		return err
	}
	alreadyHandled, err := n.checkIfCodeAlreadyHandledOrExpired(event, phoneCode.Expiry,
		es_model.UserPhoneCodeAdded, es_model.UserPhoneCodeSent,
		es_model.HumanPhoneCodeAdded, es_model.HumanPhoneCodeSent)
	if err != nil || alreadyHandled {
		return nil
	}
	user, err := n.view.NotifyUserByID(event.AggregateID)
	if err != nil {
		return err
	}
	err = types.SendPhoneVerificationCode(n.i18n, user, phoneCode, n.systemDefaults, n.AesCrypto)
	if err != nil {
		return err
	}
	return n.userEvents.PhoneVerificationCodeSent(getSetNotifyContextData(event.ResourceOwner), event.AggregateID)
}

func (n *Notification) handleDomainClaimed(event *models.Event) (err error) {
	alreadyHandled, err := n.checkIfAlreadyHandled(event.AggregateID, event.Sequence, es_model.DomainClaimed, es_model.DomainClaimedSent)
	if err != nil || alreadyHandled {
		return nil
	}
	data := make(map[string]string)
	if err := json.Unmarshal(event.Data, &data); err != nil {
		logging.Log("HANDLE-Gghq2").WithError(err).Error("could not unmarshal event data")
		return caos_errs.ThrowInternal(err, "HANDLE-7hgj3", "could not unmarshal event")
	}
	user, err := n.view.NotifyUserByID(event.AggregateID)
	if err != nil {
		return err
	}
	err = types.SendDomainClaimed(n.statikDir, n.i18n, user, data["userName"], n.systemDefaults)
	if err != nil {
		return err
	}
	return n.userEvents.DomainClaimedSent(getSetNotifyContextData(event.ResourceOwner), event.AggregateID)
}

func (n *Notification) checkIfCodeAlreadyHandledOrExpired(event *models.Event, expiry time.Duration, eventTypes ...models.EventType) (bool, error) {
	if event.CreationDate.Add(expiry).Before(time.Now().UTC()) {
		return true, nil
	}
	return n.checkIfAlreadyHandled(event.AggregateID, event.Sequence, eventTypes...)
}

func (n *Notification) checkIfAlreadyHandled(userID string, sequence uint64, eventTypes ...models.EventType) (bool, error) {
	events, err := n.getUserEvents(userID, sequence)
	if err != nil {
		return false, err
	}
	for _, event := range events {
		for _, eventType := range eventTypes {
			if event.Type == eventType {
				return true, nil
			}
		}
	}
	return false, nil
}

func (n *Notification) getUserEvents(userID string, sequence uint64) ([]*models.Event, error) {
	query, err := eventsourcing.UserByIDQuery(userID, sequence)
	if err != nil {
		return nil, err
	}

	return n.eventstore.FilterEvents(context.Background(), query)
}

func (n *Notification) OnError(event *models.Event, err error) error {
	logging.LogWithFields("SPOOL-s9opc", "id", event.AggregateID, "sequence", event.Sequence).WithError(err).Warn("something went wrong in notification handler")
	return spooler.HandleError(event, err, n.view.GetLatestNotificationFailedEvent, n.view.ProcessedNotificationFailedEvent, n.view.ProcessedNotificationSequence, n.errorCountUntilSkip)
}

func (n *Notification) OnSuccess() error {
	return spooler.HandleSuccess(n.view.UpdateNotificationSpoolerRunTimestamp)
}

func getSetNotifyContextData(orgID string) context.Context {
	return authz.SetCtxData(context.Background(), authz.CtxData{UserID: NotifyUserID, OrgID: orgID})
}

// Read organization specific colors
func (n *Notification) getLabelPolicy(ctx context.Context) (*iam_model.LabelPolicyView, error) {
	// read from Org
	policy, err := n.view.LabelPolicyByAggregateID(authz.GetCtxData(ctx).OrgID, labelPolicyTableOrg)
	if errors.IsNotFound(err) {
		// read from default
		policy, err = n.view.LabelPolicyByAggregateID(n.systemDefaults.IamID, labelPolicyTableDef)
		if err != nil {
			return nil, err
		}
		policy.Default = true
	}
	if err != nil {
		return nil, err
	}
	return iam_es_model.LabelPolicyViewToModel(policy), err
}

package view

import (
	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/iam/repository/view"
	"github.com/caos/zitadel/internal/iam/repository/view/model"
	global_view "github.com/caos/zitadel/internal/view/repository"
	"time"
)

const (
	loginPolicyTable = "auth.login_policies"
)

func (v *View) LoginPolicyByAggregateID(aggregateID string) (*model.LoginPolicyView, error) {
	return view.GetLoginPolicyByAggregateID(v.Db, loginPolicyTable, aggregateID)
}

func (v *View) PutLoginPolicy(policy *model.LoginPolicyView, sequence uint64, eventTimestamp time.Time) error {
	err := view.PutLoginPolicy(v.Db, loginPolicyTable, policy)
	if err != nil {
		return err
	}
	return v.ProcessedLoginPolicySequence(sequence, eventTimestamp)
}

func (v *View) DeleteLoginPolicy(aggregateID string, eventSequence uint64, eventTimestamp time.Time) error {
	err := view.DeleteLoginPolicy(v.Db, loginPolicyTable, aggregateID)
	if err != nil && !errors.IsNotFound(err) {
		return err
	}
	return v.ProcessedLoginPolicySequence(eventSequence, eventTimestamp)
}

func (v *View) GetLatestLoginPolicySequence() (*global_view.CurrentSequence, error) {
	return v.latestSequence(loginPolicyTable)
}

func (v *View) ProcessedLoginPolicySequence(eventSequence uint64, eventTimestamp time.Time) error {
	return v.saveCurrentSequence(loginPolicyTable, eventSequence, eventTimestamp)
}

func (v *View) UpdateLoginPolicySpoolerRunTimestamp() error {
	return v.updateSpoolerRunSequence(loginPolicyTable)
}

func (v *View) GetLatestLoginPolicyFailedEvent(sequence uint64) (*global_view.FailedEvent, error) {
	return v.latestFailedEvent(loginPolicyTable, sequence)
}

func (v *View) ProcessedLoginPolicyFailedEvent(failedEvent *global_view.FailedEvent) error {
	return v.saveFailedEvent(failedEvent)
}

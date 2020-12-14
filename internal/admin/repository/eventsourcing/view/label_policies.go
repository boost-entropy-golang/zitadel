package view

import (
	"github.com/caos/zitadel/internal/iam/repository/view"
	"github.com/caos/zitadel/internal/iam/repository/view/model"
	global_view "github.com/caos/zitadel/internal/view/repository"
	"time"
)

const (
	labelPolicyTable = "adminapi.label_policies"
)

func (v *View) LabelPolicyByAggregateID(aggregateID string) (*model.LabelPolicyView, error) {
	return view.GetLabelPolicyByAggregateID(v.Db, labelPolicyTable, aggregateID)
}

func (v *View) PutLabelPolicy(policy *model.LabelPolicyView, sequence uint64, eventTimestamp time.Time) error {
	err := view.PutLabelPolicy(v.Db, labelPolicyTable, policy)
	if err != nil {
		return err
	}
	return v.ProcessedLabelPolicySequence(sequence, eventTimestamp)
}

func (v *View) GetLatestLabelPolicySequence() (*global_view.CurrentSequence, error) {
	return v.latestSequence(labelPolicyTable)
}

func (v *View) ProcessedLabelPolicySequence(eventSequence uint64, eventTimestamp time.Time) error {
	return v.saveCurrentSequence(labelPolicyTable, eventSequence, eventTimestamp)
}

func (v *View) UpdateLabelPolicySpoolerRunTimestamp() error {
	return v.updateSpoolerRunSequence(labelPolicyTable)
}

func (v *View) GetLatestLabelPolicyFailedEvent(sequence uint64) (*global_view.FailedEvent, error) {
	return v.latestFailedEvent(labelPolicyTable, sequence)
}

func (v *View) ProcessedLabelPolicyFailedEvent(failedEvent *global_view.FailedEvent) error {
	return v.saveFailedEvent(failedEvent)
}

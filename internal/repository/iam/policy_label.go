package iam

import (
	"context"
	"github.com/caos/zitadel/internal/eventstore"

	"github.com/caos/zitadel/internal/eventstore/repository"
	"github.com/caos/zitadel/internal/repository/policy"
)

var (
	LabelPolicyAddedEventType     = iamEventTypePrefix + policy.LabelPolicyAddedEventType
	LabelPolicyChangedEventType   = iamEventTypePrefix + policy.LabelPolicyChangedEventType
	LabelPolicyActivatedEventType = iamEventTypePrefix + policy.LabelPolicyActivatedEventType
)

type LabelPolicyAddedEvent struct {
	policy.LabelPolicyAddedEvent
}

func NewLabelPolicyAddedEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	primaryColor,
	secondaryColor,
	warnColor,
	primaryColorDark,
	secondaryColorDark,
	warnColorDark string,
	hideLoginNameSuffix,
	errorMsgPopup,
	disableWatermark bool,
) *LabelPolicyAddedEvent {
	return &LabelPolicyAddedEvent{
		LabelPolicyAddedEvent: *policy.NewLabelPolicyAddedEvent(
			eventstore.NewBaseEventForPush(
				ctx,
				aggregate,
				LabelPolicyAddedEventType),
			primaryColor,
			secondaryColor,
			warnColor,
			primaryColorDark,
			secondaryColorDark,
			warnColorDark,
			hideLoginNameSuffix,
			errorMsgPopup,
			disableWatermark),
	}
}

func LabelPolicyAddedEventMapper(event *repository.Event) (eventstore.EventReader, error) {
	e, err := policy.LabelPolicyAddedEventMapper(event)
	if err != nil {
		return nil, err
	}

	return &LabelPolicyAddedEvent{LabelPolicyAddedEvent: *e.(*policy.LabelPolicyAddedEvent)}, nil
}

type LabelPolicyChangedEvent struct {
	policy.LabelPolicyChangedEvent
}

func NewLabelPolicyChangedEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	changes []policy.LabelPolicyChanges,
) (*LabelPolicyChangedEvent, error) {
	changedEvent, err := policy.NewLabelPolicyChangedEvent(
		eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			LabelPolicyChangedEventType),
		changes,
	)
	if err != nil {
		return nil, err
	}
	return &LabelPolicyChangedEvent{LabelPolicyChangedEvent: *changedEvent}, nil
}

func LabelPolicyChangedEventMapper(event *repository.Event) (eventstore.EventReader, error) {
	e, err := policy.LabelPolicyChangedEventMapper(event)
	if err != nil {
		return nil, err
	}

	return &LabelPolicyChangedEvent{LabelPolicyChangedEvent: *e.(*policy.LabelPolicyChangedEvent)}, nil
}

type LabelPolicyActivatedEvent struct {
	policy.LabelPolicyActivatedEvent
}

func NewLabelPolicyActivatedEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
) *LabelPolicyActivatedEvent {
	return &LabelPolicyActivatedEvent{
		LabelPolicyActivatedEvent: *policy.NewLabelPolicyActivatedEvent(
			eventstore.NewBaseEventForPush(
				ctx,
				aggregate,
				LabelPolicyActivatedEventType),
		),
	}
}

func LabelPolicyActivatedEventMapper(event *repository.Event) (eventstore.EventReader, error) {
	e, err := policy.LabelPolicyActivatedEventMapper(event)
	if err != nil {
		return nil, err
	}

	return &LabelPolicyActivatedEvent{LabelPolicyActivatedEvent: *e.(*policy.LabelPolicyActivatedEvent)}, nil
}

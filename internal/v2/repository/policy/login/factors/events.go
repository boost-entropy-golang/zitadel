package factors

import (
	"encoding/json"
	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore/v2"
	"github.com/caos/zitadel/internal/eventstore/v2/repository"
)

const (
	loginPolicySecondFactorPrefix           = "policy.login.secondfactor."
	LoginPolicySecondFactorAddedEventType   = loginPolicySecondFactorPrefix + "added"
	LoginPolicySecondFactorRemovedEventType = loginPolicySecondFactorPrefix + "removed"

	loginPolicyMultiFactorPrefix           = "policy.login.multifactor."
	LoginPolicyMultiFactorAddedEventType   = loginPolicyMultiFactorPrefix + "added"
	LoginPolicyMultiFactorRemovedEventType = loginPolicyMultiFactorPrefix + "removed"
)

type SecondFactorAddedEvent struct {
	eventstore.BaseEvent `json:"-"`

	MFAType SecondFactorType `json:"mfaType"`
}

func NewSecondFactorAddedEvent(
	base *eventstore.BaseEvent,
	mfaType SecondFactorType,
) *SecondFactorAddedEvent {
	return &SecondFactorAddedEvent{
		BaseEvent: *base,
		MFAType:   mfaType,
	}
}

func SecondFactorAddedEventMapper(event *repository.Event) (eventstore.EventReader, error) {
	e := &SecondFactorAddedEvent{
		BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "POLIC-Lp0dE", "unable to unmarshal policy")
	}

	return e, nil
}

func (e *SecondFactorAddedEvent) CheckPrevious() bool {
	return true
}

func (e *SecondFactorAddedEvent) Data() interface{} {
	return e
}

type SecondFactorRemovedEvent struct {
	eventstore.BaseEvent `json:"-"`
	MFAType              SecondFactorType `json:"mfaType"`
}

func NewSecondFactorRemovedEvent(
	base *eventstore.BaseEvent,
	mfaType SecondFactorType,
) *SecondFactorRemovedEvent {
	return &SecondFactorRemovedEvent{
		BaseEvent: *base,
		MFAType:   mfaType,
	}
}

func SecondFactorRemovedEventMapper(event *repository.Event) (eventstore.EventReader, error) {
	e := &SecondFactorRemovedEvent{
		BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "POLIC-5M9gd", "unable to unmarshal policy")
	}

	return e, nil
}

func (e *SecondFactorRemovedEvent) CheckPrevious() bool {
	return true
}

func (e *SecondFactorRemovedEvent) Data() interface{} {
	return e
}

type MultiFactorAddedEvent struct {
	eventstore.BaseEvent `json:"-"`

	MFAType MultiFactorType `json:"mfaType"`
}

func NewMultiFactorAddedEvent(
	base *eventstore.BaseEvent,
	mfaType MultiFactorType,
) *MultiFactorAddedEvent {
	return &MultiFactorAddedEvent{
		BaseEvent: *base,
		MFAType:   mfaType,
	}
}

func MultiFactorAddedEventMapper(event *repository.Event) (eventstore.EventReader, error) {
	e := &MultiFactorAddedEvent{
		BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "POLIC-5Ms90", "unable to unmarshal policy")
	}

	return e, nil
}

func (e *MultiFactorAddedEvent) CheckPrevious() bool {
	return true
}

func (e *MultiFactorAddedEvent) Data() interface{} {
	return e
}

type MultiFactorRemovedEvent struct {
	eventstore.BaseEvent `json:"-"`
	MFAType              MultiFactorType `json:"mfaType"`
}

func NewMultiFactorRemovedEvent(
	base *eventstore.BaseEvent,
	mfaType MultiFactorType,
) *MultiFactorRemovedEvent {
	return &MultiFactorRemovedEvent{
		BaseEvent: *base,
		MFAType:   mfaType,
	}
}

func MultiFactorRemovedEventMapper(event *repository.Event) (eventstore.EventReader, error) {
	e := &MultiFactorRemovedEvent{
		BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "POLIC-1N8sd", "unable to unmarshal policy")
	}

	return e, nil
}

func (e *MultiFactorRemovedEvent) CheckPrevious() bool {
	return true
}

func (e *MultiFactorRemovedEvent) Data() interface{} {
	return e
}

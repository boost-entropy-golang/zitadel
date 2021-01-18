package model

import (
	"encoding/json"

	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore/models"
	es_models "github.com/caos/zitadel/internal/eventstore/models"
	iam_model "github.com/caos/zitadel/internal/iam/model"
)

type LabelPolicy struct {
	models.ObjectRoot
	State          int32  `json:"-"`
	PrimaryColor   string `json:"primaryColor"`
	SecondaryColor string `json:"secondaryColor"`
	WarnColor	   string `json:"warnColor"`
	Logo 		   []byte `json:"logo"`
	Border         bool   `json:"border"`
	Font           int32  `json:"font"`
	IconFont       int32  `json:"iconFont"`      
}

func LabelPolicyToModel(policy *LabelPolicy) *iam_model.LabelPolicy {
	return &iam_model.LabelPolicy{
		ObjectRoot:     policy.ObjectRoot,
		State:          iam_model.PolicyState(policy.State),
		PrimaryColor:   policy.PrimaryColor,
		SecondaryColor: policy.SecondaryColor,
	}
}

func LabelPolicyFromModel(policy *iam_model.LabelPolicy) *LabelPolicy {
	return &LabelPolicy{
		ObjectRoot:     policy.ObjectRoot,
		State:          int32(policy.State),
		PrimaryColor:   policy.PrimaryColor,
		SecondaryColor: policy.SecondaryColor,
	}
}

func (p *LabelPolicy) Changes(changed *LabelPolicy) map[string]interface{} {
	changes := make(map[string]interface{}, 2)

	if changed.PrimaryColor != p.PrimaryColor {
		changes["primaryColor"] = changed.PrimaryColor
	}
	if changed.SecondaryColor != p.SecondaryColor {
		changes["secondaryColor"] = changed.SecondaryColor
	}
	if changed.WarnColor != p.WarnColor {
		changes["warnColor"] = changed.WarnColor
	}
	if changed.Border != p.Border {
		changes["border"] = changed.Border
	}
	if changed.Font != p.Font {
		changes["font"] = changed.Font
	}
	if changed.IconFont != p.IconFont {
		changes["iconFont"] = changed.IconFont
	}

	return changes
}

func (i *IAM) appendAddLabelPolicyEvent(event *es_models.Event) error {
	i.DefaultLabelPolicy = new(LabelPolicy)
	err := i.DefaultLabelPolicy.SetDataLabel(event)
	if err != nil {
		return err
	}
	i.DefaultLabelPolicy.ObjectRoot.CreationDate = event.CreationDate
	return nil
}

func (i *IAM) appendChangeLabelPolicyEvent(event *es_models.Event) error {
	return i.DefaultLabelPolicy.SetDataLabel(event)
}

func (p *LabelPolicy) SetDataLabel(event *es_models.Event) error {
	err := json.Unmarshal(event.Data, p)
	if err != nil {
		return errors.ThrowInternal(err, "MODEL-ikjhf", "unable to unmarshal data")
	}
	return nil
}

func (p *IDPProvider) SetDataLabel(event *es_models.Event) error {
	err := json.Unmarshal(event.Data, p)
	if err != nil {
		return errors.ThrowInternal(err, "MODEL-c41Hn", "unable to unmarshal data")
	}
	return nil
}

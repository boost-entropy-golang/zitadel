package model

import (
	"time"

	req_model "github.com/caos/zitadel/internal/auth_request/model"
	"github.com/caos/zitadel/internal/model"
)

type UserSessionView struct {
	CreationDate                 time.Time
	ChangeDate                   time.Time
	State                        req_model.UserSessionState
	ResourceOwner                string
	UserAgentID                  string
	UserID                       string
	UserName                     string
	LoginName                    string
	DisplayName                  string
	SelectedIDPConfigID          string
	PasswordVerification         time.Time
	PasswordlessVerification     time.Time
	ExternalLoginVerification    time.Time
	SecondFactorVerification     time.Time
	SecondFactorVerificationType req_model.MFAType
	MultiFactorVerification      time.Time
	MultiFactorVerificationType  req_model.MFAType
	Sequence                     uint64
}

type UserSessionSearchRequest struct {
	Offset        uint64
	Limit         uint64
	SortingColumn UserSessionSearchKey
	Asc           bool
	Queries       []*UserSessionSearchQuery
}

type UserSessionSearchKey int32

const (
	UserSessionSearchKeyUnspecified UserSessionSearchKey = iota
	UserSessionSearchKeyUserAgentID
	UserSessionSearchKeyUserID
	UserSessionSearchKeyState
	UserSessionSearchKeyResourceOwner
)

type UserSessionSearchQuery struct {
	Key    UserSessionSearchKey
	Method model.SearchMethod
	Value  interface{}
}

type UserSessionSearchResponse struct {
	Offset      uint64
	Limit       uint64
	TotalResult uint64
	Result      []*UserSessionView
}

func (r *UserSessionSearchRequest) EnsureLimit(limit uint64) {
	if r.Limit == 0 || r.Limit > limit {
		r.Limit = limit
	}
}

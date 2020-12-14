package handler

import (
	"net/http"

	http_mw "github.com/caos/zitadel/internal/api/http/middleware"
	"github.com/caos/zitadel/internal/auth_request/model"
	"github.com/caos/zitadel/internal/errors"
)

const (
	queryInitPWCode   = "code"
	queryInitPWUserID = "userID"

	tmplInitPassword     = "initpassword"
	tmplInitPasswordDone = "initpassworddone"
)

type initPasswordFormData struct {
	Code            string `schema:"code"`
	Password        string `schema:"password"`
	PasswordConfirm string `schema:"passwordconfirm"`
	UserID          string `schema:"userID"`
	Resend          bool   `schema:"resend"`
}

type initPasswordData struct {
	baseData
	profileData
	Code                      string
	UserID                    string
	PasswordPolicyDescription string
	MinLength                 uint64
	HasUppercase              string
	HasLowercase              string
	HasNumber                 string
	HasSymbol                 string
}

func (l *Login) handleInitPassword(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue(queryInitPWUserID)
	code := r.FormValue(queryInitPWCode)
	l.renderInitPassword(w, r, nil, userID, code, nil)
}

func (l *Login) handleInitPasswordCheck(w http.ResponseWriter, r *http.Request) {
	data := new(initPasswordFormData)
	authReq, err := l.getAuthRequestAndParseData(r, data)
	if err != nil {
		l.renderError(w, r, authReq, err)
		return
	}

	if data.Resend {
		l.resendPasswordSet(w, r, authReq)
		return
	}
	l.checkPWCode(w, r, authReq, data, nil)
}

func (l *Login) checkPWCode(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, data *initPasswordFormData, err error) {
	if data.Password != data.PasswordConfirm {
		err := errors.ThrowInvalidArgument(nil, "VIEW-KaGue", "Errors.User.Password.ConfirmationWrong")
		l.renderInitPassword(w, r, authReq, data.UserID, data.Code, err)
		return
	}
	userOrg := login
	if authReq != nil {
		userOrg = authReq.UserOrgID
	}
	userAgentID, _ := http_mw.UserAgentIDFromCtx(r.Context())
	err = l.authRepo.SetPassword(setContext(r.Context(), userOrg), data.UserID, data.Code, data.Password, userAgentID)
	if err != nil {
		l.renderInitPassword(w, r, authReq, data.UserID, "", err)
		return
	}
	l.renderInitPasswordDone(w, r, authReq)
}

func (l *Login) resendPasswordSet(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest) {
	userOrg := login
	if authReq != nil {
		userOrg = authReq.UserOrgID
	}
	err := l.authRepo.RequestPasswordReset(setContext(r.Context(), userOrg), authReq.LoginName)
	l.renderInitPassword(w, r, authReq, authReq.UserID, "", err)
}

func (l *Login) renderInitPassword(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, userID, code string, err error) {
	var errType, errMessage string
	if err != nil {
		errMessage = l.getErrorMessage(r, err)
	}
	if userID == "" && authReq != nil {
		userID = authReq.UserID
	}

	data := initPasswordData{
		baseData:    l.getBaseData(r, authReq, "Init Password", errType, errMessage),
		profileData: l.getProfileData(authReq),
		UserID:      userID,
		Code:        code,
	}
	policy, description, _ := l.getPasswordComplexityPolicyByUserID(r, userID)
	if policy != nil {
		data.PasswordPolicyDescription = description
		data.MinLength = policy.MinLength
		if policy.HasUppercase {
			data.HasUppercase = UpperCaseRegex
		}
		if policy.HasLowercase {
			data.HasLowercase = LowerCaseRegex
		}
		if policy.HasSymbol {
			data.HasSymbol = SymbolRegex
		}
		if policy.HasNumber {
			data.HasNumber = NumberRegex
		}
	}
	l.renderer.RenderTemplate(w, r, l.renderer.Templates[tmplInitPassword], data, nil)
}

func (l *Login) renderInitPasswordDone(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest) {
	data := l.getUserData(r, authReq, "Password Init Done", "", "")
	l.renderer.RenderTemplate(w, r, l.renderer.Templates[tmplInitPasswordDone], data, nil)
}

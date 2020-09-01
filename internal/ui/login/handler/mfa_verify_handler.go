package handler

import (
	"net/http"

	http_mw "github.com/caos/zitadel/internal/api/http/middleware"
	"github.com/caos/zitadel/internal/auth_request/model"
)

const (
	tmplMfaVerify = "MfaVerify"
)

type mfaVerifyFormData struct {
	MfaType model.MfaType `schema:"mfaType"`
	Code    string        `schema:"code"`
}

func (l *Login) handleMfaVerify(w http.ResponseWriter, r *http.Request) {
	data := new(mfaVerifyFormData)
	authReq, err := l.getAuthRequestAndParseData(r, data)
	if err != nil {
		l.renderError(w, r, authReq, err)
		return
	}
	if data.MfaType == model.MfaTypeOTP {
		userAgentID, _ := http_mw.UserAgentIDFromCtx(r.Context())
		err = l.authRepo.VerifyMfaOTP(setContext(r.Context(), authReq.UserOrgID), authReq.ID, authReq.UserID, data.Code, userAgentID, model.BrowserInfoFromRequest(r))
	}
	if err != nil {
		l.renderError(w, r, authReq, err)
		return
	}
	l.renderNextStep(w, r, authReq)
}

func (l *Login) renderMfaVerify(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, verificationStep *model.MfaVerificationStep, err error) {
	data := l.getUserData(r, authReq, tmplMfaVerify, err)
	if verificationStep != nil {
		data.MfaProviders = verificationStep.MfaProviders
		data.SelectedMfaProvider = verificationStep.MfaProviders[0]
	}
	l.renderer.RenderTemplate(w, r, l.renderer.Templates[tmplMfaVerify], data, nil)
}

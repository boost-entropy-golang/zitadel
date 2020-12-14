package handler

import (
	"net/http"

	http_mw "github.com/caos/zitadel/internal/api/http/middleware"
	"github.com/caos/zitadel/internal/auth_request/model"
)

const (
	tmplMFAVerify = "mfaverify"
)

type mfaVerifyFormData struct {
	MFAType model.MFAType `schema:"mfaType"`
	Code    string        `schema:"code"`
}

func (l *Login) handleMFAVerify(w http.ResponseWriter, r *http.Request) {
	data := new(mfaVerifyFormData)
	authReq, err := l.getAuthRequestAndParseData(r, data)
	if err != nil {
		l.renderError(w, r, authReq, err)
		return
	}
	if data.MFAType == model.MFATypeOTP {
		userAgentID, _ := http_mw.UserAgentIDFromCtx(r.Context())
		err = l.authRepo.VerifyMFAOTP(setContext(r.Context(), authReq.UserOrgID), authReq.ID, authReq.UserID, data.Code, userAgentID, model.BrowserInfoFromRequest(r))
	}
	if err != nil {
		l.renderError(w, r, authReq, err)
		return
	}
	l.renderNextStep(w, r, authReq)
}

func (l *Login) renderMFAVerify(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, verificationStep *model.MFAVerificationStep, err error) {
	if verificationStep == nil {
		l.renderError(w, r, authReq, err)
		return
	}
	provider := verificationStep.MFAProviders[len(verificationStep.MFAProviders)-1]
	l.renderMFAVerifySelected(w, r, authReq, verificationStep, provider, err)
}

func (l *Login) renderMFAVerifySelected(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, verificationStep *model.MFAVerificationStep, selectedProvider model.MFAType, err error) {
	var errType, errMessage string
	if err != nil {
		errMessage = l.getErrorMessage(r, err)
	}
	data := l.getUserData(r, authReq, "MFA Verify", errType, errMessage)
	if verificationStep == nil {
		l.renderError(w, r, authReq, err)
		return
	}
	switch selectedProvider {
	case model.MFATypeU2F:
		l.renderU2FVerification(w, r, authReq, removeSelectedProviderFromList(verificationStep.MFAProviders, model.MFATypeU2F), nil)
		return
	case model.MFATypeOTP:
		data.MFAProviders = removeSelectedProviderFromList(verificationStep.MFAProviders, model.MFATypeOTP)
		data.SelectedMFAProvider = model.MFATypeOTP
	}
	l.renderer.RenderTemplate(w, r, l.renderer.Templates[tmplMFAVerify], data, nil)
}

func removeSelectedProviderFromList(providers []model.MFAType, selected model.MFAType) []model.MFAType {
	for i := len(providers) - 1; i >= 0; i-- {
		if providers[i] == selected {
			copy(providers[i:], providers[i+1:])
			return providers[:len(providers)-1]
		}
	}
	return providers
}

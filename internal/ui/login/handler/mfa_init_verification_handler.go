package handler

import (
	"bytes"
	"net/http"

	svg "github.com/ajstarks/svgo"
	"github.com/boombuler/barcode/qr"

	"github.com/caos/zitadel/internal/auth_request/model"
	"github.com/caos/zitadel/internal/qrcode"
)

const (
	tmplMfaInitVerification = "MfaInitVerification"
	tmplMfaInitDone         = "MfaInitDone"
)

type mfaInitVerificationData struct {
	MfaType model.MfaType `schema:"mfaType"`
	Code    string        `schema:"code"`
	URL     string        `schema:"url"`
	Secret  string        `schema:"secret"`
}

func (l *Login) handleMfaInitVerification(w http.ResponseWriter, r *http.Request) {
	data := new(mfaInitVerificationData)
	authReq, err := l.getAuthRequestAndParseData(r, data)
	if err != nil {
		l.renderError(w, r, authReq, err)
		return
	}
	var verifyData *mfaVerifyData
	switch data.MfaType {
	case model.MfaTypeOTP:
		verifyData = l.otpVerificationData(w, r, authReq, data)
	}

	if verifyData != nil {
		l.renderMfaInitVerification(w, r, authReq, verifyData, err)
		return
	}

	done := &mfaDoneData{
		MfaType: data.MfaType,
	}
	l.renderMfaInitDone(w, r, authReq, done)
}

func (l *Login) otpVerificationData(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, data *mfaInitVerificationData) *mfaVerifyData {
	err := l.authRepo.VerifyMfaOTPSetup(setContext(r.Context(), authReq.UserOrgID), authReq.UserID, data.Code)
	if err == nil {
		return nil
	}
	mfadata := &mfaVerifyData{
		MfaType: data.MfaType,
		otpData: otpData{
			Secret: data.Secret,
			Url:    data.URL,
		},
	}

	return mfadata
}

func (l *Login) renderMfaInitVerification(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, data *mfaVerifyData, err error) {
	data.baseData = l.getBaseData(r, authReq, tmplMfaInitVerification, err)
	data.profileData = l.getProfileData(authReq)
	if data.MfaType == model.MfaTypeOTP {
		code, err := generateQrCode(data.otpData.Url)
		if err == nil {
			data.otpData.QrCode = code
		}
	}

	l.renderer.RenderTemplate(w, r, l.renderer.Templates[tmplMfaInitVerification], data, nil)
}

func (l *Login) renderMfaInitDone(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, data *mfaDoneData) {
	data.baseData = l.getBaseData(r, authReq, tmplMfaInitDone, nil)
	data.profileData = l.getProfileData(authReq)
	l.renderer.RenderTemplate(w, r, l.renderer.Templates[tmplMfaInitDone], data, nil)
}

func generateQrCode(url string) (string, error) {
	var b bytes.Buffer
	s := svg.New(&b)

	qrCode, err := qr.Encode(url, qr.M, qr.Auto)
	if err != nil {
		return "", err
	}
	qs := qrcode.NewQrSVG(qrCode, 5)
	qs.StartQrSVG(s)
	qs.WriteQrSVG(s)

	s.End()
	return string(b.Bytes()), nil
}

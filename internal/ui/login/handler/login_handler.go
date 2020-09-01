package handler

import (
	"net/http"

	http_mw "github.com/caos/zitadel/internal/api/http/middleware"
	"github.com/caos/zitadel/internal/auth_request/model"
)

const (
	tmplLogin = "Login"
)

type loginData struct {
	LoginName string `schema:"loginName"`
	Register  bool   `schema:"register"`
}

func (l *Login) handleLogin(w http.ResponseWriter, r *http.Request) {
	authReq, err := l.getAuthRequest(r)
	if err != nil {
		l.renderError(w, r, authReq, err)
		return
	}
	if authReq == nil {
		http.Redirect(w, r, l.zitadelURL, http.StatusFound)
		return
	}
	l.renderNextStep(w, r, authReq)
}

func (l *Login) handleLoginName(w http.ResponseWriter, r *http.Request) {
	authSession, err := l.getAuthRequest(r)
	if err != nil {
		l.renderError(w, r, authSession, err)
		return
	}
	l.renderLogin(w, r, authSession, nil)
}

func (l *Login) handleLoginNameCheck(w http.ResponseWriter, r *http.Request) {
	data := new(loginData)
	authReq, err := l.getAuthRequestAndParseData(r, data)
	if err != nil {
		l.renderError(w, r, authReq, err)
		return
	}
	if data.Register {
		l.handleRegistration(w, r)
		return
	}
	userAgentID, _ := http_mw.UserAgentIDFromCtx(r.Context())
	err = l.authRepo.CheckLoginName(r.Context(), authReq.ID, data.LoginName, userAgentID)
	if err != nil {
		l.renderLogin(w, r, authReq, err)
		return
	}
	l.renderNextStep(w, r, authReq)
}

func (l *Login) renderLogin(w http.ResponseWriter, r *http.Request, authReq *model.AuthRequest, err error) {
	data := l.getUserData(r, authReq, tmplLogin, err)
	l.renderer.RenderTemplate(w, r, l.renderer.Templates[tmplLogin], data, nil)
}

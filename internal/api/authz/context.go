package authz

import (
	"context"

	"github.com/caos/zitadel/internal/api/grpc"
	http_util "github.com/caos/zitadel/internal/api/http"
	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/telemetry/tracing"
)

type key int

const (
	requestPermissionsKey key = 1
	dataKey               key = 2
	allPermissionsKey     key = 3
)

type CtxData struct {
	UserID            string
	OrgID             string
	ProjectID         string
	AgentID           string
	PreferredLanguage string
}

func (ctxData CtxData) IsZero() bool {
	return ctxData.UserID == "" || ctxData.OrgID == ""
}

type Grants []*Grant

type Grant struct {
	OrgID string
	Roles []string
}

func VerifyTokenAndWriteCtxData(ctx context.Context, token, orgID string, t *TokenVerifier, method string) (_ context.Context, err error) {
	ctx, span := tracing.NewSpan(ctx)
	defer func() { span.EndWithError(err) }()

	if orgID != "" {
		err = t.ExistsOrg(ctx, orgID)
		if err != nil {
			return nil, errors.ThrowPermissionDenied(nil, "AUTH-Bs7Ds", "Organisation doesn't exist")
		}
	}

	userID, clientID, agentID, prefLang, err := verifyAccessToken(ctx, token, t, method)
	if err != nil {
		return nil, err
	}
	projectID, origins, err := t.ProjectIDAndOriginsByClientID(ctx, clientID)
	if err != nil {
		return nil, errors.ThrowPermissionDenied(err, "AUTH-GHpw2", "could not read projectid by clientid")
	}
	if err := checkOrigin(ctx, origins); err != nil {
		return nil, err
	}
	return context.WithValue(ctx, dataKey, CtxData{UserID: userID, OrgID: orgID, ProjectID: projectID, AgentID: agentID, PreferredLanguage: prefLang}), nil
}

func SetCtxData(ctx context.Context, ctxData CtxData) context.Context {
	return context.WithValue(ctx, dataKey, ctxData)
}

func GetCtxData(ctx context.Context) CtxData {
	ctxData, _ := ctx.Value(dataKey).(CtxData)
	return ctxData
}

func GetRequestPermissionsFromCtx(ctx context.Context) []string {
	ctxPermission, _ := ctx.Value(requestPermissionsKey).([]string)
	return ctxPermission
}

func GetAllPermissionsFromCtx(ctx context.Context) []string {
	ctxPermission, _ := ctx.Value(allPermissionsKey).([]string)
	return ctxPermission
}

func checkOrigin(ctx context.Context, origins []string) error {
	origin := grpc.GetGatewayHeader(ctx, http_util.Origin)
	if origin == "" {
		return nil
	}
	if http_util.IsOriginAllowed(origins, origin) {
		return nil
	}
	return errors.ThrowPermissionDenied(nil, "AUTH-DZG21", "Errors.OriginNotAllowed")
}

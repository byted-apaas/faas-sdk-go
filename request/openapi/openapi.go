package openapi

import (
	"context"
	"strconv"

	"github.com/byted-apaas/faas-sdk-go/common/structs"
	"github.com/byted-apaas/faas-sdk-go/common/utils"
	reqCommon "github.com/byted-apaas/faas-sdk-go/request/common"
	cConstants "github.com/byted-apaas/server-common-go/constants"
	cExceptions "github.com/byted-apaas/server-common-go/exceptions"
	cHttp "github.com/byted-apaas/server-common-go/http"
	cStructs "github.com/byted-apaas/server-common-go/structs"
	cUtils "github.com/byted-apaas/server-common-go/utils"
	"github.com/tidwall/gjson"
)

type RequestHttp struct{}

func (r *RequestHttp) InvokeFunctionWithAuth(ctx context.Context, appCtx *structs.AppCtx, apiName string, params interface{}, result interface{}) error {
	ctx = utils.SetCtx(ctx, appCtx, cConstants.InvokeFuncWithAuth)

	body, err := reqCommon.BuildInvokeParamsObj(ctx, apiName, params, appCtx == nil || appCtx.Credential == nil || appCtx.Mode != structs.AppModeOpenSDK)
	if err != nil {
		return err
	}

	namespace, err := utils.GetNamespace(ctx, appCtx)
	if err != nil {
		return err
	}
	headers := map[string][]string{
		cConstants.HttpHeaderKeyUser: {strconv.FormatInt(cUtils.GetUserIDFromCtx(ctx), 10)},
	}

	respBody, extra, err := getOpenapiClient().PostJson(ctx, GetPathInvokeFunctionWithAuth(namespace, apiName), headers, body, cHttp.AppTokenMiddleware)
	data, err := errorWrapper(respBody, extra, err)
	if err != nil {
		return err
	}

	var resp struct {
		Result string `json:"result"`
	}

	logid := cUtils.GetLogIDFromExtra(extra)
	if err := cUtils.JsonUnmarshalBytes(data, &resp); err != nil {
		return cExceptions.InternalError("InvokeFunctionWithAuth failed, err: %v, logid: %v", err, logid)
	}

	code := gjson.GetBytes([]byte(resp.Result), "code").String()
	if code != "0" {
		msg := gjson.GetBytes([]byte(resp.Result), "msg").String()
		return cExceptions.InvalidParamError("%v ([%v] %v)", msg, code, logid)
	}

	dataRaw := gjson.GetBytes([]byte(resp.Result), "data").Raw
	if len(dataRaw) > 0 {
		if err := cUtils.JsonUnmarshalBytes([]byte(dataRaw), result); err != nil {
			return cExceptions.InvalidParamError("InvokeFunctionWithAuth failed, err: %v", err)
		}
	}

	permission := cStructs.Permission{}
	permissionRaw := gjson.GetBytes([]byte(resp.Result), "permission").Raw
	if len(permissionRaw) > 0 {
		if err := cUtils.JsonUnmarshalBytes([]byte(permissionRaw), &permission); err != nil {
			return cExceptions.InvalidParamError("InvokeFunctionWithAuth failed, err: %v, logid: %v", err, logid)
		}
	}

	_, err = cHttp.AppendParamsUnauthFields(ctx, apiName, "output", result, permission.UnauthFields)
	if err != nil {
		return err
	}

	return nil
}

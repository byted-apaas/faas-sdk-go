package innerapi

type RequestRpc struct{}

//func (r *RequestRpc) pre(ctx context.Context, appCtx *structs.AppCtx, method string) (context.Context, context.CancelFunc, string, error) {
//	var err error
//	ctx, err = cHttp.RebuildRpcCtx(utils.SetCtx(ctx, appCtx, method))
//	if err != nil {
//		return nil, nil, "", err
//	}
//
//	namespace, err := utils.GetNamespace(ctx, appCtx)
//	if err != nil {
//		return nil, nil, "", err
//	}
//
//	if namespace == "" {
//		return nil, nil, "", cExceptions.InternalError("namespace is empty")
//	}
//
//	var cancel context.CancelFunc
//	ctx, cancel = cHttp.GetTimeoutCtx(ctx)
//	return ctx, cancel, namespace, nil
//}
//
//func (r *RequestRpc) post(ctx context.Context, err error, baseResp *base.BaseResp, baseReq *base.Base) error {
//	var logid string
//	if baseReq != nil {
//		logid = baseReq.LogID
//	}
//
//	if err != nil {
//		return cExceptions.InternalError("Call InnerAPI failed: %+v, logid: %s", err, logid)
//	}
//
//	if baseResp == nil {
//		return cExceptions.InternalError("Call InnerAPI resp is empty, logid: %s", logid)
//	}
//
//	if baseResp.KStatusCode != "" {
//		msg := baseResp.KStatusMessage
//		if baseResp.StatusMessage != "" {
//			msg = baseResp.StatusMessage
//		}
//		return cExceptions.NewErrWithCodeV2(baseResp.KStatusCode, msg, logid)
//	}
//	return nil
//}
//
//func (r *RequestRpc) InvokeFunctionWithAuth(ctx context.Context, appCtx *structs.AppCtx, apiName string, params interface{}, result interface{}) error {
//	sysParams, bizParams, err := reqCommon.BuildInvokeParamAndContext(ctx, params, apiName, appCtx == nil || appCtx.Credential == nil || appCtx.Mode != structs.AppModeOpenSDK)
//	if err != nil {
//		return err
//	}
//
//	req := cloudfunction.NewInvokeFunctionWithAuthRequest()
//
//	ctx, cancel, _, err := r.pre(ctx, appCtx, cConstants.InvokeFuncWithAuth)
//	if err != nil {
//		return err
//	}
//	defer cancel()
//	ctx = utils.SetUserMetaInfoToContext(ctx, appCtx)
//
//	namespace, err := utils.GetNamespace(ctx, appCtx)
//	if err != nil {
//		return err
//	}
//	req.Namespace = namespace
//	req.ApiName = cUtils.StringPtr(apiName)
//	req.Context = cUtils.StringPtr(sysParams)
//	req.Params = cUtils.StringPtr(bizParams)
//
//	cli, err := cHttp.GetInnerAPICli(ctx)
//	if err != nil {
//		return err
//	}
//
//	resp, err := cli.InvokeFunctionWithAuth(ctx, req)
//
//	var baseResp *base.BaseResp
//	if resp != nil {
//		baseResp = resp.BaseResp
//	}
//	if err = r.post(ctx, err, baseResp, req.Base); err != nil {
//		return err
//	}
//
//	var logid string
//	if req.Base != nil {
//		logid = req.Base.LogID
//	}
//
//	if resp.Result_ != nil {
//		data := []byte(*resp.Result_)
//		code := gjson.GetBytes(data, "code").String()
//		if code != "0" {
//			msg := gjson.GetBytes(data, "msg").String()
//			return cExceptions.InvalidParamError("%v ([%v] %v)", msg, code, logid)
//		}
//
//		dataRaw := gjson.GetBytes(data, "data").Raw
//		if len(dataRaw) > 0 {
//			if err := cUtils.JsonUnmarshalBytes([]byte(dataRaw), &result); err != nil {
//				return cExceptions.InvalidParamError("InvokeFunctionWithAuth failed, err: %v, logid: %v", err, logid)
//			}
//		}
//
//		permission := cStructs.Permission{}
//		permissionRaw := gjson.GetBytes(data, "permission").Raw
//		if len(permissionRaw) > 0 {
//			if err := cUtils.JsonUnmarshalBytes([]byte(permissionRaw), &permission); err != nil {
//				return cExceptions.InvalidParamError("InvokeFunctionWithAuth failed, err: %v, logid: %v", err, logid)
//			}
//		}
//
//		_, err := cHttp.AppendParamsUnauthFields(ctx, apiName, "output", result, permission.UnauthFields)
//		if err != nil {
//			return err
//		}
//		return nil
//	}
//
//	return nil
//}

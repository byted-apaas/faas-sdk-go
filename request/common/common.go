package common

import (
	"context"

	cHttp "github.com/byted-apaas/server-common-go/http"
	cUtils "github.com/byted-apaas/server-common-go/utils"
)

func BuildInvokeParamsObj(ctx context.Context, apiName string, params interface{}, needPermission bool) (map[string]interface{}, error) {
	return map[string]interface{}{
		"apiAlias":    apiName,
		"params":      params,
		"context":     BuildInvokeSysParams(ctx, params, apiName, needPermission),
		"triggerType": "default-trigger",
	}, nil
}

func BuildInvokeSysParams(ctx context.Context, params interface{}, funcAPIName string, needPermission bool) map[string]interface{} {
	v := map[string]interface{}{
		"triggertaskid":             cUtils.GetTriggerTaskIDFromCtx(ctx),
		"x-kunlun-distributed-mask": cUtils.GetDistributedMaskFromCtx(ctx),
		"x-kunlun-loop-masks":       cUtils.GetLoopMaskFromCtx(ctx),
	}

	if needPermission {
		v["permission"] = cHttp.CalcParamsNeedPermission(ctx, funcAPIName, "input", params)
	}
	return v
}

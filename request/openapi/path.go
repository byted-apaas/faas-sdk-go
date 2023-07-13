package openapi

import (
	cUtils "github.com/byted-apaas/server-common-go/utils"
)

const (
	PathInvokeFunctionWithAuth = "/api/cloudfunction/v1/namespaces/:namespace/invoke/:functionAPIName"
)

func GetPathInvokeFunctionWithAuth(namespace, functionAPIName string) string {
	return cUtils.NewPathReplace(PathInvokeFunctionWithAuth).Namespace(namespace).FunctionAPIName(functionAPIName).Path()
}

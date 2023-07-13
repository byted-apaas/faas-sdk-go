package request

import (
	"context"
	"sync"

	"github.com/byted-apaas/faas-sdk-go/common/structs"
	"github.com/byted-apaas/faas-sdk-go/request/openapi"
)

//go:generate mockery --name=IRequestOpenapi --structname=RequestOpenapi --filename=RequestOpenapi.go
type IRequestOpenapi interface {
	InvokeFunctionWithAuth(ctx context.Context, appCtx *structs.AppCtx, apiName string, params interface{}, result interface{}) error
}

var (
	reqHTTP     IRequestOpenapi
	reqHTTPOnce sync.Once
)

func GetInstance(ctx context.Context) IRequestOpenapi {
	return GetHTTPInstance()
}

func GetHTTPInstance() IRequestOpenapi {
	if reqHTTP == nil {
		reqHTTPOnce.Do(func() {
			reqHTTP = &openapi.RequestHttp{}
		})
	}
	return reqHTTP
}

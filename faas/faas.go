package faas

import (
	functionImpl "github.com/byted-apaas/faas-sdk-go/service/function/impl"
	"github.com/byted-apaas/faas-sdk-go/service/tools"
)

var (
	Tool     = tools.NewTools(nil)
	Function = functionImpl.Function
)

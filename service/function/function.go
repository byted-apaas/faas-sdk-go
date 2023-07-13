package function

import (
	"context"
)

type IFunction interface {
	Invoke(ctx context.Context, params map[string]interface{}, result interface{}) error
}

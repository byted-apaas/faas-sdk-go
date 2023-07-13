package tools

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	cUtils "github.com/byted-apaas/server-common-go/utils"
)

var (
	ctx = context.Background()
)

func Init() {

}

func TestMain(m *testing.M) {
	Init()
	m.Run()
}

var (
	tools = &Tools{}
)

func Test_Retry(t *testing.T) {
	err := tools.Retry(func() error {
		e := errors.New("err is not nil")
		fmt.Printf("err: %v\n", e)
		return e
	}, NewRetryOption(2, time.Second))
	cUtils.PrintLog(err)
}

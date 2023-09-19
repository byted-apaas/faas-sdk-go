package constants

import (
	"fmt"

	cConstants "github.com/byted-apaas/server-common-go/constants"
)

type PlatformEnvType int

const (
	PlatformEnvLR PlatformEnvType = iota + 3
	PlatformEnvPRE
	PlatformEnvOnline
)

func (p PlatformEnvType) String() string {
	switch p {
	case PlatformEnvLR:
		return cConstants.EnvTypeLr
	case PlatformEnvPRE:
		return cConstants.EnvTypeGray
	case PlatformEnvOnline:
		return cConstants.EnvTypeOnline
	}
	fmt.Printf("invalid platform env type %d", p)
	return ""
}

type OperationType int

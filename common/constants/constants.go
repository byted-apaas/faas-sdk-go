package constants

import (
	"fmt"

	cConstants "github.com/byted-apaas/server-common-go/constants"
)

type PlatformEnvType int

const (
	PlatformEnvUAT PlatformEnvType = iota + 2
	PlatformEnvLR
	PlatformEnvPRE
	PlatformEnvOnline
)

func (p PlatformEnvType) String() string {
	switch p {
	case PlatformEnvUAT:
		return cConstants.EnvTypeStaging
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

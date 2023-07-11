package constants

import (
	"fmt"

	cConstants "github.com/byted-apaas/server-common-go/constants"
)

type PlatformEnvType int

const (
	PlatformEnvDEV PlatformEnvType = iota + 1
	PlatformEnvUAT
	PlatformEnvLR
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

const (
	OpRateLimitError = "k_op_ec_20003"
	FsRateLimitError = "k_fs_ec_000004"
)

// Package version defines version of faas-sdk-go.
package version

import (
	"sync"

	cVersion "github.com/byted-apaas/server-common-go/version"
)

const Version = "v0.0.8"

const SDKName = "byted-apaas/faas-sdk-go"

type FaaSSDKInfo struct{}

func (b *FaaSSDKInfo) GetVersion() string {
	return Version
}

func (b *FaaSSDKInfo) GetSDKName() string {
	return SDKName
}

var (
	faasSDKInfoOnce sync.Once
	faasSDKInfo     cVersion.ISDKInfo
)

func GetFaaSSDKInfo() cVersion.ISDKInfo {
	if faasSDKInfo == nil {
		faasSDKInfoOnce.Do(func() {
			faasSDKInfo = &FaaSSDKInfo{}
		})
	}
	return faasSDKInfo
}

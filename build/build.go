package build

import (
	"runtime/debug"
)

// these are populated via LDFLAGS -X
var (
	BuildVersion 	string
	BuildTime		string
	BuildUser		string
)

type VersionInfo struct {
	Version		string
	Timestamp	string
	user		string
}

func GetVersion(v *VersionInfo) {
	v.Version = BuildVersion
	v.Timestamp = BuildTime

	if v.Version == "" {
		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			v.Version =  buildInfo.Main.Version
		} else {
			v.Version = "(unknown)"
		}
	}
}
package version

import (
	"fmt"
	"runtime"
)

// These variables are replaced by ldflags at build time
var (
	version   = "v0.1.0-alpha"
	gitCommit = ""
	buildDate = "1970-01-01T00:00:00Z"
)

type VersionInfo struct {
	Version   string `json:"version" yaml:"version"`
	GitCommit string `json:"gitCommit" yaml:"gitCommit"`
	BuildDate string `json:"buildDate" yaml:"buildDate"`
	GoVersion string `json:"goVersion" yaml:"goVersion"`
	Compiler  string `json:"compiler" yaml:"compiler"`
	Platform  string `json:"platform" yaml:"platform"`
}

func Get() *VersionInfo {
	return &VersionInfo{
		Version:   version,
		GitCommit: gitCommit,
		BuildDate: buildDate,
		GoVersion: runtime.Version(),
		Compiler:  runtime.Compiler,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

func (v *VersionInfo) String() string {

	version_string := fmt.Sprintf(`revisio
Version:    %s,
GitCommit:  %s,
BuildDate:  %s,
GoVersion:  %s,
Compiler:   %s,
Platform:   %s`, v.Version,
		v.GitCommit,
		v.BuildDate,
		v.GoVersion,
		v.Compiler,
		v.Platform)
	return version_string
}

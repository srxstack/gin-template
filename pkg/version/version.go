// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package version

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/gosuri/uitable"
)

var (
	// gitVersion 是语义化的版本号.
	gitVersion = "v0.0.0-master+$Format:%h$"
	// buildDate 是 ISO8601 格式的构建时间, $(date -u +'%Y-%m-%dT%H:%M:%SZ') 命令的输出.
	buildDate = "1970-01-01T00:00:00Z"
	// gitCommit 是 Git 的 SHA1 值，$(git rev-parse HEAD) 命令的输出.
	gitCommit = "$Format:%H$"
	// gitTreeState 代表构建时 Git 仓库的状态，可能的值有：clean, dirty.
	gitTreeState = ""
)

type Info struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func (info Info) String() string {
	return info.GitVersion
}

func (info Info) ToJSON() string {
	s, _ := json.Marshal(info)

	return string(s)
}

func (info Info) Text() string {
	table := uitable.New()
	table.RightAlign(0)
	table.MaxColWidth = 80
	table.Separator = " "
	table.AddRow("gitVersion:", info.GitVersion)
	table.AddRow("gitCommit:", info.GitCommit)
	table.AddRow("gitTreeState:", info.GitTreeState)
	table.AddRow("buildDate:", info.BuildDate)
	table.AddRow("goVersion:", info.GoVersion)
	table.AddRow("compiler:", info.Compiler)
	table.AddRow("platform:", info.Platform)

	return table.String()
}

func Get() Info {
	return Info{
		GitVersion:   gitVersion,
		GitCommit:    gitCommit,
		GitTreeState: gitTreeState,
		BuildDate:    buildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

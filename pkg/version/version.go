//
// Copyright (c) 2023 Red Hat, Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package version

import (
	"fmt"
	"os"
	"runtime/debug"
)

type VersionInfo struct {
	Version   string
	GitCommit string
}

// getVcsRevision fetches the build information embedded in the binary by Golang and returns the VCS Revision (e.g., Git commit hash),
// if the binary was built from a Git project.
// Note that this can be disabled when building the binary with the -buildvcs flag.
func getVcsRevision() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}
	return ""
}

func GetVersionInfo() VersionInfo {
	return VersionInfo{
		Version:   os.Getenv("VERSION"),
		GitCommit: getVcsRevision(),
	}
}

func (v VersionInfo) String() string {
	var vMsg string
	if v.Version != "" {
		vMsg = fmt.Sprintf(" version %s", v.Version)
	}
	return fmt.Sprintf("janus-idp.io/backstage-operator%s (%s)", vMsg, v.GitCommit)
}

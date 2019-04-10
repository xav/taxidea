// Copyright © 2019 Xavier Basty <xbasty@gmail.com>
//
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

package main

import (
	"os"
	"taxidea/cmd"

	"github.com/rs/zerolog"
)

const taxideaVersion = "0.0.1"

var log = zerolog.New(os.Stdout).With().Timestamp().Logger()

func main() {
	openCmd := &cmd.OpenCmd{}
	cmd.RootCmd.AddCommand(openCmd.Init())

	versionCmd := &cmd.VersionCmd{TaxideaVersion: taxideaVersion}
	cmd.RootCmd.AddCommand(versionCmd.Init())

	cmd.Execute()
}

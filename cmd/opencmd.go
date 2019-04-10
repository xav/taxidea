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

package cmd

import (
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

// OpenCmd prints the current version.
type OpenCmd struct {
	TaxideaVersion string
}

// Init returns the runnable cobra command.
func (v *OpenCmd) Init() *cobra.Command {
	return &cobra.Command{
		Use:   "open",
		Short: "Open the specified Badger store",
		Run:   v.openStore,
	}
}

func (v *OpenCmd) openStore(cmd *cobra.Command, args []string) {
	box := tview.NewBox().SetBorder(true).SetTitle("Taxidea")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}

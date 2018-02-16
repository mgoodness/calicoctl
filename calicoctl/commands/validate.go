// Copyright (c) 2017 Tigera, Inc. All rights reserved.

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

package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/docopt/docopt-go"

	"github.com/projectcalico/calicoctl/calicoctl/commands/argutils"
	"github.com/projectcalico/calicoctl/calicoctl/commands/constants"
	"github.com/projectcalico/calicoctl/calicoctl/resourcemgr"
)

func Validate(args []string) {
	doc := constants.DatastoreIntro + `Usage:
  calicoctl validate --filename=<FILENAME>

Examples:
  # Validate the contents of policy.yaml.
  calicoctl validate -f ./policy.yaml

  # Validate a policy based on the JSON passed into stdin.
  cat policy.json | calicoctl validate -f -

Options:
  -h --help                     Show this screen.
  -f --filename=<FILENAME>      Filename to use to create the resource. If set to
                                "-" loads from stdin.


Description:
  Validate config files. Both YAML and JSON formats are accepted.
`
	parsedArgs, err := docopt.Parse(doc, args, true, "", false, false)
	if err != nil {
		fmt.Printf("Invalid option: 'calicoctl %s'. Use flag '--help' to read about a specific subcommand.\n", strings.Join(args, " "))
		os.Exit(1)
	}
	if len(parsedArgs) == 0 {
		return
	}

	filename := argutils.ArgStringOrBlank(parsedArgs, "--filename")

	// Load the resource from file and convert to a slice
	// of resources for easier handling.
	r, err := resourcemgr.CreateResourcesFromFile(filename)
	if err != nil {
		fmt.Printf("Failed to validate: %v\n", err)
		os.Exit(1)
	}

	if _, err := convertToSliceOfResources(r); err != nil {
		fmt.Printf("Failed to validate: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Validation successful")
}

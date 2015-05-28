// Copyright 2015 CoreOS, Inc.
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
	"fmt"
	"os"

	"github.com/coreos/mantle/cli"
	"github.com/coreos/mantle/kola"
)

const (
	cliName        = "kolet"
	cliDescription = "Native code runner for kola"
)

// main test harness
var cmdRun = &cli.Command{
	Name:    "run",
	Summary: "Run native tests a group at a time",
	Run:     Run,
}

func init() {
	cli.Register(cmdRun)
}

func main() {
	cli.Run(cliName, cliDescription)
}

// test runner
func Run(args []string) int {
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "FAIL: Extra arguements specified. Usage: 'kolet run <test name> <func name>'\n")
		return 2
	}
	testname, funcname := args[0], args[1]

	// find test with matching name
	test, ok := kola.Tests[testname]
	if !ok {
		fmt.Fprintf(os.Stderr, "FAIL: test group not found\n")
		return 1
	}
	// find native function in test
	f, ok := test.NativeFuncs[funcname]
	if !ok {
		fmt.Fprintf(os.Stderr, "FAIL: native function not found\n")
		return 1
	}
	err := f()
	if err != nil {
		fmt.Fprintf(os.Stderr, "FAIL: on native test %v: %v", funcname, err)
		return 1
	}
	return 0
}
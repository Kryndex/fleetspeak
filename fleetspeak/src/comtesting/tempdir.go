// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package comtesting

import (
	"fmt"
	"io/ioutil"
	"os"

	"log"
)

var tempDir string

// GetTempDir creates and returns the name of a temporary directory. Multiple
// calls will return the same name. It uses the name of the test that is being
// run as part of the name the directory.
func GetTempDir(testName string) string {
	if tempDir != "" {
		return tempDir
	}

	tempDir = os.Getenv("TEST_TMPDIR")

	if tempDir == "" {
		d, err := ioutil.TempDir("", testName+"_")
		if err != nil {
			panic(fmt.Sprintf("Unable to create temp directory: %v", err))
		}
		tempDir = d
	}
	log.Printf("Created temp directory: %s", tempDir)
	return tempDir
}

/*
 * Copyright The Spinnaker Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"fmt"
	"os"

	"github.com/spinnaker/kleat/internal/fileio"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "arguments must be <halPath> <dir>\n")
		os.Exit(1)
	}

	hal := os.Args[1]
	dir := os.Args[2]
	if err := writeServiceConfigs(hal, dir); err != nil {
		panic(err)
	}
}

func writeServiceConfigs(halPath string, dir string) error {
	h, err := fileio.ParseHalConfig(halPath)
	if err != nil {
		return err
	}

	return fileio.WriteConfigs(h, dir)
}

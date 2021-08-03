//
// Copyright 2021 Three Guys Labs, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"perkeep.org/pkg/jsonsign"
)

type Options struct {
	SecretPath string
	ConfigPath string
}

func getOptions() *Options {
	opts := &Options{}
	flag.StringVar(&opts.SecretPath, "secret", "secring.gpg", "Secret key ring file to generate")
	flag.StringVar(&opts.ConfigPath, "config", "secring-config.json", "Configuration information about key")
	flag.Parse()
	return opts
}

func generateConfig(id, path string) ([]byte, error) {
	cfg := make(map[string]string)
	cfg["identity"] = id
	cfg["identitySecretRing"] = path
	return json.MarshalIndent(cfg, "", "  ")
}

func main() {
	opts := getOptions()
	if id, err := jsonsign.GenerateNewSecRing(opts.SecretPath); err != nil {
		panic(err)
	} else if data, err := generateConfig(id, opts.SecretPath); err != nil {
		panic(err)
	} else if err := ioutil.WriteFile(opts.ConfigPath, data, 0644); err != nil {
		panic(err)
	} else {
		fmt.Printf("Generated key to %s\n", opts.SecretPath)
		fmt.Printf("Generated config to %s\n", opts.ConfigPath)
		fmt.Printf("  Key Identifier: %s\n", id)
	}
}

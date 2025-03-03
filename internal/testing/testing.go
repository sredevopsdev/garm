// Copyright 2022 Cloudbase Solutions SRL
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

//go:build testing
// +build testing

package testing

import (
	"garm/config"
	"os"
	"path/filepath"
	"testing"
)

var (
	encryptionPassphrase = "bocyasicgatEtenOubwonIbsudNutDom"
)

func GetTestSqliteDBConfig(t *testing.T) config.Database {
	dir, err := os.MkdirTemp("", "garm-config-test")
	if err != nil {
		t.Fatalf("failed to create temporary directory: %s", err)
	}
	t.Cleanup(func() { os.RemoveAll(dir) })

	return config.Database{
		Debug:      false,
		DbBackend:  config.SQLiteBackend,
		Passphrase: encryptionPassphrase,
		SQLite: config.SQLite{
			DBFile: filepath.Join(dir, "garm.db"),
		},
	}
}

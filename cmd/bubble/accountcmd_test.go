// Copyright 2016 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/cespare/cp"
)

// These tests are 'smoke tests' for the account related
// subcommands and flags.
//
// For most tests, the test files from package accounts
// are copied into a temporary keystore directory.

func tmpDatadirWithKeystore(t *testing.T) string {
	datadir := tmpdir(t)
	keystore := filepath.Join(datadir, "keystore")
	source := filepath.Join("..", "..", "accounts", "keystore", "testdata", "keystore")
	if err := cp.CopyAll(keystore, source); err != nil {
		t.Fatal(err)
	}
	return datadir
}

func TestAccountListEmpty(t *testing.T) {
	bubble := runBubble(t, "account", "list")
	bubble.ExpectExit()
}

func TestAccountList(t *testing.T) {
	datadir := tmpDatadirWithKeystore(t)
	bubble := runBubble(t, "account", "list", "--datadir", datadir)
	defer bubble.ExpectExit()
	if runtime.GOOS == "windows" {
		bubble.Expect(`
Account #0: {lat10m66vy6lrlt2qfvnamwgd8rdg8vnfthcd74p32} keystore://{{.Datadir}}\keystore\UTC--2016-03-22T12-57-55.920751759Z--7ef5a6135f1fd6a02593eedc869c6d41d934aef8
Account #1: {lat173ngt84dryedws7kyt9hflq93zpwsey2m0wqp6} keystore://{{.Datadir}}\keystore\aaa
Account #2: {lat19zw5shvhw9c5en536vun6ajwzvgeq7kvh7rqmg} keystore://{{.Datadir}}\keystore\zzz
`)
	} else {
		bubble.Expect(`
Account #0: {lat10m66vy6lrlt2qfvnamwgd8rdg8vnfthcd74p32} keystore://{{.Datadir}}/keystore/UTC--2016-03-22T12-57-55.920751759Z--7ef5a6135f1fd6a02593eedc869c6d41d934aef8
Account #1: {lat173ngt84dryedws7kyt9hflq93zpwsey2m0wqp6} keystore://{{.Datadir}}/keystore/aaa
Account #2: {lat19zw5shvhw9c5en536vun6ajwzvgeq7kvh7rqmg} keystore://{{.Datadir}}/keystore/zzz
`)
	}
}

func TestAccountNew(t *testing.T) {
	bubble := runBubble(t, "account", "new", "--lightkdf")
	defer bubble.ExpectExit()
	bubble.Expect(`
Your new account is locked with a password. Please give a password. Do not forget this password.
!! Unsupported terminal, password will be echoed.
Password: {{.InputLine "foobar"}}
Repeat password: {{.InputLine "foobar"}}

Your new key was generated
`)

	bubble.ExpectRegexp(`
Public address of the key:   lat1[0-9a-z]{38}
Path of the secret key file: .*UTC--.+--[0-9a-f]{40}

- You can share your public address with anyone. Others need it to interact with you.
- You must NEVER share the secret key with anyone! The key controls access to your funds!
- You must BACKUP your key file! Without the key, it's impossible to access account funds!
- You must REMEMBER your password! Without the password, it's impossible to decrypt the key!
`)
}

func TestAccountImport(t *testing.T) {
	tests := []struct{ key, output string }{
		{
			key:    "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
			output: "Address: {lat1ljkskxdm982xw3f36mc32gm7z6huudmux9pn7j}\n",
		},
		{
			key:    "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef1",
			output: "Fatal: Failed to load the private key: invalid character '1' at end of key file\n",
		},
	}
	for _, test := range tests {
		importAccountWithExpect(t, test.key, test.output)
	}
}

func importAccountWithExpect(t *testing.T, key string, expected string) {
	dir := tmpdir(t)
	keyfile := filepath.Join(dir, "key.prv")
	if err := ioutil.WriteFile(keyfile, []byte(key), 0600); err != nil {
		t.Error(err)
	}
	passwordFile := filepath.Join(dir, "password.txt")
	if err := ioutil.WriteFile(passwordFile, []byte("foobar"), 0600); err != nil {
		t.Error(err)
	}
	bubble := runBubble(t, "account", "import", keyfile, "-password", passwordFile)
	defer bubble.ExpectExit()
	bubble.Expect(expected)
}

func TestAccountNewBadRepeat(t *testing.T) {
	bubble := runBubble(t, "account", "new", "--lightkdf")
	defer bubble.ExpectExit()
	bubble.Expect(`
Your new account is locked with a password. Please give a password. Do not forget this password.
!! Unsupported terminal, password will be echoed.
Password: {{.InputLine "something"}}
Repeat password: {{.InputLine "something else"}}
Fatal: Passwords do not match
`)
}

func TestAccountUpdate(t *testing.T) {
	datadir := tmpDatadirWithKeystore(t)
	bubble := runBubble(t, "account", "update",
		"--datadir", datadir, "--lightkdf",
		"lat173ngt84dryedws7kyt9hflq93zpwsey2m0wqp6")
	defer bubble.ExpectExit()
	bubble.Expect(`
Unlocking account lat173ngt84dryedws7kyt9hflq93zpwsey2m0wqp6 | Attempt 1/3
!! Unsupported terminal, password will be echoed.
Password: {{.InputLine "foobar"}}
Please give a new password. Do not forget this password.
Password: {{.InputLine "foobar2"}}
Repeat password: {{.InputLine "foobar2"}}
`)
}

func TestUnlockFlag(t *testing.T) {
	datadir := tmpDatadirWithKeystore(t)
	bubble := runBubble(t,
		"--datadir", datadir, "--ipcdisable", "--testnet", "--nat", "none", "--nodiscover", "--maxpeers", "60", "--port", "0",
		"--unlock", "lat10m66vy6lrlt2qfvnamwgd8rdg8vnfthcd74p32",
		"js", "testdata/empty.js")
	bubble.Expect(`
Unlocking account lat10m66vy6lrlt2qfvnamwgd8rdg8vnfthcd74p32 | Attempt 1/3
!! Unsupported terminal, password will be echoed.
Password: {{.InputLine "foobar"}}
`)
	bubble.ExpectExit()

	wantMessages := []string{
		"Unlocked account",
		"=lat10m66vy6lrlt2qfvnamwgd8rdg8vnfthcd74p32",
	}
	for _, m := range wantMessages {
		if !strings.Contains(bubble.StderrText(), m) {
			t.Errorf("stderr text does not contain %q", m)
		}
	}
}

func TestUnlockFlagWrongPassword(t *testing.T) {
	datadir := tmpDatadirWithKeystore(t)
	bubble := runBubble(t,
		"--datadir", datadir, "--nat", "none", "--nodiscover", "--maxpeers", "60", "--port", "0", "--ipcdisable", "--testnet",
		"--unlock", "lat173ngt84dryedws7kyt9hflq93zpwsey2m0wqp6")
	defer bubble.ExpectExit()
	bubble.Expect(`
Unlocking account lat173ngt84dryedws7kyt9hflq93zpwsey2m0wqp6 | Attempt 1/3
!! Unsupported terminal, password will be echoed.
Password: {{.InputLine "wrong1"}}
Unlocking account lat173ngt84dryedws7kyt9hflq93zpwsey2m0wqp6 | Attempt 2/3
Password: {{.InputLine "wrong2"}}
Unlocking account lat173ngt84dryedws7kyt9hflq93zpwsey2m0wqp6 | Attempt 3/3
Password: {{.InputLine "wrong3"}}
Fatal: Failed to unlock account lat173ngt84dryedws7kyt9hflq93zpwsey2m0wqp6 (could not decrypt key with given password)
`)
}

// https://github.com/ethereum/go-ethereum/issues/1785
func TestUnlockFlagMultiIndex(t *testing.T) {
	datadir := tmpDatadirWithKeystore(t)
	bubble := runBubble(t,
		"--datadir", datadir, "--nat", "none", "--nodiscover", "--maxpeers", "60", "--port", "0", "--ipcdisable", "--testnet",
		"--unlock", "0,2",
		"js", "testdata/empty.js")
	bubble.Expect(`
Unlocking account 0 | Attempt 1/3
!! Unsupported terminal, password will be echoed.
Password: {{.InputLine "foobar"}}
Unlocking account 2 | Attempt 1/3
Password: {{.InputLine "foobar"}}
`)
	bubble.ExpectExit()

	wantMessages := []string{
		"Unlocked account",
		"=lat10m66vy6lrlt2qfvnamwgd8rdg8vnfthcd74p32",
		"=lat19zw5shvhw9c5en536vun6ajwzvgeq7kvh7rqmg",
	}
	for _, m := range wantMessages {
		if !strings.Contains(bubble.StderrText(), m) {
			t.Errorf("stderr text does not contain %q", m)
		}
	}
}

func TestUnlockFlagPasswordFile(t *testing.T) {
	datadir := tmpDatadirWithKeystore(t)
	bubble := runBubble(t,
		"--datadir", datadir, "--nat", "none", "--nodiscover", "--maxpeers", "60", "--port", "0",
		"--password", "testdata/passwords.txt", "--unlock", "0,2", "--ipcdisable", "--testnet",
		"js", "testdata/empty.js")
	bubble.ExpectExit()

	wantMessages := []string{
		"Unlocked account",
		"=lat10m66vy6lrlt2qfvnamwgd8rdg8vnfthcd74p32",
		"=lat19zw5shvhw9c5en536vun6ajwzvgeq7kvh7rqmg",
	}
	for _, m := range wantMessages {
		if !strings.Contains(bubble.StderrText(), m) {
			t.Errorf("stderr text does not contain %q", m)
		}
	}
}

func TestUnlockFlagPasswordFileWrongPassword(t *testing.T) {
	datadir := tmpDatadirWithKeystore(t)
	bubble := runBubble(t,
		"--datadir", datadir, "--nat", "none", "--nodiscover", "--maxpeers", "60", "--port", "0", "--ipcdisable", "--testnet",
		"--password", "testdata/wrong-passwords.txt", "--unlock", "0,2")
	defer bubble.ExpectExit()
	bubble.Expect(`
Fatal: Failed to unlock account 0 (could not decrypt key with given password)
`)
}

func TestUnlockFlagAmbiguous(t *testing.T) {
	store := filepath.Join("..", "..", "accounts", "keystore", "testdata", "dupes")
	bubble := runBubble(t,
		"--keystore", store, "--nat", "none", "--nodiscover", "--maxpeers", "60", "--port", "0", "--ipcdisable", "--testnet",
		"--unlock", "lat173ngt84dryedws7kyt9hflq93zpwsey2m0wqp6",
		"js", "testdata/empty.js")
	defer bubble.ExpectExit()

	// Helper for the expect template, returns absolute keystore path.
	bubble.SetTemplateFunc("keypath", func(file string) string {
		abs, _ := filepath.Abs(filepath.Join(store, file))
		return abs
	})
	bubble.Expect(`
Unlocking account lat173ngt84dryedws7kyt9hflq93zpwsey2m0wqp6 | Attempt 1/3
!! Unsupported terminal, password will be echoed.
Password: {{.InputLine "foobar"}}
Multiple key files exist for address lat173ngt84dryedws7kyt9hflq93zpwsey2m0wqp6:
   keystore://{{keypath "1"}}
   keystore://{{keypath "2"}}
Testing your password against all of them...
Your password unlocked keystore://{{keypath "1"}}
In order to avoid this warning, you need to remove the following duplicate key files:
   keystore://{{keypath "2"}}
`)
	bubble.ExpectExit()

	wantMessages := []string{
		"Unlocked account",
		"=lat173ngt84dryedws7kyt9hflq93zpwsey2m0wqp6",
	}
	for _, m := range wantMessages {
		if !strings.Contains(bubble.StderrText(), m) {
			t.Errorf("stderr text does not contain %q", m)
		}
	}
}

func TestUnlockFlagAmbiguousWrongPassword(t *testing.T) {
	store := filepath.Join("..", "..", "accounts", "keystore", "testdata", "dupes")
	bubble := runBubble(t,
		"--keystore", store, "--nat", "none", "--nodiscover", "--maxpeers", "60", "--port", "0", "--ipcdisable", "--testnet",
		"--unlock", "lat173ngt84dryedws7kyt9hflq93zpwsey2m0wqp6")
	defer bubble.ExpectExit()

	// Helper for the expect template, returns absolute keystore path.
	bubble.SetTemplateFunc("keypath", func(file string) string {
		abs, _ := filepath.Abs(filepath.Join(store, file))
		return abs
	})
	bubble.Expect(`
Unlocking account lat173ngt84dryedws7kyt9hflq93zpwsey2m0wqp6 | Attempt 1/3
!! Unsupported terminal, password will be echoed.
Password: {{.InputLine "wrong"}}
Multiple key files exist for address lat173ngt84dryedws7kyt9hflq93zpwsey2m0wqp6:
   keystore://{{keypath "1"}}
   keystore://{{keypath "2"}}
Testing your password against all of them...
Fatal: None of the listed files could be unlocked.
`)
	bubble.ExpectExit()
}

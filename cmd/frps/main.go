// Copyright 2018 fatedier, fatedier@gmail.com
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
	"context"
	"flag"
	"fmt"
	"github.com/fatedier/frp/pkg/config"
	"github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/server"
	"github.com/fatedier/golib/crypto"
	"math/rand"
	"os"
	"time"
)

var (
	cfgFlag = flag.String("config", "", "config file of frps")
)

func main() {
	crypto.DefaultSalt = "frp"
	// TODO: remove this when we drop support for go1.19
	rand.Seed(time.Now().UnixNano())
	// Parse and ensure all needed inputs are specified
	flag.Parse()

	if *cfgFlag == "" {
		fmt.Printf("config path is null\n")
		os.Exit(-1)
	}

	var cfg config.ServerCommonConf
	var err error
	var content []byte
	content, err = config.GetRenderedConfFromFile(*cfgFlag)
	if err != nil {
		return
	}
	cfg, err = config.UnmarshalServerConfFromIni(content)
	if err != nil {
		return
	}
	cfg.Complete()
	err = cfg.Validate()
	if err != nil {
		err = fmt.Errorf("parse config error: %v", err)
		return
	}
	if err != nil {
		return
	}
	log.Info("frps uses config file: %s", *cfgFlag)
	log.InitLog(cfg.LogWay, cfg.LogFile, cfg.LogLevel, cfg.LogMaxDays, cfg.DisableLogColor)

	svr, err := server.NewService(cfg)
	if err != nil {
		return
	}
	log.Info("frps started successfully")
	svr.Run(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

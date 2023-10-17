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
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fatedier/frp/client"
	"github.com/fatedier/frp/pkg/config"
	"github.com/fatedier/frp/pkg/util/log"
)

var (
	cfgFlag = flag.String("config", "", "config file of frpc")
)

func handleTermSignal(svr *client.Service) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	svr.GracefulClose(500 * time.Millisecond)
}

func startService(cfgFile string) (err error) {
	cfg, pxyCfgs, visitorCfgs, err := config.ParseClientConfig(cfgFile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	log.InitLog(cfg.LogWay, cfg.LogFile, cfg.LogLevel,
		cfg.LogMaxDays, cfg.DisableLogColor)

	if cfgFile != "" {
		log.Info("start frpc service for config file [%s]", cfgFile)
		defer log.Info("frpc service for config file [%s] stopped", cfgFile)
	}
	svr, errRet := client.NewService(cfg, pxyCfgs, visitorCfgs, cfgFile)
	if errRet != nil {
		err = errRet
		return
	}

	shouldGracefulClose := cfg.Protocol == "kcp" || cfg.Protocol == "quic"
	// Capture the exit signal if we use kcp or quic.
	if shouldGracefulClose {
		go handleTermSignal(svr)
	}

	_ = svr.Run(context.Background())
	return
}

func main() {
	// Parse and ensure all needed inputs are specified
	flag.Parse()

	if *cfgFlag == "" {
		fmt.Printf("config path is null\n")
		os.Exit(-1)
	}

	err := startService(*cfgFlag)
	if err != nil {
		os.Exit(1)
	}
}

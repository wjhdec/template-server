/*
Copyright © 2022 wangjinghao wjhdec@163.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wjhdec/template-server/internal/api"
	"github.com/wjhdec/template-server/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Version string
	cfg     *viper.Viper
)

func GetConfig() *viper.Viper {
	return cfg
}

func init() {
	var cfgFile string
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file path")

	cfg = viper.New()
	cfg.SetConfigType("toml")
	if cfgFile != "" {
		cfg.SetConfigFile(cfgFile)
	} else {
		cfg.AddConfigPath(".")
		cfg.AddConfigPath("./configs")
		cfg.SetConfigName("config")
	}
	cfg.AutomaticEnv()
	if err := cfg.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("can not load config file: %s, %+v", cfgFile, err))
	} else {
		fmt.Printf("Using config file: %s", cfg.ConfigFileUsed())
	}
}

type ServerOptions struct {
	Port uint
	Path string
}

var rootCmd = &cobra.Command{
	Use:   "template-server",
	Short: "template",
	Long:  `template`,
	Run: func(cmd *cobra.Command, args []string) {
		o := new(ServerOptions)
		if err := cfg.UnmarshalKey("server", o); err != nil {
			panic(err)
		}
		e := echo.New()
		writer := new(lumberjack.Logger)
		if err := cfg.UnmarshalKey("logger", writer); err != nil {
			panic(err)
		}
		e.Logger.SetOutput(writer)
		e.Logger.SetLevel(getLogLevel(cfg.GetString("logger.logLevel")))
		logger.SetLogger(e.Logger)
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		routers := []api.Router{
			newBookRouter(),
		}

		for _, r := range routers {
			r.Route(e.Group(o.Path))
		}

		e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", o.Port)))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func getLogLevel(l string) log.Lvl {
	switch strings.ToLower(l) {
	case "debug":
		return log.DEBUG
	case "info":
		return log.INFO
	case "error":
		return log.ERROR
	default:
		return log.INFO
	}
}

package main

import (
	"flag"
	"fmt"
	api "github.com/H-b-IO-T-O-H/kts-backend/application"
	"github.com/H-b-IO-T-O-H/kts-backend/application/common"
	_ "github.com/H-b-IO-T-O-H/kts-backend/docs"
	yamlConfig "github.com/rowdyroad/go-yaml-config"
	"os"
	"path"
)

var listenPort = flag.String("port", "8080", "Configure server port: --port='8080'")
var serverName = flag.String("name", "backend", "Configure server name: --name='backend'")
var needLog = flag.Bool("log", false, "Enable IO logging")

// @title 10tka API
// @version 1.0
// @description This is a backend server.

// @host 10-tka.ru
// @BasePath /api/v1
// @query.collection.format multi

func main() {
	var config api.Config

	_ = yamlConfig.LoadConfig(&config, "configs/config.yaml", nil)
	flag.Parse()
	config.Listen = ":" + *listenPort
	if config.Listen != ":8080" {
		common.Domain = fmt.Sprintf("http://localhost%s/api/v1", config.Listen)
	}

	fileDir, _ := os.Getwd()
	if fileDir != "/home/vlad/10tka/10tka-backend" {
		common.PathToSaveStatic = path.Join(fileDir, "static")
	}
	config.ServerName = *serverName
	config.NeedLog = *needLog
	app := api.NewApp(config)
	defer app.Close()
	app.Run()
}

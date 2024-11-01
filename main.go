/*
 * Copyright (c) 2024 flowerinsnow
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */
package main

import (
	"database/sql"
	"fmt"
	"github.com/flowerinsnowdh/www/config"
	"github.com/flowerinsnowdh/www/controller"
	"github.com/flowerinsnowdh/www/dao"
	"github.com/flowerinsnowdh/www/object"
	"github.com/flowerinsnowdh/www/service"
	"github.com/flowerinsnowdh/www/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pelletier/go-toml/v2"
	"net/http"
	"os"
)

var db *sql.DB

func main() {
	var file *os.File
	var err error
	if file, err = os.Open("config.toml"); err != nil {
		util.ErrPrintln("failed to open config.toml")
		panic(err)
	}

	var conf *config.Config = &config.Config{}

	if err = toml.NewDecoder(file).Decode(conf); err != nil {
		util.ErrPrintln("failed to parse config.toml")
		panic(err)
	}

	var indexPage *object.IndexPageVars = &object.IndexPageVars{
		Title:        conf.SiteConfig.Title,
		StaticDomain: conf.SiteConfig.StaticDomain,
		WWWDomain:    conf.SiteConfig.WWWDomain,
		BlogURL:      conf.SiteConfig.BlogURL,
		ICPNumber:    conf.SiteConfig.ICPNumber,
		NISMSPNumber: conf.SiteConfig.NISMSPNumber,
	}

	if db, err = sql.Open("mysql", conf.MySQLConfig.User+":"+conf.MySQLConfig.Password+"@tcp("+conf.MySQLConfig.Host+")/"+conf.MySQLConfig.Schema); err != nil {
		util.ErrPrintln("failed to open connection to mysql")
		panic(err)
	}

	var d *dao.DAO = (*dao.DAO)(db)
	var s *service.Service = (*service.Service)(d)

	if err = d.SQLInitTest(); err != nil {
		util.ErrPrintln("failed to test connection to mysql")
		panic(err)
	}

	var mux *http.ServeMux = http.NewServeMux()
	controller.Control(mux, conf, indexPage, s)

	fmt.Println("Listening on " + conf.Bind + "...")
	if err = http.ListenAndServe(conf.Bind, mux); err != nil {
		util.ErrPrintln("failed to start http server")
		panic(err)
	}
}

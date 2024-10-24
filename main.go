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
        util.StderrPrintln("failed to open config.toml")
        panic(err)
    }

    var conf *config.Config = &config.Config{}

    if err = toml.NewDecoder(file).Decode(conf); err != nil {
        util.StderrPrintln("failed to parse config.toml")
        panic(err)
    }

    var indexPage *object.IndexPageConfig = &object.IndexPageConfig{
        StaticDomain: conf.StaticDomain,
        WWWDomain:    conf.WWWDomain,
        BlogURL:      conf.BlogURL,
    }

    if db, err = sql.Open("mysql", conf.MySQL.User+":"+conf.MySQL.Password+"@tcp("+conf.MySQL.Host+")/"+conf.MySQL.Schema); err != nil {
        util.StderrPrintln("failed to open connection to mysql")
        panic(err)
    }

    var d *dao.DAO = (*dao.DAO)(db)
    var s *service.Service = (*service.Service)(d)

    if err = d.SQLInitTest(); err != nil {
        util.StderrPrintln("failed to test connection to mysql")
        panic(err)
    }

    controller.Control(indexPage, s)

    fmt.Println("Listening on "+conf.Bind+"...")
    if err = http.ListenAndServe(conf.Bind, nil); err != nil {
        util.StderrPrintln("failed to start http server")
        panic(err)
    }
}

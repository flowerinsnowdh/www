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
    "github.com/flowerinsnowdh/www/config"
    "github.com/flowerinsnowdh/www/page"
    _ "github.com/go-sql-driver/mysql"
    "github.com/pelletier/go-toml/v2"
    "html/template"
    "log"
    "net/http"
    "os"
    "strconv"
)

var db *sql.DB

func main() {
    var file *os.File
    var err error
    if file, err = os.Open("config.toml"); err != nil {
        panic(err)
    }

    var conf *config.Config = &config.Config{}

    if err = toml.NewDecoder(file).Decode(conf); err != nil {
        panic(err)
    }

    var indexPage *page.IndexPage = &page.IndexPage{
        StaticDomain: conf.StaticDomain,
        WWWDomain:    conf.WWWDomain,
        BlogURL:      conf.BlogURL,
    }

    if db, err = sql.Open("mysql", conf.MySQL.User+":"+conf.MySQL.Password+"@tcp("+conf.MySQL.Host+":"+strconv.Itoa(conf.MySQL.Port)+")/"+conf.MySQL.Schema); err != nil {
        panic(err)
    }

    if err = sqlInitTest(); err != nil {
        panic(err)
    }

    http.Handle("/", middleHandler(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            http.NotFound(w, r)
            return
        }

        var t *template.Template
        if t, err = template.ParseFiles("resources/index.html"); err != nil {
            log.Println("an error occurred while parsing template")
            log.Println(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }
        if err = t.Execute(w, indexPage); err != nil {
            log.Println("an error occurred while executing template")
            log.Println(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
    }))

    http.Handle("/contact.html", middleHandler(func(w http.ResponseWriter, r *http.Request) {
        var t *template.Template
        if t, err = template.ParseFiles("resources/contact.html"); err != nil {
            log.Println("an error occurred while parsing template")
            log.Println(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }
        if err = t.Execute(w, indexPage); err != nil {
            log.Println("an error occurred while executing template")
            log.Println(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
    }))

    if err = http.ListenAndServe(conf.Bind+":"+strconv.Itoa(conf.BindPort), nil); err != nil {
        panic(err)
    }
}

func sqlInitTest() error {
    if rows, err := db.Query("SELECT 1"); err != nil {
        return err
    } else {
        _ = rows.Close()
        return nil
    }
}

func logAccess(remoteAddr string, path string, referer string, userAgent string) error {
    var nullableRemoteAddr sql.NullString = sql.NullString{
        String: remoteAddr,
        Valid:  remoteAddr != "",
    }
    var nullableReferer sql.NullString = sql.NullString{
        String: referer,
        Valid:  referer != "",
    }
    var nullableUserAgent sql.NullString = sql.NullString{
        String: userAgent,
        Valid:  userAgent != "",
    }
    _, err := db.Exec(
        "INSERT INTO `access_log` (`remote_address`, `path`, `referer`, `user_agent`, `time`) VALUES (?, ?, ?, ?, now())",
        nullableRemoteAddr, path, nullableReferer, nullableUserAgent,
    )
    return err
}

func middleHandler(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
    // 这里可读性非常差，非常容易混淆概念
    // 这段代码是将 func(http.ResponseWriter, *http.Request) 转换成了 http.HandlerFunc 类型
    // http.HandlerFunc 是个类型！类型！不是个函数！
    // 简而言之，下面就是封装了一个 HttpHandler，在执行 next（控制层的 HttpHandler）前对长度进行了判断
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if len(r.URL.Path) > 255 {
            http.Error(w, http.StatusText(http.StatusRequestURITooLong), http.StatusRequestURITooLong)
            return
        }
        if len(r.Header.Get("Referer")) > 255 {
            http.Error(w, "'Referer' too long", http.StatusRequestURITooLong)
            return
        }
        if len(r.Header.Get("User-Agent")) > 255 {
            http.Error(w, "'User-Agent' too long", http.StatusRequestURITooLong)
            return
        }
        if err := logAccess(r.Header.Get("X-Real-IP"), r.URL.Path, r.Header.Get("Referer"), r.Header.Get("User-Agent")); err != nil {
            log.Println("an error occurred while logging access log")
            log.Println(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }
        http.HandlerFunc(next).ServeHTTP(w, r)
    })
}

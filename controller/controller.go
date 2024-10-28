package controller

import (
    "github.com/flowerinsnowdh/www/config"
    "github.com/flowerinsnowdh/www/object"
    "github.com/flowerinsnowdh/www/service"
    "github.com/flowerinsnowdh/www/util"
    "net/http"
)

func Control(mux *http.ServeMux, conf *config.Config, indexPageConfig *object.IndexPageConfig, s *service.Service) {
    mux.Handle("/", MiddleHandler(s, conf.SiteConfig.WWWDomain, indexPageConfig, func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            if err := util.ExecuteTemplate(w, indexPageConfig, "resources/error_page/404.html"); err != nil {
                util.StderrPrintln("an error occurred while executing template")
                util.StderrPrintln(err)
                http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            } else {
                w.Header().Set("Content-Type", "text/html; charset=utf-8")
                w.WriteHeader(http.StatusNotFound)
            }
            return
        }

        if err := util.ExecuteTemplate(w, indexPageConfig, "resources/index.html"); err != nil {
            util.StderrPrintln("an error occurred while executing template")
            util.StderrPrintln(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "text/html; charset=utf-8")
    }))

    mux.Handle("/contact", MiddleHandler(s, conf.SiteConfig.WWWDomain, indexPageConfig, func(w http.ResponseWriter, r *http.Request) {
        if err := util.ExecuteTemplate(w, indexPageConfig, "resources/contact.html"); err != nil {
            util.StderrPrintln("an error occurred while executing template")
            util.StderrPrintln(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "text/html; charset=utf-8")
    }))
}

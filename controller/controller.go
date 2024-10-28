package controller

import (
    "github.com/flowerinsnowdh/www/object"
    "github.com/flowerinsnowdh/www/service"
    "github.com/flowerinsnowdh/www/util"
    "net/http"
)

func Control(mux *http.ServeMux, indexPageConfig *object.IndexPageConfig, s *service.Service) {
    mux.Handle("/", MiddleHandler(s, indexPageConfig, func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            http.NotFound(w, r)
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

    mux.Handle("/contact", MiddleHandler(s, indexPageConfig, func(w http.ResponseWriter, r *http.Request) {
        if err := util.ExecuteTemplate(w, indexPageConfig, "resources/contact.html"); err != nil {
            util.StderrPrintln("an error occurred while executing template")
            util.StderrPrintln(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "text/html; charset=utf-8")
    }))
}

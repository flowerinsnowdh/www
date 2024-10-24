package controller

import (
    "github.com/flowerinsnowdh/www/object"
    "github.com/flowerinsnowdh/www/service"
    "github.com/flowerinsnowdh/www/util"
    "html/template"
    "net/http"
)

func Control(indexPage *object.IndexPageConfig, s *service.Service) {
    http.Handle("/", MiddleHandler(s, func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
            http.NotFound(w, r)
            return
        }
        
        var t *template.Template
        var err error
        
        if t, err = template.ParseFiles("resources/index.html"); err != nil {
            util.StderrPrintln("an error occurred while parsing template")
            util.StderrPrintln(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }
        
        if err = t.Execute(w, indexPage); err != nil {
            util.StderrPrintln("an error occurred while executing template")
            util.StderrPrintln(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
    }))

    http.Handle("/contact", MiddleHandler(s, func(w http.ResponseWriter, r *http.Request) {
        var t *template.Template
        var err error
        
        if t, err = template.ParseFiles("resources/contact.html"); err != nil {
            util.StderrPrintln("an error occurred while parsing template")
            util.StderrPrintln(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }
        if err = t.Execute(w, indexPage); err != nil {
            util.StderrPrintln("an error occurred while executing template")
            util.StderrPrintln(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
    }))
}

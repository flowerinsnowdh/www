package controller

import (
	"github.com/flowerinsnowdh/www/config"
	"github.com/flowerinsnowdh/www/object"
	"github.com/flowerinsnowdh/www/service"
	"github.com/flowerinsnowdh/www/util"
	"net/http"
	"strings"
)

func Control(mux *http.ServeMux, conf *config.Config, indexPageVars *object.IndexPageVars, s *service.Service) {
	mux.Handle("/", MiddleHandler(s, conf.SiteConfig.WWWDomain, indexPageVars, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			util.NotFound(w, indexPageVars)
			return
		}
		util.Template(w, "resources/index.html", indexPageVars, http.StatusOK, "text/html; charset=utf-8", nil, nil)
	}))

	mux.Handle("/contact", MiddleHandler(s, conf.SiteConfig.WWWDomain, indexPageVars, func(w http.ResponseWriter, r *http.Request) {
		util.Template(w, "resources/contact.html", indexPageVars, http.StatusOK, "text/html; charset=utf-8", nil, nil)
	}))

	mux.Handle("/redirect", XSSMiddleHandler(s, conf.SiteConfig.WWWDomain, indexPageVars, func(w http.ResponseWriter, r *http.Request) {
		var to string
		if to = strings.Trim(r.URL.Query().Get("to"), " "); to == "" {
			util.BadRequest(
				w, &object.IndexPageVariables{
					IndexPageVars: *indexPageVars,
					HasError:      true,
					ErrorMessages: []string{"缺少字段 to"},
				},
			)
			return
		}

		util.Template(
			w, "resources/redirect.html",
			&object.IndexPageVariables{
				IndexPageVars: *indexPageVars,
				TargetURL:     to,
			},
			http.StatusOK, "text/html; charset=utf-8",
			nil, nil,
		)
	}, "to"))
}

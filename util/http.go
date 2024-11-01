package util

import (
	"github.com/flowerinsnowdh/www/object"
	"net/http"
)

func Template(w http.ResponseWriter, file string, data any, status int, contentType string, success func(), fail func()) {
	if err := ExecuteTemplate(w, data, file); err != nil { // failed
		ErrPrintln("an error occurred while executing template")
		ErrPrintln(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		if fail != nil {
			fail()
		}
	} else { // success
		w.Header().Set("Content-Type", contentType)
		w.WriteHeader(status)
		if success != nil {
			success()
		}
	}
}

func httpError(w http.ResponseWriter, file string, data any, status int) {
	Template(w, file, data, status, "text/html; charset=utf-8", nil, nil)
}

func BadRequest(w http.ResponseWriter, IndexPageVariables *object.IndexPageVariables) {
	httpError(w, "resources/error_page/400.html", IndexPageVariables, http.StatusBadRequest)
}

func RequestURITooLong(w http.ResponseWriter, indexPageVars *object.IndexPageVars, tooLong string) {
	httpError(w, "resources/error_page/414.html", &object.IndexPageVariables{
		IndexPageVars: *indexPageVars,
		HasError:      true,
		ErrorMessages: []string{"过长的 " + tooLong},
	}, http.StatusRequestURITooLong)
}

func InternalServerError(w http.ResponseWriter, indexPageVars *object.IndexPageVars) {
	httpError(w, "resources/error_page/500.html", indexPageVars, http.StatusInternalServerError)
}

func NotFound(w http.ResponseWriter, indexPageVars *object.IndexPageVars) {
	httpError(w, "resources/error_page/404.html", indexPageVars, http.StatusNotFound)
}

func TeaPot(w http.ResponseWriter, indexPageVars *object.IndexPageVars) {
	httpError(w, "resources/error_page/418.html", indexPageVars, http.StatusTeapot)
}

func Forbidden(w http.ResponseWriter, indexPageVars *object.IndexPageVars) {
	httpError(w, "resources/error_page/403.html", indexPageVars, http.StatusForbidden)
}

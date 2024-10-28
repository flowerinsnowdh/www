package controller

import (
    "github.com/flowerinsnowdh/www/object"
    "github.com/flowerinsnowdh/www/service"
    "github.com/flowerinsnowdh/www/util"
    "net/http"
    "strings"
)

func MiddleHandler(s *service.Service, host string, pageConfig *object.IndexPageConfig, next func(w http.ResponseWriter, r *http.Request)) http.Handler {
    // 这里可读性非常差，非常容易混淆概念
    // 这段代码是将 func(http.ResponseWriter, *http.Request) 转换成了 http.HandlerFunc 类型
    // http.HandlerFunc 是个类型！类型！不是个函数！就是以类型的方式封装成了一个函数，并实现了 Handler 接口
    // 简而言之，下面就是封装了一个 HttpHandler，在执行 next（控制层的 HttpHandler）前对长度进行了判断
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 对比 Host，不一致返回 400
        if !strings.EqualFold(r.Host, host) {
            if err := util.ExecuteTemplate(w, pageConfig, "resources/error_page/400.html"); err != nil {
                http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
                util.StderrPrintln("an error occurred while executing template")
                util.StderrPrintln(err)
                return
            } else {
                w.Header().Set("Content-Type", "text/html; charset=utf-8")
                w.WriteHeader(http.StatusBadRequest)
            }
            return
        }

        // 检查各参数长度，超出长度限制返回 414
        var tooLong string

        if len(r.URL.Host) > 255 {
            tooLong = "Host"
        }
        if len(r.URL.Path) > 255 {
            tooLong = "Path"
        }
        if len(r.Header.Get("Referer")) > 255 {
            tooLong = "Referer"
        }
        if len(r.Header.Get("User-Agent")) > 255 {
            tooLong = "User-Agent"
        }
        if len(r.Header.Get("X-Real-IP")) > 255 {
            tooLong = "X-Real-IP"
        }
        if tooLong != "" {
            if err := util.ExecuteTemplate(w, &object.IndexPageVariables{
                IndexPageConfig: *pageConfig,
                Status414:       tooLong,
            }, "resources/error_page/414.html"); err != nil {
                http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
                util.StderrPrintln("an error occurred while executing template")
                util.StderrPrintln(err)
            } else {
                w.Header().Set("Content-Type", "text/html; charset=utf-8")
                w.WriteHeader(http.StatusRequestURITooLong)
            }
            return
        }

        // 记录访问请求日志到 MySQL
        if err := s.LogAccess(r.Header.Get("X-Real-IP"), r.Host, r.URL.Path, r.Header.Get("Referer"), r.Header.Get("User-Agent")); err != nil {
            if err := util.ExecuteTemplate(w, &object.IndexPageVariables{
                IndexPageConfig: *pageConfig,
                Status414:       tooLong,
            }, "resources/error_page/500.html"); err != nil {
                http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
                util.StderrPrintln("an error occurred while executing template")
                util.StderrPrintln(err)
            } else {
                w.Header().Set("Content-Type", "text/html; charset=utf-8")
                w.WriteHeader(http.StatusRequestURITooLong)
            }
            return
        }
        http.HandlerFunc(next).ServeHTTP(w, r)
    })
}

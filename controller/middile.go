package controller

import (
    "github.com/flowerinsnowdh/www/service"
    "github.com/flowerinsnowdh/www/util"
    "net/http"
)

func MiddleHandler(s *service.Service, next func(w http.ResponseWriter, r *http.Request)) http.Handler {
    // 这里可读性非常差，非常容易混淆概念
    // 这段代码是将 func(http.ResponseWriter, *http.Request) 转换成了 http.HandlerFunc 类型
    // http.HandlerFunc 是个类型！类型！不是个函数！就是以类型的方式封装成了一个函数，并实现了 Handler 接口
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
        if err := s.LogAccess(r.Header.Get("X-Real-IP"), r.URL.Path, r.Header.Get("Referer"), r.Header.Get("User-Agent")); err != nil {
            util.StderrPrintln("an error occurred while logging access log")
            util.StderrPrintln(err)
            http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            return
        }
        http.HandlerFunc(next).ServeHTTP(w, r)
    })
}

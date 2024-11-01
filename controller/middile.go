package controller

import (
	"github.com/flowerinsnowdh/www/object"
	"github.com/flowerinsnowdh/www/service"
	"github.com/flowerinsnowdh/www/util"
	"net/http"
	"strings"
)

func MiddleHandler(s *service.Service, host string, indexPageVars *object.IndexPageVars, next func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	// 这里可读性非常差，非常容易混淆概念
	// 这段代码是将 func(http.ResponseWriter, *http.Request) 转换成了 http.HandlerFunc 类型
	// http.HandlerFunc 是个类型！类型！不是个函数！就是以类型的方式封装成了一个函数，并实现了 Handler 接口
	// 简而言之，下面就是封装了一个 HttpHandler，在执行 next（控制层的 HttpHandler）前对长度进行了判断
	return func(w http.ResponseWriter, r *http.Request) {
		// 对比 Host，不一致返回 400
		if !strings.EqualFold(r.Host, host) {
			util.BadRequest(w, &object.IndexPageVariables{
				IndexPageVars: *indexPageVars,
			})
			return
		}

		var (
			blocked bool
			err     error
		)

		if blocked, err = s.IsBlacklistAddress(r.Header.Get("X-Real-IP")); err != nil {
			util.ErrPrintln("failed to query blacklist")
			util.ErrPrintln(err)
			util.InternalServerError(w, indexPageVars)
			return
		}

		var firstBlocked bool

		if !blocked {
			if strings.HasSuffix(r.URL.Path, ".php") {
				if err = s.AddToBlacklist(r.Header.Get("X-Real-IP")); err != nil {
					util.ErrPrintln("failed to add blacklist")
					util.ErrPrintln(err)
					util.InternalServerError(w, indexPageVars)
					return
				}
				util.TeaPot(w, indexPageVars)
				blocked = true
				firstBlocked = true
			}
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
		if len(r.URL.RawQuery) > 65535 {
			tooLong = "URL 参数"
		}
		if tooLong != "" {
			util.RequestURITooLong(w, indexPageVars, tooLong)
			return
		}

		// 记录访问请求日志到 MySQL
		if err := s.LogAccess(
			r.Header.Get("X-Real-IP"),
			r.Method,
			r.Host,
			r.URL.Path,
			r.Header.Get("Referer"),
			r.Header.Get("User-Agent"),
			r.URL.RawQuery,
			blocked,
		); err != nil {
			util.ErrPrintln("failed to log access")
			util.ErrPrintln(err)
			util.InternalServerError(w, indexPageVars)
			return
		}
		if blocked && !firstBlocked {
			util.Forbidden(w, indexPageVars)
			return
		}
		http.HandlerFunc(next).ServeHTTP(w, r)
	}
}

func XSSMiddleHandler(s *service.Service, host string, indexPageVars *object.IndexPageVars, next func(w http.ResponseWriter, r *http.Request), checkParams ...string) http.HandlerFunc {
	return MiddleHandler(s, host, indexPageVars, func(w http.ResponseWriter, r *http.Request) {
		for _, param := range checkParams {
			if strings.ContainsAny(r.URL.Query().Get(param), "<>") {
				util.BadRequest(w, &object.IndexPageVariables{
					IndexPageVars: *indexPageVars,
				})
				return
			}
		}
		http.HandlerFunc(next).ServeHTTP(w, r)
	})
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"io"
	"net/url"
	"time"
)

// 提取Gin框架的日志中间件
func GinLoggerWithConfig(out io.Writer, notlogged ...string) echo.MiddlewareFunc {
	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			var (
				cerr   error
				values url.Values
				bytes  []byte
			)
			// Start timer
			start := time.Now()
			path := c.Request().URL.Path
			raw := c.Request().URL.RawQuery
			// Process request
			if err = next(c); err != nil {
				c.Error(err)
			}

			// Log only when path is not being skipped
			if _, ok := skip[path]; !ok {
				// Stop timer
				end := time.Now()
				latency := end.Sub(start)

				clientIP := c.RealIP()
				method := c.Request().Method
				statusCode := c.Response().Status
				var statusColor, methodColor, resetColor string

				comment := "" //c.Errors.ByType(gin.ErrorTypePrivate).String()

				if raw != "" {
					path = path + "?" + raw
				}
				cerr = c.Request().ParseForm()
				// 获取请求参数
				if cerr != nil {
					log.Infof("解析参数异常:%s \n", cerr.Error())
					return cerr
				}
				contentType := c.Request().Header.Get(echo.HeaderContentType)
				switch contentType {
				case echo.MIMEApplicationJSON, echo.MIMEApplicationJSONCharsetUTF8:

				case echo.MIMEApplicationForm:
					values = c.Request().Form
					break
				default:
					values = c.Request().Form
				}

				bytes, cerr = json.Marshal(values)
				if cerr != nil {
					log.Printf("转换参数异常:%s \n", cerr.Error())
					return cerr
				}

				fmt.Fprintf(out, "[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %s\n%s[param] %s \n\n",
					end.Format("2006/01/02 - 15:04:05"),
					statusColor, statusCode, resetColor,
					latency,
					clientIP,
					methodColor, method, resetColor,
					path,
					comment,
					string(bytes),
				)
			}
			return
		}
	}
}

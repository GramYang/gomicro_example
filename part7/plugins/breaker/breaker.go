package breaker

import (
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	statusCode "gomicro_example/part7/plugins/breaker/http"
	"net/http"
)

//BreakerWrapper hystrix breaker
func BreakerWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.Method + "-" + r.RequestURI
		_ = hystrix.Do(name, func() error {
			sct := &statusCode.StatusCodeTracker{ResponseWriter: w, Status: http.StatusOK}
			h.ServeHTTP(sct.WrappedResponseWriter(), r)

			if sct.Status >= http.StatusInternalServerError {
				str := fmt.Sprintf("status code %d", sct.Status)
				return errors.New(str)
			}
			return nil
		}, func(e error) error {
			if e == hystrix.ErrCircuitOpen {
				w.WriteHeader(http.StatusAccepted)
				_, _ = w.Write([]byte("请稍后重试"))
			}

			return e
		})
	})
}

package middleware

import (
	"context"
	"errors"
	"github.com/Mykola-Mateichuk/golearn/internal/loghelper"
	"github.com/Mykola-Mateichuk/golearn/internal/token"
	"net/http"
	"time"
)

// MiddlewareLogAllErrors log errors.
func MiddlewareLogAllErrors(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		// @todo log errors here.
	})
}

// MiddlewareLogAllCalls log calls.
func MiddlewareLogAllCalls(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		uri := r.RequestURI
		method := r.Method
		next.ServeHTTP(w, r)

		duration := time.Since(start)

		// Log request details.
		standardLogger := loghelper.NewLogger()
		standardLogger.LogCalls(uri, method, duration)
	})
}

// MiddlewareLogPanicsAndRecover log panics.
func MiddlewareLogPanicsAndRecover(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)

				// Log request details.
				standardLogger := loghelper.NewLogger()
				standardLogger.LogPanic(err)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// ErrorHandler handle error.
func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	uri := r.RequestURI
	method := r.Method

	standardLogger := loghelper.NewLogger()
	standardLogger.LogError(uri, method, err)
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if (r.URL.Path != "/v1/chat/ws.rtm.start") {
			next.ServeHTTP(w, r)
			return
		}

		keys := r.URL.Query()
		accessToken := keys.Get("token")
		standardLogger := loghelper.NewLogger()

		if len(accessToken) == 0 {
			standardLogger.LogPanic(errors.New("authorization token is not provided"))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		tokenMaker, err := token.NewPasetoMaker("01234567890123456789012345678912") // @todo move to config
		if err != nil {
			standardLogger.LogPanic(errors.New("can't create token"))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			standardLogger.LogPanic(errors.New("authorization token is not provided"))
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "UserName", payload.Username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
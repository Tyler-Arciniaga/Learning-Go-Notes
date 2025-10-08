package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"net/http"
	"time"
)

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
)

type LogKey int

const (
	_ LogKey = iota
	key
)

func main() {
	q2()
}

func MiddlewareFactory(m int) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx, cancel := context.WithTimeout(ctx, time.Duration(m)*time.Microsecond)
			defer cancel()
			r = r.WithContext(ctx)
			h.ServeHTTP(w, r)
		})
	}
}

func q2() {
	ctx, cancelFuncParent := context.WithTimeout(context.Background(), time.Duration(2)*time.Second)
	defer cancelFuncParent()
	ctx, cancel := context.WithCancelCause(ctx)
	defer cancel(nil)

	func() {
		i := 0
		sum := 0
		for {
			if err := context.Cause(ctx); err != nil {
				fmt.Println("done diddly!")
				fmt.Println("Number of iterations:", i)
				fmt.Println("Final sum:", sum)
				if err := ctx.Err(); err == context.Canceled {
					fmt.Println("Reason: Found 1234")
				} else {
					fmt.Println("Reason: Time limit exceeded")
				}
				return
			}

			newNum := rand.IntN(100_000_000)
			fmt.Println(newNum)
			if newNum == 1_234 {
				cancel(context.Canceled)
			}
			i += 1
			sum += newNum
		}
	}()
}

func ContextWithLevel(ctx context.Context, level Level) context.Context {
	return context.WithValue(ctx, key, level)
}

func LevelFromContext(ctx context.Context) (Level, bool) {
	l, ok := ctx.Value(key).(Level)
	return l, ok
}

func LogLevelMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		level := r.URL.Query().Get("log_level")
		ctx := r.Context()
		ctx = ContextWithLevel(ctx, Level(level))
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}

func Log(ctx context.Context, level Level, message string) {
	var inLevel Level
	inLevel, ok := LevelFromContext(ctx)
	if !ok {
		fmt.Println("could not extract level from context")
		return
	}
	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}
}

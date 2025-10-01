package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func main(){
	q1()
}

type CurrentTimeHandler struct{}

type JSONTime struct{
	DayOfWeek string `json:"day_of_week"`
	DayOfMonth int `json:"day_of_month"`
	Month string `json:"month"`
	Year int `json:"year"`
	Hour int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
}

func (c CurrentTimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	if r.Header["Accept"][0] == "JSON"{
		t := time.Now()
		j := JSONTime{
			DayOfWeek: t.Weekday().String(), 
			DayOfMonth: t.Day(), 
			Month: t.Month().String(), 
			Year: t.Year(), 
			Hour: t.Hour(), 
			Minute: t.Minute(), 
			Second: t.Second(),
		}
		json.NewEncoder(w).Encode(j)		
	} else {
		w.Write([]byte(time.Now().Format(time.RFC3339)))

	}	
}

func q1(){
	s := http.Server{Addr: ":9090", Handler: IPLoggerMiddleware(CurrentTimeHandler{})}
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed{
			panic(err)
		}
	}
}

func IPLoggerMiddleware(h http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("start of middleware checks...")
		h.ServeHTTP(w, r)
		slog.Info("request detected", "ip", r.RemoteAddr)
		fmt.Println("clean up by middleware...")
	})
}
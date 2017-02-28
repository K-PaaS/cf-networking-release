package handlers

import (
	"fmt"
	"net/http"
	"time"
)

//go:generate counterfeiter -o fakes/metrics_sender.go --fake-name MetricsSender . metricsSender
type metricsSender interface {
	SendDuration(string, time.Duration)
	IncrementCounter(string)
}

type MetricWrapper struct {
	Name          string
	MetricsSender metricsSender
}

func (mw *MetricWrapper) Wrap(handle http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		handle.ServeHTTP(w, req)
		mw.MetricsSender.SendDuration(fmt.Sprintf("%sRequestTime", mw.Name), time.Now().Sub(startTime))
		mw.MetricsSender.IncrementCounter(fmt.Sprintf("%sRequestCount", mw.Name))
	})
}

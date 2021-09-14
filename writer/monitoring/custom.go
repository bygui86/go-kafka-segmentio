package monitoring

import "github.com/prometheus/client_golang/prometheus"

var (
	opsCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ops_total",
			Help: "Total operations",
		},
		[]string{"app"},
	)
	succPublishedMsgCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "success_published_messages_total",
			Help: "Total SUCCESSFULLY published messages",
		},
		[]string{"app", "topic"},
	)
	failPublishedMsgCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "fail_published_messages_total",
			Help: "Total FAILED published messages",
		},
		[]string{"app", "topic"},
	)
	failTracingInjCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "failed_tracing_injection_total",
			Help: "Total FAILED tracing span injection into messages",
		},
		[]string{"app", "topic"},
	)
)

func RegisterCustomMetrics() {
	prometheus.MustRegister(
		opsCounter,
		succPublishedMsgCounter,
		failPublishedMsgCounter,
		failTracingInjCounter,
	)
}

func IncreaseOpsCounter(app string) {
	go opsCounter.WithLabelValues(app).Inc()
}

func IncreaseSuccPublishedMsgCounter(app, topic string) {
	go succPublishedMsgCounter.WithLabelValues(app, topic).Inc()
}

func IncreaseFailPublishedMsgCounter(app, topic string) {
	go failPublishedMsgCounter.WithLabelValues(app, topic).Inc()
}

func IncreaseFailTracingInjCounter(app, topic string) {
	go failTracingInjCounter.WithLabelValues(app, topic).Inc()
}

package mqtt

import (
	"github.com/prometheus/client_golang/prometheus"
	"go-sample/utils/go-util/log"
	"time"
)

var filedKeys = []string{`topic`}
var filedKeysErrors = []string{`topic`, `error`}

var publishedCount *prometheus.CounterVec
var publisherErrorCount *prometheus.CounterVec
var publisherLatency *prometheus.SummaryVec

var receivedCount *prometheus.CounterVec
var endToEndLatency *prometheus.SummaryVec

func initMetrics(namespace string, subsystem string) {

	publishedCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      `mqtt_message_produced_count`,
		Help:      `Number of mqtt messages produced.`,
	}, filedKeys)

	publisherErrorCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      `mqtt_messages_produce_errors_count`,
		Help:      `Number of messages produce errors count.`,
	}, filedKeysErrors)

	publisherLatency = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      `mqtt_message_produced_latency_milliseconds`,
		Help:      `Messages produced to broker latency in milliseconds`,
	}, filedKeys)

	receivedCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      `mqtt_message_count`,
		Help:      `Number of messages received.`,
	}, filedKeys)

	endToEndLatency = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      `mqtt_message_end_to_end_latency_milliseconds`,
		Help:      `Latency between message produced and received.`,
	}, filedKeys)
}

func Register(namespace string, subsystem string) {

	initMetrics(namespace, subsystem)

	prometheus.Register(publishedCount)
	prometheus.Register(publisherErrorCount)
	prometheus.Register(publisherLatency)
	prometheus.Register(receivedCount)
	prometheus.Register(endToEndLatency)

	log.Info(`Mqtt Metrics registered`)
}

func CountProduced(topic string) {
	lvs := prometheus.Labels{`topic`: ``}
	publishedCount.With(lvs).Add(1)
}

func CountProducerErrors(err error, topic string) {
	lvs := prometheus.Labels{`topic`: ``, `error`: err.Error()}
	publisherErrorCount.With(lvs).Add(1)
}

func MeasureProducerLatency(begin time.Time, topic string) {
	lvs := prometheus.Labels{`topic`: ``}
	publisherLatency.With(lvs).Observe(float64(time.Since(begin).Nanoseconds() / 1000000))
}

func CountConsumed(topic string) {
	lvs := prometheus.Labels{`topic`: ``}
	receivedCount.With(lvs).Add(1)
}

func MeasureEndToEndLatency(begin time.Time, topic string) {
	lvs := prometheus.Labels{`topic`: ``}
	endToEndLatency.With(lvs).Observe(float64(time.Since(begin).Nanoseconds() / 1000000))
}

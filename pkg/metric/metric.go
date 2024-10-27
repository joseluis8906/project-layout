package metric

import (
	"errors"

	"github.com/joseluis8906/project-layout/pkg/metric/metricopt"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	counters   map[string]*prometheus.CounterVec
	gauges     map[string]*prometheus.GaugeVec
	histograms map[string]*prometheus.HistogramVec
	summaries  map[string]*prometheus.SummaryVec
)

var (
	ErrAlreadyExists   = errors.New("metric already exists")
	ErrUnsopportedType = errors.New("unsupported metric type")
)

func Tag(key string, val any) metricopt.Tag {
	return metricopt.Tag{Key: key, Val: val}
}

type Type int

const (
	Counter   Type = 1
	Gauge     Type = 2
	Histogram Type = 3
	Summary   Type = 4
)

func Register(kind Type, name, help string, labels ...string) error {
	switch kind {
	case Counter:
		return registerCounter(name, help, labels...)
	case Gauge:
		return registerGauge(name, help, labels...)
	case Histogram:
		return registerHistogram(name, help, labels...)
	case Summary:
		return registterSummary(name, help, labels...)
	default:
		return ErrUnsopportedType
	}
}

func registerCounter(name, help string, labels ...string) error {
	if _, ok := counters[name]; ok {
		return ErrAlreadyExists
	}

	m := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name,
			Help: help,
		},
		labels,
	)
	prometheus.MustRegister(m)
	counters[name] = m

	return nil
}

func registerGauge(name, help string, labels ...string) error {
	if _, ok := gauges[name]; ok {
		return ErrAlreadyExists
	}

	gauges[name] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: name,
			Help: help,
		},
		labels,
	)

	return nil
}

func registerHistogram(name, help string, labels ...string) error {
	if _, ok := histograms[name]; ok {
		return ErrAlreadyExists
	}

	histograms[name] = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: name,
			Help: help,
		},
		labels,
	)

	return nil
}

func registterSummary(name, help string, labels ...string) error {
	if _, ok := summaries[name]; ok {
		return ErrAlreadyExists
	}

	summaries[name] = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: name,
			Help: help,
		},
		labels,
	)

	return nil
}

func Add(kind Type, name string, tags ...metricopt.Tag) error {
	switch kind {
	case Counter:
	case Gauge:
	case Histogram:
	case Summary:
	}

	return nil
}

func Sub(kind Type, name string, tags ...metricopt.Tag) error {
	switch kind {
	case Counter:
	case Gauge:
	case Histogram:
	case Summary:
	}

	return nil
}

func Set(kind Type, name string, tags ...metricopt.Tag) error {
	switch kind {
	case Counter:
	case Gauge:
	case Histogram:
	case Summary:
	}

	return nil
}

func Inc(kind Type, name string, tags ...metricopt.Tag) error {
	switch kind {
	case Counter:
	case Gauge:
	case Histogram:
	case Summary:
	}

	return nil
}

func Dec(kind Type, name string, tags ...metricopt.Tag) error {
	switch kind {
	case Counter:
	case Gauge:
	case Histogram:
	case Summary:
	}

	return nil
}

func Reset(kind Type, name string, tags ...metricopt.Tag) error {
	switch kind {
	case Counter:
	case Gauge:
	case Histogram:
	case Summary:
	}

	return nil
}

func Observe(kind Type, name string, tags ...metricopt.Tag) error {
	switch kind {
	case Counter:
	case Gauge:
	case Histogram:
	case Summary:
	}

	return nil
}

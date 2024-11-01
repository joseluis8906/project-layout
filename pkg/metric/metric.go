package metric

import (
	"errors"
	"fmt"
	stdlog "log"

	"github.com/joseluis8906/project-layout/pkg/log"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/fx"
)

type Collector struct {
	log          *stdlog.Logger
	promRegister func(prometheus.Collector) error
	counters     map[string]*prometheus.CounterVec
	gauges       map[string]*prometheus.GaugeVec
	histograms   map[string]*prometheus.HistogramVec
	summaries    map[string]*prometheus.SummaryVec
}

var (
	ErrAlreadyExists   = errors.New("metric already exists")
	ErrUnsopportedType = errors.New("unsupported metric type")
)

const (
	EndpointResult = "endpoint_result"
)

var defCollector *Collector

func Default() *Collector {
	return defCollector
}

func init() {
	defCollector = &Collector{
		log:          stdlog.Default(),
		promRegister: prometheus.DefaultRegisterer.Register,
		counters:     make(map[string]*prometheus.CounterVec),
		gauges:       make(map[string]*prometheus.GaugeVec),
		histograms:   make(map[string]*prometheus.HistogramVec),
		summaries:    make(map[string]*prometheus.SummaryVec),
	}
}

type Deps struct {
	fx.In
	Log *stdlog.Logger
}

func New(deps Deps) *Collector {
	c := &Collector{
		log:          deps.Log,
		promRegister: prometheus.Register,
		counters:     make(map[string]*prometheus.CounterVec),
		gauges:       make(map[string]*prometheus.GaugeVec),
		histograms:   make(map[string]*prometheus.HistogramVec),
		summaries:    make(map[string]*prometheus.SummaryVec),
	}

	err := c.Register(
		Counter,
		EndpointResult,
		"How many endpoint requests are success or failure",
		[]string{ServiceTagKey, MethodTagKey, ResultTagKey},
	)
	if err != nil {
		panic(err)
	}

	return c
}

func Noop() *Collector {
	return &Collector{
		log:          log.Noop(),
		promRegister: func(c prometheus.Collector) error { return nil },
		counters:     make(map[string]*prometheus.CounterVec),
		gauges:       make(map[string]*prometheus.GaugeVec),
		histograms:   make(map[string]*prometheus.HistogramVec),
		summaries:    make(map[string]*prometheus.SummaryVec),
	}
}

type tag struct {
	Key string
	Val string
}

func Tag(key string, val string) tag {
	return tag{Key: key, Val: val}
}

type Type int

const (
	Counter   Type = 1
	Gauge     Type = 2
	Histogram Type = 3
	Summary   Type = 4
)

func (c *Collector) Register(kind Type, name, help string, labels []string) error {
	switch kind {
	case Counter:
		return c.registerCounter(name, help, labels)
	case Gauge:
		return c.registerGauge(name, help, labels)
	case Histogram:
		return c.registerHistogram(name, help, labels)
	case Summary:
		return c.registterSummary(name, help, labels)
	default:
		return ErrUnsopportedType
	}
}

func (c *Collector) registerCounter(name, help string, labels []string) error {
	if _, ok := c.counters[name]; ok {
		return ErrAlreadyExists
	}

	m := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name,
			Help: help,
		},
		labels,
	)
	if err := c.promRegister(m); err != nil {
		return fmt.Errorf("registering %q counter metric: %w", name, err)
	}

	c.counters[name] = m

	return nil
}

func (c *Collector) registerGauge(name, help string, labels []string) error {
	if _, ok := c.gauges[name]; ok {
		return ErrAlreadyExists
	}

	m := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: name,
			Help: help,
		},
		labels,
	)

	if err := c.promRegister(m); err != nil {
		return fmt.Errorf("registering %q gauge metric: %w", name, err)
	}

	c.gauges[name] = m

	return nil
}

func (c *Collector) registerHistogram(name, help string, labels []string) error {
	if _, ok := c.histograms[name]; ok {
		return ErrAlreadyExists
	}

	m := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: name,
			Help: help,
		},
		labels,
	)

	if err := c.promRegister(m); err != nil {
		return fmt.Errorf("registering %q histogram metric: %w", name, err)
	}

	c.histograms[name] = m

	return nil
}

func (c *Collector) registterSummary(name, help string, labels []string) error {
	if _, ok := c.summaries[name]; ok {
		return ErrAlreadyExists
	}

	m := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: name,
			Help: help,
		},
		labels,
	)

	if err := c.promRegister(m); err != nil {
		return fmt.Errorf("registering %q summary metric: %w", name, err)
	}

	c.summaries[name] = m

	return nil
}

func (c *Collector) Add(kind Type, name string, val float64, tags ...tag) {
	switch kind {
	case Counter:
		c.counterAdd(name, val, tags...)
	case Gauge:
	case Histogram:
	case Summary:
	}
}

func (c *Collector) Sub(kind Type, name string, tags ...tag) {
	switch kind {
	case Counter:
	case Gauge:
	case Histogram:
	case Summary:
	}
}

func (c *Collector) Set(kind Type, name string, tags ...tag) {
	switch kind {
	case Counter:
	case Gauge:
	case Histogram:
	case Summary:
	}
}

func (c *Collector) Inc(kind Type, name string, tags ...tag) {
	switch kind {
	case Counter:
		c.counterInc(name, tags...)
	case Gauge:
	case Histogram:
	case Summary:
	}
}

func (c *Collector) Dec(kind Type, name string, tags ...tag) {
	switch kind {
	case Counter:
	case Gauge:
	case Histogram:
	case Summary:
	}
}

func (c *Collector) Reset(kind Type, name string, tags ...tag) {
	switch kind {
	case Counter:
	case Gauge:
	case Histogram:
	case Summary:
	}
}

func (c *Collector) Observe(kind Type, name string, tags ...tag) {
	switch kind {
	case Counter:
	case Gauge:
	case Histogram:
	case Summary:
	}
}

func (c *Collector) counterAdd(name string, val float64, tags ...tag) {
	mv, ok := c.counters[name]
	if !ok {
		c.log.Printf("counter %q doesn't exists", name)
	}

	labels := make(prometheus.Labels)
	for _, t := range tags {
		labels[t.Key] = t.Val
	}

	m, err := mv.GetMetricWith(labels)
	if err != nil {
		c.log.Printf("getting counter with labels %s: %v", labels, err)
	}

	m.Add(val)
}

func (c *Collector) counterInc(name string, tags ...tag) {
	mv, ok := c.counters[name]
	if !ok {
		c.log.Printf("counter %q doesn't exists", name)
	}

	labels := make(prometheus.Labels)
	for _, t := range tags {
		labels[t.Key] = t.Val
	}

	m, err := mv.GetMetricWith(labels)
	if err != nil {
		c.log.Printf("getting counter with labels %s: %v", labels, err)
	}

	m.Inc()
}

const (
	ServiceTagKey = "service"
	MethodTagKey  = "method"
	ResultTagKey  = "result"

	SuccessTagVal = "success"
	FailureTagVal = "failure"
)

func (c *Collector) OpsResult(err error, tags ...tag) {
	if err != nil {
		tags = append(tags, Tag(ResultTagKey, FailureTagVal))
		c.Inc(Counter, EndpointResult, tags...)
	} else {
		tags = append(tags, Tag(ResultTagKey, SuccessTagVal))
		c.Inc(Counter, EndpointResult, tags...)
	}
}

func Register(kind Type, name, help string, labels []string) error {
	return defCollector.Register(kind, name, help, labels)
}

func Add(kind Type, name string, val float64, tags ...tag) {
	defCollector.Add(kind, name, val, tags...)
}

func Sub(kind Type, name string, tags ...tag) {
	defCollector.Sub(kind, name, tags...)
}

func Set(kind Type, name string, tags ...tag) {
	defCollector.Set(kind, name, tags...)
}

func Inc(kind Type, name string, tags ...tag) {
	defCollector.Inc(kind, name, tags...)
}

func Dec(kind Type, name string, tags ...tag) {
	defCollector.Dec(kind, name, tags...)
}

func Reset(kind Type, name string, tags ...tag) {
	defCollector.Reset(kind, name, tags...)
}

func Observe(kind Type, name string, tags ...tag) {
	defCollector.Observe(kind, name, tags...)
}

func OpsResult(err error, tags ...tag) {
	defCollector.OpsResult(err, tags...)
}

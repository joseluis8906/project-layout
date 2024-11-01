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

const (
	endpointResult = "endpoint_result"
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
		endpointResult,
		"How many endpoint requests are success or failure",
		[]string{ServiceTagKey, MethodTagKey, ResultTagKey},
		nil,
		nil,
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

func (c *Collector) Register(kind Type, name, help string, labels []string, buckets []float64, objectives map[float64]float64) error {
	switch kind {
	case Counter:
		return c.registerCounter(name, help, labels)
	case Gauge:
		return c.registerGauge(name, help, labels)
	case Histogram:
		return c.registerHistogram(name, help, labels, buckets)
	case Summary:
		return c.registerSummary(name, help, labels, objectives)
	default:
		return errors.New("unsupported metric type")
	}
}

func (c *Collector) registerCounter(name, help string, labels []string) error {
	if _, ok := c.counters[name]; ok {
		return nil
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
		return nil
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

func (c *Collector) registerHistogram(name, help string, labels []string, buckets []float64) error {
	if _, ok := c.histograms[name]; ok {
		return nil
	}

	m := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    name,
			Help:    help,
			Buckets: buckets,
		},
		labels,
	)

	if err := c.promRegister(m); err != nil {
		return fmt.Errorf("registering %q histogram metric: %w", name, err)
	}

	c.histograms[name] = m

	return nil
}

func (c *Collector) registerSummary(name, help string, labels []string, objectives map[float64]float64) error {
	if _, ok := c.summaries[name]; ok {
		return nil
	}

	m := prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       name,
			Help:       help,
			Objectives: objectives,
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
		c.gaugeAdd(name, val, tags...)
	case Histogram:
		c.log.Printf("histogram %q doesn't support add operation", name)
	case Summary:
		c.log.Printf("summary %q doesn't support add operation", name)
	}
}

func (c *Collector) Sub(kind Type, name string, val float64, tags ...tag) {
	switch kind {
	case Counter:
		c.log.Printf("counter %q doesn't support sub operation", name)
	case Gauge:
		c.gaugeSub(name, val, tags...)
	case Histogram:
		c.log.Printf("histogram %q doesn't support sub operation", name)
	case Summary:
		c.log.Printf("summary %q doesn't support sub operation", name)
	}
}

func (c *Collector) Set(kind Type, name string, val float64, tags ...tag) {
	switch kind {
	case Counter:
		c.log.Printf("counter %q doesn't support set operation", name)
	case Gauge:
		c.gaugeSet(name, val, tags...)
	case Histogram:
		c.log.Printf("histogram %q doesn't support set operation", name)
	case Summary:
		c.log.Printf("summary %q doesn't support set operation", name)
	}
}

func (c *Collector) Inc(kind Type, name string, tags ...tag) {
	switch kind {
	case Counter:
		c.counterInc(name, tags...)
	case Gauge:
		c.gaugeInc(name, tags...)
	case Histogram:
		c.log.Printf("histogram %q doesn't support dec operation", name)
	case Summary:
		c.log.Printf("summary %q doesn't support dec operation", name)
	}
}

func (c *Collector) Dec(kind Type, name string, tags ...tag) {
	switch kind {
	case Counter:
		c.log.Printf("counter %q doesn't support dec operation", name)
	case Gauge:
		c.gaugeDec(name, tags...)
	case Histogram:
		c.log.Printf("histogram %q doesn't support dec operation", name)
	case Summary:
		c.log.Printf("summary %q doesn't support dec operation", name)
	}
}

func (c *Collector) Observe(kind Type, name string, val float64, tags ...tag) {
	switch kind {
	case Counter:
		c.log.Printf("counter %q doesn't support observe operation", name)
	case Gauge:
		c.log.Printf("gauge %q doesn't support observe operation", name)
	case Histogram:
		c.histogramObserve(name, val, tags...)
	case Summary:
		c.summaryObserve(name, val, tags...)
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

func (c *Collector) gaugeSet(name string, val float64, tags ...tag) {
	mv, ok := c.gauges[name]
	if !ok {
		c.log.Printf("gauge %q doesn't exists", name)
	}

	labels := make(prometheus.Labels)
	for _, t := range tags {
		labels[t.Key] = t.Val
	}

	m, err := mv.GetMetricWith(labels)
	if err != nil {
		c.log.Printf("getting gauge with labels %s: %v", labels, err)
	}

	m.Set(val)
}

func (c *Collector) gaugeInc(name string, tags ...tag) {
	mv, ok := c.gauges[name]
	if !ok {
		c.log.Printf("gauge %q doesn't exists", name)
	}

	labels := make(prometheus.Labels)
	for _, t := range tags {
		labels[t.Key] = t.Val
	}

	m, err := mv.GetMetricWith(labels)
	if err != nil {
		c.log.Printf("getting gauge with labels %s: %v", labels, err)
	}

	m.Inc()
}

func (c *Collector) gaugeDec(name string, tags ...tag) {
	mv, ok := c.gauges[name]
	if !ok {
		c.log.Printf("gauge %q doesn't exists", name)
	}

	labels := make(prometheus.Labels)
	for _, t := range tags {
		labels[t.Key] = t.Val
	}

	m, err := mv.GetMetricWith(labels)
	if err != nil {
		c.log.Printf("getting gauge with labels %s: %v", labels, err)
	}

	m.Dec()
}

func (c *Collector) gaugeAdd(name string, val float64, tags ...tag) {
	mv, ok := c.gauges[name]
	if !ok {
		c.log.Printf("gauge %q doesn't exists", name)
	}

	labels := make(prometheus.Labels)
	for _, t := range tags {
		labels[t.Key] = t.Val
	}

	m, err := mv.GetMetricWith(labels)
	if err != nil {
		c.log.Printf("getting gauge with labels %s: %v", labels, err)
	}

	m.Add(val)
}

func (c *Collector) gaugeSub(name string, val float64, tags ...tag) {
	mv, ok := c.gauges[name]
	if !ok {
		c.log.Printf("gauge %q doesn't exists", name)
	}

	labels := make(prometheus.Labels)
	for _, t := range tags {
		labels[t.Key] = t.Val
	}

	m, err := mv.GetMetricWith(labels)
	if err != nil {
		c.log.Printf("getting gauge with labels %s: %v", labels, err)
	}

	m.Sub(val)
}

func (c *Collector) histogramObserve(name string, val float64, tags ...tag) {
	mv, ok := c.histograms[name]
	if !ok {
		c.log.Printf("histogram %q doesn't exists", name)
	}

	labels := make(prometheus.Labels)
	for _, t := range tags {
		labels[t.Key] = t.Val
	}

	m, err := mv.GetMetricWith(labels)
	if err != nil {
		c.log.Printf("getting histogram with labels %s: %v", labels, err)
	}

	m.Observe(val)
}

func (c *Collector) summaryObserve(name string, val float64, tags ...tag) {
	mv, ok := c.summaries[name]
	if !ok {
		c.log.Printf("summary %q doesn't exists", name)
	}

	labels := make(prometheus.Labels)
	for _, t := range tags {
		labels[t.Key] = t.Val
	}

	m, err := mv.GetMetricWith(labels)
	if err != nil {
		c.log.Printf("getting summary with labels %s: %v", labels, err)
	}

	m.Observe(val)
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
		c.Inc(Counter, endpointResult, tags...)
	} else {
		tags = append(tags, Tag(ResultTagKey, SuccessTagVal))
		c.Inc(Counter, endpointResult, tags...)
	}
}

func Register(kind Type, name, help string, labels []string, buckets []float64, objetives map[float64]float64) error {
	return defCollector.Register(kind, name, help, labels, buckets, objetives)
}

func Add(kind Type, name string, val float64, tags ...tag) {
	defCollector.Add(kind, name, val, tags...)
}

func Sub(kind Type, name string, val float64, tags ...tag) {
	defCollector.Sub(kind, name, val, tags...)
}

func Set(kind Type, name string, val float64, tags ...tag) {
	defCollector.Set(kind, name, val, tags...)
}

func Inc(kind Type, name string, tags ...tag) {
	defCollector.Inc(kind, name, tags...)
}

func Dec(kind Type, name string, tags ...tag) {
	defCollector.Dec(kind, name, tags...)
}

func Observe(kind Type, name string, val float64, tags ...tag) {
	defCollector.Observe(kind, name, val, tags...)
}

func OpsResult(err error, tags ...tag) {
	defCollector.OpsResult(err, tags...)
}

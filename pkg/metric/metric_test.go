package metric_test

import (
	"fmt"
	"net/http"

	"github.com/joseluis8906/project-layout/pkg/metric"
)

const (
	HttpRequestsTotal = "http_requests_total"
	HttpCode          = "code"
	HttpMethod        = "method"
)

const (
	BlobStorageOpsQueued = "blob_storage_ops_queued"
)

func ExampleRegister() {
	err := metric.Register(
		metric.Counter,
		HttpRequestsTotal,
		"How many HTTP requests processed, partitioned by status code and HTTP method.",
		[]string{HttpCode, HttpMethod},
	)
	if err != nil {
		// do handle the error
	}

	err = metric.Register(
		metric.Gauge,
		BlobStorageOpsQueued,
		"Number of blob storage operations waiting to be processed",
		nil,
	)
	if err != nil {
		// do handle the error
	}
}

func ExampleInc() {
	metric.Inc(
		metric.Counter,
		HttpRequestsTotal,
		metric.Tag(HttpCode, fmt.Sprintf("%d", http.StatusNotFound)),
		metric.Tag(HttpMethod, http.MethodPost),
	)

	metric.Inc(metric.Gauge, BlobStorageOpsQueued)
}

func ExampleDec() {
	metric.Dec(metric.Gauge, BlobStorageOpsQueued)
}

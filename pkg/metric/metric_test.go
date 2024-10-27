package metric_test

import (
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
		HttpCode,
		HttpMethod,
	)
	if err != nil {
		// do handle the error
	}

	err = metric.Register(
		metric.Gauge,
		BlobStorageOpsQueued,
		"Number of blob storage operations waiting to be processed",
	)
	if err != nil {
		// do handle the error
	}
}

func ExampleInc() {
	err := metric.Inc(
		metric.Counter,
		HttpRequestsTotal,
		metric.Tag(HttpCode, http.StatusNotFound),
		metric.Tag(HttpMethod, http.MethodPost),
	)
	if err != nil {
		// do handle error
	}

	if err := metric.Inc(metric.Gauge, BlobStorageOpsQueued); err != nil {
		// do handle error
	}
}

func ExampleDec() {
	if err := metric.Dec(metric.Gauge, BlobStorageOpsQueued); err != nil {
		// do handle error
	}
}

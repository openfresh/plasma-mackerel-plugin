package metrics

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"strings"

	"strconv"

	"github.com/stretchr/testify/assert"
)

func TestGetMetrics(t *testing.T) {

	assert := assert.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"time":1498732824187993554,"connections":100,"connections_sse":30,"connections_grpc":70}`)
	}))
	defer ts.Close()

	port := strings.Replace(ts.URL, "http://127.0.0.1:", "", -1)
	port = strings.Replace(port, "/metrics/plasma", "", -1)

	portNo, err := strconv.Atoi(port)
	assert.NoError(err)

	cli := NewPlasmaMetricsClient(portNo)
	res, err := cli.GetMetrics()
	assert.NoError(err, "")

	connections, _ := res["connections"]
	connectionsSSE, _ := res["connections_sse"]
	connectionsGRPC, _ := res["connections_grpc"]

	assert.Equal(float64(100), connections)
	assert.Equal(float64(30), connectionsSSE)
	assert.Equal(float64(70), connectionsGRPC)
}

package metrics

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type PlasmaMetricsClient interface {
	GetMetrics() (map[string]interface{}, error)
}

type PlasmaMetricsClientImpl struct {
	metricsHost string
	metricsPort int
	httpCli     *http.Client
}

func NewPlasmaMetricsClient(metricsPort int) PlasmaMetricsClient {
	return &PlasmaMetricsClientImpl{
		metricsHost: "127.0.0.1",
		metricsPort: metricsPort,
		httpCli: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
}

func (c *PlasmaMetricsClientImpl) GetMetrics() (map[string]interface{}, error) {

	metricsURL := fmt.Sprintf("http://%s:%d/metrics/plasma", c.metricsHost, c.metricsPort)
	req, err := http.NewRequest("GET", metricsURL, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "Create request is failed. [%s]", metricsURL)
	}

	res, err := c.httpCli.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "Get response is failed. [%s]", metricsURL)
	}

	go func() {
		defer res.Body.Close()
	}()

	if res.StatusCode != http.StatusOK {
		return nil, errors.Errorf("StatusCode=%d [%s]", res.StatusCode, metricsURL)
	}

	result := make(map[string]interface{}, 0)
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, errors.Wrapf(err, "json parse error [%s]", metricsURL)
	}

	return result, nil
}

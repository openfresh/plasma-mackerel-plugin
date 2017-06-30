package main

import (
	"flag"
	"fmt"
	"strings"

	mp "github.com/mackerelio/go-mackerel-plugin-helper"
	"github.com/openfresh/plasma-mackerel-plugin/metrics"
)

type PlasmaPlugin struct {
	Prefix string
	Port   int
}

func (p PlasmaPlugin) GraphDefinition() map[string](mp.Graphs) {
	labelPrefix := strings.Title(p.Prefix)
	return map[string](mp.Graphs){
		p.Prefix: mp.Graphs{
			Label: labelPrefix,
			Unit:  "float",
			Metrics: [](mp.Metrics){
				mp.Metrics{Name: "connections", Label: "Connections All"},
				mp.Metrics{Name: "connections_sse", Label: "Connections SSE"},
				mp.Metrics{Name: "connections_grpc", Label: "Connections GRPC"},
			},
		},
	}
}

func (p PlasmaPlugin) FetchMetrics() (map[string]interface{}, error) {
	client := metrics.NewPlasmaMetricsClient(p.Port)
	return client.GetMetrics()
}

func main() {
	optPrefix := flag.String("metric-key-prefix", "plasma", "Metric key prefix")
	optTempfile := flag.String("tempfile", "", "Temp file name")
	optPort := flag.Int("port", 9999, "plasma mertics port")
	flag.Parse()

	p := PlasmaPlugin{
		Prefix: *optPrefix,
		Port:   *optPort,
	}
	helper := mp.NewMackerelPlugin(p)
	helper.Tempfile = *optTempfile
	if helper.Tempfile == "" {
		helper.Tempfile = fmt.Sprintf("/tmp/mackerel-plugin-%s", *optPrefix)
	}
	helper.Run()
}

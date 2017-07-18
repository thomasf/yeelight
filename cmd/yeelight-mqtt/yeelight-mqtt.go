// yeelmqtt server
package main

import (
	"flag"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/facebookgo/flagenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/thomasf/lg"
	_ "github.com/thomasf/lg/pkg/prometheus"
	"github.com/thomasf/yeelight/pkg/mqtt"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lg.SetSrcHighlight("yeelight/cmd", "yeelight/pkg")
	lg.CopyStandardLogTo("warning")

	var netifName, mqttConnStr, metricsAddr, clientID string
	flag.StringVar(&netifName, "if", "", "network interface for yl multicast network")
	flag.StringVar(&mqttConnStr, "mqtt", "mqtt://guest:guest@127.0.0.1", "mqtt connection string")
	flag.StringVar(&metricsAddr, "metricsAddr", ":8080", "metrics listening addr")
	flag.StringVar(&clientID, "client_id", "yeelight-mqtt", "mqtt client id")
	flagenv.Prefix = "YEEL_"
	flagenv.Parse()

	flag.Parse()
	if netifName == "" {
		lg.Errorln("network interface name (-if) not specified")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if metricsAddr != "" {
		http.Handle("/metrics", promhttp.Handler())
		go func(addr string) {
			lg.Fatal(http.ListenAndServe(addr, nil))
		}(metricsAddr)
	}

	service := mqtt.MQTTService{
		MQTTConnectionStr: mqttConnStr,
		MulticastIF:       netifName,
		ClientID:          clientID,
	}
	lg.Error(service.Start())

}

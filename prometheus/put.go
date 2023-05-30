package main

import (
    "fmt"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/push"
)

func main() {
    completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "db_backup_last_completion_timestamp_seconds",
        Help: "The timestamp of the last successful completion of a DB backup.",
    })
    completionTime.SetToCurrentTime()
    if err := push.New("http://172.20.62.152:9091", "cxm").
        Collector(completionTime).
        Grouping("db", "customers").
        Push(); err != nil {
        fmt.Println("Could not push completion time to Pushgateway:", err)
    } else {
        fmt.Println("success")
    }

    //var myRegistry prometheus.Gatherer
    //push.New("http://example.org/metrics", "my_job").Gatherer(myRegistry).Push()
}

package main

import (
	"encoding/json"
	"github.com/nolo-metrics/go-nolo"
	"log"
	"os"
)

func main() {
	args := os.Args

	if 1 == len(args) {
		log.Fatal("usage: nolo-json <meter-path>")
	}

	meter_paths := collectMeterPaths(args[1:])

	ml := nolo.MeterList{}
	for _, mp := range meter_paths {
		m, err := mp.Execute()
		if err != nil {
			log.Fatalf("plugin failed: %v", err)
		}

		ml = append(ml, m)
	}

	meter_map := ml.ToMeterMap()

	output, _ := json.MarshalIndent(meter_map, "", "  ")
	os.Stdout.Write(output)
	// json.MarshalIndent finishes output without a trailing newline
	os.Stdout.Write([]byte("\n"))
}

func collectMeterPaths(args []string) []nolo.MeterPath {
	meter_paths := []nolo.MeterPath{}
	for _, arg := range args {
		meters_from_path, _ := nolo.MeterPath(arg).Expand()
		for _, m := range meters_from_path {
			meter_paths = append(meter_paths, m)
		}
	}
	return meter_paths
}

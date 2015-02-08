package main

import (
	"encoding/json"
	"fmt"
	"github.com/nolo-metrics/go-nolo"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	args := os.Args

	if 1 == len(args) {
		log.Fatal("usage: nolo-json <meter-path>")
	}

	name := args[1]

	out, err := exec.Command(name).Output()
	if err != nil {
		log.Fatal(err)
	}
	input := fmt.Sprintf("%s", out)

	basename := filepath.Base(name)
	meter := nolo.Parse(basename, input)

	meter_map := meter.ToMap()
	output, _ := json.MarshalIndent(meter_map, "", "  ")
	os.Stdout.Write(output)
}

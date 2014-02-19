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
		log.Fatal("usage: nolo-json path-to-plugin")
	}

	name := args[1]

	out, err := exec.Command(name).Output()
	if err != nil {
		log.Fatal(err)
	}
	input := fmt.Sprintf("%s", out)

	basename := filepath.Base(name)
	plugin := nolo.Parse(basename, input)

	plugin_map := plugin.ToMap()
	output, _ := json.MarshalIndent(plugin_map, "", "  ")
	os.Stdout.Write(output)
}

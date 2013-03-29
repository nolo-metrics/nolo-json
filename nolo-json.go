package main

import (
	// "github.com/nolo-metrics/nolo-json/nolo"
	"./nolo"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	args := os.Args;
	
	if 1 == len(args) {
		log.Fatal("usage: nolo-json path-to-plugin")
	}

	name := args[1]

	out, err := exec.Command(name).Output()
	if err != nil {
		log.Fatal(err)
	}
	input := fmt.Sprintf("%s", out)

	plugin := nolo.Parse(name, input)

	plugin_map := plugin.ToMap()
	output, _ := json.MarshalIndent(plugin_map, "", "  ")
	os.Stdout.Write(output)
}

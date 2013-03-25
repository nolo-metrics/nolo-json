package main

import (
	// "github.com/nolo-metrics/nolo-json"
	".."
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	name := "./app-plugin"

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

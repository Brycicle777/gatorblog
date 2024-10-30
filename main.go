package main

import (
	"encoding/json"
	"fmt"
	"internal/config"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}

	cfg.SetUser("Bryce")
	cfg, err = config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}

	formattedConfig, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}

	fmt.Printf("%s\n", string(formattedConfig))
}

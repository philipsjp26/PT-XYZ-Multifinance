package utils

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadFromYAML(dst string, target any) error {
	f, err := os.ReadFile(dst)
	if err != nil {
		log.Fatalf("Read from configuration got err : %v", err)
		return err
	}
	return yaml.Unmarshal(f, target)
}

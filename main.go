package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"sigs.k8s.io/yaml"
)

func main() {
	var (
		deleteJson bool
	)

	flag.BoolVar(&deleteJson, "delete-json", false, "delete the ori json file")
	flag.Parse()

	files := os.Args

	if len(files) < 1 {
		log.Println("At least one args be provided,Please rerun yaml")
	}

	for k, v := range files {
		if k == 0 {
			continue
		}
		if strings.HasSuffix(v, ".json") {
			data, err := os.ReadFile(v)
			if err != nil {
				log.Fatal(err)
			}

			b, err1 := yaml.JSONToYAML(data)
			if err != nil {
				log.Fatal(err1)
			}

			name := strings.Replace(v, ".json", ".yaml", 1)
			if err := os.WriteFile(name, b, 0755); err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(b))

			if deleteJson {
				err := os.Remove(v)
				if err != nil {
					log.Fatal(err)
				}
			}

		}
	}

}

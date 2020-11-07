package main

import (
	"flag"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/racirx/yaya"
	"io/ioutil"
	"log"
)

func main() {
	path := flag.String("p", "conf", "path for config file")
	file := flag.String("f", "config.yml", "filename for config file")
	flag.Parse()

	b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", *path, *file))
	if err != nil {
		log.Fatalf("server: error loading config: %v\n", err)
	}

	conf := new(yaya.Config)
	err = yaml.Unmarshal(b, conf)
	if err != nil {
		log.Fatalf("server: error unmarshaling configg: %v\n", err)
	}

	server := new(yaya.Server)
	server.Initialize(conf)
	server.Run()
}

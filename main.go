package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"

	"github.com/MichaelS11/go-ads"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var conf config

func main() {
	b, err := ioutil.ReadFile("config.yaml")
	check(err)

	err = yaml.Unmarshal(b, &conf)
	check(err)

	err = ads.HostInit()
	check(err)

	// start the sensor read loop
	// based on the sensors in conf
	// go read()

	http.Handle("/", http.FileServer(http.Dir("static")))

	go func() {
		err = http.ListenAndServe("localhost:8080", nil)
		check(err)
	}()

	log.Println("get em wet shaggy")

	// wait indefinitely
	select {}
}

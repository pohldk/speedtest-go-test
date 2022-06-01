package main

import (
	"flag"

	"github.com/librespeed/speedtest/config"
	"github.com/librespeed/speedtest/web"
)

var (
	optConfig = flag.String("c", "", "config file to be used, defaults to settings.toml in the same directory")
)

func main() {
	flag.Parse()
	conf := config.Load(*optConfig)
	web.SetServerLocation(&conf)
	web.ListenAndServe(&conf)
}

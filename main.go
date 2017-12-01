package main

import (
	"flag"
	"github.com/oyuk/APNsGoServer/APNsGoServer"
)

func main()  {
	confPath := flag.String("c", "", "configuration file path for APNsGoServer")
	flag.Parse()
	APNsGoServer.StartAPNsGoServer(confPath)
}

package main

import (
	"flag"
	"github.com/oyuk/APNsGoServer/APNsGoServer"
)

func main()  {
	confPath := flag.String("c", "", "configuration file path for gaurun")
	flag.Parse()
	APNsGoServer.StartAPNsGoServer(confPath)
}

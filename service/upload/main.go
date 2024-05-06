package main

import (
	"Distributed-cloud-storage/config"
	"Distributed-cloud-storage/route"
)

func main() {
	// gin framework
	router := route.Router()
	router.Run(config.UploadServiceHost)
}

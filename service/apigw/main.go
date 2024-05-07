package main

import "Distributed-cloud-storage/route"

func main() {
	r := route.Router()
	r.Run(":8080")
}

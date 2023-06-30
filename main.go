package main

import (
	"github.com/DavidWhelan/iotzoo/datastore"
	"github.com/DavidWhelan/iotzoo/endpoint"
)

func main() {
	var d datastore.Interface = &datastore.Terminal{}
	var e endpoint.Interface = &endpoint.Dummy{}

	e.Start(d)
}

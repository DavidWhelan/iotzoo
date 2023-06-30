package endpoint

import (
	"time"

	"github.com/DavidWhelan/iotzoo/datastore"
)

type Dummy struct {
	datastore datastore.Interface
}

func (d *Dummy) Start(datastore datastore.Interface) {
	d.datastore = datastore
	for {
		time.Sleep(5 * time.Second)
		d.datastore.DataRecieve([]byte("Hello World"))
	}
}

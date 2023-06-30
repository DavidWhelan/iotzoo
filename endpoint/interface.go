package endpoint

import "github.com/DavidWhelan/iotzoo/datastore"

type Interface interface {
	Start(datastore.Interface)
}

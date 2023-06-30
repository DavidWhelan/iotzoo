package datastore

import (
	"bytes"
	"os"
)

type Terminal struct {
	data []byte
}

func (t *Terminal) DataRecieve(data []byte) {
	t.data = data
	b := new(bytes.Buffer)
	b.Write(data)
	b.WriteTo(os.Stdout)
}

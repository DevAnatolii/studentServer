package generator

import (
	"fmt"
	"crypto/rand"
)

const id_size = 4

func RandomId() string {
	b := make([]byte, id_size)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%X", b)
}

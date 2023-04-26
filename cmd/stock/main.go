package main

import (
	"fmt"

	"github.com/libp2p/go-libp2p"
)

func main() {

	h, err := libp2p.New()
	if err != nil {
		panic(err)
	}

	defer h.Close()

	fmt.Printf("Hello World, my hosts ID is %s\n", h.ID())
}

package main

import (
	"log"
	"os"

	"github.com/ugorji/go/codec"
)

func main() {
	var (
		jh codec.JsonHandle
		mh codec.MsgpackHandle
	)
	dec := codec.NewDecoder(os.Stdin, &jh)
	enc := codec.NewEncoder(os.Stdout, &mh)

	for {
		var data interface{}
		if err := dec.Decode(&data); err != nil {
			if data == nil {
				break
			}
			log.Fatal(err)
		}

		if err := enc.Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

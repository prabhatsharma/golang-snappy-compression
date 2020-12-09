package main

import (
	"fmt"
	"log"

	"github.com/golang/snappy"
)

func main() {

	FP := "034a0256e42493f8276246506eddb6e1ff37c18a7443939025a07avd944d8fd021"
	src := FP

	// BlobID := "51d2/595134743616/full/a862ef2998e9750b/1/AIDATTGBLF6ZCLDDX7D62/L0"
	// src := BlobID

	fmt.Println("uncompressed data: ", src)

	srcByte := []byte(src)

	encoded := snappy.Encode(nil, srcByte)
	fmt.Println("compressed data..: ", string(encoded)) // ABCF

	decoded, err := snappy.Decode(nil, encoded)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("uncompressed data: ", string(decoded))

}

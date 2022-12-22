package main

import (
	"log"

	"github.com/fthrslntgy/go-wol/wol"
)

func main() {

	macAddr := "aa:bb:cc:dd:ee:ff"
	bcastAddr := "255.255.255.255"
	err := wol.Wake(macAddr, bcastAddr, nil)
	if err != nil {
		log.Println(err.Error())
	}
}

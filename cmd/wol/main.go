package main

import (
	"log"

	"github.com/fthrslntgy/go-wol/wol"
)

func main() {

	mac := "aa:bb:cc:dd:ee:ff"
	broadcast := "255.255.255.255:9"
	err := wol.Wake(mac, broadcast)
	if err != nil {
		log.Println(err.Error())
	}
}

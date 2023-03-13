package main

import (
	"log"

	"github.com/fthrslntgy/go-wol/pkg/helpers"
	"github.com/fthrslntgy/go-wol/pkg/wol"
)

func main() {

	// wol.Wake can be called with empty broadcast port. It's default 9
	macAddr := "aa:bb:cc:dd:ee:ff"
	bcastAddr := "255.255.255.255"
	err := wol.Wake(macAddr, bcastAddr, nil)
	if err != nil {
		log.Println(err.Error())
	}

	// broadcast IP address can be calculated from client's ip address if it's known
	ipAddr := "10.0.0.2"
	mask := "24"
	bcastAddr, err = helpers.CalculateBroadcastIP(ipAddr, mask)
	if err != nil {
		log.Println(err.Error())
	} else {
		err = wol.Wake(macAddr, bcastAddr, 7)
		if err != nil {
			log.Println(err.Error())
		}
	}

	// also clients CIDR address can be parsed and broadcast address can be calculated from these
	cidr := "10.0.0.2/24"
	ipAddr, mask, err := helpers.ParseCIDR(cidr)
	if err != nil {
		log.Println(err.Error())
	} else {
		bcastAddr, err = helpers.CalculateBroadcastIP(ipAddr, mask)
		if err != nil {
			log.Println(err.Error())
		} else {
			err = wol.Wake(macAddr, bcastAddr, 7)
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}

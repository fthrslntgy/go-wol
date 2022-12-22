package wol

import (
	"fmt"
	"net"

	"github.com/fthrslntgy/go-wol/internal/connection"
	"github.com/fthrslntgy/go-wol/internal/magicpacket"
)

func Wake(macAddr string, bcastAddr string, bcastPort any) error {

	if bcastPort == nil {
		bcastPort = 9
	}

	host := fmt.Sprintf("%s:%d", bcastAddr, bcastPort)
	udpAddr, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		return err
	}

	mp, err := magicpacket.New(macAddr)
	if err != nil {
		return err
	}

	bs, err := mp.Marshal()
	if err != nil {
		return err
	}

	fmt.Printf("Broadcasting to: %s\n", host)

	err = connection.Connection(udpAddr, bs)
	if err != nil {
		return err
	}

	fmt.Printf("Magic packet sent successfully to %s\n", macAddr)
	return nil
}

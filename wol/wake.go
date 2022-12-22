package wol

import (
	"fmt"
	"net"

	"github.com/fthrslntgy/go-wol/internal/magicpacket"
)

func Wake(macAddr string, bcastAddr string, bcastPort any) error {

	mp, err := magicpacket.New(macAddr)
	if err != nil {
		return err
	}

	bs, err := mp.Marshal()
	if err != nil {
		return err
	}

	if bcastPort == nil {
		bcastPort = 9
	}

	host := fmt.Sprintf("%s:%d", bcastAddr, bcastPort)
	udpAddr, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Printf("Attempting to send a magic packet to MAC %s\n", macAddr)
	fmt.Printf("... Broadcasting to: %s\n", host)
	n, err := conn.Write(bs)
	if err == nil && n != 102 {
		err = fmt.Errorf("magic packet sent was %d bytes (expected 102 bytes sent)", n)
	}
	if err != nil {
		return err
	}

	fmt.Printf("Magic packet sent successfully to %s\n", macAddr)
	return nil
}

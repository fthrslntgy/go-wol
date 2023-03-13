package helpers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func ParseCIDR(address string) (string, string, error) {

	_, ipNet, err := net.ParseCIDR(address)
	if err != nil {
		return "", "", err
	}

	ip := ipNet.IP.String()
	maski, err := mtoi(ipNet.Mask)
	if err != nil {
		return "", "", err
	}
	mask := fmt.Sprintf("%d", maski)
	return ip, mask, nil
}

// Converts IP mask to 16 bit unsigned integer.
func mtoi(mask net.IPMask) (uint16, error) {
	var i uint16
	buf := bytes.NewReader(mask)
	err := binary.Read(buf, binary.BigEndian, &i)
	return i, err
}

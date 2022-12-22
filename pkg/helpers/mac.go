package helpers

import (
	"fmt"
	"net"
	"regexp"
)

var (
	delims = ":-"
	reMAC  = regexp.MustCompile(`^([0-9a-fA-F]{2}[` + delims + `]){5}([0-9a-fA-F]{2})$`)
)

type MACAddress [6]byte

func CheckMac(macAddr string) (net.HardwareAddr, error) {

	hwAddr, err := net.ParseMAC(macAddr)
	if err != nil {
		return nil, err
	}

	if !reMAC.MatchString(macAddr) {
		return nil, fmt.Errorf("%s is not a IEEE 802 MAC-48 address", macAddr)
	}

	return hwAddr, nil
}

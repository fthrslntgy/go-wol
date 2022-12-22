package magicpacket

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"regexp"
)

var (
	delims = ":-"
	reMAC  = regexp.MustCompile(`^([0-9a-fA-F]{2}[` + delims + `]){5}([0-9a-fA-F]{2})$`)
)

type MACAddress [6]byte
type MagicPacket struct {
	header  [6]byte
	payload [16]MACAddress
}

func New(mac string) (*MagicPacket, error) {
	var packet MagicPacket
	var macAddr MACAddress

	hwAddr, err := net.ParseMAC(mac)
	if err != nil {
		return nil, err
	}

	if !reMAC.MatchString(mac) {
		return nil, fmt.Errorf("%s is not a IEEE 802 MAC-48 address", mac)
	}

	for idx := range macAddr {
		macAddr[idx] = hwAddr[idx]
	}

	for idx := range packet.header {
		packet.header[idx] = 0xFF
	}

	for idx := range packet.payload {
		packet.payload[idx] = macAddr
	}

	return &packet, nil
}

func (mp *MagicPacket) Marshal() ([]byte, error) {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.BigEndian, mp); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

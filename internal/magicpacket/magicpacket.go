package magicpacket

import (
	"bytes"
	"encoding/binary"

	"github.com/fthrslntgy/go-wol/helpers/mac"
)

type MACAddress [6]byte

type MagicPacket struct {
	header  [6]byte
	payload [16]MACAddress
}

func New(macAddr string) (*MagicPacket, error) {
	var packet MagicPacket
	var address MACAddress

	hwAddr, err := mac.CheckMac(macAddr)
	if err != nil {
		return nil, err
	}

	for idx := range address {
		address[idx] = hwAddr[idx]
	}

	for idx := range packet.header {
		packet.header[idx] = 0xFF
	}

	for idx := range packet.payload {
		packet.payload[idx] = address
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

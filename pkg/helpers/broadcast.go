package helpers

import (
	"fmt"
	"strings"

	"github.com/seancfoley/ipaddress-go/ipaddr"
)

func CalculateBroadcastIP(ip, mask string) (string, error) {
	addr := ipaddr.NewIPAddressString(fmt.Sprintf("%s/%s", ip, mask)).GetAddress()
	bcast, err := addr.ToIPv4().ToBroadcastAddress()
	if err != nil {
		return "", err
	}
	return strings.Split(bcast.String(), "/")[0], nil
}

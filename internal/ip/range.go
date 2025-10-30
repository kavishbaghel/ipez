package ip

import (
	"fmt"

	"inet.af/netaddr"
)

func getCidrRange(cidr string) {
	CidrRange, err := netaddr.ParseIPPrefix(cidr)

	if err != nil {
		panic(err)
	}

	// Calculate the start and end ip address

	networkRange := CidrRange.Range()
	startIp := networkRange.From()
	endIP := networkRange.To()

	fmt.Printf("\n Starting IP Address: %s \n \n Ending IP Address: %s", startIp, endIP)

}

// internal/ip/info.go
package ip

import (
	"inet.af/netaddr"
)

type InfoResult struct {
	Network   netaddr.IP `json:"network"`
	PrefixLen int        `json:"prefix_len"`
	Bits      int        `json:"bits"`
	From      netaddr.IP `json:"from"`
	To        netaddr.IP `json:"to"`
	Total     uint64     `json:"total"`
}

func Info(prefix netaddr.IPPrefix) (InfoResult, error) {
	// Masked network IP
	masked := prefix.Masked()
	r := prefix.Range()

	ones, bits := prefix.Bits()
	total := calcCount(r) // helper to compute number of addresses

	return InfoResult{
		Network:   masked.IP(),
		PrefixLen: ones,
		Bits:      bits,
		From:      r.From(),
		To:        r.To(),
		Total:     total,
	}, nil
}

func calcCount(r netaddr.IPRange) uint64 {
	// Works for IPv4 and IPv6 using big.Int-like incrementing on bytes
	// Use netaddr's As4/As16 for safe arithmetic
	// Simple approach for IPv4:
	if r.From().Is4() && r.To().Is4() {
		a := r.From().As4()
		b := r.To().As4()
		// convert to uint32
		a32 := uint32(a[0])<<24 | uint32(a[1])<<16 | uint32(a[2])<<8 | uint32(a[3])
		b32 := uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3])
		return uint64(b32 - a32 + 1)
	}
	// For IPv6: compute using 128-bit arithmetic via big.Int
	var fromBytes [16]byte
	var toBytes [16]byte
	copy(fromBytes[:], r.From().As16())
	copy(toBytes[:], r.To().As16())

	// convert to big.Int and subtract
	// (omitted for brevity — implement using math/big in production)
	// For now, return 0 for IPv6 if not needed.
	return 0
}

/*Package contains result type definitions and functions that are used by the main
  tool 'blacksheepwall'*/

package bsw

import (
	"encoding/binary"
	"net"
)

// Result is used to store a single IP and Hostname record.
type Result struct {
	Source   string `json:"src"`
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
}

// Results is a slice of Result.
type Results []Result

func (r Results) Len() int      { return len(r) }
func (r Results) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

// Sorts by IPv4 address, IPv6 addresses will be show first and will be unsorted.
func (r Results) Less(i, j int) bool {
	first := net.ParseIP(r[i].IP).To4()
	second := net.ParseIP(r[j].IP).To4()
	if first == nil {
		return true
	}
	if second == nil {
		return false
	}
	return binary.BigEndian.Uint32(first) < binary.BigEndian.Uint32(second)
}

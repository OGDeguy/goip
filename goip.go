package goip

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

// IPv4CIDR is a useful data structure for holding IPv4 network data.
type IPv4CIDR struct {
	Network *net.IPNet
	Mask net.IPMask
	Broadcast string
	Hosts []string
	NetworkIP string
}
// NewCIDR is constructor function that init populates an IPv4CIDR struct.
func NewCIDR(cidr string) (*IPv4CIDR, error) {
	if len(strings.Split(cidr, "/")) < 2 {
		return nil, errors.New("address is not a CIDR")
	}
	objCIDR := new(IPv4CIDR)
	ip, network, err := net.ParseCIDR(cidr)
	objCIDR.Network = network
	objCIDR.Mask = network.Mask
	objCIDR.NetworkIP = network.IP.String()
	if err != nil {
		return nil, err
		}
	ips := []string{}
	for ip:= ip.Mask(network.Mask); network.Contains(ip); inc(ip) {
		if ip.String() == network.IP.String() {
			continue // Skip adding the network address to the array
		}
		ips = append(ips, ip.String())
	}
	broadcast := ips[len(ips)-1]
	objCIDR.Broadcast = broadcast
	ips = ips[:len(ips)-1] // Remove the broadcast address
	objCIDR.Hosts = ips
	return objCIDR, nil
}

// inc Increments an IPv4 address
func inc(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] > 0 {
			break
		}
	}
}

func main() {
	test, err := NewCIDR("166.168.0.0/24")
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Valid Hosts: ", len(test.Hosts))
		fmt.Println("Broadcast Address: ", test.Broadcast)
		fmt.Println("Network Address: ", test.NetworkIP)
	}
}
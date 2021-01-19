package goip

import (
	"fmt"
	"goip"
)

func ExampleNewCIDR () {
	cidr, _ := goip.NewCIDR("192.168.1.0/24")
	fmt.Println(len(cidr.Hosts))
	// Output: 254
}

func ExampleIsPrivate() {
	fmt.Println(goip.IsPrivate("192.168.0.1"))
	// Output: true
}

func ExampleIsCIDR() {
	fmt.Println(goip.IsCIDR("192.168.0.0/24"))
	// Output: true
}
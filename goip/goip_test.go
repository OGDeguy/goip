package goip

import "fmt"

func ExampleNewCIDR () {
	cidr, _ := NewCIDR("192.168.1.0/24")
	fmt.Println(len(cidr.Hosts))
	// Output: 254
}

func ExampleIsPrivate() {
	fmt.Println(IsPrivate("192.168.0.1"))
	// Output: true
}

func ExampleIsCIDR() {
	fmt.Println(IsCIDR("192.168.0.0/24"))
	// Output: true
}
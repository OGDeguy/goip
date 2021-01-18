package main

func ExampleNewCIDR () {
	cidr, _ := NewCIDR("192.168.1.0/24")
	fmt.Println(cidr.Hosts)
	// Output: 254
}
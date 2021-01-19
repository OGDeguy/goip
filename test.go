package main

import (
	"fmt"
	"github.com/OGDeguy/goip/goip"
)

func main() {
	test, _ := goip.NewCIDR("192.168.0.0/24")
	fmt.Println(test)
}
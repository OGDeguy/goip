[![Build Status](https://travis-ci.org/OGDeguy/goip.svg?branch=main)](https://travis-ci.org/OGDeguy/goip) 
# GoIP
Simple library for working with IPv4 ranges in GoLang. Currently only supports IPv4 CIDRs.

## Documentation

Documentation and additional examples are available at godoc:
https://pkg.go.dev/github.com/OGDeguy/goip

[![Go Reference](https://pkg.go.dev/badge/github.com/OGDeguy/goip.svg)](https://pkg.go.dev/github.com/OGDeguy/goip)

## Example

```go
package main

import (

"github.com/OGDeguy/goip/goip"
"log"
)

func main() {
    network, err := goip.NewCIDR("192.168.1.0/24")
    if err != nil {
        log.Fatalln(err)
    }
    log.Println("Network address: ", network.NetworkIP)
    log.Println("Broadcast address: ", network.Broadcast)
    log.Println("Total hosts: ", len(network.Hosts))
}
```
Which would print:
```bash
2021/01/17 21:17:24 Network address:  192.168.1.0
2021/01/17 21:17:24 Broadcast address:  192.168.1.255
2021/01/17 21:17:24 Total hosts:  254
```

Thanks to @kotakanbe for your gist which helped me put this together:
https://gist.github.com/kotakanbe/d3059af990252ba89a82
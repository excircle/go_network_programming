# Go Network Programming!

Go Network Tomfoolery: How-To's and Whaddaya-Do's

This is my repository to publish basic `net` and `net/http` code. These code samples serve as quick reminders in case I forget anything.

Below, you will find a list of programs and their respective source code.

# Code Samples

<details><summary>print_laddr.go - Print Local Address</summary>
<p>

<code>print_laddr.go</code> loops through locally available network interfaces and prints the IP Address found on the interface that has a <code>192.168</code> address.

```go
package main

import (
	"fmt"
	"net"
	"log"
)

var localNetworkCIDR string = "192.168.0.0/16"

func main() {
	//Create IP objects for comparison
	_, localNetwork, err := net.ParseCIDR(localNetworkCIDR)
	if err != nil {
		log.Fatalf("ERROR: Could not parse localNetworkCIDR: %+v\n", err)
	}

	//Obtain all local network interfaces
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Printf("ERROR: Could not obtain interfaces: %+v\n", err)
	}

	//Search for 192.168 address
	for _, v := range ifaces {
		addrs, err := v.Addrs()
		if err != nil {
			log.Printf("ERROR: Could not obtain addresses: %+v\n", err)
		}
		ifaceNetworkIP, _, err := net.ParseCIDR(addrs[0].String())
		if err != nil {
			log.Printf("ERROR: Could not parse ifaceNetworkIP: %+v\n", err)
		}
		if localNetwork.Contains(ifaceNetworkIP) {
			fmt.Printf("Interface '%v' has IP='%+v'\n", v.Name, ifaceNetworkIP)
		}
	}
}
```

</p>
</details>
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
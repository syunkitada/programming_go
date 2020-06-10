package main

import (
	"fmt"
	"sort"
)

type Network struct {
	Name string
	Ips  int
}

func main() {
	sortSlice()
}

func sortSlice() {
	networks := []Network{
		Network{Name: "net2", Ips: 1},
		Network{Name: "net4", Ips: 2},
		Network{Name: "net1", Ips: 4},
		Network{Name: "net3", Ips: 2},
	}

	sort.Slice(networks, func(i, j int) bool {
		return networks[i].Name < networks[j].Name
	})
	fmt.Println("Networks", networks)

	sort.Slice(networks, func(i, j int) bool {
		return networks[i].Ips < networks[j].Ips
	})
	fmt.Println("Networks", networks)
}

// Networks [{net1 4} {net2 1} {net3 2} {net4 2}]
// Networks [{net2 1} {net3 2} {net4 2} {net1 4}]

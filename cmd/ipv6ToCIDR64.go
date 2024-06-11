package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net"
)

var Ipv6ToCIDR64Cmd = &cobra.Command{
	Use:   "toCidr64",
	Short: "Convert a IPV6 to CIDR64",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ipv6ToCIDR64(args[0])
	},
}

func ipv6ToCIDR64(ipv6 string) {
	fmt.Println("Converting " + ipv6)
	ip := net.ParseIP(ipv6)
	if ip == nil {
		fmt.Errorf("invalid IPv6 address")
		return
	}

	if ip.To4() != nil {
		fmt.Errorf("the provided address is not an IPv6 address")
		return
	}

	cidr := fmt.Sprintf("%s/64", ipv6)
	fmt.Println(cidr)
}

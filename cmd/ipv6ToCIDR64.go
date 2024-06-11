package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
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
	ip := net.ParseIP(ipv6)
	if ip == nil {
		fmt.Errorf("The provided address is not a valid IP address")
		return
	}

	if ip.To4() != nil {
		fmt.Errorf("The provided address is not an IPv6 address")
		return
	}

	cidr := fmt.Sprintf("%s/64", ipv6)
	_, networkIp, _ := net.ParseCIDR(cidr)

	fmt.Println("IPV6", ip, "=> CIDR64", networkIp)
}

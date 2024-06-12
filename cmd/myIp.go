package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

type ipVersion uint8

const (
	ipv4 ipVersion = iota
	ipv6
)

var MyIpCmd = &cobra.Command{
	Use:   "myIp",
	Short: "Get your public IP address",
	Run: func(cmd *cobra.Command, args []string) {
		myIp()
	},
}

func myIp() {
	ipv4 := getIp(ipv4)
	ipv6 := getIp(ipv6)

	fmt.Println("IPv4:", ipv4)
	fmt.Println("IPv6:", ipv6)
}

func getIp(version ipVersion) string {
	url := "https://api.ipify.org?format=text"
	if version == ipv6 {
		url = "https://api6.ipify.org?format=text"
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return "Could not get IP"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return "Could not get IP"
	}

	return string(body)
}

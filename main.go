package main

import (
	"correa-cli/cmd"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "correa",
		Short: "Correa is a personal assistant CLI",
		Long:  `Correa is a personal assistant CLI`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Correa-CLI!")
		},
	}

	rootCmd.AddCommand(cmd.Ipv6ToCIDR64Cmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

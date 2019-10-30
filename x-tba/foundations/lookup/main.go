// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	missingHost = "Please provide at least one domain. --help for more information."

	help = `
Host to IP Lookup:
------------------

It finds the ip addresses of the given hosts. You can provide hosts by separating them with spaces.

Example:

  host google.com
  host google.com uber.com`
)

func main() {
	// url := "google.com"
	var message string

	args := os.Args
	switch l := len(args); {
	// case len(args) == 1:
	case l == 1:
		message = missingHost
	case l == 2 && args[1] == "--help":
		message = strings.TrimSpace(help)
	}

	if message != "" {
		fmt.Println(message)
		return
	}

	// for i := 0; i < len(args); i++ {}
	// for i, url := range args {
	for _, url := range args[1:] {
		// if i == 0 {
		// 	continue
		// }

		ips, err := net.LookupIP(url)
		if err != nil {
			fmt.Printf("%-20s => %s\n", url, err)
			// break
			continue
		}

		for _, ip := range ips {
			if ip = ip.To4(); ip != nil {
				fmt.Printf("%-20s => %s\n", url, ip)
			}
		}
	}
}

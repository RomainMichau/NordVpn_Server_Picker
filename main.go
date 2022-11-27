package main

import (
	"flag"
	"fmt"
	"nordvpn_srv_picker/srvpicker"
)

type Params struct {
	country string
	feature string
	detail  bool
}

func main() {
	params, err := parseParams()
	if err != nil {
		panic(err)
	}
	options := srvpicker.Options{
		Country: params.country,
		Feature: params.feature,
	}
	picker := srvpicker.Init(&options)
	selectedServers, err := picker.GetServers()

	if params.detail {
		fmt.Printf("%d server found\n", len(selectedServers))
	}
	for i, server := range selectedServers {
		if params.detail {
			fmt.Printf("%d: Server ID: %s. country: %s.  Hostname: %s\n", i, server.IPAddress, server.Country, server.Domain)
		} else {
			fmt.Printf("%s\n", server.Domain)

		}
	}

}

func parseParams() (*Params, error) {
	country := flag.String("country", "", "country containing server")
	feature := flag.String("feature", "", "feature")
	details := flag.Bool("v", false, "Will display details about servers")
	flag.Parse()

	return &Params{
		country: *country,
		feature: *feature,
		detail:  *details,
	}, nil
}

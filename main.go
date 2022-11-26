package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type params struct {
	country string
	feature string
	details bool
}

type NordVpnServers struct {
	ID             int      `json:"id"`
	IPAddress      string   `json:"ip_address"`
	SearchKeywords []string `json:"search_keywords"`
	Categories     []struct {
		Name string `json:"name"`
	} `json:"categories"`
	Name     string `json:"name"`
	Domain   string `json:"domain"`
	Price    int    `json:"price"`
	Flag     string `json:"flag"`
	Country  string `json:"country"`
	Location struct {
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
	} `json:"location"`
	Load     int             `json:"load"`
	Features map[string]bool `json:"features"`
}

func main() {
	params, err := parseParams()
	if err != nil {
		panic(err)
	}

	resp, err := http.Get("https://api.nordvpn.com/server")
	if err != nil {
		panic(fmt.Errorf("Failed to query https://api.nordvpn.com/server. %w ", err))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(fmt.Errorf("Failed to parse https://api.nordvpn.com/server response. %w ", err))
	}
	var respJson []NordVpnServers
	var selectedServer []NordVpnServers
	_ = json.Unmarshal(body, &respJson)
	for _, server := range respJson {
		if params.country != "" && strings.ToLower(server.Country) != strings.ToLower(params.country) {
			continue
		}
		if params.feature != "" && !server.Features[params.feature] {
			continue
		}
		selectedServer = append(selectedServer, server)
	}

	if params.details {
		fmt.Printf("%d server found\n", len(selectedServer))
	}
	for i, server := range selectedServer {
		if params.details {
			fmt.Printf("%d: Server ID: %s. Country: %s.  Hostname: %s\n", i, server.IPAddress, server.Country, server.Domain)
		} else {
			fmt.Printf("%s\n", server.Domain)

		}
	}

}

func parseParams() (*params, error) {
	country := flag.String("country", "", "Country containing server")
	feature := flag.String("feature", "", "feature")
	details := flag.Bool("v", false, "Will display details about servers")
	flag.Parse()

	return &params{
		country: *country,
		feature: *feature,
		details: *details,
	}, nil
}

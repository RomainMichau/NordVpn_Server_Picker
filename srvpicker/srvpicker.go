package srvpicker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Options struct {
	Country string
	Feature string
}

type NordVpnServer struct {
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

type SrvPicker struct {
	options *Options
}

func Init(option *Options) *SrvPicker {
	return &SrvPicker{
		options: option,
	}
}

func (picker SrvPicker) GetServers() ([]NordVpnServer, error) {
	resp, err := http.Get("https://api.nordvpn.com/server")
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("Failed to query https://api.nordvpn.com/server. %w ", err)
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected response code from nord: %d. Body: %s", resp.StatusCode, string(body))
	}
	var respJson []NordVpnServer
	err = json.Unmarshal(body, &respJson)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse https://api.nordvpn.com/server response. %w ", err)
	}
	var selectedServer []NordVpnServer
	for _, server := range respJson {
		if picker.options.Country != "" && strings.ToLower(server.Country) != strings.ToLower(picker.options.Country) {
			continue
		}
		if picker.options.Feature != "" && !server.Features[picker.options.Feature] {
			continue
		}
		selectedServer = append(selectedServer, server)
	}
	return selectedServer, nil
}

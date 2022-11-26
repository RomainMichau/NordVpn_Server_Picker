# nordvpn_srv_picker
Fetch a list of server matching some filters  
Use https://api.nordvpn.com/server
# filters
- Country (Optional, Case insensitive)
- feature (Optional, case sensitive)
# Example
```
go run ./main.go --country Argentina --feature proxy_ssl
ar50.nordvpn.com
ar51.nordvpn.com
ar52.nordvpn.com
ar53.nordvpn.com
ar54.nordvpn.com
```

```
go run ./main.go --country Argentina --feature proxy_ssl -v
5 server found
0: Server ID: 103.50.33.48. Country: Argentina.  Hostname: ar50.nordvpn.com
1: Server ID: 103.50.33.64. Country: Argentina.  Hostname: ar51.nordvpn.com
2: Server ID: 103.50.33.74. Country: Argentina.  Hostname: ar52.nordvpn.com
3: Server ID: 103.50.33.84. Country: Argentina.  Hostname: ar53.nordvpn.com
4: Server ID: 103.50.33.94. Country: Argentina.  Hostname: ar54.nordvpn.com
```

```
go run ./main.go --country Belgium -v  
60 server found
0: Server ID: 82.102.19.137. Country: Belgium.  Hostname: be148.nordvpn.com
1: Server ID: 77.243.191.250. Country: Belgium.  Hostname: be149.nordvpn.com
2: Server ID: 185.210.217.115. Country: Belgium.  Hostname: be150.nordvpn.com
3: Server ID: 185.210.217.120. Country: Belgium.  Hostname: be151.nordvpn.com
4: Server ID: 82.102.19.131. Country: Belgium.  Hostname: be152.nordvpn.com
5: Server ID: 82.102.19.211. Country: Belgium.  Hostname: be153.nordvpn.com
...
59: Server ID: 146.70.55.43. Country: Belgium.  Hostname: be207.nordvpn.com
```
# Feature existing in nordVpn: 
- ikev2 
- openvpn_udp 
- openvpn_tcp 
- socks 
- proxy 
- pptp 
- l2tp 
- openvpn_xor_udp 
- openvpn_xor_tcp 
- proxy_cybersec 
- proxy_ssl 
- proxy_ssl_cybersec 
- ikev2_v6 
- openvpn_udp_v6 
- openvpn_tcp_v6 
- wireguard_udp 
- openvpn_udp_tls_crypt 
- openvpn_tcp_tls_crypt 
- openvpn_dedicated_udp 
- openvpn_dedicated_tcp 
- skylark 
- mesh_relay false
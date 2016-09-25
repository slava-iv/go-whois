package main

import (
	"fmt"
	"net"
	"regexp"
)

const (
	defaultWhoisServer = "com.whois-servers.net"
)

//Run getting whois information.
func Run(domain string, servers ...string) (string, error) {
	var responce string
	var err error

	for _, server := range servers {
		responce, err = request(domain, server)
		if err == nil && responce != "" {
			break
		}
	}

	if responce == "" || len(servers) == 0 {
		server, err := getServer(domain, defaultWhoisServer)
		if err != nil {
			return "", err
		}
		responce, err = request(domain, server)
		if err != nil {
			return "", err
		}
	}

	if responce == "" {
		return "", fmt.Errorf("whois information for %s not found", domain)
	}

	return responce, nil
}

func getServer(domain, server string) (string, error) {
	responce, err := request(domain, server)
	if err != nil {
		return "", err
	}
	serverRegexp := regexp.MustCompile("(?i)whois server:\\s*([^\\s]+)")
	serverWhois := serverRegexp.FindStringSubmatch(responce)[1]

	return serverWhois, nil
}

func request(domain, server string) (string, error) {
	conn, err := net.Dial("tcp", server+":43")

	if err != nil {
		return "", err
	}

	defer conn.Close()

	conn.Write([]byte(domain + "\r\n"))

	buf := make([]byte, 1024)

	result := []byte{}

	for {
		numBytes, err := conn.Read(buf)
		sbuf := buf[0:numBytes]
		result = append(result, sbuf...)

		if err != nil {
			break
		}
	}

	return string(result), nil
}

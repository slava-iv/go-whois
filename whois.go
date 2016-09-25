package main

import "net"

//Run getting whois information.
func Run(domain, server string) (string, error) {
    return whoisRequest(domain, server)
}

func whoisRequest(domain, server string) (string, error) {
	conn, err := net.Dial("tcp", server + ":43")

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
package utils

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"time"
	"os"
)

var PATTERN = regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)

func IsFoundHost(host string, port uint16) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return false // Connection failed
	}
	conn.Close()
	return true // Connection successful
}

func FindNeighbors(myHost string, myPort uint16, startIp uint8, endIp uint8, startPort uint16, endPort uint16) []string {
	address := fmt.Sprintf("%s:%d", myHost, myPort)

	m := PATTERN.FindStringSubmatch(myHost)
	if m == nil {
		return nil
	}

	prefixHost := fmt.Sprintf("%s.%s.%s", m[1], m[2], m[3])
	lastIp, _ := strconv.Atoi(m[4])
	neighbors := make([]string, 0)

	for port := startPort; port <= endPort; port++ {
		for ip := startIp; ip <= endIp; ip++ {
			guessHost := fmt.Sprintf("%s.%d", prefixHost, lastIp+int(ip))
			guessTarget := fmt.Sprintf("%s:%d", guessHost, port)

			if guessTarget != address && IsFoundHost(guessHost, port) {
				neighbors = append(neighbors, guessTarget)
			}
		}
	}

	return neighbors
}


func GetHost() string {
	hostname, err := os.Hostname()
	if err != nil {
         return "127.0.0.1"
	}

	address, err := net.LookupHost(hostname)
	if err != nil {
		return "127.0.0.1"
   }

   return address[0]

}
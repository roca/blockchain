package utils

import (
	"fmt"
	"net"
	"os"
	"regexp"
	"time"
)

func IsFoundHost(host string, port int) bool {
	target := fmt.Sprintf("%s:%d", host, port)

	_, err := net.DialTimeout("tcp", target, 1*time.Second)
	if err != nil {
		fmt.Printf("%s %v\n", target, err)
		return false
	}
	return true
}

/*

((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3})(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)

192.168.0.10:5003
*/

var PATTERN = regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?\.){3})(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)`)

func FindNeighbors(
	myHost string, myPort int,
	startIp int, endIp int,
	startPort int, endPort int,
) []string {
	address := fmt.Sprintf("%s:%d", myHost, myPort)

	m := PATTERN.FindStringSubmatch(myHost)
	if m == nil {
		return nil
	}
	prefixHost := m[1]
	//lastIp, _ := strconv.Atoi(m[len(m)-1])

	neighbors := make([]string, 0)

	for port := startPort; port <= endPort; port += 1 {
		for ip := startIp; ip <= endIp; ip += 1 {
			guessHost := fmt.Sprintf("%s%d", prefixHost, ip)
			guessTarget := fmt.Sprintf("%s:%d", guessHost, port)
			if ip != 1 && guessTarget != address && IsFoundHost(guessHost, port) {
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
	fmt.Println("HOSTNAME:", hostname)
	address, err := net.LookupHost(hostname)
	if err != nil {
		return "127.0.0.1"
	}
	fmt.Println("ADDRESS:", address)
	return address[0]
}

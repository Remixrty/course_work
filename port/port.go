package port

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

var results []ScanResult
var resultsv1 []ScanResult

type ScanResult struct {
	Port     string `json:"port"`
	Protocol string `json:"protocol"`
	State    string `json:"state"`
	Service  string
}

var wg sync.WaitGroup
var count int

func ScanUDPPort(protocol, hostname string, port int) {
	defer wg.Done()
	result := ScanResult{Port: strconv.Itoa(port)}
	result.Protocol = protocol
	address := hostname + ":" + strconv.Itoa(port)
	_, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed"
		results = append(results, result)
		fmt.Println(result)
		count++

	} else {
		result.State = "Open"
		results = append(results, result)
		fmt.Println(result)
		count++
	}

}
func ScanTCPPort(protocol, hostname string, port int) {
	defer wg.Done()
	result := ScanResult{Port: strconv.Itoa(port)}
	result.Protocol = protocol
	address := hostname + ":" + strconv.Itoa(port)
	_, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed"
		resultsv1 = append(resultsv1, result)
		fmt.Println(result)
		count++

	} else {
		result.State = "Open"
		resultsv1 = append(resultsv1, result)
		fmt.Println(result)
		count++
	}

}
func WideScan(hostname string) []ScanResult {

	if len(results) != 0 {
		fmt.Print(count)
		return results
	}

	wg.Add(60001)
	for i := 0; i <= 60000; i++ {
		go ScanUDPPort("udp", hostname, i)
	}

	wg.Wait()
	fmt.Print(count)
	count = 0

	return results
}

func WideScan1(hostname string) []ScanResult {

	if len(resultsv1) != 0 {
		fmt.Print(count)
		return resultsv1
	}
	wg.Add(60001)

	for i := 0; i <= 60000; i++ {
		go ScanTCPPort("tcp", hostname, i)
	}

	wg.Wait()
	fmt.Print(count)
	count = 0
	return resultsv1

}

package main

import (
	"os"
	"fmt"
	"net"
	"log"

)
func getMacAddr() ([]string, error) {
    ifas, err := net.Interfaces()
    if err != nil {
        return nil, err
    }
    var as []string
    for _, ifa := range ifas {
        a := ifa.HardwareAddr.String()
        if a != "" {
            as = append(as, a)
        }
    }
    return as, nil
}

func main() {
	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hostname:", host)
	as, err := getMacAddr()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("MAC address:", as[1])
}
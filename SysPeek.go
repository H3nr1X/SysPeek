package main

import (
	"os"
	"fmt"
	"net"
	"log"
	"net/http"
    	"io/ioutil"

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
func getip() string {
    url := "http://api.ipify.org/"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
    }
    return string(html)

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
    fmt.Print("Public ip: ", getip())
}

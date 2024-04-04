package main

import (
	"fmt"
	"net"
	"time"
)

func whats_serv_redirect(count int) int {
	if config_map[2] == "round-robin" {
		count += 1
	} else {
		count += 1
	}
	if (count >= len(ip_map)){
		count = 0
	}
	return count
}

func sort_ip_map(ip_map map[int]string) map[int]string {
	for key, value := range ip_map {
		if !checkHealth(value) {
			delete(ip_map, key)
		}
	}
	return (ip_map)
}

func changeKeysIndex(inputMap map[string]string) map[int]string {
	outputMap := make(map[int]string)
	i := 0
	for _, value := range inputMap {
		outputMap[i] = value
		i++
	}
	return outputMap
}

func checkHealth(ip string) bool {
	adress := fmt.Sprintf(ip)
	conn, err := net.DialTimeout("tcp", adress, 3*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

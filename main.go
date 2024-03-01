package main

import (
	"fmt"
	"net/http"
	"os"
	"gopkg.in/yaml.v2"
)

var count int
var ip_map map[int]string
var config_map map[int]string

func take_conf(str string) map[int]string {
	yfile, err := os.ReadFile(str)
	if err != nil {
		fmt.Println("error to get yaml")
	}
	data := make(map[string]string)
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		fmt.Printf("error to get map \n")
	}
	return changeKeysIndex(data)
}

func collector(w http.ResponseWriter, r *http.Request) {
	if (count > len(ip_map)){
		count = 0
	}
	fmt.Println("IP TEST ["+ip_map[count]+"]")
	if (!checkHealth(ip_map[count])) {
		delete(ip_map,count)
	}else {
		fmt.Println("request redirect [https://"+ip_map[count]+"]" , " status :",count)
		http.Redirect(w, r, "http://"+ip_map[count],http.StatusMovedPermanently)
	}
	count = count + 1
}

func main() {
	count = 0
	ip_map = take_conf("config_ip.yaml")
	config_map = take_conf("config_lb.yaml")
	fmt.Println("Debug MAP IP :",ip_map)
    http.HandleFunc("/", collector)
	ip_map = sort_ip_map(ip_map)
	fmt.Println("Debug MAP IP AFTER SORT:",ip_map)
	if (len(ip_map) == 0){
		fmt.Println("No ip Valide ")
	} else {
    	fmt.Println("start load_balancing on port "+config_map[0])
    	err := http.ListenAndServe(":"+config_map[0], nil)
    	if err != nil {
        	fmt.Println("Fail to start the load_balancing :", err)
		}
	}
}

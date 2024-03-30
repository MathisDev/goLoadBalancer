package main

import (
	"fmt"
	"net/http"
	"os"
	"gopkg.in/yaml.v2"
	"time"
    "net/url"
    "net/http/httputil"
	"log"
)

var count int
var number_retry int
var ip_map map[int]string
var config_map map[int]string
var remote *url.URL

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

func req(w http.ResponseWriter, r *http.Request) {
	if (count >= len(ip_map)){
		count = 0
	}
	currentTime := time.Now()
	if (!checkHealth(ip_map[count])) {
		delete(ip_map,count)
	}else {
		fmt.Println(currentTime.Format("2006-01-02 15:04:05"),"redirect [https://"+ip_map[count]+"]" , " index of server :",count)
		http.Redirect(w, r, "https://"+ip_map[count],http.StatusMovedPermanently)
	}
	count = count + 1
}

func req_proxy(w http.ResponseWriter, r *http.Request) {
	if (count >= len(ip_map)){
		count = 0
	}
	log.Println("Redirect to : http://"+ip_map[count])
	remote, _ = url.Parse("http://"+ip_map[count])
	p := httputil.NewSingleHostReverseProxy(remote)
	r.Host = remote.Host
	w.Header().Set("X-Ben", "Rad")
	p.ServeHTTP(w, r)
	count += 1
}

func main() {
	count = 0
	ip_map = take_conf("config_ip.yaml")
	config_map = take_conf("config_lb.yaml")
	ip_map = sort_ip_map(ip_map)
	fmt.Println("MAP CONFIG :",config_map)
	fmt.Println("MAP IP VERIF:",ip_map)
	if (len(ip_map) == 0){
		fmt.Println("No Config IP Valide | number retry",number_retry)
		if (number_retry != 5){
			number_retry += 1
			time.Sleep(6 * time.Second)
			main()
		} else{
			os.Exit(2)
		}
	} else {
		fmt.Println("start load_balancing on port "+config_map[0])
		if (config_map[1] == "redirect"){
			http.HandleFunc("/", req)
		} else if config_map[1] == "proxy"{
			http.HandleFunc("/", req_proxy)
		} else {
			fmt.Println("config_lb.yaml not valid")
			os.Exit(1)
		}
    	err := http.ListenAndServe(":"+config_map[0], nil)
    	if err != nil {
        	fmt.Println("Fail to start the load_balancing :", err)
			os.Exit(1)
		}
	}
}

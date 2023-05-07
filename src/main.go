package main

import (
	"clash-config-plug/service"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"net/http"
)

func responseErr(w http.ResponseWriter, err error) {
	log.Printf("handle err %v", err)
	io.WriteString(w, fmt.Sprintf("500 %s", err))
}

func proxyClashConfig(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	proxyUrl := r.Form.Get("url")
	if proxyUrl == "" {
		responseErr(w, errors.New("url empty"))
		return
	}

	config, err := service.GetConfig(proxyUrl)
	if err != nil {
		responseErr(w, err)
		return
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		responseErr(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream; charset=utf-8")
	w.Header().Set("Content-Disposition", "attachment; filename=clash-config-plug.yaml")
	io.WriteString(w, string(data))
}

/**
sample: http://127.0.0.1:9876/clash_config/get_config?url=${subscribe url}
*/
func main() {
	http.HandleFunc("/clash_config/get_config", proxyClashConfig)
	err := http.ListenAndServe(":9876", nil)
	if err != nil {
		log.Fatal("listen serve fail: ", err)
	}
}

package service

import (
	"log"
	"testing"
)

func TestGetConfig(t *testing.T) {
	config, err := GetConfig("${sub url}")
	if err != nil {
		t.Errorf("get config fail %s", err)
	}
	log.Printf("%v", config)
}

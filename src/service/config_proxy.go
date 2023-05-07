package service

import (
	common2 "clash-config-plug/common"
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
)

func getConfig(subUrl string) ([]byte, error) {
	resp, err := http.Get(subUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	return data, err
}

func parseConfig(data []byte, ty string) (config common2.ClashConfig, err error) {
	switch ty {
	case common2.Yaml:
		err = yaml.Unmarshal(data, &config)
	default:
		err = errors.New("un support config type")
	}

	return config, err
}

func rulesConfig() (config common2.RulesConfig, err error) {
	var data []byte
	data, err = ioutil.ReadFile("/usr/share/clash-config-plug/rules.yaml") // TODO
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &config)
	return
}

// only modify proxy groups and rules
func convertConfig(srcConfig *common2.ClashConfig) error {
	srcConfig.ProxyGroups, srcConfig.Rules = make([]common2.ProxyGroups, 0), make([]string, 0)
	ruleConfig, err := rulesConfig()
	if err != nil {
		return err
	}
	srcConfig.Rules = append(srcConfig.Rules, ruleConfig.Rules...)

	proxyNames := make([]string, 0)
	for _, proxy := range srcConfig.Proxies {
		proxyNames = append(proxyNames, proxy.Name)
	}

	for _, group := range ruleConfig.ProxyGroups {
		group.Proxies = append(group.Proxies, proxyNames...)
		srcConfig.ProxyGroups = append(srcConfig.ProxyGroups, group)
	}
	return nil
}

func GetConfig(subUrl string) (*common2.ClashConfig, error) {
	data, err := getConfig(subUrl)
	if err != nil {
		return nil, err
	}

	config, err := parseConfig(data, common2.Yaml)
	if err != nil {
		return nil, err
	}

	if err := convertConfig(&config); err != nil {
		return &config, err
	}
	return &config, nil
}

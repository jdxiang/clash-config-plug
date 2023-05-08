package service

import (
	"clash-config-plug/common"
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
)

func getConfig(subUrl string) (config *common.ClashConfig, err error) {
	resp, err := http.Get(subUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return parseConfig(data, common.Yaml, resp.Header.Get("Subscription-Userinfo"))
}

func parseConfig(data []byte, ty string, subInfo string) (config *common.ClashConfig, err error) {
	config = &common.ClashConfig{}
	switch ty {
	case common.Yaml:
		err = yaml.Unmarshal(data, config)
	default:
		err = errors.New("un support config type")
	}

	config.SubscriptionInfo = subInfo
	return config, err
}

func rulesConfig() (config common.RulesConfig, err error) {
	var data []byte
	data, err = ioutil.ReadFile("/usr/share/clash-config-plug/rules.yaml") // TODO
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &config)
	return
}

// only modify proxy groups and rules
func convertConfig(srcConfig *common.ClashConfig) error {
	srcConfig.ProxyGroups, srcConfig.Rules = make([]common.ProxyGroups, 0), make([]string, 0)
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

func GetConfig(subUrl string) (*common.ClashConfig, error) {
	config, err := getConfig(subUrl)
	if err != nil {
		return nil, err
	}

	if err := convertConfig(config); err != nil {
		return config, err
	}
	return config, nil
}

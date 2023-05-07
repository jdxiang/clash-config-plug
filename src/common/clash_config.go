package common

type ClashConfig struct {
	Port               int           `yaml:"port"`
	SocksPort          int           `yaml:"socks-port"`
	RedirPort          int           `yaml:"redir-port"`
	AllowLan           bool          `yaml:"allow-lan"`
	Mode               string        `yaml:"mode"`
	LogLevel           string        `yaml:"log-level"`
	ExternalController string        `yaml:"external-controller"`
	Secret             string        `yaml:"secret"`
	DNS                DNS           `yaml:"dns"`
	Proxies            []Proxies     `yaml:"proxies"`
	ProxyGroups        []ProxyGroups `yaml:"proxy-groups"`
	Rules              []string      `yaml:"rules"`
}

type FallbackFilter struct {
	Geoip  bool     `yaml:"geoip"`
	Ipcidr []string `yaml:"ipcidr"`
}

type DNS struct {
	Enable         bool           `yaml:"enable"`
	Ipv6           bool           `yaml:"ipv6"`
	Listen         string         `yaml:"listen"`
	EnhancedMode   string         `yaml:"enhanced-mode"`
	FakeIPRange    string         `yaml:"fake-ip-range"`
	Nameserver     []string       `yaml:"nameserver"`
	Fallback       []string       `yaml:"fallback"`
	FallbackFilter FallbackFilter `yaml:"fallback-filter"`
}

type PluginOpts struct {
	Mode string `yaml:"mode"`
	Host string `yaml:"host"`
}

type Proxies struct {
	Name       string     `yaml:"name"`
	Type       string     `yaml:"type"`
	Server     string     `yaml:"server"`
	Port       int        `yaml:"port"`
	Password   string     `yaml:"password"`
	Sni        string     `yaml:"sni,omitempty"`
	UDP        bool       `yaml:"udp"`
	Cipher     string     `yaml:"cipher,omitempty"`
	Plugin     string     `yaml:"plugin,omitempty"`
	PluginOpts PluginOpts `yaml:"plugin-opts,omitempty"`
}

type ProxyGroups struct {
	Name    string   `yaml:"name"`
	Type    string   `yaml:"type"`
	Proxies []string `yaml:"proxies"`
}

type RulesConfig struct {
	ProxyGroups []ProxyGroups `yaml:"proxy-groups"`
	Rules       []string      `yaml:"rules"`
}

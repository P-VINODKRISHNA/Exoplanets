package config

// Logging client

var (
	Params   = make(map[string]string)
	RoutesV1 = make(map[string]string)
	Regex    = make(map[string]string)
)

type Routing struct {
	V1 []*KeyValue `json:"V1"`
}

type KeyValue struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

type General struct {
	Logging     string `json:"Logging"`
	Port        string `json:"Port"`
	Environment string `json:"Environment"`
	SSLEnable   bool   `json:"SSLEnable"`
	CORSAllowed bool   `json:"CORSAllowed"`
	CORSMaxAge  string `json:"CORSMaxAge"`
}

type Caching struct {
	RedisHost       string `json:"RedisHost"`
	RedisPassword   string `json:"RedisPassword"`
	SessionLifeTime string `json:"SessionLifeTime"`
	SelectDB        []int  `json:"SelectDB"`
}

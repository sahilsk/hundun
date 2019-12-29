package config

/**
 * @var		type	HundunConfig
 * @global
 */
type HundunConfig struct {
	Pagerduty PagerdutyConfiguration `mapstructure:"pagerduty"`
}

/**
 * @var		type	PagerdutyConfiguration
 * @global
 */
type PagerdutyConfiguration struct {
	ApiKey string `mapstructure:"api_key"`
	Url    string `mapstructure:"url"`
}

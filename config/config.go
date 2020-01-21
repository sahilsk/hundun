package config

/**
 * @var		type	HundunConfig
 * @global
 */
type HundunConfig struct {
	Pagerduty  PagerdutyConfiguration   `mapstructure:"pagerduty"`
	Alienvault AlientVaultConfiguration `mapstructure:"alienvault"`
}

/**
 * @var		type	PagerdutyConfiguration
 * @global
 */
type PagerdutyConfiguration struct {
	ApiKey string `mapstructure:"api_key"`
	Url    string `mapstructure:"url"`
	Email  string `mapstructure:"email"`
}

type AlientVaultConfiguration struct {
	ApiKey string `mapstructure:"api_key"`
	Url    string `mapstructure:"url"`
	Cookie string `mapstructure:"cookie"`
}

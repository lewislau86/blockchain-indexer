package config

import (
	"time"
)

// Default is a config instance.
var Default Config //nolint:gochecknoglobals // config must be global

type Config struct {
	LogLevel string `mapstructure:"log_level"`

	Port string `mapstructure:"port"`

	Gin struct {
		Mode string `mapstructure:"mode"`
	} `mapstructure:"gin"`

	Swagger struct {
		Hostname string `mapstructure:"hostname"`
	} `mapstructure:"swagger"`

	Sentry struct {
		DSN        string  `mapstructure:"dsn"`
		SampleRate float32 `mapstructure:"sample_rate"`
	} `mapstructure:"sentry"`

	Database struct {
		URL string `mapstructure:"url"`
		Log bool   `mapstructure:"log"`
	} `mapstructure:"database"`

	Kafka struct {
		Brokers           []string `mapstructure:"brokers"`
		BlocksTopicPrefix string   `mapstructure:"blocks_topic_prefix"`
		MaxAttempts       int      `mapstructure:"max_attempts"`
		MessageMaxBytes   int      `mapstructure:"message_max_bytes"`
	} `mapstructure:"kafka"`

	Prometheus struct {
		NameSpace string `mapstructure:"namespace"`
		SubSystem string `mapstructure:"subsystem"`

		PushGateway struct {
			URL          string        `mapstructure:"url"`
			Key          string        `mapstructure:"key"`
			PushInterval time.Duration `mapstructure:"push_interval"`
		} `mapstructure:"pushgateway"`
	} `mapstructure:"prometheus"`

	BlockProducer struct {
		FetchBlocksMax     int64         `mapstructure:"fetch_blocks_max"`
		FetchInterval      time.Duration `mapstructure:"fetch_interval"`
		BlockRetryNum      int           `mapstructure:"block_retry"`
		BlockRetryInterval time.Duration `mapstructure:"block_retry_interval"`
	} `mapstructure:"block_producer"`

	Platforms struct {
		Smartchain struct {
			Node string `mapstructure:"node"`
		}
	} `mapstructure:"platforms"`
}

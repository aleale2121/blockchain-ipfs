package model

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment       string `mapstructure:"ENVIRONMENT"`
	HTTPServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	HTTPHost          string `mapstructure:"HTTP_HOST"`
	HTTPPort          string `mapstructure:"HTTP_PORT"`
	IPFSServerAddress string `mapstructure:"IPFS_SERVER_ADDRESS"`

	FileDownloadDir string `mapstructure:"FILE_DOWNLOAD_DIR"`
}

package app

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultHomeDir = ".gintemplate"

	defaultConfigName = "gintpl-apiserver.yaml"
)

func onInitialize() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		for _, dir := range searchDirs() {
			viper.AddConfigPath(dir)
		}

		viper.SetConfigType("yaml")

		viper.SetConfigName(defaultConfigName)
	}

	setupEnvironmentVariables()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Failed to read viper configuration file, err: %v", err)
	}

	log.Printf("Using config file: %s", viper.ConfigFileUsed())
}

func setupEnvironmentVariables() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("GINTPL")
	// 替换环境变量 key 中的分隔符 '.' 和 '-' 为 '_'
	replacer := strings.NewReplacer(".", "_", "-", "_")
	viper.SetEnvKeyReplacer(replacer)
}

func searchDirs() []string {
	homeDir, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return []string{filepath.Join(homeDir, defaultHomeDir), "."}
}

func filePath() string {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return filepath.Join(home, defaultHomeDir, defaultConfigName)
}

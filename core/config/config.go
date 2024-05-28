package config

import (
    "github.com/spf13/viper"
)

type Config struct {
    ProjectPath string
    SeverityLevels []string
}

func Load() *Config {
    viper.SetDefault("ProjectPath", ".")
    viper.SetDefault("SeverityLevels", []string{"info", "warning", "major", "critical", "blocker"})
    viper.AutomaticEnv()

    return &Config{
        ProjectPath: viper.GetString("ProjectPath"),
        SeverityLevels: viper.GetStringSlice("SeverityLevels"),
    }
}

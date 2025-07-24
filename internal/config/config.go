package config

import (
    "log"
    
    "github.com/joho/godotenv"
    "github.com/spf13/viper"
)

type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    Redis    RedisConfig    `mapstructure:"redis"`
    GRPC     GRPCConfig     `mapstructure:"grpc"`
    JWT      JWTConfig      `mapstructure:"jwt"`
    Log      LogConfig      `mapstructure:"log"`
}

type ServerConfig struct {
    Port string `mapstructure:"port"`
    Host string `mapstructure:"host"`
}

type DatabaseConfig struct {
    Driver   string `mapstructure:"driver"`
    Host     string `mapstructure:"host"`
    Port     string `mapstructure:"port"`
    Username string `mapstructure:"username"`
    Password string `mapstructure:"password"`
    Database string `mapstructure:"database"`
    SSLMode  string `mapstructure:"ssl_mode"`
}

type RedisConfig struct {
    Host     string `mapstructure:"host"`
    Port     string `mapstructure:"port"`
    Password string `mapstructure:"password"`
    DB       int    `mapstructure:"db"`
}

type GRPCConfig struct {
    Port string `mapstructure:"port"`
}

type JWTConfig struct {
    Secret string `mapstructure:"secret"`
    Expire int    `mapstructure:"expire"`
}

type LogConfig struct {
    Level string `mapstructure:"level"`
}

var globalConfig *Config

func Load() *Config {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }
    
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("./configs")
    viper.AddConfigPath(".")
    
    // Set defaults
    setDefaults()
    
    // Read environment variables
    viper.AutomaticEnv()
    
    if err := viper.ReadInConfig(); err != nil {
        log.Printf("Error reading config file: %v", err)
    }
    
    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        log.Fatal("Unable to decode into struct:", err)
    }
    
    globalConfig = &config
    return &config
}

func Get() *Config {
    if globalConfig == nil {
        return Load()
    }
    return globalConfig
}

func setDefaults() {
    viper.SetDefault("server.port", "8080")
    viper.SetDefault("server.host", "localhost")
    viper.SetDefault("database.driver", "postgres")
    viper.SetDefault("database.host", "localhost")
    viper.SetDefault("database.port", "5432")
    viper.SetDefault("redis.host", "localhost")
    viper.SetDefault("redis.port", "6379")
    viper.SetDefault("redis.db", 0)
    viper.SetDefault("grpc.port", "9090")
    viper.SetDefault("jwt.expire", 24)
    viper.SetDefault("log.level", "info")
}

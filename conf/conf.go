package conf

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tiant-go/golib/pkg/env"
	"github.com/tiant-go/golib/pkg/http"
	"github.com/tiant-go/golib/pkg/middleware"
	"github.com/tiant-go/golib/pkg/orm"
	"github.com/tiant-go/golib/pkg/redis"
	"github.com/tiant-go/golib/pkg/zlog"
	"log"
	"path/filepath"
	"strings"
)

type SWebConf struct {
	Port    int            `yaml:"port"`
	AppName string         `yaml:"appName"`
	Log     zlog.LogConfig `yaml:"log"`

	Mysql map[string]orm.MysqlConf        `yaml:"mysql"`
	Redis map[string]redis.RedisConf      `yaml:"redis"`
	Api   map[string]*http.HttpClientConf `yaml:"api"` // 调用三方后台

	accessConf middleware.AccessLoggerConfig `yaml:"accessConf"`
}

var WebConf *SWebConf

func InitConf() {
	filePath := filepath.Join(env.GetConfDirPath(), "mount", "default.yaml")
	confViper := viper.New()
	confViper.SetConfigFile(filePath)
	confViper.SetEnvPrefix("TIANT")
	confViper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	err := confViper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	confViper.AutomaticEnv()

	if err := confViper.Unmarshal(&WebConf); err != nil {
		log.Fatal("Unmarshal config failed, ", err)
	}
}

func (s *SWebConf) GetZlogConf() zlog.LogConfig {
	return s.Log
}

func (s *SWebConf) GetAccessLogConf() middleware.AccessLoggerConfig {
	return s.accessConf
}

func (s *SWebConf) GetHandleRecoveryFunc() gin.RecoveryFunc {
	return nil
}

func (s *SWebConf) GetAppName() string {
	return s.AppName
}

func (s *SWebConf) GetPort() int {
	return s.Port
}

package config

import (
	"github.com/stevenroose/gonfig"
	"go.uber.org/fx"
	"golang-ddd-boilerplate/pkg/fxmysql/gorm"
)

var Provide = fx.Provide(get)

// Config 配置结构
type Config struct {
	ConfigFile string
	Domain     string     `id:"domain" default:"access" desc:"领域服务名"`
	MySQL      gorm.MySQL `id:"mysql" desc:"MySQL连接信息"`
}

var config *Config

// 配置文件相关配置
const (
	FileVariable    = "configfile"  // 配置文件变量
	EnvPrefix       = "MK_BIZ_"     // ConfigEnvPrefix 配置文件，环境变量前缀
	FileDefaultName = "config.json" // 默认配置文件名

)

// Init 初始化配置
func Init() error {
	config = &Config{}
	// 加载 config ......
	return gonfig.Load(config, gonfig.Conf{
		ConfigFileVariable:  FileVariable,
		FileDefaultFilename: FileDefaultName,
		FileDecoder:         gonfig.DecoderJSON,
		EnvPrefix:           EnvPrefix,
	})
}

// Get 获取配置
func get() Config {
	return *config
}

func Get() Config {
	return get()
}

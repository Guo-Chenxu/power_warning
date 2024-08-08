package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	RoomConfig       RoomConfig `yaml:"RoomConfig"`
	MailConfig       MailConfig `yaml:"MailConfig"`
	WarningThreshold float64    `yaml:"WarningThreshold"` // 报警阈值
}

type RoomConfig struct {
	PartmentId string `yaml:"partmentId"`
	FloorId    string `yaml:"floorId"`
	DromNumber string `yaml:"dromNumber"`
	AreaId     string `yaml:"areaId"`
	Cookie     string `yaml:"cookie"`
}

type MailConfig struct {
	From     string   `yaml:"from"`
	EnvTo    string   `yaml:"envTo"`
	To       []string `yaml:"to"`
	Secret   string   `yaml:"secret"`
	Host     string   `yaml:"host"`
	Port     int      `yaml:"port"`
	Nickname string   `yaml:"nickname"`
	Subject  string   `yaml:"subject"`
	Body     string   `yaml:"body"`
	Ssl      bool     `yaml:"ssl"`
}

var config Config
var env string = os.Getenv("ENV")

func getProjectPath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		info, err := os.Stat(filepath.Join(cwd, "main.go"))
		if err == nil && !info.IsDir() {
			return cwd, nil
		} else if !os.IsNotExist(err) {
			return "", err
		}

		cwd = filepath.Dir(cwd)
		if cwd == "/" || cwd == "" {
			return "", fmt.Errorf("无法找到项目根目录")
		}
	}
}

func getConfigPath() string {
	projectPath, err := getProjectPath()
	if err != nil {
		panic(fmt.Sprintf("Error get project path: %+v", err))
	}

	configFile := "config.yml"
	if env != "" {
		configFile = fmt.Sprintf("config-%s.yml", env)
	}

	return filepath.Join(projectPath, "conf", configFile)
}

// 从环境变量读取一些配置
func readConfigFromEnv(conf *Config) {
	conf.RoomConfig.Cookie = os.Getenv(conf.RoomConfig.Cookie)

	conf.MailConfig.From = os.Getenv(conf.MailConfig.From)
	conf.MailConfig.To = strings.Split(os.Getenv(conf.MailConfig.EnvTo), ",")
	conf.MailConfig.Secret = os.Getenv(conf.MailConfig.Secret)
}

func init() {
	dataBytes, err := os.ReadFile(getConfigPath())
	if err != nil {
		panic(fmt.Sprintf("Error reading config file: %+v", err))
	}

	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshalling config: %+v", err))
	}
	if env != "local" {
		readConfigFromEnv(&config)
	}

	fmt.Printf("config init success, config: %+v\n", config)
}

func GetConfig() Config {
	return config
}

package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	once sync.Once
	Cfg  Config
)

type Config struct {
	Env      string   `yaml:"env" env-default:"local"`
	DBconfig DBconfig `yaml:"dbconfig" env-required:"true"`
	Fcfg     FilesConfig
}

type DBconfig struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	Database string `yaml:"database" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
}

type FilesConfig struct {
	MeteoPath     string `yaml:"meteo_path" env-default:"./public/meteo"`
	MeteoTemplate string `yaml:"meteo_template" env-default:"public/meteo/*.txt"`
	EnPath        string `yaml:"en_path" env-default:"./public/en"`
	EnTemplate    string `yaml:"en_template" env-default:"public/en/*.csv"`
}

func (db DBconfig) ToString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", db.User, db.Password, db.Host, db.Port, db.Database)
}

func MustLoadCfg() Config {

	once.Do(func() {
		path := filepath.FromSlash(fetchConfigPath())

		MustLoadByPath(path)
	})

	return Cfg

}
func fetchConfigPath() string {

	var res string

	flag.StringVar(&res, "cfg", "./config/config.yaml", "path to config yaml file")
	flag.Parse()

	return res
}
func MustLoadByPath(path string) Config {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file is not exists: " + path)
	}

	if err := cleanenv.ReadConfig(path, &Cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return Cfg
}

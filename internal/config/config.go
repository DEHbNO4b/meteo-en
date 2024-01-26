package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

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
	ResCfg   ResearchConfig
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
type ResearchConfig struct {
	Dur   time.Duration
	Begin time.Time
	End   time.Time
}

func (db DBconfig) ToString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", db.User, db.Password, db.Host, db.Port, db.Database)
}

func MustLoadCfg() Config {

	once.Do(func() {
		path := filepath.FromSlash(fetchConfigPath())

		MustLoadByPath(path)
		Cfg.ResCfg = parse()
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

func parse() ResearchConfig {

	rc := ResearchConfig{}

	var begin, end string

	flag.DurationVar(&rc.Dur, "dur", 30*time.Minute, "research duration")
	flag.StringVar(&begin, "begin", "2022-01-01", "research begin time")
	flag.StringVar(&end, "end", "2022-12-31", "research end time")

	flag.Parse()

	t, err := time.Parse("2006-01-02", begin)
	if err != nil {
		fmt.Println("unable to parse begin time: ", err)
	}
	rc.Begin = t

	t, err = time.Parse("2006-01-02", end)
	if err != nil {
		fmt.Println("unable to parse end time: ", err)
	}
	rc.End = t

	return rc
}

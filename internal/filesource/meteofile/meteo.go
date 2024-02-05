package meteofile

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"meteo-lightning/internal/config"
	"meteo-lightning/internal/domain/models"
	"meteo-lightning/internal/filesource"
	"os"
	"path/filepath"
	"strings"
)

// Metea data source
type MeteoSource struct {
	path     string
	template string
}

func NewMeteo(p, t string) (MeteoSource, error) {

	if p == "" || t == "" {
		return MeteoSource{}, filesource.ErrEmptyDataSource
	}

	return MeteoSource{path: filepath.FromSlash(p), template: filepath.FromSlash(t)}, nil

}

// Search new files with meteo data in path
func (m *MeteoSource) Search() ([]string, error) {

	var names []string

	err := filepath.WalkDir(m.path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			matched, err := filepath.Match(m.template, filepath.FromSlash(path))
			if err != nil {
				return err
			}
			if matched {
				//читаем файлы совпавшие с заданной строкой в требуемой директории
				names = append(names, path)
			}
		}
		return nil
	})

	return names, err
}

func Data(path string) ([]models.MeteoData, error) {

	data := make([]models.MeteoData, 0, 10000)

	file, err := os.Open(path)
	if err != nil {
		return data, err
	}
	defer file.Close()

	fileName := filepath.Base(path)
	names := strings.Split(fileName, ".")
	name := names[0]

	r := csv.NewReader(file)
	r.Comma = '\t'

	strings, err := r.ReadAll()
	if err != nil {
		fmt.Println("\t somthing went wrong", err)
	}
	if len(strings) <= 2 {
		fmt.Println("empty data on path: ", path)
		return data, ErrEmtyData
	}
	for _, el := range strings[2:] {
		if len(el) != 30 {
			fmt.Println("wrong string data: ", el)
			continue
		}
		d := makeData(el)
		dmd, err := meteoToDomain(d)
		if err != nil {
			fmt.Printf("unable to convert meteo data to domain: %v\n", err)
			continue
		}

		dmd.StName = name
		data = append(data, dmd)
	}

	return data, nil
}

func makeData(rec []string) meteoData {

	md := meteoData{}

	md.Date = rec[0]
	md.Time = rec[1]
	md.TempOut = rec[2]
	md.WindSpeed = rec[7]
	md.WindDir = rec[8]
	md.WindRun = rec[9]
	md.HiSpeed = rec[10]
	md.WindChill = rec[12]
	md.Bar = rec[15]
	md.Rain = rec[16]
	md.RainRate = rec[17]

	return md

}

func Files() ([]string, error) {

	//load config
	cfg := config.MustLoadCfg()

	// Create meteo file sorce struct
	ms, err := NewMeteo(cfg.Fcfg.MeteoPath, cfg.Fcfg.MeteoTemplate)
	if err != nil {
		return nil, err
	}

	// search files with meteo data
	files, err := ms.Search()
	if err != nil {
		return nil, err
	}
	return files, nil
}

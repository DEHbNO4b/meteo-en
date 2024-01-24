package enfile

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"meteo-lightning/internal/config"
	"meteo-lightning/internal/domain/models"
	"meteo-lightning/internal/filesource"
	"os"
	"path/filepath"
)

// Metea data source
type EnSource struct {
	path     string
	template string
}

func NewEn(p, t string) (EnSource, error) {

	if p == "" || t == "" {
		return EnSource{}, filesource.ErrEmptyDataSource
	}

	return EnSource{path: filepath.FromSlash(p), template: filepath.FromSlash(t)}, nil

}

// Search new files with meteo data in path
func (e *EnSource) Search() ([]string, error) {

	var names []string

	err := filepath.WalkDir(e.path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			matched, err := filepath.Match(e.template, filepath.FromSlash(path))
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

func Files() ([]string, error) {

	//load config
	cfg := config.MustLoadCfg()

	// Create meteo file sorce struct
	ms, err := NewEn(cfg.Fcfg.EnPath, cfg.Fcfg.EnTemplate)
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

func Data(path string) ([]models.StrokeEN, error) {

	data := make([]models.StrokeEN, 0, 10000)

	file, err := os.Open(path)
	if err != nil {
		return data, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ';'

	strings, err := r.ReadAll()
	if err != nil {
		fmt.Println("\t somthing went wrong", err)
	}

	if len(strings) <= 1 {
		fmt.Println("empty data on path: ", path)
		return data, filesource.ErrEmptyData
	}

	for _, el := range strings[1:] {

		if len(el) != 8 {
			fmt.Println("wrong string data: ", el)
			continue
		}

		stroke, err := makeStroke(el)
		if err != nil {
			fmt.Println("wrong string data: ", el, err)
			continue
		}

		ds := enToDomain(stroke)

		data = append(data, ds)
	}

	return data, nil
}

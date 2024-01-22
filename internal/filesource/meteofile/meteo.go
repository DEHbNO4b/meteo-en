package meteofile

import (
	"bufio"
	"fmt"
	"io/fs"
	"meteo-lightning/internal/domain/models"
	source "meteo-lightning/internal/filesource"
	"os"
	"path/filepath"
	"strings"
)

// Metea data source
type MeteoSource struct {
	path     string
	template string
}

func New(p, t string) (MeteoSource, error) {

	if p == "" || t == "" {
		return MeteoSource{}, source.ErrEmptyDataSource
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

func (m *MeteoSource) Read(path string) ([]models.MeteoData, error) {

	data := make([]models.MeteoData, 100)
	file, err := os.Open(path)
	if err != nil {
		return data, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		meteoLine, err := parseLine(line)
		if err != nil {
			fmt.Println(line)
			fmt.Println("\t", err)
			continue
		}
		fmt.Printf("%v\n", meteoLine)

		// data = append(data, meteoLine)
	}
	return data, nil
}

func parseLine(l string) (meteoData, error) {

	rec := strings.Split(l, "\t")
	if len(rec) != 30 {
		return meteoData{}, source.ErrInvalidDataString
	}

	md := meteoData{}

	md.Date = rec[0]
	md.Time = rec[1]
	md.TempOut = rec[2]
	md.WindSpeed = rec[7]
	md.WindRun = rec[9]
	md.Bar = rec[15]
	md.Rain = rec[16]
	md.RainRate = rec[17]

	return md, nil

}

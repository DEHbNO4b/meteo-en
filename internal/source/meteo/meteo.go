package meteo

import (
	"io/fs"
	"meteo-lightning/internal/source"
	"path/filepath"
)

// Metea data source
type MeteoSource struct {
	path     string
	template string
}

func New(p, t string) (MeteoSource, error) {

	if p == "" || t == "" {
		return MeteoSource{}, source.EmptyDataSource
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

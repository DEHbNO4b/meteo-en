package postgres

import (
	"meteo-lightning/internal/domain/models"
	"regexp"
	"strconv"
	"strings"
)

func domainMeteoParToLoc(mp *models.MeteoParams) MeteoParams {

	if mp == nil {
		return MeteoParams{}
	}

	lmp := MeteoParams{}

	lmp.MaxRain = mp.MaxRain
	lmp.MaxRainRate = mp.MaxRainRate
	lmp.MaxWindSpeed = mp.HiSpeed
	lmp.Rain = mp.Rain
	lmp.RainRate = mp.RainRate
	lmp.WindSpeed = mp.WindSpeed

	return lmp
}

func parseWind(w string) (int, int, error) {
	max, avg := 0, 0

	reMax := regexp.MustCompile(`{\d+}`)
	reAll := regexp.MustCompile(`\d+`)

	maxStr := reMax.FindString(w)
	all := reAll.FindAll([]byte(w), -1)
	if maxStr != "" {
		maxStr, _ := strings.CutPrefix(maxStr, "{")
		maxStr, _ = strings.CutSuffix(maxStr, "}")

		lMax, err := strconv.Atoi(maxStr)
		if err != nil {
			return 0, 0, err
		}
		max = lMax
	}

	for i, el := range all {
		if maxStr != "" && i >= len(all)-1 {
			break
		}
		w, err := strconv.Atoi(string(el))
		if err != nil {
			continue
		}
		if i == 0 {
			avg = w
			continue
		}
		avg = (avg + w) / 2
	}
	return avg, max, nil

}

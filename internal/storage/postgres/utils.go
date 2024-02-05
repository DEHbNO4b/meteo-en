package postgres

import "meteo-lightning/internal/domain/models"

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

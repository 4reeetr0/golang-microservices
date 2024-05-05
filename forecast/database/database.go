package database

import (
	"database/sql"
	"fmt"
	config "forecast-service/config"
	models "forecast-service/models"
	"time"

	owm "github.com/briandowns/openweathermap"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("mysql", "root:rootroot@/local")
	if err != nil {
		panic(err.Error())
	}
}

func addForecast() error {
	var forecastReq models.ForecastRequest
	forecastReq.Units = "C"
	forecastReq.Lang = "EN"

	rows, err := DB.Query("SELECT id, lattitude, longitude FROM forecast_locations")
	if err != nil {
		fmt.Println(err)
		return err
	}

	var locationId int
	for rows.Next() {
		err = rows.Scan(&locationId, &forecastReq.Lattitude, &forecastReq.Longitude)
		if err != nil {
			fmt.Println(err)
			return err
		}

		w, err := owm.NewCurrent(forecastReq.Units, forecastReq.Lang, config.API_KEY)
		if err != nil {
			return err
		}

		w.CurrentByCoordinates(&owm.Coordinates{Latitude: forecastReq.Lattitude, Longitude: forecastReq.Longitude})

		res := models.ForecastResponse{
			Country:             w.Sys.Country,
			City:                w.Name,
			WheatherMain:        w.Weather[0].Main,
			WheatherDescription: w.Weather[0].Description,
			Temperature:         w.Main.Temp,
			TemperatureMin:      w.Main.TempMin,
			TemperatureMax:      w.Main.TempMax,
			FeelsLike:           w.Main.FeelsLike,
			Humidity:            w.Main.Humidity,
			Pressure:            w.Main.Pressure,
			Visibility:          w.Visibility,
			Cloudiness:          w.Clouds.All,
			WindSpeed:           w.Wind.Speed,
			WindDeg:             w.Wind.Deg,
			Rain:                w.Rain.OneH,
			Date:                time.Unix(int64(w.Dt), 0),
		}

		_, err = DB.Exec("INSERT INTO forecasts (country, city, weather_main, weather_description, temperature, temperature_min, temperature_max, feels_like, humidity, pressure, visibility, cloudiness, wind_speed, wind_deg, rain, date, location_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", res.Country, res.City, res.WheatherMain, res.WheatherDescription, res.Temperature, res.TemperatureMin, res.TemperatureMax, res.FeelsLike, res.Humidity, res.Pressure, res.Visibility, res.Cloudiness, res.WindSpeed, res.WindDeg, res.Rain, w.Dt, locationId)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	defer rows.Close()

	return nil
}

func AddForecastsHourly() {
	ticker := time.Tick(time.Hour)

	addForecast()

	go func() {
		for {
			<-ticker
			addForecast()
		}
	}()

	select {}
}

func GetForecastHistory() ([]models.ForecastHistoryResponse, error) {
	rows, err := DB.Query("SELECT country, city, weather_main, weather_description, temperature, temperature_min, temperature_max, feels_like, humidity, pressure, visibility, cloudiness, wind_speed, wind_deg, rain, date, location_id FROM forecasts")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var curForecast models.ForecastResponse
	var curHistory models.ForecastHistoryResponse
	forecasts := make([]models.ForecastResponse, 0)
	result := make([]models.ForecastHistoryResponse, 0)
	var curLocationId int = -1
	var locationId int

	for rows.Next() {
		var dateUnix int

		err = rows.Scan(&curForecast.Country, &curForecast.City, &curForecast.WheatherMain, &curForecast.WheatherDescription, &curForecast.Temperature, &curForecast.TemperatureMin, &curForecast.TemperatureMax, &curForecast.FeelsLike, &curForecast.Humidity, &curForecast.Pressure, &curForecast.Visibility, &curForecast.Cloudiness, &curForecast.WindSpeed, &curForecast.WindDeg, &curForecast.Rain, &dateUnix, &locationId)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		curForecast.Date = time.Unix(int64(dateUnix), 0)
		fmt.Println(curForecast.Date)
		curHistory.Country = curForecast.Country
		curHistory.City = curForecast.City

		if curLocationId != locationId {
			if curLocationId != -1 {
				curHistory.History = forecasts
				result = append(result, curHistory)
			}
			curLocationId = locationId
			forecasts = make([]models.ForecastResponse, 0)
		}

		forecasts = append(forecasts, curForecast)
	}

	defer rows.Close()

	return result, nil
}

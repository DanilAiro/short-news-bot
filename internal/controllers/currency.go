package controllers

import (
	"time"

	"short-news-bot/internal/initializers"
	"short-news-bot/internal/models"

	"gorm.io/gorm"
)

var (
	AllCurrencies = make(map[string]float64)
	rub float64
)

func UpdateCurrencies() {
	currentDate := time.Now().Format("2006-01-02")

	lastCurrency := []models.Currency{}
	initializers.DB.Find(&lastCurrency)

	if len(lastCurrency) > 0 && lastCurrency[0].LastUpdate == currentDate {
		for _, v := range lastCurrency {
			AllCurrencies[v.Name] = v.Cost
		}

		return
	}

	latestRates, err := initializers.CUR_API.GetLatest([]string{"USD", "RUB"})

	if err != nil {
		panic("Failed to connect to currency api: " + err.Error())
	}

	for k, v := range latestRates.Rates {
		if (k == "RUB") {
			rub = v
			AllCurrencies["EUR"] = v
			break
		}
	}
	
	for k, v := range latestRates.Rates {
		if (k == "RUB") {
			k = "EUR"
			v = 1
		}

		AllCurrencies[k] = rub / v

		currencyModel := models.Currency{}
		err := initializers.DB.Where("name = ?", k).First(&currencyModel).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// Запись не найдена, создаем новую запись
				currencyModel = models.Currency{
					Name:       k,
					Cost:       rub / v,
					LastUpdate: currentDate,
				}
				err := initializers.DB.Create(&currencyModel).Error
				if err != nil {
					initializers.Log.Fatalf("Error creating currency in database: %v", err)
				}
			} else {
				initializers.Log.Fatalf("Error checking currency existence: %v", err)
			}
		} else {
			// Запись найдена, проверяем дату и обновляем при необходимости
			currencyModel.Cost = rub / v
			currencyModel.LastUpdate = currentDate
			err := initializers.DB.Save(&currencyModel).Error
			if err != nil {
				initializers.Log.Fatalf("Error updating currency in database: %v", err)
			}
		}
	}

}

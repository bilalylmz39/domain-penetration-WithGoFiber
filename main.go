package main

import (
	"domain-penetration-testing/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func OpenPageRank(domain string) (model.Data, error) {
	// API'ye bağlanmak için gereken URL
	url := fmt.Sprintf("%s?domains[]=%s", model.OpenPageRankURL, domain)

	// API anahtarı
	apiKey := "w4w4k0w4kggg888sgs40w4w8sk4o4k84co0o8488 "

	// HTTP GET isteği
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.Data{}, err
	}
	// API anahtarını header'a ekleme
	req.Header.Set("API-OPR", apiKey)

	// İstek gönderme
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.Data{}, err
	}
	defer resp.Body.Close()

	// API'den gelen JSON verisini okuma
	data := model.Data{}
	domainResult := model.DomainResult{}
	domainResult.Domain = domain

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return model.Data{}, err
	}

	// API'den dönen sonuçları geri döndürüyoruz
	return data, nil
}

func main() {
	api := fiber.New()

	api.Get("/api/pagerank", func(ctx *fiber.Ctx) error {

		domain := ctx.Query("domain", "") // "domain" parametresini alıyoruz

		if domain == "" { // Domain boşsa 400 Bad Request hatası döndürüyoruz
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Missing domain parameter",
			})
		}

		results, err := OpenPageRank(domain) // API'den sonuçları alıyoruz
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return ctx.JSON(fiber.Map{ // Sonuçları JSON formatında döndürüyoruz
			"results": results,
		})
	})

	// Uygulamayı 8080 portundan başlatıyoruz
	fmt.Println("Server listening on port 8080...")
	api.Listen(":8080")
}

package handlers

import (
	"time"

	"github.com/dkr290/go-projects/httpchecker/data"
	"github.com/gofiber/fiber/v2"
)

func HandleGetDashboard(c *fiber.Ctx) error {

	sslTrackings := []data.SSLTracking{
		{
			ID:         1,
			DomainName: "sendit.sh",
			Issuer:     "Let's Encrypt",
			Status:     "OK",
			Expires:    time.Now().AddDate(0, 0, 4),
		},
		{
			ID:         2,
			DomainName: "microsoft.com",
			Issuer:     "Let's Encrypt",
			Status:     "OK",
			Expires:    time.Now().AddDate(0, 0, 20),
		},
		{
			ID:         3,
			DomainName: "thetotalcoder.com",
			Issuer:     "Let's Encrypt",
			Status:     "OK",
			Expires:    time.Now().AddDate(0, 10, 20),
		},
	}
	data := fiber.Map{"trackings": sslTrackings}
	return c.Render("dashboard/index", data)

}

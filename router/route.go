package router

import (
	"api/handler/jadwal"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/jadwal", jadwal.TambahJadwal)
	api.Get("/jadwal", jadwal.GetMatkul)
	api.Get("/jadwal/:kode", jadwal.GetMatkulByKode)

}

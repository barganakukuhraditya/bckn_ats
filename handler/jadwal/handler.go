package jadwal

import (
	"api/model"
	"api/repository/db"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func TambahJadwal(c *fiber.Ctx) error {
	var requestData model.JadwalMatkul
	if err := c.BodyParser(&requestData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	if err := db.InsertDataJadwal(requestData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert data into the database",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data inserted successfully",
		"data":    requestData,
	})
}

func GetMatkul(c *fiber.Ctx) error {
	filter := bson.M{}

	requestData, err := db.GetDataJadwalFilter(filter)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    requestData,
	})
}

func GetMatkulByKode(c *fiber.Ctx) error {
	kode := c.Params("kode")
	filter := bson.M{"kode": kode}
	requestData, err := db.GetOneDataJadwalFilter(filter)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "Data retrieved successfully",
		"data":    requestData,
	})
}

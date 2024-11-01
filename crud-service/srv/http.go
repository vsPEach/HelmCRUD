package srv

import (
	"encoding/json"
	"net"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vsPEach/HelmCRUD/crud-service/config"
	"github.com/vsPEach/HelmCRUD/crud-service/models"
	"github.com/vsPEach/HelmCRUD/crud-service/storage"
)

func Run(storage *storage.Storage, config *config.Config) error {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{
			"status": "ok",
		})
	})

	app.Post("/user", func(c *fiber.Ctx) error {
		var user models.User
		err := json.Unmarshal(c.Body(), &user)
		if err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}

		raw, err := storage.Create(user)
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}
		return c.JSON(map[string]int64{
			"rawAffected": raw,
		})
	})

	app.Get("/user/:user_id", func(c *fiber.Ctx) error {
		idParam, ok := c.AllParams()["user_id"]
		if !ok {
			return c.SendString("can't get user id")
		}
		id, err := strconv.ParseInt(idParam, 10, 0)
		if err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}

		user, err := storage.Get(uint64(id))
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}

		return c.JSON(user)
	})

	app.Put("/user/:user_id", func(c *fiber.Ctx) error {
		var user models.User
		idParam, ok := c.AllParams()["user_id"]
		if !ok {
			return c.SendString("can't get user id")
		}

		id, err := strconv.ParseInt(idParam, 10, 0)
		if err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}

		err = json.Unmarshal(c.Body(), &user)
		if err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}

		user.ID = uint64(id)

		raw, err := storage.Update(user)
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}

		return c.JSON(map[string]int64{
			"rawAffected": raw,
		})
	})

	app.Delete("/user/:user_id", func(c *fiber.Ctx) error {
		idParam, ok := c.AllParams()["user_id"]
		if !ok {
			return c.SendString("can't get user id")
		}

		id, err := strconv.ParseInt(idParam, 10, 0)
		if err != nil {
			return c.SendStatus(http.StatusBadRequest)
		}

		err = storage.Delete(uint64(id))
		if err != nil {
			return c.SendStatus(http.StatusInternalServerError)
		}

		return c.JSON(id)
	})

	return app.Listen(net.JoinHostPort(config.Host, config.Port))
}

package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type response struct {
	Msg string
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		resp := response{Msg: "Hello from server!"}
		jsonx, err := json.Marshal(resp)
		if err != nil {
			return c.SendString("There was an error!")
		}
		return c.SendString(string(jsonx))
	})

	app.Post("/", func(c *fiber.Ctx) error {
		msgFromClient := struct {
			Msg string `json:"Msg"`
		}{}
		if err := c.BodyParser(&msgFromClient); err != nil {
			return err
		}
		jsonBody := struct {
			Msg string `json:"Msg"`
		}{Msg: "Hello from server"}
		return c.JSON(jsonBody)
	})

	app.Listen(":3000")
}

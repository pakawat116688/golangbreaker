package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main()  {
	app := fiber.New()
	count := 0 
	app.Get("/api", func(c *fiber.Ctx) error {
		count++
		if count >= 5 && count <= 10 {
			time.Sleep(time.Millisecond * 1000)
		}
		msg := fmt.Sprintf("api server %v",count)
		fmt.Println(msg)
		return c.SendString(msg)
	})
	app.Listen(":8000")
}
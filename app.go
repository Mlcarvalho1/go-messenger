package main

import (
	"log"

	"go.messenger/database"
	"go.messenger/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

const gopherDraw = `
       ,_---~~~~~----._
_,,_,*^____      _____''*g*\"*,          Welcome to go message!
/ __/ /'     ^.  /      \ ^@q   f     /
[  @f | @))    |  | @))   l  0 _/    /
\'/   \~____ / __ \_____/    \      /
|           _l__l_           I     /
}          [______]           I   /
]            | | |            |
]             ~ ~             |
|                            |
|                           |`

func main() {
	app := fiber.New()
	fireAuth := database.InitFirebaseAuth()

	routes.SetupRoutes(app, fireAuth)

	log.Println(gopherDraw)

	app.Listen(":3000")
}

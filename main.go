package main

import(
	"go-api-posts/database"
	"go-api-posts/routes"
)

func main()  {
	database.ConnectDatabase()
	r := routes.SetupRouter()
	r.Run(":8080")
}

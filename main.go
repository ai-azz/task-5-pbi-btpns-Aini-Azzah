package main

import (
	"github.com/ai-azz/task-5-btpns-Aini-Azzah/database"
	"github.com/ai-azz/task-5-btpns-Aini-Azzah/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}

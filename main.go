package main

import (
	"Ecsite/app/controllers"
	"Ecsite/app/models"
	"fmt"
)

func main() {

	fmt.Println(models.Db)
	controllers.StarMainSerever()
}

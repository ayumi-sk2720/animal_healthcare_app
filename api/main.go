package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/myantyuWorld/animal_healthcate/handler"
	"github.com/myantyuWorld/animal_healthcate/injector"
)

// Go言語：文法基礎まとめ | https://qiita.com/HiromuMasuda0228/items/65b9a593275f769f6b69
func main() {
	fmt.Println("sever start")
	petHandler := injector.InjectPetHandler()
	e := echo.New()

	// pet
	handler.InitRouting(e, petHandler)
	e.Logger.Fatal(e.Start(":8080"))
}

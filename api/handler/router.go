package handler

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// middleware
func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("Request Body: %v\n", string(reqBody))
	fmt.Printf("Response Body: %v\n", string(resBody))
	fmt.Println("==================================================================")
}

func InitRouting(e *echo.Echo, petHandler PetHandler) {
	e.Use(middleware.Logger())
	// TODO: ？？？　GolangでXMLHttpRequestLevel2+CORSのプリフライトが通るサーバーを立てる(BasicAuth付き) | https://qiita.com/romot/items/b49f7d9e28101daaa99e
	e.Use(middleware.CORS()) // 以下を追加
	e.Use(middleware.BodyDump(bodyDumpHandler))

	e.GET("/pet/:id", petHandler.TopView())
	e.POST("/pet/:id/schedule", petHandler.PostSchedule())
}

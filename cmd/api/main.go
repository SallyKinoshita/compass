package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SallyKinoshita/compass/internal/application/usecases"
	openapicompass "github.com/SallyKinoshita/compass/internal/gen/openapi"
	"github.com/SallyKinoshita/compass/internal/infrastructure/db"
	"github.com/SallyKinoshita/compass/internal/infrastructure/persistence"
	"github.com/SallyKinoshita/compass/internal/interfaces/controllers"
	"github.com/labstack/echo/v4"
)

func main() {
	// DB接続
	dbConn, err := db.NewMySQLConnection()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// DI設定
	studentRepo := persistence.NewStudentRepo(dbConn)
	studentUsecase := usecases.NewStudentUsecase(studentRepo)
	studentController := controllers.NewStudentController(studentUsecase)

	// Echoインスタンス作成
	e := echo.New()

	// カスタムエラーハンドラー
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		// HTTPエラーを確認
		if he, ok := err.(*echo.HTTPError); ok {
			log.Printf("HTTP Error: %v, Path: %s", he.Message, c.Path())
			c.NoContent(he.Code)
			return
		}
		// その他のエラー
		log.Printf("Internal Server Error: %v", err)
		c.NoContent(http.StatusInternalServerError)
	}

	// OpenAPIのハンドラー登録
	openapicompass.RegisterHandlers(e, studentController)

	// サーバー起動
	port := os.Getenv("API_PORT")
	if port == "" {
		port = ":8080" // デフォルトポート
	}
	log.Printf("Starting server on %v", port)
	if err := e.StartTLS(port, "/app/server.crt", "/app/server.key"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

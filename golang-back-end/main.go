package main

import (
	"github.com/arcade-stick-store/database"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"net/http"
	"os"
)

type PingResponse struct {
	Value string `json:"value"`
}

func setupRouter(pr *database.ProductRepo) *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		// Add CORS configuration
		c.Header("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, PingResponse{Value: "pong"})
	})

	r.GET("/products", func(c *gin.Context) {
		products, err := pr.Products()
		if err != nil {
			_ = c.Error(err)
		} else {
			c.JSON(http.StatusOK, products)
		}
	})

	return r
}

func main() {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "whiffpunish",
		AllowNativePasswords: true,
	}

	db, err := database.ConnectSQL(cfg.FormatDSN())
	if err != nil {
		panic("Could not connect to database")
	}

	r := setupRouter(database.NewRepo(db.SQL))
	r.SetTrustedProxies([]string{"127.0.0.1"})
	r.Run(":8080")
}

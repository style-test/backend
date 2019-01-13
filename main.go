package main // import "github.com/style-test/backend"

import (
	"database/sql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

type ProductValue struct {
	Id        int            `json:"id"`
	ImageUrl  string         `json:"image_url"`
	MobileUrl sql.NullString `json:"mobile_url"`
	Price     int            `json:"price"`
	ShopEn    string         `json:"shop_en"`
	ShopKo    string         `json:"shop_ko"`
	Title     string         `json:"title"`
	Url       string         `json:"url"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderAccept,
			echo.HeaderContentType,
			echo.HeaderXRequestedWith,
		},
	}))

	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.GET("/products", func(c echo.Context) error {
		rows, err := db.Query("SELECT * FROM product")
		if err != nil {
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: "Not Query",
			}
		}
		defer rows.Close()

		var Lists []*ProductValue

		for rows.Next() {
			var Id int
			var ImageUrl string
			var MobileUrl sql.NullString
			var Price int
			var ShopEn string
			var ShopKo string
			var Title string
			var Url string

			err = rows.Scan(
				&Id, &ImageUrl, &MobileUrl, &Price, &ShopEn, &ShopKo, &Title, &Url,
			)

			if err != nil {
				return &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: err.Error(),
				}
			}

			Lists = append(Lists, &ProductValue{
				Id:        Id,
				ImageUrl:  ImageUrl,
				MobileUrl: MobileUrl,
				Price:     Price,
				ShopEn:    ShopEn,
				ShopKo:    ShopKo,
				Title:     Title,
				Url:       Url,
			})
		}

		return c.JSON(http.StatusOK, Lists)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

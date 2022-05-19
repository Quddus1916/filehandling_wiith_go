package controller

import (
	"example.com/database"
	"path"
	//"example.com/models"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"net/http"
	"os"
)

var imagenames []string

type customcontext struct {
	echo.Context
}

var filecollection *mongo.Collection = database.OpenCollection(database.CLient, "userfiles")

func Sendfile(c echo.Context) error {

	//var userdata = new(models.User)
	form, err := c.MultipartForm()

	if err != nil {
		//msg := "binding failed"
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	files := form.File["image"]

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, file.Header)
		}
		defer src.Close()

		uploadedfilename := file.Filename
		upname := string(file.Filename)
		imagenames = append(imagenames, upname)
		uploadedfilepath := path.Join("./images", uploadedfilename)
		dst, err := os.Create(uploadedfilepath)
		if err != nil {

			return c.JSON(http.StatusBadRequest, err.Error())
		}
		defer dst.Close()
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

	}
	return c.JSON(http.StatusOK, imagenames)
}

func Getfiles(c echo.Context) error {
	return nil
}

package models

import (
	"mime/multipart"
)

type User struct {
	ID    string                `form:"id"`
	Name  *string               `form:"name"`
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

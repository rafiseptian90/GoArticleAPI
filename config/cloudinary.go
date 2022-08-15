package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"os"
)

func InitCLD() (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUDINARY_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"))

	return cld, err
}

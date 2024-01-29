package workmodel

import (
	"gorm.io/gorm"

	"dbpkg"
)

type Work struct {
  gorm.Model
	Name string `json:"name"`
	Url string `json:"url"`
	Description string `json:"description"`
	EncodedImg string `json:"encodedImg"`
	Tech string `json:"tech"`
}

func Create(newWork Work) (*Work, error) {
	db := dbpkg.GormConnect()
	result := db.Create(&newWork)

	if result.Error != nil {
		return nil, result.Error
	}
	return &newWork, nil
}

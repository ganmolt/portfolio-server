package workmodel

import (
	"gorm.io/gorm"

	"dbpkg"

	"log"
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

func Update(id string, data Work) (*Work, error) {
	db := dbpkg.GormConnect()
  var work Work
  db.First(&work, "id = ?", id)
  if data.Name != "" { work.Name = data.Name }
  if data.Description != "" { work.Description = data.Description }
  if data.Tech != "" { work.Tech = data.Tech }
  if data.Url != "" { work.Url = data.Url }
  if data.EncodedImg != "" { work.EncodedImg = data.EncodedImg }
  res := db.Updates(&work)

	if res.Error != nil {
    log.Println(res.Error)
    return nil, res.Error
  }
	return &work, nil
}

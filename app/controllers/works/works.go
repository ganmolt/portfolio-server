package works

import (
  "gorm.io/gorm"
	"github.com/gin-gonic/gin"

	"controllers/dbpkg"
  
  "models/user"
)

func Create(c *gin.Context) {
  access_token := c.Request.Header.Get("access-token")
	_, errMessage := usermodel.Session(access_token)

  if errMessage != "" {
    c.JSON(401, gin.H{"err": errMessage})
  }

  type Work struct {
    gorm.Model
    Name string `json:"name"`
    Url string `json:"url"`
    Description string `json:"description"`
    EncodedImg string `json:"encodedImg"`
    Tech string `json:"tech"`
  }
	db := dbpkg.GormConnect()

	var newWork Work
	if err := c.ShouldBindJSON(&newWork); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	db.Create(&newWork)
	c.JSON(200, newWork)
}

func Works(c *gin.Context) {  
  type Work struct {
    Name string `json:"name"`
    Url string `json:"url"`
    Description string `json:"description"`
    EncodedImg string `json:"encodedImg"`
    Tech string `json:"tech"`
  }

  db := dbpkg.GormConnect()

  var works []Work
  db.Unscoped().Find(&works)
  c.JSON(200, works)
}

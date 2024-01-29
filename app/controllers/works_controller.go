package controllers

import (
	"github.com/gin-gonic/gin"

	"dbpkg"
  
  "models/user"
  "models/work"
)

type WorksController struct{}

func (wc WorksController) Create(c *gin.Context) {
  access_token := c.Request.Header.Get("access-token")
	_, errMessage := usermodel.Session(access_token)
  if errMessage != "" {
    c.JSON(401, gin.H{"err": errMessage})
    return
  }

  var newWork workmodel.Work
  if err := c.ShouldBindJSON(&newWork); err != nil {
    c.JSON(400, gin.H{"error": err.Error()})
    return
  }

  work, err := workmodel.Create(newWork)
  if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, work)
}

func (wc WorksController) Show(c *gin.Context) {  
  db := dbpkg.GormConnect()

  var works []workmodel.Work
  db.Unscoped().Find(&works)
  c.JSON(200, works)
}

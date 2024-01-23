package dbpkg

import (
	"gorm.io/gorm"
	"os"
	"fmt"
  "gorm.io/driver/mysql"
)

func GormConnect() *gorm.DB {
  dsn := fmt.Sprintf(
    "%s:%s@tcp(%s)/%s?tls=true",
    os.Getenv("DB_USERNAME"),
    os.Getenv("DB_PASSWORD"),
    os.Getenv("DB_HOST"),
    os.Getenv("DB_DATABASE"),
  )

  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }
  return db
}

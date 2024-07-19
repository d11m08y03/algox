package main

import (
	"log"

	"github.com/d11m08y03/algox/cmd/api"
	"github.com/d11m08y03/algox/config"
	"github.com/d11m08y03/algox/db"
	"github.com/go-sql-driver/mysql"
)

func main()  {
  db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
  })
  if err != nil {
    log.Fatal(err.Error())
  }

  err = db.Ping()
  if err != nil {
    log.Fatal(err.Error())
  }
  log.Println("DB connected")

  server := api.NewAPIServer(":8080", db)
  if err = server.Run(); err != nil {
    log.Fatal(err.Error())
  }
}

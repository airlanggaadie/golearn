package main

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gosuri/uiprogress"
	"golearn/m/v1/common"
	"golearn/m/v1/versions"
	"log"
)

func main() {
	// set default timezone to Asia/Jakarta
	if err := common.SetTimezone(""); err != nil {
		log.Fatal(err)
	}

	// initialization database config
	println("database init..")
	db := common.Init()
	defer db.Close()

	// setup migration
	m, errMigration := migrate.New(
		"file://database/migrations",
		common.ViperEnvVariable(common.ViperParameters{
			Key:        "DATABASE_URL",
		}))
	if errMigration != nil {
		log.Fatal(errMigration)
	}

	// refresh db
	println("refresh db..")
	if errMigrationUp := m.Down(); errMigrationUp != nil {
		log.Fatal(errMigrationUp)
	}

	println("migrate..")
	if errMigrationUp := m.Up(); errMigrationUp != nil {
		log.Fatal(errMigrationUp)
	}

	seed()

	// routing management
	r := gin.Default()
	v1 := r.Group("/api/v1")
	versions.VersionRegister(v1)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Fatal(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// seed insert dummy data for versions table
func seed()  {
	println("seeding process is running..")
	count := 4
	uiprogress.Start()
	bar := uiprogress.AddBar(count).AppendCompleted().AppendElapsed()
	for i := 1; i <= count; i++ {
		if _, err := versions.InsertVersion(versions.NewVersionsParams{
			App:         faker.FirstName(),
			Version:     fmt.Sprintf("1.%d",i),
			Code:        i,
			Description: faker.DomainName(),
		}); err != nil {
			log.Fatal("failed seeding")
		}
		bar.Incr()
	}
	uiprogress.Stop()
	println("seeding process success")
}
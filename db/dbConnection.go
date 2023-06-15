package db

import (
	"database/sql"
	"log"
	"time"

	constants "group-chat/constants"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB
var err error

func Connect(c *gin.Context) {
	var POSTGRES_URL_WITH_DATASBE string = "postgres://" + constants.DB_USERNAME + ":" + constants.DB_PASSWORD + "@" + constants.DB_HOST + "/" + constants.DB_NAME + "?sslmode=" + constants.DB_SSL_MODE + ""

	var POSTGRES_URL_WITHOUT_DATASBE string = "postgres://" + constants.DB_USERNAME + ":" + constants.DB_PASSWORD + "@" + constants.DB_HOST + "/?sslmode=" + constants.DB_SSL_MODE + ""

	log.Println(POSTGRES_URL_WITH_DATASBE)
	log.Println(POSTGRES_URL_WITHOUT_DATASBE)

	Createpgdatabase(c, POSTGRES_URL_WITHOUT_DATASBE)

	DB, err = sql.Open("postgres", POSTGRES_URL_WITH_DATASBE)
	if err != nil {
		log.Println("Error To Connect Databae")
		c.JSON(constants.SUCCESS, gin.H{
			"error": "Error To Connect Databae",
		})
	}

	// Set connection pool settings
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(time.Minute * 5)

}
func Createpgdatabase(c *gin.Context, url string) {
	pgdatabaseCon, err := sql.Open("postgres", url)
	if err != nil {
		log.Println("Error To Connect Databae")
		c.JSON(constants.SUCCESS, gin.H{
			"error": "Error To Connect Databae",
		})
	}
	if _, dbExecErr := pgdatabaseCon.Exec(TABLE_CREATION); dbExecErr != nil {
		log.Println("**********pgdatabase********************")
		log.Println(dbExecErr.Error())
		log.Println("********* pgdatabase********************")
	}
}

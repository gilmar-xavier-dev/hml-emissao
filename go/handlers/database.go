package handlers

import (
	"database/sql"
	"fmt"
	"os"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
func initDB() (*sql.DB, error) {

	host := os.Getenv("DBHOST")
	port := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASS")
	dbname := os.Getenv("DBNAME")

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	fmt.Println(conn)

	db, err := sql.Open("postgres", conn)
	CheckError(err)

	err = db.Ping()
	CheckError(err)

	fmt.Println("Conectado ao banco..")

	return db, err

}

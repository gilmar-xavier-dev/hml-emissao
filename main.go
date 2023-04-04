package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"pagai/teste12/go/handlers"
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

	//	fmt.Println(conn)

	db, err := sql.Open("postgres", conn)
	CheckError(err)

	err = db.Ping()
	CheckError(err)

	//	fmt.Println("Conectado ao banco..")

	return db, err

}

func main() {

	conn, _ := net.Dial("ip:icmp", "google.com")

	port := ":9181"

	fmt.Printf("Server rodando em %s%s...\n", conn.LocalAddr(), port)

	router := handlers.NewRouter()

	log.Fatal(http.ListenAndServe(port, router))
}

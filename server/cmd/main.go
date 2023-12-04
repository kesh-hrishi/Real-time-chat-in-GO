package main

import (
	"fmt"
	"log"
	"server/db"
	"server/internal/user"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection : %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	fmt.Println(userSvc)
	// userHandler := user.NewHandler(userSvc)
}

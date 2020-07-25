package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
	models "reinvest/pkg"
	"reinvest/pkg/models/postgres"
)


type application struct {
	// Define database and error object
	fmg *postgres.AppModel
}


func main() {
	// Create database and check if error in creating the database
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=reinvestor dbname=reinvest sslmode=disable password=postgres")
	if err != nil {
		panic("Unable to connect to database")
	}

	// Close database at the end of the work
	defer db.Close()

	app := &application{
		fmg: &postgres.AppModel{DB: db},
	}

	// Migrate the schema
	app.fmg.DB.AutoMigrate(&models.FmgProfile{})
	app.fmg.DB.AutoMigrate(&models.FmgListSnp{})
	app.fmg.DB.AutoMigrate(&models.FmgIncomeStatement{})
	app.fmg.DB.AutoMigrate(&models.FmgBalanceSheet{})
	app.fmg.DB.AutoMigrate(&models.FmgCashFlow{})
	app.fmg.DB.AutoMigrate(&models.FmgFinancialRatios{})

	fmt.Println("Server started on localhost:4000")
	log.Fatal(http.ListenAndServe("localhost:4000", app.routes()))
}

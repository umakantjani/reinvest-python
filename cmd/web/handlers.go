package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	models "reinvest/pkg"
	"strconv"
	"strings"
)

//Variables
var ApiLinks = map[string]string{
	"profile": 				"profile",
	"snp":					"sp500_constituent",
	"symbols":				"stock/list",
	"is":					"income-statement",
	"income-statement":		"income-statement",
	"bs":					"balance-sheet-statement",
	"balance-sheet":		"balance-sheet-statement",
	"cf":					"cash-flow-statement",
	"cash-flow":				"cash-flow-statement",
	"ratios":				"ratios",
	"fr":					"ratios",
	"financial-ratios":		"ratios",
}

// Welcome Handler
func (app *application) welcome(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Hi there, I love go!")
}


/*	Updates 	the Profile data for the provided symbol
	Route: 		"/collect/profile/{symbol}"
	Action: 	collects the data from the FMG API data and updates the database
	Returns: 	Message	*/
func (app *application) collectKpiData(w http.ResponseWriter, r *http.Request)  {
	// Collect parameters
	params := mux.Vars(r)
	kpi := params["kpi"]
	symbol := strings.ToUpper(params["symbol"])
	period := params["period"]
	//fmt.Println("KPI:", kpi, "Symbol:", symbol, "Period:", period)
	midString, periodType := "", "A"

	if strings.ToUpper(period) == "QTR" {
		midString, periodType =  "?period=quarter&apikey=", "Q"
	} else {
		midString, periodType =  "?apikey=", "A"
	}

	// check if provided KPI present in the kpi dictionary
	if val, ok := ApiLinks[kpi]; ok {
		// Execute this when input kpi & symbol is not null
		if symbol != "" {
			
			// Handling various kpi query types
			switch strings.ToUpper(kpi) {
			// Case statement for Profile data
			case "PROFILE":
				responseString := CollectData(ApiLinks[val] + "/" + symbol + "?apikey=")
				byt := []byte(responseString)
				var dat []models.FmgProfile
				if err := json.Unmarshal(byt, &dat); err != nil {
					panic(err)
				}
				for k, v := range dat {
					v.Pk = GetMD5Hash(v.Symbol)
					app.fmg.DB.Set("gorm:insert_option", "ON CONFLICT DO NOTHING").Create(&v)
					fmt.Println(strconv.Itoa(k+1) + ". " + strings.ToTitle(val) + " has been updated for " + v.Symbol)
				}
				fmt.Fprint(w, strings.ToTitle(val) + " updated in the database for [" + symbol + "]")
			// Case statement for Income Statement data
			case "IS", "INCOME-STATEMENT":
				var dataMod []models.FmgIncomeStatement
				responseString := CollectData(val + "/" + symbol + midString)
				byt := []byte(responseString)
				if err := json.Unmarshal(byt, &dataMod); err != nil {
					panic(err)
				}
				for k, v := range dataMod {
					v.Pk, v.PeriodType = GetMD5Hash(v.Symbol + v.Date + periodType), periodType
					app.fmg.DB.Set("gorm:insert_option", "ON CONFLICT DO NOTHING").Create(&v)
					fmt.Println(strconv.Itoa(k+1) + ". " + strings.ToTitle(val) + " has been updated for " + v.Symbol + " " + v.Date)
				}
				fmt.Fprint(w, strings.ToTitle(val) + " updated in the database for [" + symbol + "]")
			// Case statement for Balance Sheet data
			case "BS", "BALANCE-SHEET":
				var dataMod []models.FmgBalanceSheet
				responseString := CollectData(val + "/" + symbol + midString)
				byt := []byte(responseString)
				if err := json.Unmarshal(byt, &dataMod); err != nil {
					panic(err)
				}
				for k, v := range dataMod {
					v.Pk, v.PeriodType = GetMD5Hash(v.Symbol + v.Date + periodType), periodType
					app.fmg.DB.Set("gorm:insert_option", "ON CONFLICT DO NOTHING").Create(&v)
					fmt.Println(strconv.Itoa(k+1) + ". " + strings.ToTitle(val) + " has been updated for " + v.Symbol + " " + v.Date)
				}
				fmt.Fprint(w, strings.ToTitle(val) + " updated in the database for [" + symbol + "]")
			// Case statement for Balance Sheet data
			case "CF", "CASH-FLOW":
				var dataMod []models.FmgCashFlow
				responseString := CollectData(val + "/" + symbol + midString)
				byt := []byte(responseString)
				if err := json.Unmarshal(byt, &dataMod); err != nil {
					panic(err)
				}
				for k, v := range dataMod {
					v.Pk, v.PeriodType = GetMD5Hash(v.Symbol + v.Date + periodType), periodType
					app.fmg.DB.Set("gorm:insert_option", "ON CONFLICT DO NOTHING").Create(&v)
					fmt.Println(strconv.Itoa(k+1) + ". " + strings.ToTitle(val) + " has been updated for " + v.Symbol + " " + v.Date)
				}
				fmt.Fprint(w, strings.ToTitle(val) + " updated in the database for [" + symbol + "]")
			// Case statement for Financial Ratios data
			case "FR", "FINANCIAL-RATIOS", "RATIOS":
				var dataMod []models.FmgFinancialRatios
				responseString := CollectData(val + "/" + symbol + midString)
				byt := []byte(responseString)
				if err := json.Unmarshal(byt, &dataMod); err != nil {
					panic(err)
				}
				for k, v := range dataMod {
					v.Pk, v.PeriodType = GetMD5Hash(v.Symbol + v.Date + periodType), periodType
					app.fmg.DB.Set("gorm:insert_option", "ON CONFLICT DO NOTHING").Create(&v)
					fmt.Println(strconv.Itoa(k+1) + ". " + strings.ToTitle(val) + " has been updated for " + v.Symbol + " " + v.Date)
				}
				fmt.Fprint(w, strings.ToTitle(val) + " updated in the database for [" + symbol + "]")
			}

		} else {
			fmt.Fprint(w, "Please provide a symbol to get " + kpi + " information.")
		}
	} else {
		fmt.Println("KPI: " + kpi + " doesn't exist in the dictionary. Please add valid KPI Name.")
		fmt.Fprint(w, "KPI: " + kpi + " doesn't exist in the dictionary. Please add valid KPI Name.")
	}

}

/*	Updates 	the Lists data for the provided list name
	Route: 		"/collect/lists/{listName}"
	Action: 	collects the list data from the FMG API data and updates the database
	Returns: 	Message	*/
func (app *application) collectLists(w http.ResponseWriter, r *http.Request)  {
	// take the input parameter for symbol to get data from the FMG Profile
	params := mux.Vars(r)
	listName := params["listName"]
	fmt.Println(listName)

	// Execute this when input symbol is not nol
	if listName != "" {
		responseString := CollectData(listName)

		byt := []byte(responseString)
		var dat []models.FmgListSnp

		if err := json.Unmarshal(byt, &dat); err != nil {
			panic(err)
		}

		for k, v := range dat {
			v.Pk = GetMD5Hash(v.Symbol)
			app.fmg.DB.Set("gorm:insert_option", "ON CONFLICT DO NOTHING").Create(&v)
			fmt.Println(strconv.Itoa(k+1) + ". Document has been updated for " + v.Symbol)
		}

		fmt.Fprint(w, listName + " has been updated into the database!")

	} else {
		fmt.Fprint(w, "Please provide a symbol to get profile information.")
	}

}


// Find all function
func (app *application)  findAllDoc(w http.ResponseWriter, r *http.Request)  {
	var products []models.FmgListSnp
	app.fmg.DB.Find(&products)
	json.NewEncoder(w).Encode(&products)
}


// Find Single row
func (app *application)  findSingleDoc(w http.ResponseWriter, r *http.Request)  {
	var product models.FmgListSnp
	params := mux.Vars(r)
	app.fmg.DB.First(&product, params["code"])
	json.NewEncoder(w).Encode(&product)
}
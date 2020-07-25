package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"net/http"
	models "reinvest/pkg"
)

// Define a SnippetModel type which wraps a sql.DB connection pool.
type AppModel struct {
	DB *gorm.DB
}


// This will insert a new snippet into the database.
func (model *AppModel) Insert(profile []models.FmgProfile) (int, error) {
	return 0, nil
}


// This will return a specific snippet based on its id.
func (model *AppModel) Get(id int) (*models.FmgProfile, error) {
	return nil, nil
}


// This will return the 10 most recently created snippets.
func (model *AppModel) Latest() ([]*models.FmgProfile, error) {
	return nil, nil
}


// Function to get data from FMG API
func GetData() {
	url := "https://financialmodelingprep.com/api/v3/profile/AAPL?apikey=70c55be6005581c638125a74ec3c5d54"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)

	fmt.Println(responseString)
}
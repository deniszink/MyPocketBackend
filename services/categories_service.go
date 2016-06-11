package services

import (
	"backend/models"
	"backend/core/store"
	"gopkg.in/mgo.v2/bson"
	"log"
	"fmt"
)

var incomeCategories = []string{"Salary", "Selling", "Gifts", "Award", "Interest Money", "Other"}
var expenseCategories = []string{"Food & Beverage", "Shopping/Clothes", "Health/Medicine", "Gym", "Fees and Charges",
	"Insurance", "Education", "Gifts and presents", "Travel", "Love", "Entertainment",
	"Investment", "Repairs", "Transport", "Bills and Utilities", "Rent", "Family", "Other"}

func CreateCategoies() {
	mongo := store.ConnectMongo()
	var categories []models.Category
	mongo.FindAll(store.TableCategories,bson.M{},&categories)

	fmt.Println(len(categories))
	fmt.Println(len(incomeCategories) + len(expenseCategories))

	if len(categories) != (len(incomeCategories) + len(expenseCategories)) {

		if err := mongo.DropTable(store.TableCategories); err != nil{
			log.Fatal(err)
		}

		createExpenseCategories(mongo)
		createIncomeCategories(mongo)
	}
}

func createIncomeCategories(mongo *store.MongoDB) {
	for _, title := range incomeCategories {
		c := new(models.Category)
		c.Name = title
		c.Type = "income"
		if err := mongo.WriteDataTo(store.TableCategories, c); err != nil {
			panic(err)
		}

	}

}

func createExpenseCategories(mongo *store.MongoDB) {
	for _, title := range expenseCategories {
		c := new(models.Category)
		c.Name = title
		c.Type = "expense"
		if err := mongo.WriteDataTo(store.TableCategories, c); err != nil {
			panic(err)
		}
	}
}



package seed

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/joho/godotenv"

	"uwe/db"
	"uwe/types"
)

func main() {
	fmt.Println("running seeds...")
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	db := db.Create()

	customer := types.Customer{
		Id: uuid.New(),
	}

	if err := db.CreateCustomer(&customer); err != nil {
		log.Fatal(err)
	}

	fileUpload := types.FileUpload{
		ID:         uuid.New(),
		CustomerId: customer.Id,
		Mapping:    map[string]int{"amount": 0},
	}

	if err := db.CreateFileUpload(&fileUpload); err != nil {
		log.Fatal(err)
	}
}

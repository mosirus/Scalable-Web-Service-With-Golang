package main

import (
	"fmt"
	"sesi7-gorm/database"
	"sesi7-gorm/models"
	"sesi7-gorm/repository"
	"strings"

	"gorm.io/gorm"
)

func main() {
	db := database.StartDB()
	user(db)
	// product(db)
}

func user(db *gorm.DB) {
	userRepo := repository.NewUserRepo(db)

	// user := models.User{
	// 	Email: "yoyo@gmail.com",
	// }

	// err := userRepo.CreateUser(&user)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }

	// fmt.Println("Created success")

	employees, err := userRepo.GetUsersWithProducts()

	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	for k, emp := range *employees {
		fmt.Println("User :", k+1)
		emp.Print()
		fmt.Println(strings.Repeat("=", 10))
	}

	// emp, err := userRepo.GetUserByID(3)

	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }

	// emp.Print()

	//update
	// usrUpdate := models.User{
	// 	Email: "lestari@gmail.com",
	// 	ID:    2,
	// }

	// errUpdate := userRepo.UpdateUser(&usrUpdate)
	// if errUpdate != nil {
	// 	fmt.Println("error :", errUpdate.Error())
	// 	return
	// }

	// fmt.Println("Update User", usrUpdate)

	//delete
	// usrDelete, err := userRepo.DeleteUser(3)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }

	// fmt.Println("Delete User ", usrDelete)

}

func product(db *gorm.DB) {
	productRepo := repository.NewProductRepo(db)

	// product := models.Product{
	// 	Name:   "Boxer",
	// 	Brand:  "Adidas",
	// 	UserID: 2,
	// }

	// // create product
	// err := productRepo.CreateProduct(&product)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("Created Product Success !")

	products, err := productRepo.GetAllProduct()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for k, product := range *products {
		fmt.Println("Product :", k+1)
		product.Print()
		fmt.Println(strings.Repeat("=", 10))
	}

	//update
	productUpdate := models.Product{
		ID:     1,
		Name:   "Topi",
		Brand:  "Nike",
		UserID: 1,
	}

	errUpdate := productRepo.UpdateProduct(&productUpdate)
	if errUpdate != nil {
		fmt.Println("error :", errUpdate.Error())
		return
	}

	fmt.Println("Update Product", productUpdate)

	//delete
	productDelete, err := productRepo.DeleteProduct(2)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	fmt.Println("Delete Product ", productDelete)
}

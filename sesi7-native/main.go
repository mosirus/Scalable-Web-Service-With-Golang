package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "monang"
	password = "monang"
	dbname   = "hacktiv8"
)

func (e *Employee) Print() {
	fmt.Println("ID :", e.ID)
	fmt.Println("FullName :", e.FullName)
	fmt.Println("Email :", e.Email)
	fmt.Println("Age :", e.Age)
	fmt.Println("Division :", e.Division)
	fmt.Println()
}

var (
	db  *sql.DB
	err error
)

type Employee struct {
	ID       int    `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Division string `json:"division"`
}

func main() {
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// db, err = sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }

	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Successfully connected to database")

	db, err := connectDB()
	if err != nil {
		panic(err)
	}

	// Create Employee
	// emp := Employee{
	// 	Email:    "Saut@koinworks.com",
	// 	FullName: "Saut Sihotang",
	// 	Age:      22,
	// 	Division: "Project Manager",
	// }

	// fmt.Println(strings.Repeat("=", 20))

	// err = createEmployee(db, &emp)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }

	//emp.Print()

	employees, err := getAllEmployees(db)
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	for _, employee := range *employees {
		employee.Print()
	}

	fmt.Println(strings.Repeat("=", 20))
	// fmt.Println("===== Get Employee By ID =====")

	// employee, err := getEmployeeById(db, 6)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// }

	// employee.Print()

	// fmt.Println("===== Update Employee By ID =====")

	// empUpdate := Employee{
	// 	ID:       1,
	// 	FullName: "Reyhan",
	// 	Email:    "reyhan@dana.id",
	// 	Age:      23,
	// 	Division: "Back-End",
	// }

	// err = updateEmployee(db, &empUpdate)
	// if err != nil {
	// 	fmt.Println("Error :", err.Error())
	// 	return
	// }

	// fmt.Println("Update data successfully", empUpdate)

	// fmt.Println("===== Delete Employee By ID =====")

	// isDelete, err := deleteEmployeeById(db, 4)
	// if err != nil {
	// 	fmt.Println("Error", err.Error())
	// 	return
	// }

	// fmt.Println("Delete employee successfully", isDelete)

}

func connectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	//connection pool
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(10 * time.Second)
	db.SetConnMaxLifetime(10 * time.Second)

	fmt.Println("Successfully connected to database")

	return db, nil
}

func createEmployee(db *sql.DB, request *Employee) error {
	query := `
	 INSERT INTO employees (full_name, email, age, division)
	 VALUES($1, $2, $3, $4)
	`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(request.FullName, request.Email, request.Age, request.Division)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}
	return tx.Commit()
}

func getAllEmployees(db *sql.DB) (*[]Employee, error) {
	query := `
	 	SELECT id, full_name, email, age, division
		FROM employees
		ORDER BY id ASC
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	var employees []Employee

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var employee Employee
		err := rows.Scan(
			&employee.ID, &employee.FullName,
			&employee.Email, &employee.Age, &employee.Division,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, employee)
	}

	return &employees, nil
}

func getEmployeeById(db *sql.DB, id int) (*Employee, error) {
	query := `
	 	SELECT id, full_name, email, age, division
		FROM employees
		Where id=$1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var emp Employee

	row := stmt.QueryRow(id)

	err = row.Scan(
		&emp.ID, &emp.FullName, &emp.Email, &emp.Age, &emp.Division,
	)
	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func updateEmployee(db *sql.DB, request *Employee) error {
	query := `
	 UPDATE employees 
	 SET full_name = $2, email = $3, age = $4, division = $5
	 WHERE id = $1
	`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(request.ID, request.FullName, request.Email, request.Age, request.Division)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

func deleteEmployeeById(db *sql.DB, id int) (bool, error) {
	query := `
	 	DELETE FROM employees
		Where id=$1
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		return false, nil
	}

	defer stmt.Close()

	_ = stmt.QueryRow(id)

	return true, nil
}

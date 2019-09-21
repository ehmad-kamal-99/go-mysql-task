package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// User structure that stores values retrieved from csv
type User struct {
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Age        int    `json:"age"`
	BloodGroup string `json:"bloodgroup"`
}

func main() {
	FindCsv()
}

// FindCsv file in a directory
func FindCsv() {
	var root string

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("# To exit the program type 'exit'")
	fmt.Println("File Structure: first_name, last_name, age, blood_group")
	fmt.Println("Enter the file path to search for csv:")

	dir, _ := reader.ReadString('\n')

	root = strings.TrimRight(dir, "\r\n")

	if root == "exit" {
		os.Exit(1)
	} else {
		files, err := FilePathWalkDir(root)
		if err != nil {
			panic(err)
		}

		for _, file := range files {
			if filepath.Ext(file) == ".csv" {
				fmt.Println("File found at: " + filepath.Base(file))
				ReadCsvFile(filepath.Clean(file))
			} else if filepath.Ext(file) != ".csv" {

			} else {
				fmt.Println("No file found.")
				reader := bufio.NewReader(os.Stdin)
				fmt.Println("Try again? [y/n]")

				answer, _ := reader.ReadString('\n')
				ans := strings.TrimRight(answer, "\r\n")

				if ans == "y" || ans == "Y" {
					FindCsv()
				} else if ans == "n" || ans == "N" {
					os.Exit(1)
				}
			}
		}
	}
}

// ReadCsvFile function that reads csv files
func ReadCsvFile(file string) {

	f, _ := os.Open(file)
	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		age, err := strconv.Atoi(record[2])
		if err != nil {
			panic(err)
		}

		if record[0] != "" && record[1] != "" && age != 0 && record[3] != "" {
			user := new(User)
			user.FirstName = record[0]
			user.LastName = record[1]
			user.Age = age
			user.BloodGroup = record[3]
			InsertToDb(user)
		} else {
			fmt.Println("File not compatible for storage into database.")
			fmt.Println("Try another file or set the structure of file to: ")
			fmt.Println("first_name, last_name, age, blood_group")
			fmt.Println("Exiting")
			os.Exit(1)
		}
	}
}

// InsertToDb function
func InsertToDb(user *User) {

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3308)/testdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	sqlstatement := `
	INSERT INTO user (firstname, lastname, age, bloodgroup)
	VALUES (?, ?, ?, ?)
	`

	_, err = db.Exec(sqlstatement, user.FirstName, user.LastName, user.Age, user.BloodGroup)

	if err != nil {
		panic(err.Error())
	}

}

// FilePathWalkDir function
func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

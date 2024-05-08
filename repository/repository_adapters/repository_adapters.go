package repositoryadapters

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/akshay0074700747/book_store/configurations"
	customerrormessages "github.com/akshay0074700747/book_store/custom_errormessages"
	"github.com/akshay0074700747/book_store/entities"
	"github.com/akshay0074700747/book_store/helpers"
)

//repositoryadapter implements the repositoryport interface
type RepositoryAdapter struct {
	UsersDataPath      string
	UserBooksDataPath  string
	AdminBooksDataPath string
}

func NewRepositoryAdapter(config configurations.Configurations) *RepositoryAdapter {
	return &RepositoryAdapter{
		UsersDataPath:      config.UserDataPath,
		UserBooksDataPath:  config.UserBooksDataPath,
		AdminBooksDataPath: config.AdminBooksDataPath,
	}
}

func (repo *RepositoryAdapter) LoginUser(user entities.User) (entities.User, error) {

	// opening the json file
	file, err := os.Open(repo.UsersDataPath)
	if err != nil {
		fmt.Println("error opening file: ", err)
		return entities.User{}, err
	}
	defer file.Close()

	// creatiing a json decoder from the file
	decoder := json.NewDecoder(file)

	// reading the users array in the json data
	if _, err = decoder.Token(); err != nil && err != io.EOF {
		fmt.Println("error in reading the json token:", err)
		return entities.User{}, err
	}
	if _, err = decoder.Token(); err != nil && err != io.EOF {
		fmt.Println("error in reading the json token:", err)
		return entities.User{}, err
	}
	if delim, err := decoder.Token(); err != nil || delim != '[' {
		if err != nil && err != io.EOF {
			fmt.Println("invalid json format ", err)
			return entities.User{}, err
		}
	}

	var result entities.User
	for decoder.More() {
		if err := decoder.Decode(&result); err != nil {
			fmt.Println("error in decoding the json object: ", err)
			return entities.User{}, err
		}

		if result.Username == user.Username && result.Password == user.Password {
			return result, nil
		}
	}

	return entities.User{}, errors.New(customerrormessages.UserNotFound)
}

func (repo *RepositoryAdapter) GetBooks(isAdmin bool) ([]entities.Book, error) {

	var path string
	if isAdmin {
		path = repo.UserBooksDataPath
	} else {
		path = repo.AdminBooksDataPath
	}

	// opening the csv file
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("error in opening the file: ", err)
		return nil, err
	}
	defer file.Close()

	// creating a new csv reader
	reader := csv.NewReader(file)

	// reading all records from the csv
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error reading csv file:", err)
		return nil, err
	}

	var books []entities.Book

	for i, record := range records {

		// skipping the header row
		if i == 0 {
			continue
		}

		books = append(books, entities.Book{
			BookName:        record[0],
			Author:          record[1],
			PublicationYear: helpers.Parse(record[2]),
		})
	}

	return books, nil
}

func (repo *RepositoryAdapter) AddBook(book entities.Book) error {

	// opening hte csv file in an append only mode
	file, err := os.OpenFile(repo.UserBooksDataPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("eror in opening the  file:", err)
		return err
	}
	defer file.Close()

	// creating a new csv writer
	writer := csv.NewWriter(file)

	// writing the new record to the csv
	if err := writer.Write([]string{""}); err != nil {
		fmt.Println("Error writing new line:", err)
		return err
	}
	record := []string{book.BookName, book.Author, fmt.Sprintf("%d", book.PublicationYear)}
	if err := writer.Write(record); err != nil {
		fmt.Println("error writing record: ", err)
		return err
	}

	// ensuring every data is written to the file
	writer.Flush()

	return nil
}

func (repo *RepositoryAdapter) DeleteBook(bookName string) error {

	// opening the csv file for reading n writing
	file, err := os.OpenFile(repo.UserBooksDataPath, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	// creating a new csv reader
	reader := csv.NewReader(file)

	// creating a buffer to store the contents in the csv file
	var buffer strings.Builder

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		// checking if the record have to be deleted or not
		if !strings.EqualFold(record[0], bookName) {
			// writing the records which doesnt have to be deleted to the buffer
			buffer.WriteString(strings.Join(record, ",") + "\n")
		}
	}

	// removing all the existing contents in the file
	if err := file.Truncate(0); err != nil {
		fmt.Println("error in truncating the file:", err)
		return err
	}

	// seeking to the beginning of the file
	if _, err := file.Seek(0, 0); err != nil {
		fmt.Println("error seeking in the file:", err)
		return err
	}

	// writing the updated daata into the file
	if _, err := file.WriteString(buffer.String()); err != nil {
		fmt.Println("error writing to file: ", err)
		return err
	}

	return nil
}

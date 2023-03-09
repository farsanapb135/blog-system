package model

import (
	"app/blog-system/config"
	"fmt"
	"log"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// ArticleData for store article details
type Article struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []ArticleData `json:"data"`
}
type ArticleData struct {
	ID      string `db:"id"`
	Title   string `db:"title" `
	Content string `db:"content"`
	Author  string `db:"author"`
}
type ArticleResult struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID string `json:"id"`
	} `json:"data"`
}

var db *sqlx.DB

func DbConnection() *sqlx.DB {

	// for connect server on other host
	//dbinfo := fmt.Sprintf(config.Data.DbUser + ":" + config.Data.DbPass + "@tcp(" + config.Data.DbHost + ":" + config.Data.DbPort + ")/" + config.Data.DbName + "?parseTime=true")

	// for connect localserver
	dbinfo := fmt.Sprintf(config.Data.DbUser + ":" + config.Data.DbPass + "@/" + config.Data.DbName + "?parseTime=true")
	var err error
	db, err = sqlx.Open("mysql", dbinfo)
	if err != nil {
		log.Println("db error", err)
		log.Fatal("Error: The data source arguments are not valid", err)
	}

	return db

}

// Add new article details to DB
func AddArticleToDb(data ArticleData) (error, string) {

	dbConn := DbConnection()
	defer dbConn.Close()
	stmt, err := dbConn.Prepare("INSERT INTO article(title, content, author,created_on)VALUES(?,?,?,?);")
	if err != nil {
		log.Println("error", err)
		return err, ""
	}
	defer stmt.Close()
	res, err := stmt.Exec(data.Title, data.Content, data.Author, time.Now())
	if err != nil {
		log.Println("error", err)

		return err, ""
	}

	id, err := res.LastInsertId()
	return nil, strconv.FormatInt(id, 10)
	//return err, ""

	// ------------------------------------------

}

//GetArticleDetailsFromDB for read all article data from database
func GetArticleDetailsFromDB() (Article, error) {

	dbConn := DbConnection()
	defer dbConn.Close()

	sqlString := `
				SELECT id,
						title,
						content,
						author
				FROM   article 
				ORDER  BY id DESC; 
				`

	var allArticleDetails Article
	rows, err := dbConn.Queryx(sqlString)
	if err != nil {
		//log.Error(err)
		return allArticleDetails, err
	} else {
		defer rows.Close()
		var articleDetails ArticleData
		for rows.Next() {
			err := rows.StructScan(&articleDetails)
			if err != nil {
				//log.Error(err)
				return allArticleDetails, err
			}

			allArticleDetails.Data = append(allArticleDetails.Data, articleDetails)

		}
	}
	return allArticleDetails, err
}

//GetArticleDetailsByID for  read a specific article data using ID from database
func GetArticleDetailsByID(ID string) (Article, error) {

	dbConn := DbConnection()
	defer dbConn.Close()
	sqlString := `
			     SELECT id,
					title,
					content,
					author
				FROM   article 
				where id= '` + ID + `'; 
		`

	var allArticleDetails Article
	rows := dbConn.QueryRow(sqlString)

	var articleDetails ArticleData
	err := rows.Scan(&articleDetails.ID, &articleDetails.Title, &articleDetails.Content, &articleDetails.Author)
	if err != nil {
		//log.Error(err)
		return allArticleDetails, err
	}
	allArticleDetails.Data = append(allArticleDetails.Data, articleDetails)

	return allArticleDetails, err
}

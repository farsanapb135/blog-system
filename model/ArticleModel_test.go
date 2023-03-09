package model

import (
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
)

var testDb *sqlx.DB
var TestId string

func TestAddArticleToDb(t *testing.T) {

	dbinfo := "user1:password1@/blog_system?parseTime=true"
	testDbConn, err := sqlx.Open("mysql", dbinfo)
	if err != nil {
		log.Println("db error", err)
		t.Error("could not connect with database")
	}

	stmt, err := testDbConn.Prepare("INSERT INTO article(title, content, author)VALUES(?,?,?);")
	if err != nil {
		t.Error("error", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec("testing in go", "unit testing", "anu")
	if err != nil {
		t.Error("error", err)

	}

}

func TestGetArticleDetailsFromDB(t *testing.T) {

	dbinfo := "user1:password1@/blog_system?parseTime=true"
	testDbConn, err := sqlx.Open("mysql", dbinfo)
	if err != nil {
		log.Println("db error", err)
		t.Error("could not connect with database")
	}
	sqlString := `
				SELECT id
				FROM   article 
				ORDER  BY id DESC; 
				`

	rows, err := testDbConn.Queryx(sqlString)
	if err != nil {
		t.Error(err)
	} else {
		defer rows.Close()
		var articleDetails ArticleData
		for rows.Next() {
			err := rows.StructScan(&articleDetails)
			if err != nil {
				t.Error(err)

			}
		}
	}
}
func TestGetArticleDetailsByID(t *testing.T) {

	dbinfo := "user1:password1@/blog_system?parseTime=true"
	testDbConn, err := sqlx.Open("mysql", dbinfo)
	if err != nil {
		log.Println("db error", err)
		t.Error("could not connect with database")
	}
	sqlString := `
			     SELECT id
				 FROM   article 
				where id= '456'; 
		`

	rows := testDbConn.QueryRow(sqlString)
	var articleId string
	err = rows.Scan(&articleId)
	if err != nil {
		t.Error(err)
	}

}

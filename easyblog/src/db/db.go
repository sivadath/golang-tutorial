package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type Article struct {
	ID int64	`json:"id"`
	Title string	`json:"title"`
	Content string	`json:"content"`
	Author string	`json:"author"`
}

type Blogger interface {
	GetArticle(id int64) ([]Article, error)
	PushArticle(article Article)(ID int64, err error)
	GetAllArticles()([]Article, error)
}

type SQLite3 struct {
	*sql.DB
}

func (sqlDB *SQLite3) GetArticle(id int64) (articleArr []Article,err error) {
	if sqlDB == nil ||sqlDB.DB == nil {
		return nil, errors.New("failed to find db")
	}
	rows, err := sqlDB.Query(fmt.Sprintf(`select * from articles where id=%d;`,id))
	if err != nil {
		log.Printf("Failed to fetch articles with id %d, %v", id, err)
	}
	for rows.Next() {
		a := Article{}
		err = rows.Scan(&a.ID, &a.Title, &a.Content, &a.Author)
		if err != nil {
			log.Printf("Failed to scan row. %v", err)
			continue
		}
		articleArr = append(articleArr, a)
	}
	if len(articleArr) == 0 {
		return nil, errors.New("no such resource found")
	}
	return
}

func (sqlDB *SQLite3) PushArticle(article Article) (ID int64, err error) {
	if sqlDB == nil ||sqlDB.DB == nil {
		return -1, errors.New("failed to find db")
	}
	result, err := sqlDB.Exec(fmt.Sprintf(`INSERT INTO articles(title, author, content) VALUES ("%s", "%s", "%s");`,article.Title,article.Author,article.Content))
	if err == nil {
		return result.LastInsertId()
	}
	return -1, err
}

func (sqlDB *SQLite3) GetAllArticles() (articleArr []Article, err error) {
	if sqlDB == nil ||sqlDB.DB == nil {
		return nil, errors.New("failed to find db")
	}
	rows, err := sqlDB.Query(`select * from articles;`)
	if err != nil {
		log.Printf("Failed to fetch articles. %v", err)
	}
	for rows.Next() {
		a := Article{}
		err = rows.Scan(&a.ID, &a.Title, &a.Content, &a.Author)
		if err != nil {
			log.Printf("Failed to scan row. %v", err)
			continue
		}
		articleArr = append(articleArr, a)
	}
	return
}

const dbName = "easyBlogs.db"

var SQLiteDB Blogger

func init() {
	var err error
	db := &SQLite3{}
	db.DB, err = createDatabase()
	SQLiteDB = db
	if err != nil {
		log.Fatal("Failed to create database: %v", err)
	}
}

func createDatabase() (sqlDB *sql.DB,err error) {
	//Removing any previous instances.
	_ = os.Remove(dbName)
	sqlDB, err = sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}
	_, err = sqlDB.Exec("create table if not exists articles(id integer NOT NULL PRIMARY KEY AUTOINCREMENT,title varchar(255), content BLOB, author varchar(255) )")
	if err != nil {
		errClose := sqlDB.Close()
		if errClose != nil {
			return sqlDB, errors.New(err.Error() + errClose.Error())
		} else {
			return sqlDB, err
		}
	}
	return
}

func NewBlogger() Blogger{
	return SQLiteDB
}










/*type dummy struct {

}

func (*dummy) GetArticle(id int64) ([]Article, error) {
	return []Article{
		{ID: id,
			Title:   "MyTitle",
			Content: "MyContent",
			Author:  "Me",
		},
	}, nil
}

func (*dummy) PushArticle(article Article) (ID int64, err error) {
	return 101
}

func (*dummy) GetAllArticles() ([]Article, error) {
	return []Article{
		{
			ID:101,
			Title:"MyTitle1",
			Content:"MyContent1",
			Author:"Me",
		},
		{
			ID:102,
			Title:"MyTitle2",
			Content:"MyContent2",
			Author:"Me",
		},
	}, nil
}*/
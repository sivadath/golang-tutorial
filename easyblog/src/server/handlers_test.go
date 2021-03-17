package server

import (
	"db"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type BloggerMock struct {
	Articles []db.Article
	DummyPushArticle func()(int64, error)
}

func (BM *BloggerMock) GetArticle(id int64) (arr []db.Article,err error) {
	for _, article := range BM.Articles {
		if article.ID == id {
			arr = append(arr, article)
		}
	}
	if len(arr) == 0 {
		return nil, errors.New("resource not found")
	}
	return
}

func (BM *BloggerMock) PushArticle(article db.Article) (ID int64, err error) {
	return BM.DummyPushArticle()
}

func (BM *BloggerMock) GetAllArticles() ([]db.Article, error) {
	if len(BM.Articles) == 0 {
		return nil, errors.New("resource not found")
	}
	return BM.Articles, nil
}

func TestCreateArticle(t *testing.T) {
	bm := BloggerMock{}
	db.SQLiteDB = &bm

	bm.DummyPushArticle = func() (ID int64, err error) {
		return 1, nil
	}
	req, err := http.NewRequest("POST", "/article", strings.NewReader(`{"title": "Hello World","content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.","author": "John"}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateArticle)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	expected := `{"status": 201,"message": Success,"data": {"id": 1}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetArticleById(t *testing.T) {
	bm := BloggerMock{}
	db.SQLiteDB = &bm

	bm.Articles = []db.Article{{ID:3,Title:"Title",Content:"sample",Author:"nobody"},}
	req, err := http.NewRequest("GET", "/articles/3", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetArticleById)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"status": 200,"message": Success,"data": [[{"id":3,"title":"Title","content":"sample","author":"nobody"}]]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetAllArticles(t *testing.T) {
	bm := BloggerMock{}
	db.SQLiteDB = &bm

	bm.Articles = []db.Article{{ID:1,Title:"Title",Content:"sample",Author:"nobody"},}
	req, err := http.NewRequest("GET", "/articles", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllArticles)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"status": 200,"message": Success,"data": [{"id":1,"title":"Title","content":"sample","author":"nobody"}]}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	if len(alist) != len(articleList) {
		t.Fail()
	}

	for i, v:= range alist {
		if v.Content != articleList[i].Content || v.ID != articleList[i].ID || v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}

func TestCreateNewArticle(t *testing.T) {
	originalLength := len(getAllArticles())

	a, err := createNewArticle("New test title", "New test content")

	allArticles := getAllArticles()
	newLength := len(allArticles)

	if err != nil || newLength != originalLength + 1 || a.Title != "New test title" || a.Content != "New test content" {
		t.Fail()
	}
}
package goshopify

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"testing"
)

func articleTests(t *testing.T, article Article) {
	// Check that ID is assigned to the returned product
	var expectedInt int64 = 555885330570
	if article.ID != expectedInt {
		t.Errorf("Article.ID returned %+v, expected %+v", article.ID, expectedInt)
	}
}

func TestArticleCount(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", fmt.Sprintf("https://fooshop.myshopify.com/%s/blogs/%v/articles/count.json", client.pathPrefix, 77673496714),
		httpmock.NewStringResponder(200, `{"count": 3}`))

	cnt, err := client.Article.GetCountByBlogID(77673496714, nil)
	if err != nil {
		t.Errorf("Article.Count returned error: %v", err)
	}

	expected := 3
	if cnt != expected {
		t.Errorf("Article.Count returned %d, expected %d", cnt, expected)
	}
}

func TestArticleCreate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("POST", fmt.Sprintf("https://fooshop.myshopify.com/%s/blogs/%v/articles.json", client.pathPrefix, 77604257930),
		httpmock.NewBytesResponder(200, loadFixture("article.json")))

	article := Article{
		Title:             "hihihi",
		BodyHtml:          "",
		BlogID:            77604257930,
		Author:            "tu tran",
		UserID:            71812808842,
		SummaryHtml:       "",
		TemplateSuffix:    "",
		Handle:            "hihihi",
		Tags:              "",
		AdminGraphqlAPIID: "gid://shopify/OnlineStoreArticle/555885330570",
		Image: &ImageArticle{
			Alt:    "",
			Width:  720,
			Height: 960,
			Src:    "https://cdn.shopify.com/s/files/1/0552/8403/9818/articles/FED0BC96-DA34-4ECD-AE51-7B4C1EA06CE3.jpg?v=1636002218",
		},
	}

	returned, err := client.Article.Create(77604257930, &article)
	if err != nil {
		t.Errorf("Article.Create returned error: %v", err)
	}

	articleTests(t, *returned)
}

func TestArticleUpdate(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("PUT", fmt.Sprintf("https://fooshop.myshopify.com/%s/blogs/%v/articles/%v.json", client.pathPrefix, 77604257930, 555885330570),
		httpmock.NewBytesResponder(200, loadFixture("article.json")))

	article := Article{
		ID:       555885330570,
		BodyHtml: "Alo alo",
	}

	returned, err := client.Article.Update(77604257930, 555885330570, &article)
	if err != nil {
		t.Errorf("Article.Update returned error: %v", err)
	}

	articleTests(t, *returned)
}

func TestArticleDelete(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://fooshop.myshopify.com/%s/blogs/%v/articles/%v.json", client.pathPrefix, 77604257930, 555885330570),
		httpmock.NewStringResponder(200, "{}"))

	err := client.Article.Delete(77604257930, 555885330570)
	if err != nil {
		t.Errorf("Article.Delete returned error: %v", err)
	}
}

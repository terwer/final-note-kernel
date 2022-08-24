package controller

import (
	"encoding/json"
	"github.com/88250/gulu"
	"github.com/terwer/final-note-kernel/model"
	"github.com/terwer/final-note-kernel/service"
	"github.com/terwer/final-note-kernel/util"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Logger
var logger = gulu.Log.NewLogger(os.Stdout)

// GetArticlesAction gets articles.
func GetArticlesAction(r *http.Request) *string {
	result := gulu.Ret.NewResult()

	// BIDStr := util.GetSession(r, "BID")
	// BID, _ := strconv.ParseUint(*BIDStr, 0, 64)
	BID := uint64(1)
	articleModels, pagination := service.Article.ConsoleGetArticles(r.URL.Query().Get("key"), util.GetPage(r), BID)

	blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, BID)

	var articles []*ConsoleArticle
	for _, articleModel := range articleModels {
		var consoleTags []*ConsoleTag
		tagStrs := strings.Split(articleModel.Tags, ",")
		for _, tagStr := range tagStrs {
			consoleTag := &ConsoleTag{
				Title: tagStr,
				URL:   blogURLSetting.Value + util.PathTags + "/" + tagStr,
			}
			consoleTags = append(consoleTags, consoleTag)
		}

		authorModel := service.User.GetUser(articleModel.AuthorID)
		author := &ConsoleAuthor{
			Name:      authorModel.Name,
			URL:       blogURLSetting.Value + util.PathAuthors + "/" + authorModel.Name,
			AvatarURL: authorModel.AvatarURL,
		}

		article := &ConsoleArticle{
			ID:           articleModel.ID,
			Author:       author,
			CreatedAt:    articleModel.CreatedAt.Format("2006-01-02"),
			Title:        articleModel.Title,
			Tags:         consoleTags,
			URL:          blogURLSetting.Value + articleModel.Path,
			Topped:       articleModel.Topped,
			ViewCount:    articleModel.ViewCount,
			CommentCount: articleModel.CommentCount,
		}

		articles = append(articles, article)
	}

	data := map[string]interface{}{}

	data["articles"] = articles
	data["pagination"] = pagination
	result.Data = data

	jsonData, _ := json.Marshal(result)
	ret := string(jsonData)
	return &ret
}

func AddArticleAction() *string {
	result := gulu.Ret.NewResult()

	var lastArticleID string

	const (
		articleRecordSize = 99
	)
	for i := 0; i < articleRecordSize; i++ {
		article := &model.Article{AuthorID: 1,
			Title:       "Test 文章" + strconv.Itoa(i),
			Tags:        "Tag1, 标签2",
			Content:     "正文部分",
			Topped:      false,
			Commentable: true,
			BlogID:      1,
		}

		if err := service.Article.AddArticle(article); nil != err {
			logger.Fatal("add article failed: " + err.Error())
		}

		lastArticleID = strconv.FormatUint(article.ID, 10)
	}

	result.Data = lastArticleID

	jsonData, _ := json.Marshal(result)
	ret := string(jsonData)
	return &ret
}

func UpdateArticleAction(lastArticleID *string) *string {
	result := gulu.Ret.NewResult()

	var flag bool

	AID, _ := strconv.ParseUint(*lastArticleID, 0, 64)
	updatedTitle := "Updated title"
	article := service.Article.ConsoleGetArticle(AID)
	article.Title = updatedTitle
	if err := service.Article.UpdateArticle(article); nil != err {
		logger.Fatal("update article failed: " + err.Error())
		flag = false
	} else {
		// article = service.Article.ConsoleGetArticle(AID)
		flag = true
	}

	result.Data = flag

	jsonData, _ := json.Marshal(result)
	ret := string(jsonData)
	return &ret
}

func RemoveArticleAction(lastArticleID *string) *string {
	result := gulu.Ret.NewResult()

	var flag bool

	AID, _ := strconv.ParseUint(*lastArticleID, 0, 64)
	article := service.Article.ConsoleGetArticle(AID)
	if nil == article {
		flag = false
	} else if err := service.Article.RemoveArticle(AID, 1); nil != err {
		flag = false
	} else {
		article = service.Article.ConsoleGetArticle(AID)
		if nil != article {
			flag = false
			logger.Fatal("remove article failed")
		} else {
			flag = true
		}
	}

	result.Data = flag

	jsonData, _ := json.Marshal(result)
	ret := string(jsonData)
	return &ret
}

func GetArticleAction(lastArticleID *string) *string {
	result := gulu.Ret.NewResult()

	AID, _ := strconv.ParseUint(*lastArticleID, 0, 64)
	articleModel := service.Article.ConsoleGetArticle(AID)

	// BIDStr := util.GetSession(r, "BID")
	// BID, _ := strconv.ParseUint(*BIDStr, 0, 64)
	BID := uint64(1)
	blogURLSetting := service.Setting.GetSetting(model.SettingCategoryBasic, model.SettingNameBasicBlogURL, BID)
	var consoleTags []*ConsoleTag
	tagStrs := strings.Split(articleModel.Tags, ",")
	for _, tagStr := range tagStrs {
		consoleTag := &ConsoleTag{
			Title: tagStr,
			URL:   blogURLSetting.Value + util.PathTags + "/" + tagStr,
		}
		consoleTags = append(consoleTags, consoleTag)
	}

	authorModel := service.User.GetUser(articleModel.AuthorID)
	author := &ConsoleAuthor{
		Name:      authorModel.Name,
		URL:       blogURLSetting.Value + util.PathAuthors + "/" + authorModel.Name,
		AvatarURL: authorModel.AvatarURL,
	}

	article := &ConsoleArticle{
		ID:           articleModel.ID,
		Author:       author,
		CreatedAt:    articleModel.CreatedAt.Format("2006-01-02"),
		Title:        articleModel.Title,
		Tags:         consoleTags,
		URL:          blogURLSetting.Value + articleModel.Path,
		Topped:       articleModel.Topped,
		ViewCount:    articleModel.ViewCount,
		CommentCount: articleModel.CommentCount,
	}

	result.Data = article

	jsonData, _ := json.Marshal(result)
	ret := string(jsonData)
	return &ret
}

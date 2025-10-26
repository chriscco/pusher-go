package function

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"pusherGo/domain"
	"pusherGo/global"
	"strconv"
	"time"
)

func getNews(request *domain.NewsRequest) (*domain.News, error) {
	var news = &domain.News{}
	news.Category = request.Category

	if global.Configs == nil {
		return nil, fmt.Errorf("configs not initialized")
	}

	params := url.Values{}
	params.Add("category", string(request.Category))
	params.Add("country", string(request.Country))
	params.Add("max", strconv.Itoa(request.Max))
	params.Add("apikey", global.Configs.GNews.ApiKey)

	reqURL := global.Configs.GNews.Endpoint + params.Encode()

	log.Printf("Request gnews: %v\n", reqURL)

	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, news); err != nil {
		return nil, err
	}
	return news, nil
}

func GetNews() (*domain.NewsResponse, error) {
	var resp = &domain.NewsResponse{}
	resp.News = []*domain.News{}

	for _, category := range domain.Categories {
		for _, country := range domain.Countries {
			var req = &domain.NewsRequest{
				Category: category,
				Country:  country,
				Max:      global.Configs.GNews.MaxResults,
			}
			if country == domain.Us {
				req.Lang = domain.En
			}
			news, err := getNews(req)
			if err != nil {
				return nil, err
			}

			log.Printf("Received news: %v\n", news)
			if len(news.Articles) == 0 {
				continue
			}

			resp.News = append(resp.News, news)
		}
		time.Sleep(time.Second * 5)
	}

	return resp, nil
}

func FormatNews(news *domain.NewsResponse) (string, error) {
	data, err := json.MarshalIndent(news.News, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

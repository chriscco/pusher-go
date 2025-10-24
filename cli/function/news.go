package function

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"pusherGo/domain"
	"pusherGo/global"
	"strconv"
)

func getNews(request *domain.NewsRequest) (*domain.News, error) {
	var news = &domain.News{}
	news.Category = request.Category
	news.Country = request.Country

	if global.Configs == nil {
		return nil, fmt.Errorf("configs not initialized")
	}

	params := url.Values{}
	params.Add("category", string(request.Category))
	params.Add("lang", string(request.Lang))
	params.Add("country", string(request.Country))
	params.Add("max", strconv.Itoa(request.Max))
	params.Add("apikey", global.Configs.GNews.ApiKey)

	reqURL := global.Configs.GNews.Endpoint + params.Encode()

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
	resp.News = make([]*domain.News, 0)

	for _, category := range domain.Categories {
		for _, country := range domain.Countries {
			news, err := getNews(&domain.NewsRequest{
				Category: category,
				Country:  country,
				Max:      10,
				Lang:     domain.En,
			})
			if err != nil {
				return nil, err
			}

			resp.News = append(resp.News, news)
		}
	}

	return resp, nil
}

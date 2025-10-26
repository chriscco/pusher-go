package function

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"pusherGo/domain"
	"pusherGo/global"
)

func buildSystemPromptForNews() string {
	return "你是一个股票投资助手, 你将要理解输入的当日新闻并生成总结, 并在最后给出当前板块的投资建议, 如果新闻不足以得出足够有用的建议, 可以直接忽略" +
		"新闻会给出板块, 国家, 概览和标题, 无论语言种类, 你的输出都应该是简体中文\n 你的回答应该遵循下面的格式, 其中使用bullet points区分不同板块的不同新闻 \n" +
		"## 标题 \n" +
		"## 总结 \n" +
		"## 建议 \n"
}

func buildPromptForStock(stock string) string {
	return fmt.Sprintf(
		"你是一个股票投资助手, 你将要理解输入的当日新闻并生成总结, 并在最后给出当前板块的投资建议, "+
			"新闻会给出板块, 国家, 概览和标题, 无论语言种类, 你的输出都应该是简体中文\n 你的回答应该遵循下面的格式: \n"+
			"## 标题 \n"+
			"## 总结 \n"+
			"## 建议 \n"+
			"新闻内容: %s\n", stock,
	)
}

func CallModel(request *domain.ModelCallRequest) (*domain.ModelResponse, error) {
	prompt := buildSystemPromptForNews()

	req := domain.ModelRequest{
		Model: request.Model,
		Messages: []domain.Message{
			{Role: "system", Content: prompt},
			{Role: "user", Content: request.Content},
		},
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest("POST", "https://api.deepseek.com/chat/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+global.Configs.Model.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Printf("Received model response: %v\n", string(body))

	var chatResp *domain.ModelResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return nil, err
	}

	if len(chatResp.Choices) == 0 && chatResp.Error != nil {
		errMsg, err := json.MarshalIndent(chatResp.Error, "", " ")
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("no response from model: %v", errMsg)
	}

	return &domain.ModelResponse{
		Answer: chatResp.Choices[0].Message.Content,
	}, nil
}

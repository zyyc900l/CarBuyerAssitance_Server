package utils

import (
	"CarBuyerAssitance/biz/service/model"
	"CarBuyerAssitance/config"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// OpenAIRequest 定义 OpenAI 请求体
type OpenAIRequest struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens,omitempty"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIResponse 定义响应体
type OpenAIResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// CallOpenAIWithConsult 调用 OpenAI API 进行购车咨询
func CallOpenAIWithConsult(ctx context.Context, consult *model.Consult) (*model.ConsultResult, error) {
	prompt := fmt.Sprintf(`根据以下购车需求，提供专业的购车建议：
预算范围: %s
偏好车型: %s
主要使用场景: %s
燃料类型偏好: %s
品牌偏好: %s

请分析用户需求并返回 JSON 格式的购车建议，格式如下：
{
  "analysis": "对用户需求的分析",
  "result": [
    {
      "image_url": "车辆图片URL（请在网上搜索合适的图片）",
      "car_name": "推荐汽车名字",
      "fuel_consumption": "油耗",
      "power": "动力",
      "seat": "座位数",
      "drive": "驱动方式",
      "recommended_reason": "推荐理由"
    }
  ],
  "proposal": "总的购车建议"
}
请严格按照上述 JSON 格式返回，不要包含任何额外说明或 Markdown 格式。`,
		consult.BudgetRange,
		consult.PreferredType,
		consult.UseCase,
		consult.FuelType,
		consult.BrandPreference)

	reqBody := OpenAIRequest{
		Model: config.OpenAI.ApiModel,
		Messages: []Message{{
			Role:    "user",
			Content: prompt,
		}},
		MaxTokens: 2000,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("序列化请求体失败: %w", err)
	}

	client := &http.Client{Timeout: 60 * time.Second}
	apiURL := config.OpenAI.ApiUrl
	req, err := http.NewRequestWithContext(ctx, "POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.OpenAI.ApiKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API 返回错误 %d: %s", resp.StatusCode, string(b))
	}

	var result OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}

	if len(result.Choices) == 0 || result.Choices[0].Message.Content == "" {
		return nil, fmt.Errorf("API 返回空内容")
	}

	jsonText := result.Choices[0].Message.Content

	// 清理 Markdown
	cleaned := strings.TrimSpace(jsonText)
	if strings.HasPrefix(cleaned, "```json") {
		cleaned = strings.TrimPrefix(cleaned, "```json")
		cleaned = strings.TrimSuffix(cleaned, "```")
	} else if strings.HasPrefix(cleaned, "```") {
		cleaned = strings.TrimPrefix(cleaned, "```")
		cleaned = strings.TrimSuffix(cleaned, "```")
	}
	cleaned = strings.TrimSpace(cleaned)

	var consultResult model.ConsultResult
	if err := json.Unmarshal([]byte(cleaned), &consultResult); err != nil {
		return nil, fmt.Errorf("解析 JSON 失败（原始响应: %q）: %w", jsonText, err)
	}

	return &consultResult, nil
}

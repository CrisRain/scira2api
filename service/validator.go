package service

import (
	"fmt"
	"scira2api/models"
	"slices"
)

// chatParamCheck 检查聊天请求参数的有效性
func (h *ChatHandler) chatParamCheck(request models.OpenAIChatCompletionsRequest) error {
	if request.Model == "" {
		return fmt.Errorf("model is required")
	}

	// 检查请求的模型是否在 ModelMapping 中定义为外部模型名
	// ModelMapping 的键是外部模型名
	var availableExternalModels []string
	for externalName := range h.config.GetModelMapping() { // 使用辅助函数获取 ModelMapping
		availableExternalModels = append(availableExternalModels, externalName)
	}

	if !slices.Contains(availableExternalModels, request.Model) {
		return fmt.Errorf("model '%s' is not supported. Available models: %v",
			request.Model, availableExternalModels)
	}

	if len(request.Messages) == 0 {
		return fmt.Errorf("messages is required")
	}

	// 验证消息内容
	for i, message := range request.Messages {
		if message.Role == "" {
			return fmt.Errorf("message[%d].role is required", i)
		}
		if message.Content == "" {
			return fmt.Errorf("message[%d].content is required", i)
		}
	}

	return nil
}

package dto_gpt

import "oneGPT/components"

type GeneralErrorResponse struct {
	Error    components.OpenAIError `json:"error"`
	Message  string                 `json:"message"`
	Msg      string                 `json:"msg"`
	Err      string                 `json:"err"`
	ErrorMsg string                 `json:"error_msg"`
	Header   struct {
		Message string `json:"message"`
	} `json:"header"`
	Response struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	} `json:"response"`
}

func (e GeneralErrorResponse) ToMessage() string {
	if e.Error.Message != "" {
		return e.Error.Message
	}
	if e.Message != "" {
		return e.Message
	}
	if e.Msg != "" {
		return e.Msg
	}
	if e.Err != "" {
		return e.Err
	}
	if e.ErrorMsg != "" {
		return e.ErrorMsg
	}
	if e.Header.Message != "" {
		return e.Header.Message
	}
	if e.Response.Error.Message != "" {
		return e.Response.Error.Message
	}
	return ""
}

type ChatCompletionResp struct {
	Id                string                 `json:"id,omitempty"`
	Object            string                 `json:"object,omitempty"`
	Created           int                    `json:"created,omitempty"`
	Model             string                 `json:"model,omitempty"`
	Choices           []ChatCompletionChoice `json:"choices,omitempty"` // 聊天完成选项的列表
	Usage             *ChatCompletionUsage   `json:"usage,omitempty"`
	SystemFingerprint string                 `json:"system_fingerprint,omitempty"`
	BaseResp          *struct {
		StatusCode int    `json:"status_code"`
		StatusMsg  string `json:"status_msg"`
	} `json:"base_resp,omitempty"` // minimax 错误
}

type ChatCompletionChoice struct {
	Index        int                    `json:"index"`
	FinishReason string                 `json:"finish_reason"`
	Message      *ChatCompletionMessage `json:"message,omitempty"`
	Delta        *ChatCompletionMessage `json:"delta,omitempty"`
	Usage        *ChatCompletionUsage   `json:"usage,omitempty"`
}

type EmbeddingResp struct {
	Object string          `json:"object"`
	Data   []EmbeddingData `json:"data"`
	Model  string          `json:"model"`
	Usage  EmbeddingUsage  `json:"usage"`
}

type EmbeddingData struct {
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
	Index     int       `json:"index"`
}

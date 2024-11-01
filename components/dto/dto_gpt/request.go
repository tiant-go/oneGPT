package dto_gpt

type ChatCompletionRequest struct {
	Model       string                  `json:"model"`
	Messages    []ChatCompletionMessage `json:"messages"`
	MaxTokens   int                     `json:"max_tokens,omitempty"`
	Temperature float64                 `json:"temperature,omitempty"`
	TopP        float64                 `json:"top_p,omitempty"`
	Stream      bool                    `json:"stream,omitempty"`
}

type EmbeddingRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
	User  string `json:"user,omitempty"`
}

package dto_gpt

type ChatCompletionRole string

const (
	ChatCompletionUser      ChatCompletionRole = "user"
	ChatCompletionSystem    ChatCompletionRole = "system"
	ChatCompletionAssistant ChatCompletionRole = "assistant"
)

type ChatCompletionUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ChatCompletionMessage struct {
	Role      ChatCompletionRole `json:"role"`
	Content   string             `json:"content"`
	ToolCalls *[]ToolCall        `json:"tool_calls,omitempty"`
}

type ToolCall struct {
	Index    int          `json:"index"`
	Id       string       `json:"id"`
	Type     string       `json:"type"` //function
	Function FunctionInfo `json:"function"`
}

type FunctionInfo struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type EmbeddingUsage struct {
	PromptTokens int `json:"prompt_tokens"`
	TotalTokens  int `json:"total_tokens"`
}

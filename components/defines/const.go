package defines

type GPTSource string

const (
	GPTSourceOpenAI   GPTSource = "openai"
	GPTSourceKimi     GPTSource = "kimi"
	GPTSourceBaiChuan GPTSource = "baichuan"
	GPTSourceQwen     GPTSource = "qwen"
	GPTSourceGlm      GPTSource = "glm"
	GPTSourceMiniMax  GPTSource = "minimax"
	GPTSourceSkylark  GPTSource = "skylark" // 字节豆包大模型 v3接口版本
	GPTSourceAzure    GPTSource = "azure"
)

type GPTMethodType string

const (
	MethodTypeChat      GPTMethodType = "chat"
	MethodTypeEmbedding GPTMethodType = "embedding"
)

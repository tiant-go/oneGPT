package gpt

import (
	"github.com/gin-gonic/gin"
	"github.com/tiant-go/golib/flow"
	"github.com/tiant-go/golib/pkg/errors"
	"github.com/tiant-go/golib/pkg/http"
	"oneGPT/api/gpt/source"
	"oneGPT/components/defines"
	"oneGPT/components/dto/dto_gpt"
)

type IGPT interface {
	flow.IApi
	Init(endpoint string, ak string, proxyUrl string)
	GetClient() *http.HttpClientConf
	GetPath(methodType defines.GPTMethodType, model string) string
	HandleRequestHeader() map[string]string
	HandleChatResponse(responseByte []byte) (*dto_gpt.ChatCompletionResp, error)
	HandleEmbeddingResponse(responseByte []byte) (*dto_gpt.EmbeddingResp, error)
}

func GenerateThirdGpt(ctx *gin.Context, GPTSource defines.GPTSource) (IGPT, error) {
	switch GPTSource {
	case defines.GPTSourceOpenAI, defines.GPTSourceKimi, defines.GPTSourceBaiChuan, defines.GPTSourceQwen:
		return flow.Create(ctx, new(source.CommonGPT)), nil
	case defines.GPTSourceAzure:
		return flow.Create(ctx, new(source.AzureGPT)), nil
	case defines.GPTSourceGlm:
		return flow.Create(ctx, new(source.GlmGPT)), nil
	case defines.GPTSourceMiniMax:
		return flow.Create(ctx, new(source.MiniMaxGPT)), nil
	case defines.GPTSourceSkylark:
		return flow.Create(ctx, new(source.SkylarkGPT)), nil
	default:
		return nil, errors.NewError(errors.SYSTEM_ERROR, "GPT来源不支持")
	}
}

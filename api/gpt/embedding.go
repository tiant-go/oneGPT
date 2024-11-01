package gpt

import (
	http2 "github.com/tiant-go/golib/pkg/http"
	"net/http"
	"oneGPT/components"
	"oneGPT/components/defines"
	"oneGPT/components/dto/dto_gpt"
)

func Embedding(gpt IGPT, req *dto_gpt.EmbeddingRequest) (resp *dto_gpt.EmbeddingResp, errO *components.OpenAIErrorWithCode) {
	path := gpt.GetPath(defines.MethodTypeEmbedding, "")
	header := gpt.HandleRequestHeader()
	requestOpt := http2.HttpRequestOptions{
		RequestBody: req,
		Encode:      http2.EncodeJson,
		Headers:     header,
	}
	httpResult, err := gpt.GetClient().HttpPost(gpt.GetCtx(), path, requestOpt)
	if err != nil {
		return nil, components.OpenAIErrorWrapper(err.Error(), "do request failed", http.StatusInternalServerError)
	}
	if httpResult.HttpCode != http.StatusOK {
		errA := RelayErrorHandler(httpResult)
		return nil, errA
	}
	resp, err = gpt.HandleEmbeddingResponse(httpResult.Response)
	if err != nil {
		return nil, components.OpenAIErrorWrapper(err.Error(), "handle response failed", http.StatusInternalServerError)
	}
	return
}

package source

import (
	"encoding/json"
	"github.com/tiant-go/golib/pkg/errors"
	"oneGPT/components/defines"
	"oneGPT/components/dto/dto_gpt"
)

type MiniMaxGPT struct {
	CommonGPT
}

func (gpt *MiniMaxGPT) GetPath(methodType defines.GPTMethodType, model string) string {
	switch methodType {
	case defines.MethodTypeEmbedding:
		return "/v1/embeddings"
	case defines.MethodTypeChat:
		return "/v1/text/chatcompletion_v2"
	default:
		return "/v1/text/chatcompletion_v2"
	}
}

func (gpt *MiniMaxGPT) HandleChatResponse(responseByte []byte) (resp *dto_gpt.ChatCompletionResp, err error) {
	err = json.Unmarshal(responseByte, &resp)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp != nil && resp.BaseResp.StatusCode != 0 {
		return nil, errors.NewError(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp, nil
}

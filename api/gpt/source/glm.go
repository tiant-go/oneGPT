package source

import (
	"oneGPT/components/defines"
)

type GlmGPT struct {
	CommonGPT
}

func (gpt *GlmGPT) GetPath(methodType defines.GPTMethodType, model string) string {
	switch methodType {
	case defines.MethodTypeEmbedding:
		return "/api/paas/v4/embeddings"
	case defines.MethodTypeChat:
		return "/api/paas/v4/chat/completions"
	default:
		return "/api/paas/v4/chat/completions"
	}
}

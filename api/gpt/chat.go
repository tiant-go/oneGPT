package gpt

import (
	"encoding/json"
	"fmt"
	http2 "github.com/tiant-go/golib/pkg/http"
	"net/http"
	"oneGPT/components"
	"oneGPT/components/defines"
	"oneGPT/components/dto/dto_gpt"
	"strconv"
)

func ChatCompletion(gpt IGPT, req *dto_gpt.ChatCompletionRequest) (resp *dto_gpt.ChatCompletionResp, errO *components.OpenAIErrorWithCode) {
	path := gpt.GetPath(defines.MethodTypeChat, req.Model)
	header := gpt.HandleRequestHeader()
	requestOpt := http2.HttpRequestOptions{
		RequestBody: req,
		Encode:      http2.EncodeJson,
		Headers:     header,
	}
	if req.Stream == true {
		err := gpt.GetClient().HttpPostStream(gpt.GetCtx(), path, requestOpt, func(data string) error {
			if len(data) < 6 { // ignore blank line or wrong format
				return nil
			}
			if data[:6] != "data: " && data[:6] != "[DONE]" {
				return nil
			}
			if data[:6] == "data: " {
				tmpData := data[6:]
				resp, err := gpt.HandleChatResponse([]byte(tmpData))
				if err != nil {
					return err
				}
				respStr, _ := json.Marshal(resp)
				fmt.Println(string(respStr))
			}
			return nil
		})
		if err != nil {
			return nil, components.OpenAIErrorWrapper(err.Error(), "do request failed", http.StatusInternalServerError)
		}
	} else {
		httpResult, err := gpt.GetClient().HttpPost(gpt.GetCtx(), path, requestOpt)
		if err != nil {
			return nil, components.OpenAIErrorWrapper(err.Error(), "do request failed", http.StatusInternalServerError)
		}
		if httpResult.HttpCode != http.StatusOK {
			errA := RelayErrorHandler(httpResult)
			return nil, errA
		}
		resp, err = gpt.HandleChatResponse(httpResult.Response)
		if err != nil {
			return nil, components.OpenAIErrorWrapper(err.Error(), "handle response failed", http.StatusInternalServerError)
		}

	}
	return
}

func RelayErrorHandler(respResult *http2.HttpResult) (errWithStatusCode *components.OpenAIErrorWithCode) {
	errWithStatusCode = &components.OpenAIErrorWithCode{
		StatusCode: respResult.HttpCode,
		Error: components.OpenAIError{
			Type:  "upstream_error",
			Code:  "bad_response_status_code",
			Param: strconv.Itoa(respResult.HttpCode),
		},
	}

	var errResponse dto_gpt.GeneralErrorResponse
	err := json.Unmarshal(respResult.Response, &errResponse)
	if err != nil {
		return
	}
	if errResponse.Error.Message != "" {
		// OpenAI format error, so we override the default one
		errWithStatusCode.Error = errResponse.Error
	} else {
		errWithStatusCode.Error.Message = errResponse.ToMessage()
	}
	if errWithStatusCode.Error.Message == "" {
		errWithStatusCode.Error.Message = fmt.Sprintf("bad response status code %d", respResult.HttpCode)
	}
	return
}

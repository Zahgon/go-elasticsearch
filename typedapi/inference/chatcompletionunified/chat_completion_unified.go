package chatcompletionunified

import (
	gobytes "bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

const (
	inferenceidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ChatCompletionUnified struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	inferenceid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewChatCompletionUnified func(inferenceid string) *ChatCompletionUnified

func NewChatCompletionUnifiedFunc(tp elastictransport.Interface) NewChatCompletionUnified {
	_ = "STUB: not implemented"
	return *new(NewChatCompletionUnified)
}

func New(tp elastictransport.Interface) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) Raw(raw io.Reader) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) Request(req *Request) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ChatCompletionUnified) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r ChatCompletionUnified) Do(providedCtx context.Context) (Response, error) {
	_ = "STUB: not implemented"
	return *new(Response), nil
}

func (r *ChatCompletionUnified) Header(key, value string) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) _inferenceid(inferenceid string) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) Timeout(duration string) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) ErrorTrace(errortrace bool) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) FilterPath(filterpaths ...string) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) Human(human bool) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) Pretty(pretty bool) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) MaxCompletionTokens(maxcompletiontokens int64) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) Messages(messages ...types.MessageVariant) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) MessagesValues(messagesvalues []types.Message) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) Model(model string) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) Reasoning(reasoning types.ReasoningVariant) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) Stop(stops ...string) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) Temperature(temperature float32) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) ToolChoice(completiontooltype types.CompletionToolTypeVariant) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) Tools(tools ...types.CompletionToolVariant) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) ToolsValues(toolsvalues []types.CompletionTool) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

func (r *ChatCompletionUnified) TopP(topp float32) *ChatCompletionUnified {
	_ = "STUB: not implemented"
	return nil
}

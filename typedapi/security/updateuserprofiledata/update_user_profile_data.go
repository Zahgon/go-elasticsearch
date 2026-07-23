package updateuserprofiledata

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

const (
	uidMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateUserProfileData struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	uid string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewUpdateUserProfileData func(uid string) *UpdateUserProfileData

func NewUpdateUserProfileDataFunc(tp elastictransport.Interface) NewUpdateUserProfileData {
	_ = "STUB: not implemented"
	return *new(NewUpdateUserProfileData)
}

func New(tp elastictransport.Interface) *UpdateUserProfileData {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateUserProfileData) Raw(raw io.Reader) *UpdateUserProfileData {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateUserProfileData) Request(req *Request) *UpdateUserProfileData {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateUserProfileData) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateUserProfileData) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r UpdateUserProfileData) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *UpdateUserProfileData) Header(key, value string) *UpdateUserProfileData {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateUserProfileData) _uid(uid string) *UpdateUserProfileData {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateUserProfileData) IfSeqNo(sequencenumber string) *UpdateUserProfileData {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateUserProfileData) IfPrimaryTerm(ifprimaryterm string) *UpdateUserProfileData {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateUserProfileData) Refresh(refresh refresh.Refresh) *UpdateUserProfileData {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateUserProfileData) ErrorTrace(errortrace bool) *UpdateUserProfileData {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateUserProfileData) FilterPath(filterpaths ...string) *UpdateUserProfileData {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateUserProfileData) Human(human bool) *UpdateUserProfileData {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateUserProfileData) Pretty(pretty bool) *UpdateUserProfileData {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateUserProfileData) Data(data map[string]json.RawMessage) *UpdateUserProfileData {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateUserProfileData) AddDatum(key string, value json.RawMessage) *UpdateUserProfileData {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateUserProfileData) Labels(labels map[string]json.RawMessage) *UpdateUserProfileData {
	_ = "STUB: not implemented"
	return nil
}

func (r *UpdateUserProfileData) AddLabel(key string, value json.RawMessage) *UpdateUserProfileData {
	_ = "STUB: not implemented"
	return nil
}

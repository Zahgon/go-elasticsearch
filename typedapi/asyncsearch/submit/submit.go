package submit

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/searchtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/suggestmode"
)

const (
	indexMask = iota + 1
)

var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Submit struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

type NewSubmit func() *Submit

func NewSubmitFunc(tp elastictransport.Interface) NewSubmit {
	_ = "STUB: not implemented"
	return *new(NewSubmit)
}

func New(tp elastictransport.Interface) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) Raw(raw io.Reader) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) Request(req *Request) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Submit) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Submit) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Submit) Header(key, value string) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) Index(index string) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) WaitForCompletionTimeout(duration string) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) KeepAlive(duration string) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) KeepOnCompletion(keeponcompletion bool) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) AllowNoIndices(allownoindices bool) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) AllowPartialSearchResults(allowpartialsearchresults bool) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) Analyzer(analyzer string) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) AnalyzeWildcard(analyzewildcard bool) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) BatchedReduceSize(batchedreducesize string) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) CcsMinimizeRoundtrips(ccsminimizeroundtrips bool) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) DefaultOperator(defaultoperator operator.Operator) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) Df(df string) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) IgnoreThrottled(ignorethrottled bool) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) IgnoreUnavailable(ignoreunavailable bool) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) Lenient(lenient bool) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) MaxConcurrentShardRequests(maxconcurrentshardrequests int) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) Preference(preference string) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) RequestCache(requestcache bool) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) Routing(routings ...string) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) SearchType(searchtype searchtype.SearchType) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) SuggestField(field string) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) SuggestMode(suggestmode suggestmode.SuggestMode) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) SuggestSize(suggestsize string) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) SuggestText(suggesttext string) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) TypedKeys(typedkeys bool) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) RestTotalHitsAsInt(resttotalhitsasint bool) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) SourceExcludes_(fields ...string) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) SourceIncludes_(fields ...string) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) Q(q string) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) ErrorTrace(errortrace bool) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) FilterPath(filterpaths ...string) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) Human(human bool) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) Pretty(pretty bool) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) Aggregations(aggregations map[string]types.Aggregations) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) AddAggregation(key string, value types.AggregationsVariant) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) Collapse(collapse types.FieldCollapseVariant) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) DocvalueFields(docvaluefields ...types.FieldAndFormatVariant) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) DocvalueFieldsValues(docvaluefieldsvalues []types.FieldAndFormat) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) Explain(explain bool) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) Ext(ext map[string]json.RawMessage) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) AddExt(key string, value json.RawMessage) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) Fields(fields ...types.FieldAndFormatVariant) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) FieldsValues(fieldsvalues []types.FieldAndFormat) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) From(from int) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) Highlight(highlight types.HighlightVariant) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) IndicesBoost(indicesboost []map[string]types.Float64) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) Knn(knns ...types.KnnSearchVariant) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) MinScore(minscore types.Float64) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) Pit(pit types.PointInTimeReferenceVariant) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) PostFilter(postfilter types.QueryVariant) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) Profile(profile bool) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) ProjectRouting(projectrouting string) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) Query(query types.QueryVariant) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) Rescore(rescores ...types.RescoreVariant) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) ScriptFields(scriptfields map[string]types.ScriptField) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) AddScriptField(key string, value types.ScriptFieldVariant) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) SearchAfter(sortresults ...types.FieldValueVariant) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) SearchAfterValues(sortresultsvalues []types.FieldValue) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) SeqNoPrimaryTerm(seqnoprimaryterm bool) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) Size(size int) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) Slice(slice types.SlicedScrollVariant) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) Sort(sorts ...types.SortCombinationsVariant) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) SortValues(sortvalues []types.SortCombinations) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) Source_(sourceconfig types.SourceConfigVariant) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) Stats(stats ...string) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) StoredFields(fields ...string) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) Suggest(suggest types.SuggesterVariant) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) TerminateAfter(terminateafter int64) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) Timeout(timeout string) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) TrackScores(trackscores bool) *Submit { _ = "STUB: not implemented"; return nil }

func (r *Submit) TrackTotalHits(trackhits types.TrackHitsVariant) *Submit {
	_ = "STUB: not implemented"
	return nil
}

func (r *Submit) Version(version bool) *Submit { _ = "STUB: not implemented"; return nil }

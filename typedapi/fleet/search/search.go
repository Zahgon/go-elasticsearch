package search

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

type Search struct {
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

type NewSearch func(index string) *Search

func NewSearchFunc(tp elastictransport.Interface) NewSearch {
	_ = "STUB: not implemented"
	return *new(NewSearch)
}

func New(tp elastictransport.Interface) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Raw(raw io.Reader) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Request(req *Request) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) HttpRequest(ctx context.Context) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Search) Perform(providedCtx context.Context) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r Search) Do(providedCtx context.Context) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Search) Header(key, value string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) _index(index string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) AllowNoIndices(allownoindices bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Analyzer(analyzer string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) AnalyzeWildcard(analyzewildcard bool) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) BatchedReduceSize(batchedreducesize string) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) CcsMinimizeRoundtrips(ccsminimizeroundtrips bool) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) DefaultOperator(defaultoperator operator.Operator) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) Df(df string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) IgnoreThrottled(ignorethrottled bool) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) IgnoreUnavailable(ignoreunavailable bool) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) Lenient(lenient bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) MaxConcurrentShardRequests(maxconcurrentshardrequests int) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) Preference(preference string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) PreFilterShardSize(prefiltershardsize string) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) RequestCache(requestcache bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Routing(routings ...string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Scroll(duration string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) SearchType(searchtype searchtype.SearchType) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) SuggestField(field string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) SuggestMode(suggestmode suggestmode.SuggestMode) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) SuggestSize(suggestsize string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) SuggestText(suggesttext string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) TypedKeys(typedkeys bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) RestTotalHitsAsInt(resttotalhitsasint bool) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) SourceExcludes_(fields ...string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) SourceIncludes_(fields ...string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Q(q string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) WaitForCheckpoints(waitforcheckpoints ...int64) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) AllowPartialSearchResults(allowpartialsearchresults bool) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) ErrorTrace(errortrace bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) FilterPath(filterpaths ...string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Human(human bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Pretty(pretty bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Aggregations(aggregations map[string]types.Aggregations) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) AddAggregation(key string, value types.AggregationsVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) Collapse(collapse types.FieldCollapseVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) DocvalueFields(docvaluefields ...types.FieldAndFormatVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) DocvalueFieldsValues(docvaluefieldsvalues []types.FieldAndFormat) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) Explain(explain bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Ext(ext map[string]json.RawMessage) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) AddExt(key string, value json.RawMessage) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) Fields(fields ...types.FieldAndFormatVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) FieldsValues(fieldsvalues []types.FieldAndFormat) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) From(from int) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Highlight(highlight types.HighlightVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) IndicesBoost(indicesboost []map[string]types.Float64) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) MinScore(minscore types.Float64) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Pit(pit types.PointInTimeReferenceVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) PostFilter(postfilter types.QueryVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) Profile(profile bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Query(query types.QueryVariant) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Rescore(rescores ...types.RescoreVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) ScriptFields(scriptfields map[string]types.ScriptField) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) AddScriptField(key string, value types.ScriptFieldVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) SearchAfter(sortresults ...types.FieldValueVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) SearchAfterValues(sortresultsvalues []types.FieldValue) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) SeqNoPrimaryTerm(seqnoprimaryterm bool) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) Size(size int) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Slice(slice types.SlicedScrollVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) Sort(sorts ...types.SortCombinationsVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) SortValues(sortvalues []types.SortCombinations) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) Source_(sourceconfig types.SourceConfigVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) Stats(stats ...string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) StoredFields(fields ...string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) Suggest(suggest types.SuggesterVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) TerminateAfter(terminateafter int64) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) Timeout(timeout string) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) TrackScores(trackscores bool) *Search { _ = "STUB: not implemented"; return nil }

func (r *Search) TrackTotalHits(trackhits types.TrackHitsVariant) *Search {
	_ = "STUB: not implemented"
	return nil
}

func (r *Search) Version(version bool) *Search { _ = "STUB: not implemented"; return nil }

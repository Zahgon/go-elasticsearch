package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _textSimilarityReranker struct {
	v *types.TextSimilarityReranker
}

func NewTextSimilarityReranker(field string, inferencetext string, retriever types.RetrieverContainerVariant) *_textSimilarityReranker {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textSimilarityReranker) ChunkRescorer(chunkrescorer types.ChunkRescorerVariant) *_textSimilarityReranker {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textSimilarityReranker) Field(field string) *_textSimilarityReranker {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textSimilarityReranker) InferenceId(inferenceid string) *_textSimilarityReranker {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textSimilarityReranker) InferenceText(inferencetext string) *_textSimilarityReranker {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textSimilarityReranker) RankWindowSize(rankwindowsize int) *_textSimilarityReranker {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textSimilarityReranker) Retriever(retriever types.RetrieverContainerVariant) *_textSimilarityReranker {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textSimilarityReranker) Filter(filters ...types.QueryVariant) *_textSimilarityReranker {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textSimilarityReranker) MinScore(minscore float32) *_textSimilarityReranker {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textSimilarityReranker) Name_(name_ string) *_textSimilarityReranker {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textSimilarityReranker) RetrieverContainerCaster() *types.RetrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_textSimilarityReranker) TextSimilarityRerankerCaster() *types.TextSimilarityReranker {
	_ = "STUB: not implemented"
	return nil
}

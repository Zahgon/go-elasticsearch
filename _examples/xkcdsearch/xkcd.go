package xkcdsearch

const baseURL = "https://xkcd.com"

type Document struct {
	ID        string `json:"id"`
	ImageURL  string `json:"image_url"`
	Published string `json:"published"`

	Title      string `json:"title"`
	Alt        string `json:"alt"`
	Transcript string `json:"transcript"`
	Link       string `json:"link,omitempty"`
	News       string `json:"news,omitempty"`
}

func (d *Document) URL() string { _ = "STUB: not implemented"; return "" }

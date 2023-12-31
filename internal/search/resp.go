package search

import "encoding/json"

type TraceMoeResponse struct {
	FrameCount int       `json:"frameCount"`
	Error      string    `json:"error"`
	Result     []*Result `json:"result"`
}

type Result struct {
	Anilist    *Anilist `json:"anilist"`
	Filename   string   `json:"filename"`
	Episode    Episode  `json:"episode"`
	From       float64  `json:"from"`
	To         float64  `json:"to"`
	Similarity float64  `json:"similarity"`
	Video      string   `json:"video"`
	Image      string   `json:"image"`
}

type Episode string

var _ json.Unmarshaler = (*Episode)(nil)

func (ep *Episode) UnmarshalJSON(data []byte) error {
	*ep = Episode(data)
	return nil
}

type Anilist struct {
	ID       int      `json:"id"`
	IDMal    int      `json:"idMal"`
	Title    *Title   `json:"title"`
	Synonyms []string `json:"synonyms"`
	IsAdult  bool     `json:"isAdult"`
}

type Title struct {
	Native  string `json:"native"`
	Romaji  string `json:"romaji"`
	English string `json:"english"`
}

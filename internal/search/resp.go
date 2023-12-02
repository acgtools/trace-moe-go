package search

type TraceMoeResponse struct {
	FrameCount int       `json:"frameCount"`
	Error      string    `json:"error"`
	Result     []*Result `json:"result"`
}

type Result struct {
	Anilist    int     `json:"anilist"`
	Filename   string  `json:"filename"`
	Episode    int     `json:"episode"`
	From       float64 `json:"from"`
	To         float64 `json:"to"`
	Similarity float64 `json:"similarity"`
	Video      string  `json:"video"`
	Image      string  `json:"image"`
}

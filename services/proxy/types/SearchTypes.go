package Types

type SearchResult struct {
	Title     string `json:"title"`
	VideoUrl  string `json:"videoUrl"`
	Thumbnail string `json:"thumbnail"`
	Channel   string `json:"channel"`
	Duration  int    `json:"duration"`
}
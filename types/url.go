package types

type URL struct {
	Url string
}

type ShortenResponse struct {
	ShortURL    string
	OriginalURL string
}

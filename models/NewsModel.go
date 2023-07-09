package models

type NewsResponseModel struct {
	Status     string `json:"status"`
	TotalHits  int    `json:"total_hits"`
	Page       int    `json:"page"`
	TotalPages int    `json:"total_pages"`
	PageSize   int    `json:"page_size"`
	Articles   []struct {
		Title                  string      `json:"title"`
		Author                 string      `json:"author"`
		PublishedDate          string      `json:"published_date"`
		PublishedDatePrecision string      `json:"published_date_precision"`
		Link                   string      `json:"link"`
		CleanURL               string      `json:"clean_url"`
		Excerpt                string      `json:"excerpt"`
		Summary                string      `json:"summary"`
		Rights                 string      `json:"rights"`
		Rank                   int         `json:"rank"`
		Topic                  string      `json:"topic"`
		Country                string      `json:"country"`
		Language               string      `json:"language"`
		Authors                interface{} `json:"authors"`
		Media                  string      `json:"media"`
		IsOpinion              bool        `json:"is_opinion"`
		TwitterAccount         string      `json:"twitter_account"`
		Score                  float64     `json:"_score"`
		ID                     string      `json:"_id"`
	} `json:"articles"`
	UserInput struct {
		Q                      string      `json:"q"`
		SearchIn               []string    `json:"search_in"`
		Lang                   interface{} `json:"lang"`
		NotLang                interface{} `json:"not_lang"`
		Countries              []string    `json:"countries"`
		NotCountries           interface{} `json:"not_countries"`
		From                   string      `json:"from"`
		To                     interface{} `json:"to"`
		RankedOnly             string      `json:"ranked_only"`
		FromRank               interface{} `json:"from_rank"`
		ToRank                 interface{} `json:"to_rank"`
		SortBy                 string      `json:"sort_by"`
		Page                   int         `json:"page"`
		Size                   int         `json:"size"`
		Sources                interface{} `json:"sources"`
		NotSources             interface{} `json:"not_sources"`
		Topic                  interface{} `json:"topic"`
		PublishedDatePrecision interface{} `json:"published_date_precision"`
	} `json:"user_input"`
}

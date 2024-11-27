package data

import "time"

type RelatedLinks struct {
	Previous string `json:"prev ,omitempty"`
	Next     string `json:"next ,omitempty"`
}

type Article struct {
	ID           int64        `json:"id"`
	CreatedAt    time.Time    `json:"-"`
	Title        string       `json:"title"`
	PublishDate  time.Time    `json:"publish_date"`
	ReadTime     ReadTime     `json:"read_time ,omitempty"`
	Authors      []string     `json:"authors"`
	Excerpt      string       `json:"excerpt"`
	Tags         []string     `json:"tags"`
	SourceURL    string       `json:"source_url"`
	RelatedLinks RelatedLinks `json:"related_links"`
	LastUpdated  time.Time    `json:"last_updated"`
	Version      int32        `json:"version"`
}

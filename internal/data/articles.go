package data

import "time"

type RelatedLinks struct {
	Previous string `json:"prev"`
	Next     string `json:"next"`
}

type Article struct {
	ID           int64        `json:"id"`
	CreatedAt    time.Time    `json:"created_at"`
	Title        string       `json:"title"`
	PublishDate  time.Time    `json:"publish_date"`
	ReadTime     int          `json:"read_time"`
	Authors      []string     `json:"authors"`
	Excerpt      string       `json:"excerpt"`
	Tags         []string     `json:"tags"`
	SourceURL    string       `json:"source_url"`
	RelatedLinks RelatedLinks `json:"related_links"`
	LastUpdated  time.Time    `json:"last_updated"`
	Version      int32        `json:"version"`
}

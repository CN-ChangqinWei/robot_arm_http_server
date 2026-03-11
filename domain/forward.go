package domain

type Forward struct {
	Topic       string   `json:"topic"`
	Subscribers []string `json:"subscribers"`
	Publisher   string   `json:"publisher"`
}

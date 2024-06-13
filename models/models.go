package models

type Feed struct {
	Nickname    string   `json:"nickname"`
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	UpdateURL   string   `json:"updatedURL"`
	Categories  []string `json:"categories"`
	Items       []Item   `json:"items"`
}

type Item struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Link    string `json:"link"`
}

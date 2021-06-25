package internal

type Ad struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Ads struct {
	Topic string `json:"topic"`
	Ads   []Ad   `json:"ads"`
}

type Data struct {
	Data []Ads `json:"data"`
}

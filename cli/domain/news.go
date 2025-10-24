package domain

type Category string
type Lang string
type Country string

const (
	Technology Category = "technology"
	Business   Category = "business"
)

const (
	Zh Lang = "zh"
	En Lang = "en"
)

const (
	Cn Country = "cn"
	Us Country = "us"
	Hk Country = "hk"
)

type News struct {
	Category    Category `json:"category"`
	Country     Country  `json:"country"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Lang        Lang     `json:"lang"`
}

var (
	Categories = []Category{Technology, Business}
	Langs      = []Lang{En}
	Countries  = []Country{Cn, Us, Hk}
)

type NewsRequest struct {
	Category Category `json:"category"`
	Lang     Lang     `json:"lang"`
	Country  Country  `json:"country"`
	Max      int      `json:"max"`
}

type NewsResponse struct {
	News []*News
}

package domain

type Category string
type Lang string
type Country string

const (
	Technology Category = "technology"
	Business   Category = "business"
	World      Category = "world"
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

type Source struct {
	Country Country `json:"country"`
}

type Articles struct {
	Source      Source `json:"source"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Lang        Lang   `json:"lang"`
	Url         string `json:"url"`
}

type News struct {
	Category Category    `json:"category"`
	Articles []*Articles `json:"articles"`
}

var (
	Categories = []Category{Technology, Business}
	Langs      = []Lang{En, Zh}
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

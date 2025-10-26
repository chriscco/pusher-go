package domain

type Category string
type Lang string
type Country string

const (
	Technology Category = "technology"
	Business   Category = "business"
	World      Category = "world"
	Nation     Category = "nation"
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

type Articles struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}

type News struct {
	Category Category    `json:"category"`
	Articles []*Articles `json:"articles"`
}

var (
	Categories = []Category{Business, Technology}
	Langs      = []Lang{En, Zh}
	Countries  = []Country{Cn, Us}
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

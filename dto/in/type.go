package in

type in struct {
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
	OrderBy string `json:"order_by"`
	Range   struct {
		Unit string `json:"unit"`
		From int64  `json:"from"`
		To   int64  `json:"to"`
	} `json:"range"`
}

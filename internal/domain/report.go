package domain

type Period struct {
	Name  int     `json:"name"`
	Value float32 `json:"value"`
}

type Report struct {
	Stock   Stock    `json:"stock"`
	Periods []Period `json:"period"`
}

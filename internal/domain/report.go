package domain

type Period struct {
	Name  int
	Value float32
}

type ReportStock struct {
	Stock   Stock
	Periods []Period
}

type Report struct {
	Stocks []ReportStock
}

func (r *Report) AddStock(reportStock ReportStock) {
	if r.Stocks == nil {
		r.Stocks = make([]ReportStock, 0)
	}

	r.Stocks = append(r.Stocks, reportStock)
}

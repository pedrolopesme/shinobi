package domain

import "github.com/shopspring/decimal"

type Period struct {
	Name  int
	Value decimal.Decimal
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

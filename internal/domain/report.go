package domain

import "github.com/shopspring/decimal"

type Period struct {
	Name  int
	Value decimal.Decimal
}

type ReportStock struct {
	Stock          Stock
	BaseValue      decimal.Decimal
	Periods        []Period
	Trend          string
	RecentVolumeMA int32
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

func (rs ReportStock) IsEmpty() bool {
	return rs.Stock.Symbol == "" || len(rs.Periods) == 0
}

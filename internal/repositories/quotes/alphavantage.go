package quotes

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"sort"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/pedrolopesme/shinobi/internal/domain"
// 	"github.com/pedrolopesme/shinobi/internal/domain/application"
// 	"go.uber.org/zap"
// )

// type AlphaVantageQuoteRepository struct {
// 	application application.Application
// }

// func NewAlphaVantageQuoteRepository(application application.Application) AlphaVantageQuoteRepository {
// 	return AlphaVantageQuoteRepository{
// 		application: application,
// 	}
// }

// func (a AlphaVantageQuoteRepository) GetQuotes(symbol string) ([]domain.Quote, error) {
// 	logger := a.application.Logger()
// 	logger.Info("Retrieving Quotes from AlphaVantage API", zap.String("symbol", symbol))

// 	defer func() {
// 		if r := recover(); r != nil {
// 			logger.Error("Impossible to get quotes", zap.String("symbol", symbol))
// 		}
// 	}()

// 	endpoint := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&outputsize=full&apikey=%s", symbol, a.application.AlphaVantageKey())
// 	resp, err := http.Get(endpoint)
// 	if err != nil {
// 		return nil, err
// 	}

// 	logger.Info("Parsing Data Points", zap.String("symbol", symbol))
// 	quotes := make([]domain.Quote, 0)

// 	buffer := new(strings.Builder)
// 	io.Copy(buffer, resp.Body)
// 	responseBody := buffer.String()

// 	var rawResult map[string]interface{}
// 	json.Unmarshal([]byte(responseBody), &rawResult)

// 	rawTimeSeries := rawResult["Time Series (Daily)"].(map[string]interface{})
// 	for index := range rawTimeSeries {
// 		date, err := time.Parse("2006-01-02", index)
// 		if err != nil {
// 			logger.Error("Impossible to parse date", zap.String("date", index), zap.Error(err), zap.String("symbol", symbol))
// 		}

// 		rawDataPoint := rawTimeSeries[index].(map[string]interface{})
// 		open, _ := strconv.ParseFloat(rawDataPoint["1. open"].(string), 32)
// 		high, _ := strconv.ParseFloat(rawDataPoint["2. high"].(string), 32)
// 		low, _ := strconv.ParseFloat(rawDataPoint["3. low"].(string), 32)
// 		close, _ := strconv.ParseFloat(rawDataPoint["4. close"].(string), 32)
// 		volume, _ := strconv.Atoi(rawDataPoint["5. volume"].(string))

// 		// quote := domain.Quote{
// 		// 	Date:   date,
// 		// 	Open:   float32(open),
// 		// 	High:   float32(high),
// 		// 	Low:    float32(low),
// 		// 	Close:  float32(close),
// 		// 	Volume: int32(volume),
// 		// }

// 		// quotes = append(quotes, quote)
// 	}

// 	// unfortunnately, I have to loop through the quotes twice
// 	// since all have to sort it by date
// 	sort.Slice(quotes, func(i, j int) bool {
// 		return quotes[i].Date.After(quotes[j].Date)
// 	})

// 	return quotes, nil
// }

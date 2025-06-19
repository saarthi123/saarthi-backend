package models

import "time"

type StockPrediction struct {
	Symbol     string `json:"symbol"`
	Prediction string `json:"prediction"`
	Confidence int    `json:"confidence"`
}

type TradeRequest struct {
	Symbol   string `json:"symbol" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

type Exchange struct {
	Name         string   `json:"name"`
	Code         string   `json:"code"`
	Country      string   `json:"country"`
	TradingHours string   `json:"trading_hours"`
	Holidays     []string `json:"holidays"`
	Region       string   `json:"region"`
	LogoUrl      string   `json:"logo_url"`
}

type MarketIndex struct {
	Symbol        string  `json:"symbol"`
	Name          string  `json:"name"`
	Price         float64 `json:"price"`
	Change        float64 `json:"change"`
	PercentChange float64 `json:"percent_change"`
}

type NewsArticle struct {
	Title          string `json:"title"`
	Source         string `json:"source"`
	Date           string `json:"date"`
	Sentiment      string `json:"sentiment"`
	SentimentEmoji string `json:"sentiment_emoji"`
}

type DummyTradingService struct {
	balance   float64
	portfolio map[string]int
}

func NewDummyTradingService() *DummyTradingService {
	return &DummyTradingService{
		balance:   10000.0,
		portfolio: make(map[string]int),
	}
}

func (d *DummyTradingService) GetPortfolio() map[string]int {
	return d.portfolio
}

func (d *DummyTradingService) GetBalance() float64 {
	return d.balance
}

func (d *DummyTradingService) Buy(symbol string, quantity int) bool {
	price := 100.0 // dummy price
	cost := price * float64(quantity)
	if d.balance < cost {
		return false
	}
	d.balance -= cost
	d.portfolio[symbol] += quantity
	return true
}

func (d *DummyTradingService) Sell(symbol string, quantity int) bool {
	if d.portfolio[symbol] < quantity {
		return false
	}
	price := 100.0 // dummy price
	d.portfolio[symbol] -= quantity
	d.balance += price * float64(quantity)
	return true
}

func (d *DummyTradingService) Reset() {
	d.balance = 10000.0
	d.portfolio = make(map[string]int)
}

type Trade struct {
	ID        int       `json:"id"`
  Symbol    string    `json:"symbol"` // ✅ this is the correct field name
	Ticker    string    `json:"ticker"` // <-- probably Ticker, not Symbol
	Qty       int       `json:"qty"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
	UserID    uint      `json:"user_id"` // Foreign key
	Quantity  int       `json:"quantity"`
	
}



type Holding struct {
Symbol       string  `json:"symbol"`
Name         string  `json:"name"`
Quantity     int     `json:"quantity"`
AvgPrice     float64 `json:"avg_price"`
CurrentPrice float64 `json:"current_price"`
ProfitLoss   float64 `json:"profit_loss" gorm:"-"`
Sector       string  `json:"sector"`
}



type Stock struct {
	Symbol  string    `json:"symbol"`
	Company string    `json:"company,omitempty"`
	Price   float64   `json:"price"`
	Change  float64   `json:"change"`
	Trend   []float64 `json:"trend"`
	AITag   string    `json:"ai_tag"`
	Quantity int    `json:"quantity"`
}


type TradeOrder struct {
    Symbol     string  `json:"symbol"`
    Quantity   int     `json:"quantity"`
    OrderType  string  `json:"orderType"`
    Price      float64 `json:"price"`
    Side       string  `json:"side"`
}



type AISuggestionResponse struct {
    Suggestion string `json:"suggestion"`
}


type PortfolioItem struct {
	Symbol        string  `json:"symbol"`
	Name          string  `json:"name"`
	Quantity      int     `json:"quantity"`
	AvgPrice      float64 `json:"avg_price"`
	CurrentPrice  float64 `json:"current_price"`
	ProfitLoss    float64 `json:"profit_loss"`
	Sector        string  `json:"sector"`
}


type AnalyticsData struct {
	Label string  `json:"label"`
	Value float64 `json:"value"`
}


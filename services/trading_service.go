package services

import (
    "sync"
    "github.com/saarthi123/saarthi-backend/models"
)

type DummyTradingService struct {
    mu        sync.Mutex
    balance   int
    portfolio map[string]int
}

func NewDummyTradingService() *DummyTradingService {
    return &DummyTradingService{
        balance:   100000,
        portfolio: make(map[string]int),
    }
}

func (s *DummyTradingService) GetBalance() int {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.balance
}

func (s *DummyTradingService) GetPortfolio() []models.PortfolioItem {
    s.mu.Lock()
    defer s.mu.Unlock()

    items := []models.PortfolioItem{}
    for sym, qty := range s.portfolio {
        items = append(items, models.PortfolioItem{Symbol: sym, Quantity: qty})
    }
    return items
}

func (s *DummyTradingService) Buy(symbol string, qty int) bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    cost := qty * 100
    if cost > s.balance {
        return false // insufficient funds
    }
    s.balance -= cost
    s.portfolio[symbol] += qty
    return true
}

func (s *DummyTradingService) Sell(symbol string, qty int) bool {
    s.mu.Lock()
    defer s.mu.Unlock()
    currentQty := s.portfolio[symbol]
    if qty > currentQty {
        return false // not enough stocks to sell
    }
    s.portfolio[symbol] -= qty
    s.balance += qty * 100
    if s.portfolio[symbol] == 0 {
        delete(s.portfolio, symbol)
    }
    return true
}

func (s *DummyTradingService) Reset() {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.balance = 100000
    s.portfolio = make(map[string]int)
}

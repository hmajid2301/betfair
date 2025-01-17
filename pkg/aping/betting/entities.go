package betting

import (
	"time"
)

type MarketFilter struct {
	TextQuery          string    `json:"textQuery,omitempty"`
	ExchangeIds        []string  `json:"exchangeIds,omitempty"`
	EventTypeIds       []string  `json:"eventTypeIds,omitempty"`
	EventIds           []string  `json:"eventIds,omitempty"`
	CompetitionIds     []string  `json:"competitionIds,omitempty"`
	MarketIds          []string  `json:"marketIds,omitempty"`
	Venues             []string  `json:"venues,omitempty"`
	BSPOnly            *bool     `json:"bspOnly,omitempty"`
	TurnInPlayEnabled  *bool     `json:"turnInPlayEnabled,omitempty"`
	InPlayOnly         *bool     `json:"inPlayOnly,omitempty"`
	MarketBettingTypes []string  `json:"marketBettingTypes,omitempty"`
	MarketCountries    []string  `json:"marketCountries,omitempty"`
	MarketTypeCodes    []string  `json:"marketTypeCodes,omitempty"`
	MarketStartTime    TimeRange `json:"marketStartTime,omitempty"`
	WithOrders         []string  `json:"withOrders,omitempty"`
	RaceTypes          []string  `json:"raceTypes,omitempty"`
}

type TimeRange struct {
	From *time.Time `json:"from,omitempty"`
	To   *time.Time `json:"to,omitempty"`
}

type MarketCatalogue struct {
	MarketID        string          `json:"marketId"`
	MarketName      string          `json:"marketName"`
	MarketStartTime *time.Time      `json:"marketStartTime"`
	Runners         []RunnerCatalog `json:"runners"`
}

type RunnerCatalog struct {
	SelectionID  uint    `json:"selectionId"`
	RunnerName   string  `json:"runnerName"`
	Handicap     float64 `json:"handicap"`
	SortPriority uint    `json:"sortPriority"`
}

type BetfairAPIError struct {
	Detail      BetfairDetailError `json:"detail"`
	FaultCode   string             `json:"faultCode"`
	FaultString string             `json:"faultstring"`
}

type BetfairDetailError struct {
	APINGException APINGException `json:"APINGException"`
	ExceptionName  string         `json:"exceptionname"`
}

type APINGException struct {
	ErrorCode    APINGExceptionCode `json:"errorCode"`
	ErrorDetails string             `json:"errorDetails"`
	RequestUUID  string             `json:"requestUUID"`
}

type MarketBook struct {
	MarketID string   `json:"marketId"`
	Runners  []Runner `json:"runners"`
}

type Runner struct {
	SelectionID uint `json:"selectionId"`
	// Change this to an ENUM
	Status          string  `json:"status"`
	LastPriceTraded float64 `json:lastPriceTraded`
}

type PlaceInstruction struct {
	OrderType   string     `json:"orderType"`
	SelectionID string     `json:"selectionId"`
	Side        string     `json:"side"`
	LimitOrder  LimitOrder `json:"limitOrder"`
}

type LimitOrder struct {
	Size            string `json:"size"`
	Price           string `json:"price"`
	PersistenceType string `json:"persistenceType"`
}

type PlaceExecutionReport struct {
	MarketID           string                   `json:"marketId"`
	Status             string                   `json:"status"`
	InstructionReports []PlaceInstructionReport `json:"instructionReports"`
}

type PlaceInstructionReport struct {
	Status              string  `json:"status"`
	AveragePriceMatched float64 `json:"averagePriceMatched"`
	SizeMatched         float64 `json:"sizeMatched"`
	OrderStatus         string  `json:"orderStatus"`
}

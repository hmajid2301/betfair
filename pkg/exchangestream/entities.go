package exchangestream

type ConnectionMessage struct {
	ConnectionID string `json:"connectionId"`
}

type AuthenticationMessage struct {
	AppKey       string `json:"appKey"`
	SessionToken string `json:"session"`
}

type StatusMessage struct {
	ErrorMessage     string     `json:"errorMessage"`
	ErrorCode        ErrorCode  `json:"errorCode"`
	ConnectionID     string     `json:"connectionId"`
	ConnectionClosed *bool      `json:"connectionClosed"`
	StatusCode       StatusCode `json:"statusCode"`
}

type OrderSubscriptionMessage struct {
	SegmentationEnabled bool        `json:"segmentationEnabled"`
	OrderFilter         OrderFilter `json:"orderFilter"`
	Clk                 string      `json:"clk,omitempty"`
	HeartbeatMs         uint        `json:"heartbeatMs,omitempty"`
	InitialClk          string      `json:"initialClk,omitempty"`
	ConflateMs          uint        `json:"conflateMs,omitempty"`
}

type OrderFilter struct {
	IncludeOverallPosition        *bool    `json:"includeOverallPosition"`
	CustomerStrategyRefs          []string `json:"customerStrategyRefs"`
	PartitionMatchedByStrategyRef *bool    `json:"partitionMatchedByStrategyRef"`
}

type MarketSubscriptionMessage struct {
	SegmentationEnabled *bool            `json:"segmentationEnabled,omitempty"`
	Clk                 string           `json:"clk,omitempty"`
	HeartbeatMs         uint             `json:"heartbeatMs,omitempty"`
	InitialClk          string           `json:"initialClk,omitempty"`
	MarketFilter        MarketFilter     `json:"marketFilter,omitempty"`
	ConflateMs          uint             `json:"conflateMs,omitempty"`
	MarketDataFilter    MarketDataFilter `json:"marketDataFilter,omitempty"`
}

type MarketFilter struct {
	CountryCodes      []string      `json:"countryCodes,omitempty"`
	BettingTypes      []BettingType `json:"bettingTypes,omitempty"`
	TurnInPlayEnabled *bool         `json:"turnInPlayEnabled,omitempty"`
	MarketTypes       []string      `json:"marketTypes,omitempty"`
	Venues            []string      `json:"venues,omitempty"`
	MarketIDs         []string      `json:"marketIds,omitempty"`
	EventTypeIDs      []string      `json:"eventTypeIds,omitempty"`
	EventIDs          []string      `json:"eventIds,omitempty"`
	BSPMarket         *bool         `json:"bspMarket,omitempty"`
	RaceTypes         []string      `json:"raceTypes,omitempty"`
}

type MarketDataFilter struct {
	LadderLevels uint        `json:"ladderLevels,omitempty"`
	Fields       []PriceData `json:"fields,omitempty"`
}

type MarketChangeMessage struct {
	Ct          *ChangeType    `json:"ct"`
	Clk         string         `json:"clk"`
	HeartbeatMs uint           `json:"heartbeatMs"`
	Pt          uint           `json:"pt"`
	InitialClk  string         `json:"initialClk"`
	Mc          []MarketChange `json:"mc"`
	ConflateMs  uint           `json:"conflateMs"`
	SegmentType *SegmentType   `json:"segmentType"`
	Status      uint           `json:"status"`
}

type MarketChange struct {
	Rc               []RunnerChange   `json:"rc"`
	Img              *bool            `json:"img"`
	Tv               *float64         `json:"tv"`
	Con              *bool            `json:"con"`
	MarketDefinition MarketDefinition `json:"marketDefinition"`
	ID               string           `json:"id"`
}

type MarketDefinition struct {
	Venue                 string                `json:"venue"`
	RaceType              string                `json:"raceType"`
	SettledTime           string                `json:"settledTime"`
	Timezone              string                `json:"timezone"`
	EachWayDivisor        float64               `json:"eachWayDivisor"`
	Regulators            []string              `json:"regulators"`
	MarketType            string                `json:"marketType"`
	MarketBaseRate        float64               `json:"marketBaseRate"`
	NumberOfWinners       uint                  `json:"numberOfWinners"`
	CountryCode           string                `json:"countryCode"`
	LineMaxUnit           float64               `json:"lineMaxUnit"`
	InPlay                *bool                 `json:"inPlay"`
	BetDelay              uint                  `json:"betDelay"`
	BSPMarket             *bool                 `json:"bspMarket"`
	BettingType           BettingType           `json:"bettingType"`
	NumberOfActiveRunners uint                  `json:"numberOfActiveRunners"`
	LineMinUnit           float64               `json:"lineMinUnit"`
	EventID               string                `json:"eventId"`
	CrossMatching         *bool                 `json:"crossMatching"`
	RunnersVoidable       *bool                 `json:"runnersVoidable"`
	TurnInPlayEnabled     *bool                 `json:"turnInPlayEnabled"`
	PriceLadderDefinition PriceLadderDefinition `json:"priceLadderDefinition"`
	KeyLineDefinition     KeyLineDefinition     `json:"keyLineDefinition"`
	SuspendTime           string                `json:"suspendTime"`
	DiscountAllowed       *bool                 `json:"discountAllowed"`
	PersistenceEnabled    *bool                 `json:"persistenceEnabled"`
	Runners               []RunnerDefinition    `json:"runners"`
	Version               uint                  `json:"version"`
	EventTypeID           string                `json:"eventTypeId"`
	Complete              *bool                 `json:"complete"`
	OpenDate              string                `json:"openDate"`
	MarketTime            string                `json:"marketTime"`
	BSPReconciled         *bool                 `json:"bspReconciled"`
	LineInterval          float64               `json:"lineInterval"`
	Status                RaceStatus            `json:"status"`
}

type RunnerDefinition struct {
	SortPriority     uint       `json:"sortPriority"`
	RemovalDate      string     `json:"removalDate"`
	ID               uint       `json:"id"`
	Hc               float64    `json:"hc"`
	AdjustmentFactor float64    `json:"adjustmentFactor"`
	BSP              float64    `json:"bsp"`
	Status           RaceStatus `json:"status"`
}

type PriceLadderDefinition struct {
	Type PriceLadderType `json:"type"`
}

type KeyLineDefinition struct {
	Kl []KeyLineSelection `json:"kl"`
}

type KeyLineSelection struct {
	Id uint    `json:"id"`
	Hc float64 `json:"hc"`
}

type RunnerChange struct {
	Tv    float64     `json:"tv"`
	Batb  [][]float64 `json:"batb"`
	Spb   [][]float64 `json:"spb"`
	Bdatl [][]float64 `json:"bdatl"`
	Trd   [][]float64 `json:"trd"`
	Spf   float64     `json:"spf"`
	Ltp   float64     `json:"ltp"`
	Atb   [][]float64 `json:"atb"`
	Spl   [][]float64 `json:"spl"`
	Spn   float64     `json:"spn"`
	Atl   [][]float64 `json:"atl"`
	Batl  [][]float64 `json:"batl"`
	ID    uint        `json:"id"`
	Hc    float64     `json:"hc"`
	Bdatb [][]float64 `json:"bdatb"`
}

type OrderChangeMessage struct {
	Ct          *ChangeType         `json:"ct"`
	Clk         string              `json:"clk"`
	HeartbeatMs uint                `json:"heartbeatMs"`
	Pt          uint                `json:"pt"`
	Oc          []OrderMarketChange `json:"oc"`
	InitialClk  string              `json:"initialClk"`
	ConflateMs  uint                `json:"conflateMs"`
	SegmentType *SegmentType        `json:"segmentType"`
	Status      uint                `json:"status"`
}

type OrderMarketChange struct {
	AccountID uint                `json:"accountId"`
	Orc       []OrderRunnerChange `json:"orc"`
	Closed    *bool               `json:"closed"`
	ID        string              `json:"id"`
}

type OrderRunnerChange struct {
	Mb        [][]float64                    `json:"mb"`
	Smc       map[string]StrategyMatchChange `json:"smc"`
	Uo        []Order                        `json:"uo"`
	ID        uint                           `json:"id"`
	Hc        float64                        `json:"hc"`
	FullImage *bool                          `json:"fullImage"`
	Ml        [][]float64                    `json:"ml"`
}

type Order struct {
	Side   OrderSide       `json:"side"`
	Sv     float64         `json:"sv"`
	Pt     PersistenceType `json:"pt"`
	Ot     OrderType       `json:"ot"`
	P      float64         `json:"p"`
	Sc     float64         `json:"sc"`
	Rc     string          `json:"rc"`
	S      float64         `json:"s"`
	Pd     uint            `json:"pd"`
	Rac    string          `json:"rac"`
	Md     uint            `json:"md"`
	Ld     uint            `json:"ld"`
	Sl     float64         `json:"sl"`
	Avp    float64         `json:"avp"`
	Sm     float64         `json:"sm"`
	Rfo    string          `json:"rfo"`
	ID     string          `json:"id"`
	BSP    float64         `json:"bsp"`
	Rfs    string          `json:"rfs"`
	Status OrderStatus     `json:"status"`
	Sr     float64         `json:"sr"`
}

type StrategyMatchChange struct {
	Mb [][]float64 `json:"mb"`
	Ml [][]float64 `json:"ml"`
}

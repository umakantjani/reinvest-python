package models

// Profile model from the Financial Modeling Rep API
type FmgProfile struct {
	Pk					string `gorm:"primary_key"`
	Symbol            string  `json:"symbol"`
	Price             float64 `json:"price"`
	Beta              float64 `json:"beta"`
	VolAvg            float64     `json:"volAvg"`
	MktCap            float64   `json:"mktCap"`
	LastDiv           float64 `json:"lastDiv"`
	Range             string  `json:"range"`
	Changes           float64 `json:"changes"`
	CompanyName       string  `json:"companyName"`
	Exchange          string  `json:"exchange"`
	ExchangeShortName string  `json:"exchangeShortName"`
	Industry          string  `json:"industry"`
	Website           string  `json:"website"`
	Description       string  `json:"description"`
	Ceo               string  `json:"ceo"`
	Sector            string  `json:"sector"`
	Country           string  `json:"country"`
	FullTimeEmployees string  `json:"fullTimeEmployees"`
	Phone             string  `json:"phone"`
	Address           string  `json:"address"`
	City              string  `json:"city"`
	State             string  `json:"state"`
	Zip               string  `json:"zip"`
	DcfDiff           float64 `json:"dcfDiff"`
	Dcf               float64 `json:"dcf"`
	Image             string  `json:"image"`
}


// S&P 500 Stock lists
type FmgListSnp struct {
	Pk				string `gorm:"primary_key"`
	Symbol         string `json:"symbol"`
	Name           string `json:"name"`
	Sector         string `json:"sector"`
	SubSector      string `json:"subSector"`
	HeadQuarter    string `json:"headQuarter"`
	DateFirstAdded string `json:"dateFirstAdded"`
	Cik            string `json:"cik"`
	Founded        string `json:"founded"`
}


//List of stocks in the market
type FmgListSymbols struct {
	Pk			string `gorm:"primary_key"`
	Symbol   string  `json:"symbol"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Exchange string  `json:"exchange"`
}


// Financial Statements
// Income Statement: Annual, Quarterly
type FmgIncomeStatement struct {
	Pk								string `gorm:"primary_key"`
	PeriodType						string
	Date                             string  `json:"date"`
	Symbol                           string  `json:"symbol"`
	FillingDate                      string  `json:"fillingDate"`
	AcceptedDate                     string  `json:"acceptedDate"`
	Period                           string  `json:"period"`
	Revenue                          float64   `json:"revenue"`
	CostOfRevenue                    float64   `json:"costOfRevenue"`
	GrossProfit                      float64   `json:"grossProfit"`
	GrossProfitRatio                 float64 `json:"grossProfitRatio"`
	ResearchAndDevelopmentExpenses   float64   `json:"researchAndDevelopmentExpenses"`
	GeneralAndAdministrativeExpenses float64   `json:"generalAndAdministrativeExpenses"`
	SellingAndMarketingExpenses      float64 `json:"sellingAndMarketingExpenses"`
	OtherExpenses                    float64     `json:"otherExpenses"`
	OperatingExpenses                float64   `json:"operatingExpenses"`
	CostAndExpenses                  float64   `json:"costAndExpenses"`
	InterestExpense                  float64   `json:"interestExpense"`
	DepreciationAndAmortization      float64   `json:"depreciationAndAmortization"`
	Ebitda                           float64   `json:"ebitda"`
	Ebitdaratio                      float64 `json:"ebitdaratio"`
	OperatingIncome                  float64   `json:"operatingIncome"`
	OperatingIncomeRatio             float64 `json:"operatingIncomeRatio"`
	TotalOtherIncomeExpensesNet      float64     `json:"totalOtherIncomeExpensesNet"`
	IncomeBeforeTax                  float64   `json:"incomeBeforeTax"`
	IncomeBeforeTaxRatio             float64 `json:"incomeBeforeTaxRatio"`
	IncomeTaxExpense                 float64   `json:"incomeTaxExpense"`
	NetIncome                        float64   `json:"netIncome"`
	NetIncomeRatio                   float64 `json:"netIncomeRatio"`
	Eps                              float64 `json:"eps"`
	Epsdiluted                       float64 `json:"epsdiluted"`
	WeightedAverageShsOut            float64   `json:"weightedAverageShsOut"`
	WeightedAverageShsOutDil         float64   `json:"weightedAverageShsOutDil"`
	Link                             string  `json:"link"`
	FinalLink                        string  `json:"finalLink"`
}


// Financial Statements
// Balance Sheet: Annual, Quarterly
type FmgBalanceSheet struct {
	Pk										string `gorm:"primary_key"`
	PeriodType								string
	Date                                    string  `json:"date"`
	Symbol                                  string  `json:"symbol"`
	FillingDate                             string  `json:"fillingDate"`
	AcceptedDate                            string  `json:"acceptedDate"`
	Period                                  string  `json:"period"`
	CashAndCashEquivalents                  float64   `json:"cashAndCashEquivalents"`
	ShortTermInvestments                    float64   `json:"shortTermInvestments"`
	CashAndShortTermInvestments             float64   `json:"cashAndShortTermInvestments"`
	NetReceivables                          float64   `json:"netReceivables"`
	Inventory                               float64   `json:"inventory"`
	OtherCurrentAssets                      float64   `json:"otherCurrentAssets"`
	TotalCurrentAssets                      float64   `json:"totalCurrentAssets"`
	PropertyPlantEquipmentNet               float64   `json:"propertyPlantEquipmentNet"`
	Goodwill                                float64 `json:"goodwill"`
	IntangibleAssets                        float64 `json:"intangibleAssets"`
	GoodwillAndIntangibleAssets             float64 `json:"goodwillAndIntangibleAssets"`
	LongTermInvestments                     float64   `json:"longTermInvestments"`
	TaxAssets                               float64 `json:"taxAssets"`
	OtherNonCurrentAssets                   float64   `json:"otherNonCurrentAssets"`
	TotalNonCurrentAssets                   float64   `json:"totalNonCurrentAssets"`
	OtherAssets                             float64   `json:"otherAssets"`
	TotalAssets                             float64   `json:"totalAssets"`
	AccountPayables                         float64   `json:"accountPayables"`
	ShortTermDebt                           float64   `json:"shortTermDebt"`
	TaxPayables                             float64 `json:"taxPayables"`
	DeferredRevenue                         float64   `json:"deferredRevenue"`
	OtherCurrentLiabilities                 float64   `json:"otherCurrentLiabilities"`
	TotalCurrentLiabilities                 float64   `json:"totalCurrentLiabilities"`
	LongTermDebt                            float64   `json:"longTermDebt"`
	DeferredRevenueNonCurrent               float64 `json:"deferredRevenueNonCurrent"`
	DeferredTaxLiabilitiesNonCurrent        float64 `json:"deferredTaxLiabilitiesNonCurrent"`
	OtherNonCurrentLiabilities              float64   `json:"otherNonCurrentLiabilities"`
	TotalNonCurrentLiabilities              float64   `json:"totalNonCurrentLiabilities"`
	OtherLiabilities                        float64 `json:"otherLiabilities"`
	TotalLiabilities                        float64   `json:"totalLiabilities"`
	CommonStock                             float64   `json:"commonStock"`
	RetainedEarnings                        float64   `json:"retainedEarnings"`
	AccumulatedOtherComprehensiveIncomeLoss float64     `json:"accumulatedOtherComprehensiveIncomeLoss"`
	OthertotalStockholdersEquity            float64   `json:"othertotalStockholdersEquity"`
	TotalStockholdersEquity                 float64   `json:"totalStockholdersEquity"`
	TotalLiabilitiesAndStockholdersEquity   float64   `json:"totalLiabilitiesAndStockholdersEquity"`
	TotalInvestments                        float64   `json:"totalInvestments"`
	TotalDebt                               float64   `json:"totalDebt"`
	NetDebt                                 float64   `json:"netDebt"`
	Link                                    string  `json:"link"`
	FinalLink                               string  `json:"finalLink"`
}


// Financial Statements
// Cash Flow: Annual, Quarterly
type FmgCashFlow struct {
	Pk										string `gorm:"primary_key"`
	PeriodType								string
	Date                                     string  `json:"date"`
	Symbol                                   string  `json:"symbol"`
	FillingDate                              string  `json:"fillingDate"`
	AcceptedDate                             string  `json:"acceptedDate"`
	Period                                   string  `json:"period"`
	NetIncome                                float64   `json:"netIncome"`
	DepreciationAndAmortization              float64   `json:"depreciationAndAmortization"`
	DeferredIncomeTax                        float64     `json:"deferredIncomeTax"`
	StockBasedCompensation                   float64   `json:"stockBasedCompensation"`
	ChangeInWorkingCapital                   float64   `json:"changeInWorkingCapital"`
	AccountsReceivables                      float64     `json:"accountsReceivables"`
	Inventory                                float64     `json:"inventory"`
	AccountsPayables                         float64     `json:"accountsPayables"`
	OtherWorkingCapital                      float64   `json:"otherWorkingCapital"`
	OtherNonCashItems                        float64 `json:"otherNonCashItems"`
	NetCashProvidedByOperatingActivities     float64   `json:"netCashProvidedByOperatingActivities"`
	InvestmentsInPropertyPlantAndEquipment   float64   `json:"investmentsInPropertyPlantAndEquipment"`
	AcquisitionsNet                          float64     `json:"acquisitionsNet"`
	PurchasesOfInvestments                   float64   `json:"purchasesOfInvestments"`
	SalesMaturitiesOfInvestments             float64   `json:"salesMaturitiesOfInvestments"`
	OtherInvestingActivites                  float64   `json:"otherInvestingActivites"`
	NetCashUsedForInvestingActivites         float64   `json:"netCashUsedForInvestingActivites"`
	DebtRepayment                            float64   `json:"debtRepayment"`
	CommonStockIssued                        float64     `json:"commonStockIssued"`
	CommonStockRepurchased                   float64   `json:"commonStockRepurchased"`
	DividendsPaid                            float64   `json:"dividendsPaid"`
	OtherFinancingActivites                  float64     `json:"otherFinancingActivites"`
	NetCashUsedProvidedByFinancingActivities float64   `json:"netCashUsedProvidedByFinancingActivities"`
	EffectOfForexChangesOnCash               float64 `json:"effectOfForexChangesOnCash"`
	NetChangeInCash                          float64   `json:"netChangeInCash"`
	CashAtEndOfPeriod                        float64   `json:"cashAtEndOfPeriod"`
	CashAtBeginningOfPeriod                  float64   `json:"cashAtBeginningOfPeriod"`
	OperatingCashFlow                        float64   `json:"operatingCashFlow"`
	CapitalExpenditure                       float64   `json:"capitalExpenditure"`
	FreeCashFlow                             float64   `json:"freeCashFlow"`
	Link                                     string  `json:"link"`
	FinalLink                                string  `json:"finalLink"`
}



// Financial Statements
// Financial Ratios, Annual
type FmgFinancialRatios struct {
	Pk									string `gorm:"primary_key"`
	PeriodType							string
	Symbol                             string  `json:"symbol"`
	Date                               string  `json:"date"`
	CurrentRatio                       float64 `json:"currentRatio"`
	QuickRatio                         float64 `json:"quickRatio"`
	CashRatio                          float64 `json:"cashRatio"`
	DaysOfSalesOutstanding             float64 `json:"daysOfSalesOutstanding"`
	DaysOfInventoryOutstanding         float64 `json:"daysOfInventoryOutstanding"`
	OperatingCycle                     float64 `json:"operatingCycle"`
	DaysOfPayablesOutstanding          float64 `json:"daysOfPayablesOutstanding"`
	CashConversionCycle                float64 `json:"cashConversionCycle"`
	GrossProfitMargin                  float64 `json:"grossProfitMargin"`
	OperatingProfitMargin              float64 `json:"operatingProfitMargin"`
	PretaxProfitMargin                 float64 `json:"pretaxProfitMargin"`
	NetProfitMargin                    float64 `json:"netProfitMargin"`
	EffectiveTaxRate                   float64 `json:"effectiveTaxRate"`
	ReturnOnAssets                     float64 `json:"returnOnAssets"`
	ReturnOnEquity                     float64 `json:"returnOnEquity"`
	ReturnOnCapitalEmployed            float64 `json:"returnOnCapitalEmployed"`
	NetIncomePerEBT                    float64 `json:"netIncomePerEBT"`
	EbtPerEbit                         float64  `json:"ebtPerEbit"`
	EbitPerRevenue                     float64 `json:"ebitPerRevenue"`
	DebtRatio                          float64 `json:"debtRatio"`
	DebtEquityRatio                    float64 `json:"debtEquityRatio"`
	LongTermDebtToCapitalization       float64 `json:"longTermDebtToCapitalization"`
	TotalDebtToCapitalization          float64 `json:"totalDebtToCapitalization"`
	InterestCoverage                   float64 `json:"interestCoverage"`
	CashFlowToDebtRatio                float64 `json:"cashFlowToDebtRatio"`
	CompanyEquityMultiplier            float64 `json:"companyEquityMultiplier"`
	ReceivablesTurnover                float64 `json:"receivablesTurnover"`
	PayablesTurnover                   float64 `json:"payablesTurnover"`
	InventoryTurnover                  float64 `json:"inventoryTurnover"`
	FixedAssetTurnover                 float64 `json:"fixedAssetTurnover"`
	AssetTurnover                      float64 `json:"assetTurnover"`
	OperatingCashFlowPerShare          float64 `json:"operatingCashFlowPerShare"`
	FreeCashFlowPerShare               float64 `json:"freeCashFlowPerShare"`
	CashPerShare                       float64 `json:"cashPerShare"`
	PayoutRatio                        float64 `json:"payoutRatio"`
	OperatingCashFlowSalesRatio        float64 `json:"operatingCashFlowSalesRatio"`
	FreeCashFlowOperatingCashFlowRatio float64 `json:"freeCashFlowOperatingCashFlowRatio"`
	CashFlowCoverageRatios             float64 `json:"cashFlowCoverageRatios"`
	ShortTermCoverageRatios            float64 `json:"shortTermCoverageRatios"`
	CapitalExpenditureCoverageRatio    float64 `json:"capitalExpenditureCoverageRatio"`
	DividendPaidAndCapexCoverageRatio  float64 `json:"dividendPaidAndCapexCoverageRatio"`
	DividendPayoutRatio                float64 `json:"dividendPayoutRatio"`
	PriceBookValueRatio                float64 `json:"priceBookValueRatio"`
	PriceToBookRatio                   float64 `json:"priceToBookRatio"`
	PriceToSalesRatio                  float64 `json:"priceToSalesRatio"`
	PriceEarningsRatio                 float64 `json:"priceEarningsRatio"`
	PriceToFreeCashFlowsRatio          float64 `json:"priceToFreeCashFlowsRatio"`
	PriceToOperatingCashFlowsRatio     float64 `json:"priceToOperatingCashFlowsRatio"`
	PriceCashFlowRatio                 float64 `json:"priceCashFlowRatio"`
	PriceEarningsToGrowthRatio         float64 `json:"priceEarningsToGrowthRatio"`
	PriceSalesRatio                    float64 `json:"priceSalesRatio"`
	DividendYield                      float64 `json:"dividendYield"`
	EnterpriseValueMultiple            float64 `json:"enterpriseValueMultiple"`
	PriceFairValue                     float64 `json:"priceFairValue"`
}
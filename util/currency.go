package util

const (
	// USD currency
	USD = "USD"
	// EUR currency
	EUR = "EUR"
	// CAD currency
	CAD = "CAD"
)

// IsSupportedCurrency returns true if currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}

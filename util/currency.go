package util

const (
	USD = "USD"
	EUR = "EUR"
	MXN = "MXN"
	CAD = "CAD"
	GBP = "GBP"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, MXN, CAD, GBP:
		return true
	}
	return false
}

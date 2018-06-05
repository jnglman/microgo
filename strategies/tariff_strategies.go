package strategies

type TariffStrategy interface {
	CalculatePrice(basePrice float64) float64
}

type StrategyHolder struct {
	CurrentStrategy TariffStrategy
}

type MultiplierTariffStrategy struct {
	multiplier float64
}

func (s *MultiplierTariffStrategy) CalculatePrice(basePrice float64) float64 {
	return basePrice * s.multiplier
}

func NewMultiplierStrategy(coefficient float64) *MultiplierTariffStrategy {
	s := new(MultiplierTariffStrategy)
	s.multiplier = coefficient
	return s
}

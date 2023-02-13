package main

type calculator struct {
	minimumpurchaseamount int
	discountamount        int
}

func NewdiscountCalculator(minimumpurchaseamount, discountamount int) *calculator {
	return &calculator{
		minimumpurchaseamount: minimumpurchaseamount,
		discountamount:        discountamount,
	}
}

func (c *calculator) Calculate(Purchaseamount int) int {
	if Purchaseamount > c.minimumpuchaseamount {
		return purchaseamount - c.discountamount
	}
	return Purchaseamount
}

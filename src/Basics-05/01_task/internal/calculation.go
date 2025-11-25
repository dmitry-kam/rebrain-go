package internal

func CalcPrice(c *Customer, price int) (int, error) {
	discount, err := c.CalcDiscount()
	if err != nil {
		return 0, err
	}

	if price > discount {
		return price - discount, nil
	}

	return 0, nil
}

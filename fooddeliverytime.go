package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var time food
	switch order {
	case "burger":
		time.preptime = 15
	case "chips":
		time.preptime = 10
	case "nuggets":
		time.preptime = 12
	default:
		time.preptime = 404
	}
	return time.preptime
}

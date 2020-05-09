package main

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) <= 1 {
		return true
	}

	var firstRoc float64

	for i := 1; i < len(coordinates); i++ {
		roc := rateOfChange(coordinates[i], coordinates[i-1])

		if firstRoc == 0 {
			firstRoc = roc
			continue
		}

		if firstRoc != roc {
			return false
		}
	}

	return true
}

func rateOfChange(c2, c1 []int) float64 {
	return (float64(c2[1]) - float64(c1[1])) / (float64(c2[0]) - float64(c1[0]))
}

package template

func isPerfectSquare(num int) bool {
	for i := 1; i <= num; {
		square := i * i

		if square == num {
			return true
		}

		if square < num {
			i *= 2
			continue
		}

		for j := i/2 + 1; j < i; j++ {
			square := j * j

			if square == num {
				return true
			}

			if square > num {
				return false
			}
		}

		break
	}

	return false
}

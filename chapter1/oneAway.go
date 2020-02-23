package chapter1


func intAbs(integer int) int {
	if integer < 0 {
		return -integer
	}
	return integer
}

func shorterLonger(s1, s2 string) (string, string) {
	if len(s1) > len(s2) {
		return s2, s1
	}
	return s1, s2
}

func OneEditAway(s1, s2 string) bool {
	if intAbs(len(s1)-len(s2)) > 1 {
		return false
	}

	shorter, longer := shorterLonger(s1, s2)

	sIdx := 0
	lIdx := 0
	foundDifference := false

	for lIdx < len(longer) && sIdx < len(shorter) {
		if shorter[sIdx] != longer[lIdx] {
			if foundDifference {
				return false
			}
			foundDifference = true

			if len(shorter) == len(longer) {
				sIdx++
			}
		} else {
			sIdx++
		}
		lIdx++
	}
	return true
}
package rules

func contains(needle string, haystack []string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func containsAny(needles []string, haystack []string) bool {
	for _, item := range haystack {
		for _, needle := range needles {
			if item == needle {
				return true
			}
		}
	}

	return false
}

func containsAll(needles []string, haystack []string) bool {
OUTER:
	for _, needle := range needles {
		for _, item := range haystack {
			if needle == item {
				continue OUTER
			}
		}
		return false
	}
	return true
}

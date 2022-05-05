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

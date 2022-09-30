package eval

func LongestStrings(tabstring []string) []string {
	max := -1

	var result []string

	for _, s := range tabstring {
		if len(s) < max {
			continue
		}
		if len(s) > max {
			max = len(s)
			result = result[:0]
		}
		result = append(result, s)
	}
	return result
}

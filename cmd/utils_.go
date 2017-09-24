package cmd

func getDefaultConfig() string {
	return "./konnect.yml"
}

func getVersion() string {
	version := "0.0.1"
	return version
}

// Remove duplicate elements from a string slice.
// https://goo.gl/ttDAg2
func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}
	for _, host := range elements {
		if encountered[host] == false {
			encountered[host] = true
			result = append(result, host)
		}
	}
	return result
}

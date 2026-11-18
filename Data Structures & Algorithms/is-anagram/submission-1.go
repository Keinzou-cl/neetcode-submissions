func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sStorage := make(map[byte]int)
	tStorage := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		sStorage[s[i]] += 1
		tStorage[t[i]] += 1
	}
	for k, v := range sStorage {
		if tStorage[k] != v {
			return false
		}
	}
	return true
}
package leetcode

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		a, b := s[i], t[i]
		if v, ok := m1[a]; ok {
			if v != b {
				return false
			}
		}
		if v, ok := m2[b]; ok {
			if v != a {
				return false
			}
		}
		m1[a] = b
		m2[b] = a

	}
	return true
}

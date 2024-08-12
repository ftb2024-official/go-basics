package tasks

func spamMasker(url string) string {
	bUrl := []byte(url)
	subArr := []byte{104, 116, 116, 112, 58, 47, 47}
	startIdx := 0

	for i := 0; i < len(bUrl)-len(subArr); i++ {
		if match(bUrl[i:i+len(subArr)], subArr) {
			startIdx = i + 6
			break
		}
	}

	for bUrl[startIdx] != 32 {
		bUrl[startIdx] = 42
		startIdx++
	}

	return string(bUrl)
}

func match(a, b []byte) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// " " -> [32]
// * -> [42]
// http:// -> [104 116 116 112 58 47 47]

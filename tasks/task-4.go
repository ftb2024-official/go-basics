package tasks

import "fmt"

func isArrsEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func findSubArrStartIdx(mainArr, subArr []byte) int {
	subArrLen := len(subArr)
	for i := 0; i < len(mainArr)-subArrLen; i++ {
		// extract current subarray from mainArr
		if isArrsEqual(mainArr[i:i+subArrLen], subArr) {
			return i
		}
	}

	return -1
}

func spamMasker(url string) string {
	bUrl := []byte(url)
	bSubArr := []byte{104, 116, 116, 112, 58, 47, 47} // http://
	startIdx := findSubArrStartIdx(bUrl, bSubArr) + 7

	for bUrl[startIdx] != 32 {
		bUrl[startIdx] = 42
		startIdx++
		if startIdx == len(bUrl) {
			break
		}
	}

	return string(bUrl)
}

func Example4() {
	fmt.Println(spamMasker("Here's my spammy page: http://hehefouls.netHAHAHA see you."))
	fmt.Println(spamMasker("http://example.com/home"))
	fmt.Println(spamMasker("http://mywebsite.net/about-us"))
	fmt.Println(spamMasker("http://coolapp.org/contact"))
	fmt.Println(spamMasker("http://fakesite.com/blog"))
	fmt.Println(spamMasker("http://randomdomain.io/services"))

	fmt.Println(spamMasker("http://shoponline.com/products?id=1234&category=shoes"))
	fmt.Println(spamMasker("http://socialmedia.net/user/profile?username=johndoe"))
	fmt.Println(spamMasker("http://travelguru.com/search?destination=paris&checkin=2024-09-01"))
	fmt.Println(spamMasker("http://moviesite.org/watch?movieid=5678&quality=hd"))
	fmt.Println(spamMasker("http://newsportal.com/articles?topic=tech&author=jane_smith"))
}

// " " -> [32]
// * -> [42]
// http:// -> [104 116 116 112 58 47 47]

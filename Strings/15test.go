package strings

import "fmt"

func MainFunction15() {
	ans := strStr("arjunmalhotra", "alh")
	ans1 := str("arjunmalhotra", "")
	fmt.Println(ans)
	fmt.Println(ans1)
}

func strStr(haystack string, needle string) string {
	needleByte := []byte(needle)
	haystackByte := []byte(haystack)
	ansArray := []byte{}

	lenNeedle := len(needle)
	lenHaystack := len(haystack)

	for i := 0; i < lenHaystack; i++ {
		if haystackByte[i] == needleByte[0] {
			// fmt.Println(i)
			ansArray = haystackByte[i : i+lenNeedle]
			tempString := string(ansArray)
			if tempString != needle {
				continue
			} else {
				return string(ansArray)
			}
		}
	}
	return ""
}

func str(haystack string, needle string) int {
	needleByte := []byte(needle)
	haystackByte := []byte(haystack)
	ansArray := []byte{}

	lenNeedle := len(needle)
	lenHaystack := len(haystack)
	if lenNeedle == 0 {
		return 0
	}

	for i := 0; i < lenHaystack; i++ {
		if haystackByte[i] == needleByte[0] {
			// fmt.Println(i)
			ansArray = haystackByte[i : i+lenNeedle]
			tempString := string(ansArray)
			if tempString != needle {
				continue
			} else {
				return i
			}
		}
	}
	return -1
}

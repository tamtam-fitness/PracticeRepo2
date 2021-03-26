package rumaNum

import "fmt"

func ConvertRumanToInt(numeral string) int {
	// 1.文字列を分ける
	rumanMap := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	// 2.分けた文字列をパースする
	// +1した理由:前と後ろの値を見比べる際にindex out of rangeが起きないようにする
	arabicVals := make([]int, len(numeral)+1)

	// ローマ数字一つ一つをforでまわす
	for idx, digit := range numeral {
		// 入れた文字がちゃんとローマ数字であるか判定
		if val, present := rumanMap[digit]; present {
			arabicVals[idx] = val
		} else {
			fmt.Printf("Error: The roman numeral %s has a bad digit: %c\n", numeral, digit)
			return 0
		}
	}

	sumVal := 0

	for idx := 0; idx < len(numeral); idx++ {
		if arabicVals[idx] < arabicVals[idx+1] {
			arabicVals[idx] = -arabicVals[idx]
		}
		sumVal += arabicVals[idx]
	}

	return sumVal

	//forで回し和を出す
}

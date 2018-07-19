package hangul


func ExampleHasConsonantSuffix() {
	HasConsonantSuffix("사과")
	HasConsonantSuffix("바나나")
	HasConsonantSuffix("토마토")
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
}

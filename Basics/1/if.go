func main() {
	fmt.Println("Understanding conditional statements")

	// Problem Statement:
	// Grade a student according to score as follows:
	// 0-40  -> F
	// 41-80 -> B
	// 81-90 -> A
	// 91+   -> A+

	if score := 100; score >= 0 && score <= 40 {
		fmt.Println("F")
	} else if score >= 41 && score <= 80 {
		fmt.Println("B")
	} else if score >= 81 && score <= 90 {
		fmt.Println("A")
	} else if score >= 91 {
		fmt.Println("A+")
	}

}
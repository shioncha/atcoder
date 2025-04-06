package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	// ans := 0
	// tmp := 0
	// for i := 2; i <= n; i++ {
	// 	tmp = i
	// 	if tmp % 2 != 0 {
	// 		continue
	// 	}

	// 	a := math.Sqrt(float64(tmp))
	// 	if a == float64(int(a)) {
	// 		ans++
	// 		continue
	// 	}
	// 	if tmp % 2 == 0 {
	// 		b := math.Sqrt(float64(tmp * 2))
	// 		if b == float64(int(b)) {
	// 			ans++
	// 			continue
	// 		}
	// 	}
	// }

	a := 60
	ans := 0
	for i := 1; a > 0; i+=2 {
		for a > 0 && (1<<a)*i*i > n {
			a--
		}
		ans+=a
	}

	
	fmt.Println(ans)
}


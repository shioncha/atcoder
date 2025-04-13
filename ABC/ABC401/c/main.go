// 正整数 N,K が与えられます。長さ N+1 の数列 A=(A0,A1,...,AN) の各要素の値を、以下の方法で定義します。0≤i<K のときAi​=1、K≤iのとき、Ai​=Ai−K​+Ai−K+1+...+Ai−1​。AN​を10^9で割ったあまりを求めてください。
package main

import "fmt"

func main() {
  var n, k int
  fmt.Scanf("%d %d\n", &n, &k)

	if k > n {
		fmt.Println(1)
		return
	}

	a := make([]int64, n+1)
	tmp := int64(0)
	for i := range a[:k] {
		a[i] = 1
		tmp = (tmp + a[i])%1000000000
	}
	for i := k; i < n+1; i++ {
		a[i] = tmp
		if tmp - a[i-k] + a[i] < 0 {
			tmp = (tmp - a[i-k] + a[i] + 1000000000) % 1000000000
		} else {
			tmp = (tmp - a[i-k] + a[i]) % 1000000000
		}
	}
	fmt.Println(a[n])
}

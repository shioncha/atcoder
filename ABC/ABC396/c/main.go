package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// var sc = bufio.NewScanner(os.Stdin)
var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

func readLine() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	// buf := make([]byte, 4096)
	// sc.Buffer(buf, 2e6)
	var n, m int
	// fmt.Scanf("%d %d\n", &n, &m)
	// b := make([]int, n)
	// w := make([]int, m)
	// for i := 0; i < n; i++ {
	// 	fmt.Scanf("%d", &b[i])
	// }
	// fmt.Scanf("\n")
	// for i := 0; i < m; i++ {
	// 	fmt.Scanf("%d", &w[i])
	// }

	// sc.Scan()
	// fmt.Sscanf(sc.Text(), "%d %d", &n, &m)
	// b := make([]int, n)
	// w := make([]int, m)
	// sc.Scan()
	// bs := strings.Split(sc.Text(), " ")
	// for i := 0; i < n; i++ {
	// 	b[i], _ = strconv.Atoi(bs[i])
	// }
	// sc.Scan()
	// ws := strings.Split(sc.Text(), " ")
	// for i:=0; i<m; i++ {
	// 	w[i], _ = strconv.Atoi(ws[i])
	// }

	fmt.Fscan(reader, &n, &m)
	b := make([]int, n)
	w := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &b[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &w[i])
	}

	// fmt.Sscanf(readLine(), "%d %d", &n, &m)
	// b := make([]int, n)
	// w := make([]int, m)
	// bs := strings.Split(readLine(), " ")
	// for i := 0; i < n; i++ {
	// 	b[i], _ = strconv.Atoi(string(bs[i]))
	// }
	// ws := strings.Split(readLine(), " ")
	// for i := 0; i < m; i++ {
	// 	w[i], _ = strconv.Atoi(string(ws[i]))
	// }

	sort.Slice(b, func(i, j int) bool {
		return b[i] > b[j]
	})
	sort.Slice(w, func(i, j int) bool {
		return w[i] > w[j]
	})

	fmt.Println(b)
	fmt.Println(w)
	
	ans := 0
	bp := 0
	wp := 0
	for bp < n && b[bp] > 0 {
		ans += b[bp]
		bp++
	}
	for wp < m && wp < bp && w[wp] > 0 {
		ans += w[wp]
		wp++
	}
	for i := 0; i < n && bp < n && wp < m && bp >= wp; i++ {
		if ans + b[bp] > ans {
			ans += b[bp]
			bp++
		} else if ans + b[bp] + w[wp] > ans {
			ans += b[bp] + w[wp]
			bp++
			wp++
		}
	}

	fmt.Println(ans)
}

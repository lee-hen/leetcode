package cracking_the_safe

import (
	"math"
	"strings"
)

func CrackSafe(n int, k int) string {
	str := new(strings.Builder)
	for i := 0; i < n; i++ {
		str.WriteString("0")
	}

	visited := make(map[string]bool)
	visited[str.String()] = true

	limit := math.Pow(float64(k) , float64(n))
	dfs(str, byte(k), n, int(limit), visited)
	return str.String()
}

func dfs(str *strings.Builder, k byte, n, limit int, visited map[string]bool) bool {
	if len(visited) == limit {
		return true
	}

	last := str.String()[str.Len()-n+1:]
	var c byte
	for c = '0'; c < '0' + k; c++ {
		newStr := last + string(c)

		if !visited[newStr]  {
			visited[newStr] = true
			str.WriteByte(c)

			if dfs(str, k, n, limit, visited) {
				return true
			}

			delete(visited, newStr)

			tmp := str.String()[:str.Len()-1]
			str.Reset()
			str.WriteString(tmp)
		}
	}

	return false
}

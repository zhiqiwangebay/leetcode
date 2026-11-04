package leetcode

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	weights := make(map[string]map[string]float64)
	for i, eq := range equations {
		if _, ok := weights[eq[0]]; !ok {
			weights[eq[0]] = make(map[string]float64)
		}
		if _, ok := weights[eq[1]]; !ok {
			weights[eq[1]] = make(map[string]float64)
		}
		weights[eq[0]][eq[1]] = values[i]
		weights[eq[1]][eq[0]] = 1.0 / values[i]
	}
	res := make([]float64, len(queries))
	for i, q := range queries {
		a, b := q[0], q[1]
		if _, ok := weights[a]; !ok {
			res[i] = -1
			continue
		}
		if _, ok := weights[b]; !ok {
			res[i] = -1
			continue
		}
		if a == b {
			res[i] = 1
			continue
		}
		res[i] = dfs(a, b, weights, make(map[string]bool))
	}
	return res
}

func dfs(a, b string, weights map[string]map[string]float64, visited map[string]bool) float64 {
	if visited[a] {
		return -1
	}
	visited[a] = true
	if w, ok := weights[a][b]; ok {
		return w
	}
	for k, v := range weights[a] {
		res := dfs(k, b, weights, visited)
		if res != -1 {
			return v * res
		}
	}
	return -1
}

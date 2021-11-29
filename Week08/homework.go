package Week08

//684. 冗余连接
func findRedundantConnection(edges [][]int) []int {
	parent := make([]int, len(edges)+1)
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(from, to int) bool {
		x, y := find(from), find(to)
		if x == y {
			return false
		}
		parent[x] = y
		return true
	}
	for _, e := range edges {
		if !union(e[0], e[1]) {
			return e
		}
	}
	return nil
}

//200. 岛屿数量
func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	ans := 0
	dx := []int{-1, 0, 0, 1}
	dy := []int{0, -1, 1, 0}

	fa := make([]int, m*n+1)
	for i := 0; i <= m*n; i++ {
		fa[i] = i
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				continue
			}
			for k := 0; k < 4; k++ {
				nx := i + dx[k]
				ny := j + dy[k]
				if nx >= m || ny >= n || nx < 0 || ny < 0 {
					continue
				} else if grid[nx][ny] == '1' {
					unionSet(fa, nums(n, i, j), nums(n, nx, ny))
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && Find(fa, nums(n, i, j)) == nums(n, i, j) {
				ans++
			}
		}
	}

	return ans

}

func nums(n, i, j int) int {
	return i*n + j
}

func Find(fa []int, x int) int {
	if fa[x] == x {
		return x
	}
	fa[x] = Find(fa, fa[x])
	return fa[x]
}

func unionSet(fa []int, x, y int) {
	x, y = Find(fa, x), Find(fa, y)
	if x != y {
		fa[x] = y
	}
}

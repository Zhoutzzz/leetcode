package hard

type pair struct {
	x, y int
}

func containVirus(isInfected [][]int) (ans int) {
	m, n := len(isInfected), len(isInfected[0])
	dirs := []pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for {
		neighbors := []map[pair]struct{}{}
		firewalls := []int{}
		// 广度优先，每一次找出所有有毒的区域
		for x, row := range isInfected {
			for y, col := range row {
				if col != 1 {
					continue
				}
				idx, firewall := len(neighbors)+1, 0
				row[y] = -idx
				neighbor := map[pair]struct{}{}
				neighbor[pair{x, y}] = struct{}{}
				// 当前位置有毒时，找出与它相关的其他中毒和没中毒的位置
				queue := []pair{{x, y}}
				for len(queue) > 0 {
					p := queue[0]
					queue = queue[1:]
					for _, d := range dirs {
						if x, y := d.x+p.x, d.y+p.y; 0 <= x && x < m && y >= 0 && y < n {
							// 如果找出的是中毒的位置，标记这个位置，下次就从这个位置接着找，没中毒就表示可以处理
							if isInfected[x][y] == 1 {
								queue = append(queue, pair{x, y})
								isInfected[x][y] = -idx
							} else if isInfected[x][y] == 0 {
								neighbor[pair{x, y}] = struct{}{}
								firewall++
							}
						}
					}
				}
				// 把此次找出的区域中没中毒的位置和中毒位置需要都防火墙数量都保存起来
				neighbors = append(neighbors, neighbor)
				firewalls = append(firewalls, firewall)
			}
		}

		// 没有病毒
		if len(neighbors) == 0 {
			break
		}

		// 找出病毒影响最大的区域
		idx := 0
		for i, v := range neighbors {
			if len(v) > len(neighbors[idx]) {
				idx = i
			}
		}

		// 把这个区域处理了，然后把有病毒的地方都复位下，标记该区域已经复位，其他区域还没有处理
		ans += firewalls[idx]
		for i := 0; i < len(isInfected); i++ {
			for j := 0; j < len(isInfected[i]); j++ {
				if isInfected[i][j] < 0 {
					if isInfected[i][j] != -idx-1 {
						isInfected[i][j] = 1
					} else {
						isInfected[i][j] = 2
					}
				}
			}
		}

		// 这一次病毒扩散后，要把除了该区域以外的所有未被毒的区域标记成被毒
		for i, v := range neighbors {
			if i != idx {
				for p := range v {
					isInfected[p.x][p.y] = 1
				}
			}
		}

		// 只剩下最后一个区域表示处理完
		if len(neighbors) == 1 {
			break
		}
	}

	return
}

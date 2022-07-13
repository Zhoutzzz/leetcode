package medium

func AsteroidCollision(asteroids []int) (res []int) {
	for _, v := range asteroids {
		flag := true
		for flag && v < 0 && len(res) > 0 && res[len(res)-1] > 0 {
			flag = res[len(res)-1] >= -v
			if flag && res[len(res)-1] < -v {
				res = res[:len(res)-1]
			}
		}
		if flag {
			res = append(res, v)
		}
	}
	return
}

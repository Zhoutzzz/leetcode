package medium

func alphabetBoardPath(target string) string {
	cx, cy := 0, 0
	res := []byte{}
	for _, c := range target {
		nx := int(c-'a') / 5
		ny := int(c-'a') % 5
		if nx < cx {
			for j := 0; j < cx-nx; j++ {
				res = append(res, 'U')
			}
		}
		if ny < cy {
			for j := 0; j < cy-ny; j++ {
				res = append(res, 'L')
			}
		}
		if nx > cx {
			for j := 0; j < nx-cx; j++ {
				res = append(res, 'D')
			}
		}
		if ny > cy {
			for j := 0; j < ny-cy; j++ {
				res = append(res, 'R')
			}
		}
		res = append(res, '!')
		cx = nx
		cy = ny
	}
	return string(res)
}

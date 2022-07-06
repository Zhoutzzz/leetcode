package medium

func reorganizeString(S string) string {
	arr := ['z']int{}
	prev := 'a' - 1
	res := ""
	for _, v := range S {
		arr[v]++
	}
	for j := 0; j < len(S); j++ {
		for i := 'a'; i < 'z'+1; i++ {
			if arr[i] == 0 {
				continue
			}
			if prev == i {
				return ""
			}
			res += string(i)
			prev = i
			arr[i]--
			break
		}
	}
	return res
}

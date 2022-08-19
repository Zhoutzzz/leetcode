package easy

func busyStudent(startTime []int, endTime []int, queryTime int) (ans int) {
	l := len(startTime)
	for i := 0; i < l; i++ {
		if startTime[i] < queryTime && queryTime <= endTime[i] {
			ans++
		}
	}
	return
}

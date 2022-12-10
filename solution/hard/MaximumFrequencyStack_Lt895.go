package hard

import "leetcode/solution/common"

type FreqStack struct {
	freq    map[int]int
	group   map[int][]int
	maxFreq int
}

func ConstructorFS() FreqStack {
	return FreqStack{map[int]int{}, map[int][]int{}, 0}
}

func (f *FreqStack) Push(val int) {
	f.freq[val]++
	f.group[f.freq[val]] = append(f.group[f.freq[val]], val)
	f.maxFreq = common.Max(f.maxFreq, f.freq[val])
}

func (f *FreqStack) Pop() int {
	val := f.group[f.maxFreq][len(f.group[f.maxFreq])-1]
	f.group[f.maxFreq] = f.group[f.maxFreq][:len(f.group[f.maxFreq])-1]
	f.freq[val]--
	if len(f.group[f.maxFreq]) == 0 {
		f.maxFreq--
	}
	return val
}

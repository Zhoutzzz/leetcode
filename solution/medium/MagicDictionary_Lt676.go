package medium

type MagicDictionary struct {
	dict []string
}

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	this.dict = dictionary
}

func (this *MagicDictionary) Search(searchWord string) bool {
	for _, v := range this.dict {
		if len(v) == len(searchWord) {
			diff := 0
			for i := 0; i < len(v); i++ {
				if v[i] != searchWord[i] {
					diff++
				}
				if diff > 1 {
					break
				}
			}
			if diff == 1 {
				return true
			}
		}
	}
	return false
}

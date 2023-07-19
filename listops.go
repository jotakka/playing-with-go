package exercism

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	prev := initial
	for val := range s {
		prev = fn(val, prev)
	}
	return prev
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {

	prev := initial
	for val := range s {
		prev = fn(prev, val)
	}
	return prev
}

func (s IntList) Filter(fn func(int) bool) IntList {
	outList := make(IntList, len(s))
	cnt := 0

	for val := range s {
		if fn(val) {
			outList[cnt] = val
			cnt++
		}
	}

	return outList[:cnt]
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	outList := make(IntList, len(s))

	idx := 0
	for val := range s {
		outList[idx] = fn(val)
		idx++
	}

	return outList
}

func (s IntList) Reverse() IntList {
	outList := make(IntList, len(s))

	j := 0
	for i := len(s) - 1; i >= 0; i-- {
		outList[j] = s[i]
		j++
	}

	return outList
}

func (s IntList) Append(lst IntList) IntList {
	outList := append(s, lst...)
	return outList
}

func (s IntList) Concat(lists []IntList) IntList {
	outList := s
	for list := range lists {
		outList = append(outList, list...)
	}
	return outList
}

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	prev := initial
	for _, val := range s {
		prev = fn(prev, val)
	}
	return prev
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {

	prev := initial
	var reverseS = s.Reverse()
	for _, val := range reverseS {
		prev = fn(val, prev)
	}
	return prev
}

func (s IntList) Filter(fn func(int) bool) IntList {
	outList := make(IntList, len(s))
	cnt := 0

	for _, val := range s {
		if fn(val) {
			outList[cnt] = val
			cnt++
		}
	}

	return outList[:cnt]
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	outList := make(IntList, len(s))

	idx := 0
	for _, val := range s {
		outList[idx] = fn(val)
		idx++
	}

	return outList
}

func (s IntList) Reverse() IntList {
	outList := make(IntList, len(s))

	j := 0
	for i := len(s) - 1; i >= 0; i-- {
		outList[j] = s[i]
		j++
	}

	return outList
}

func (s IntList) Append(lst IntList) IntList {
	outList := append(s, lst...)
	return outList
}

func (s IntList) Concat(lists []IntList) IntList {
	outList := s
	for _, list := range lists {
		outList = append(outList, list...)
	}
	return outList
}

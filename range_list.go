package main

import (
    "sort"
	"fmt"
)


type Range struct {
	Left, Right int
}
type RangeList struct {
	SortedRanges []Range
}

func NewRangeList() RangeList {
	return RangeList{SortedRanges: make([]Range, 0)}
}

func (this *RangeList) Add(arr [2]int) {
	var left = arr[0]
	var right = arr[1]
	if left > right{
		return
	}
	if this.SortedRanges == nil{
		//fmt.Println("SortedRanges is nil")
		this.SortedRanges = make([]Range,0)
	} 

	newRange := []Range{}
	i := 0
	var mergeRangeLeft = left
	var mergeRangeRight = right
	for i < len(this.SortedRanges) {
		if this.SortedRanges[i].Right < left {
			newRange = append(newRange, this.SortedRanges[i])
			i++
		} else {
			mergeRangeLeft = min(this.SortedRanges[i].Left, mergeRangeLeft)
			break
		}
	}
	j := i
	for j < len(this.SortedRanges) {
		if this.SortedRanges[j].Left > right {
			break
		} else {
			mergeRangeRight = max(mergeRangeRight, this.SortedRanges[j].Right)
			j++
		}
	}
	newRange = append(newRange, Range{mergeRangeLeft, mergeRangeRight})
	for i := j; i < len(this.SortedRanges); i++ {
		newRange = append(newRange, this.SortedRanges[i])
	}
	this.SortedRanges = newRange
}

func (this *RangeList) Query(arr [2]int) bool {
	var left = arr[0]
	var right = arr[1]

	index := sort.Search(len(this.SortedRanges), func(i int) bool {
		return this.SortedRanges[i].Left > left
	})
	if index == 0 {
		return false
	}

	shouldIncludeSet := this.SortedRanges[index-1]
	if shouldIncludeSet.Left <= left && shouldIncludeSet.Right >= right {
		return true
	} else {
		return false
	}
}

func (this *RangeList) Print()  {
	for i := 0; i < len(this.SortedRanges); i++ {
		fmt.Printf("[%d, %d) ",this.SortedRanges[i].Left,this.SortedRanges[i].Right)
	}
	fmt.Printf("\n")
}

func (this *RangeList) Remove(arr [2]int) {

	var left = arr[0]
	var right = arr[1]
	if left > right{
		return
	}
	
	newRange := []Range{}
	i := 0
	for i < len(this.SortedRanges) {
		if this.SortedRanges[i].Right < left {
			newRange = append(newRange, this.SortedRanges[i])
			i++
		} else {
			break
		}
	}
	j := i
	for j < len(this.SortedRanges) {
		if this.SortedRanges[j].Left > right {
			break
		} else {
			j++
		}
	}
	for p := i; p < j; p++ {
		curRange := this.SortedRanges[p]
		if left <= curRange.Left && right >= curRange.Right {
			continue
		}
		if curRange.Left < left {
			newRange = append(newRange, Range{curRange.Left, left})
			if curRange.Right > right {
				// split two
				newRange = append(newRange, Range{right, curRange.Right})
			}
		}else {
			newRange = append(newRange, Range{right, curRange.Right})
		}
	}
	for i := j; i < len(this.SortedRanges); i++ {
		newRange = append(newRange, this.SortedRanges[i])
	}
	this.SortedRanges = newRange
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main(){
	rl := RangeList{}
	//rl := NewRangeList();
	rl.Add([2]int{1, 5})
    rl.Print()
	// Should display: [1, 5)
	rl.Add([2]int{10, 20})
	rl.Print()
	// Should display: [1, 5) [10, 20)
	rl.Add([2]int{20, 20})
	rl.Print()
	// Should display: [1, 5) [10, 20)
	rl.Add([2]int{20, 21})
	rl.Print()
	// Should display: [1, 5) [10, 21)
	rl.Add([2]int{2, 4})
	rl.Print()
	// Should display: [1, 5) [10, 21)
	rl.Add([2]int{3, 8})
	rl.Print()
	// Should display: [1, 8) [10, 21)
	rl.Remove([2]int{10, 10})
	rl.Print()
	// Should display: [1, 8) [10, 21)
	rl.Remove([2]int{10, 11})
	rl.Print()
	// Should display: [1, 8) [11, 21)
	rl.Remove([2]int{15, 17})
	rl.Print()
	// Should display: [1, 8) [11, 15) [17, 21)
	rl.Remove([2]int{3, 19})
	rl.Print()
	// Should display: [1, 3) [19, 21)
}



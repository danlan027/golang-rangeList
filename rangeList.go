package RangeList

import (
	"errors"
	"fmt"
)

type RangeList struct {
	rl []section
}

type section struct {
	left int
	right int
}

func (r *RangeList) Add(rangeElement [2]int) error {
	left, right := rangeElement[0], rangeElement[1]
	if left > right {
		return errors.New("element error")
	}
	if len(r.rl) == 0 {
		r.rl = append(r.rl, section{left, right})
		return nil
	}

	leftLoc, leftLocRes := r.locate(left)
	rightLoc, rightLocRes := r.locate(right)

	if leftLoc == rightLoc {
		//一个子区间内操作，只需改变子区间边界，无需裁剪切片
		if leftLoc == -1 {
			r.rl = append(r.rl, section{})
			for i:= len(r.rl)-1; i>0; i-- {
				r.rl[i] = r.rl[i-1]
			}
			r.rl[0] = section{left, right}
		} else if leftLoc == len(r.rl) {
			r.rl = append(r.rl, section{left, right})
		} else {
			r.rl[leftLoc].left = min(left, r.rl[leftLoc].left)
		}
	} else {
		//跨子区间，改变边界同时需要裁剪切片
		if leftLocRes && rightLocRes {
			r.rl[leftLoc].right = r.rl[rightLoc].right
			tmp := append([]section{}, r.rl[0:leftLoc+1]...)
			tmp = append(tmp, r.rl[rightLoc+1:]...)
			r.rl = tmp
		} else if leftLocRes && !rightLocRes {
			r.rl[leftLoc].right = right
			tmp := append([]section{}, r.rl[0:leftLoc+1]...)
			tmp = append(tmp, r.rl[rightLoc:]...)
			r.rl = tmp
		} else if !leftLocRes && rightLocRes {
			r.rl[rightLoc].left = left
			tmp := append([]section{}, r.rl[0:leftLoc]...)
			tmp = append(tmp, r.rl[rightLoc:]...)
			r.rl = tmp
		} else {
			r.rl[leftLoc].left = left
			r.rl[leftLoc].right = right
			tmp := append([]section{}, r.rl[0:leftLoc+1]...)
			tmp = append(tmp, r.rl[rightLoc:]...)
			r.rl = tmp
		}
	}

	return nil
}

func (r *RangeList) Remove(rangeElement [2]int) error {
	left, right := rangeElement[0], rangeElement[1]
	if left > right {
		return errors.New("element error")
	}
	leftLoc, leftLocRes := r.locate(left)
	rightLoc, rightLocRes := r.locate(right)
	if leftLoc == rightLoc {
		//一个子区间内操作，只需改变子区间边界，无需裁剪切片
		if leftLoc == -1 || leftLoc == len(r.rl) || left == right {
			return nil
		} else {
			if leftLocRes && rightLocRes {
				if left > r.rl[leftLoc].left && right < r.rl[leftLoc].right {
					r.rl = append(r.rl, section{right, r.rl[leftLoc].right})
					r.rl[leftLoc].right = left
				} else if left == r.rl[leftLoc].left && right < r.rl[leftLoc].right {
					r.rl[leftLoc].left = right
				} else if left > r.rl[leftLoc].left && right == r.rl[leftLoc].right {
					r.rl[leftLoc].right = left
				}
			} else {
				r.rl[leftLoc].left = right
			}
		}
	} else {
		//跨子区间，改变边界同时需要裁剪切片
		if leftLocRes && rightLocRes {
			r.rl[leftLoc].right = left
			r.rl[rightLoc].left = right
			tmp := append([]section{}, r.rl[0:leftLoc+1]...)
			tmp = append(tmp, r.rl[rightLoc:]...)
			r.rl = tmp
		} else if leftLocRes && !rightLocRes {
			r.rl[leftLoc].right = left
			tmp := append([]section{}, r.rl[0:leftLoc+1]...)
			tmp = append(tmp, r.rl[rightLoc:]...)
			r.rl = tmp
		} else if !leftLocRes && rightLocRes {
			r.rl[rightLoc].left = right
			tmp := append([]section{}, r.rl[0:leftLoc]...)
			tmp = append(tmp, r.rl[rightLoc:]...)
			r.rl = tmp
		} else {
			r.rl[leftLoc].right = left
			r.rl[rightLoc].left = right
			tmp := append([]section{}, r.rl[0:leftLoc+1]...)
			tmp = append(tmp, r.rl[rightLoc:]...)
			r.rl = tmp
		}
	}
	return nil
}

func (r *RangeList) Print() error {
	output := ""
	for _, v := range r.rl {
		output += fmt.Sprintf("[%d,%d)\n", v.left, v.right)
	}
	fmt.Println(output)
	return nil
}

// 定位元素子区间位置
func (r *RangeList) locate(e int) (int, bool) {
	curL := len(r.rl)
	if e < r.rl[0].left {
		return -1, false
	}
	if e > r.rl[curL-1].right {
		return curL, false
	}

	for i, v := range r.rl {
		if e >= v.left && e <= v.right {
			return i, true
		}
		if e > v.right && e < r.rl[i+1].left {
			return i+1, false
		}
	}
	return -1, false
}

func min(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}
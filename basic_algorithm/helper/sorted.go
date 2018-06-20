package helper

import (
	"time"
	"math/rand"
	"errors"
	"fmt"
	"math"
)

type RandomArray struct {
	Array []interface{}
	SortedFunc func([]interface{}) time.Duration
}

var GlobalArray = RandomArray{}

func (r *RandomArray) GenarateRandomArray(n , start int ) {

	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i:=0;i<=n;i++ {
		r.Array = append(r.Array , random.Intn(n) + start)
	}
}

func (r *RandomArray) GenaratNearlySortedmArray(n , swapNum int ) {

	rand.Seed(time.Now().UnixNano())

	for i:=0;i<=n;i++ {
		r.Array = append(r.Array , i)
	}

	for  i := 0 ; i < swapNum ; i ++ {
		x:=rand.Intn(n)
		y:=rand.Intn(n)
		r.Array[x] , r.Array[y] = r.Array[y] , r.Array[x]
	}

}

func MergeSortedArray (arr [] interface{}) time.Duration{
	t1 := time.Now()
	mergeSortedArray(arr , 0 , len(arr) - 1)
	return time.Since(t1)
}

func mergeSortedArray(arr []interface{} , start , end int) {
	if start >= end {
		return
	}
	mid := (start + end ) /2

	mergeSortedArray(arr , start , mid)
	mergeSortedArray(arr , mid+1 , end)
	// 必须定义end  因为arr传递的是全部的数组
	// 如果左边的数组 已经大于右边了 可以不进行merge的操作
	b, _ := CheckVal(arr[mid], arr[mid+1])
	if !b{
		merge(arr , start , mid , end)
	}

}

// 归并排序的方法就是 首先需要先定义两个一样的数组  相当于copy 一份 然后 进行比较
func merge(arr []interface{} , start ,mid, end int) {
	res := make([]interface{} , end-start + 1)
	// 如果数组数据比较少的话  可以使用其他的排序算法

	for  j := start ; j <= end ; j ++  {
		res[j-start] = arr[j]
	}

	// 定义判断两个函数的起始位置
	n:=start
	m:=mid+1

	for i:=start;i<=end;i++ {
		// 先判断 左右两边的数组是否越界
		if n > mid {
			arr[i] = res[m-start]
			m++
			continue
		}

		if m > end {
			arr[i] = res[n-start]
			n++
			continue
		}

		b, e := CheckVal(res[n-start], res[m-start])
		if e != nil {
			panic(e)
		}

		if b {
			arr[i] = res[n-start]
			n++
			continue
		}

		arr[i] = res[m-start]
		m++
	}

}


func (r *RandomArray) findPos(pos , end int)int {
	// 随机选取一个作为坐标点
	r.Array[rand.Intn(end-pos+1)+pos] , r.Array[pos] =  r.Array[pos] , r.Array[rand.Intn(end-pos+1)+pos]

	val := r.Array[pos]
	i := pos+ 1
	for i<=end {
		b, _ := CheckVal(r.Array[i] , val)
		if b {
			r.Array[i] , r.Array[pos] = r.Array[pos], r.Array[i]
			pos++
			i++
		} else {
			r.Array[i], r.Array[end] = r.Array[end], r.Array[i]
			end--
		}

	}
	return pos
}


func (r *RandomArray) QuickSort( start , end int)([]interface{}) {
	if start < end {
		pos := r.findPos(start , end)
		r.QuickSort(start, pos-1)
		r.QuickSort(pos+1, end)
	}
	return r.Array
}


func (r *RandomArray) MergeButtomArray(n int) {
	var step int
	for step = 1 ; step <= n ; step+=step {
		for start:=0;start<n;start+=step+step {
			merge(r.Array , start , start + step -1 , Min(start + step+step-1 , n-1))
		}
	}
}

func Min(a , b int)int {
	if a < b {
		return a
	}

	return b
}

func CheckVal(a,b interface{}) (bool , error) {
	switch a.(type) {
	case float64:
		return a.(float64) <= b.(float64) , nil
	case float32:
		return a.(float32) <= b.(float32) , nil
	case int:
		return a.(int) <= b.(int) , nil
	case int64:
		return a.(int64) <= b.(int64) , nil
	case int32:
		return a.(int32) <= b.(int32) , nil
	case string:
		return a.(string) <= b.(string) , nil
	default:
		return false , errors.New("类型有误")
	}
}

func (r *RandomArray) CheckSorted()bool {
	for i:=0;i<len(r.Array)-1;i++ {
		b, _ := CheckVal(r.Array[i], r.Array[i+1])
		if !b{
			return false
		}
	}

	return true
}

// 堆的实现
type MaxHeap struct {
	HeapData map[int]interface{} // 存储空间
	Count int  // 堆里面元素的数量
	Indexes map[int]int
}

func (h *MaxHeap) InitHeap() {
	h.Count = 0
	h.HeapData = make(map[int]interface{})
}

func (h* MaxHeap) InitHeap2(arr [] interface{} , n int) {
	h.HeapData = make(map[int]interface{})

	for i:=0;i<n;i++ {
		h.HeapData[h.Count+1] = arr[i]
		h.Count++
	}
	// 每次都是从父节点 开始进行调整 将子堆先调整为最大堆  然后依次调整 最后得到的就是最大堆
	for j:=h.Count /2 ; j >= 1;j-- {
		h.shiftDown(j)
	}

}

func (h *MaxHeap) Insert(val int) {
	h.HeapData[h.Count+1] = val
	h.Count ++
	h.shiftUp(h.Count)
}

func  (h *MaxHeap) shiftUp (k int) {
	for k > 1 && h.HeapData[k].(int) > h.HeapData[k/2].(int) {
		h.HeapData[k] , h.HeapData[k/2] = h.HeapData[k/2] , h.HeapData[k]
		k = k/2
	}
}

// 向下调整 直到堆成为一个最大堆
func (h *MaxHeap) shiftDown(k int) {
	// 判断元素是否有子节点
	for 2 * k <= h.Count {
		left := 2*k
		if left +1 <= h.Count && h.HeapData[left].(int) < h.HeapData[left+1].(int) {
			left ++
		}

		if h.HeapData[k].(int) >= h.HeapData[left].(int) {
			break
		}
		// TODO 一次性替换
		h.HeapData[k] , h.HeapData[left] = h.HeapData[left] , h.HeapData[k]

		k = left
	}


}

func (h *MaxHeap) PrintVal(start int) int {
	val := h.HeapData[start]
	h.HeapData[start] , h.HeapData[h.Count] = h.HeapData[h.Count] , h.HeapData[start]
	h.Count--
	if start == 1 {
		h.shiftDown(start)

	} else {
		h.shiftDown2(h.Count-1 , start)

	}
	return val.(int)
}

func (h *MaxHeap) CheckIsMaxHeap(start int) bool {
	for i:=start;i<=h.Count;i++ {
		for 2*i <= h.Count {
			val := h.HeapData[2*i]
			if 2*i +1 <= h.Count && h.HeapData[2*i +1].(int) < h.HeapData[2*i].(int) {
				val = h.HeapData[2*i+1]
			}
			if h.HeapData[i].(int) < val.(int) {
				return false
			}
			i = 2*i
		}
	}

	return true
}

func (h *MaxHeap) shiftDown2(length , k int) {
	val := GlobalArray.Array[k]
	// 判断元素是否有子节点
	for 2 * k +1 < length {
		left := 2*k +1
		if left +1 < length && GlobalArray.Array[left].(int) <GlobalArray.Array[left+1].(int) {
			left ++
		}

		if val.(int) >= GlobalArray.Array[left].(int) {
			break
		}
		// TODO 一次性替换
		GlobalArray.Array[k] = GlobalArray.Array[left]

		k = left
	}

	GlobalArray.Array[k] = val
}


func (h *MaxHeap) Heap3( n int) {
	for i:=n-1/2;i>=0;i-- {
		h.shiftDown2(n, i)
	}

	for i:=n-1;i>=0;i-- {
		GlobalArray.Array[0] , GlobalArray.Array[i] = GlobalArray.Array[i] , GlobalArray.Array[0]

		h.shiftDown2(i, 0)
	}
}

func (h *MaxHeap) InitIndexMaxHeap(arr [] interface{} , n int) {
	h.HeapData = make(map[int]interface{})
	h.Indexes = make(map[int]int)

	for i:=0;i< n;i++ {
		h.HeapData[h.Count+1] = arr[i]
		h.Indexes[i+1] = i+1
		h.Count++
	}
	// 每次都是从父节点 开始进行调整 将子堆先调整为最大堆  然后依次调整 最后得到的就是最大堆
	for j:=h.Count /2 ; j >= 1;j-- {
		h.shiftDownIndex(j)
	}
}

func (h *MaxHeap) shiftDownIndex(k int) {
	// 判断元素是否有子节点
	for 2 * k <= h.Count {
		left := 2*k
		if left +1 <= h.Count && h.HeapData[h.Indexes[left]].(int) < h.HeapData[h.Indexes[left+1]].(int) {
			left ++
		}
		if h.HeapData[h.Indexes[k]].(int) >= h.HeapData[h.Indexes[left]].(int) {
			break
		}
		// TODO 一次性替换
		h.Indexes[k] , h.Indexes[left] = h.Indexes[left] , h.Indexes[k]

		k = left
	}
}

func  (h *MaxHeap) shiftUpIndex (k int) {
	for k > 1 && h.HeapData[h.Indexes[k]].(int) > h.HeapData[h.Indexes[k/2]].(int) {
		h.Indexes[k] , h.Indexes[k/2] = h.Indexes[k/2] , h.Indexes[k]
		k = k/2
	}
}

func (h *MaxHeap) ChangeIndexVal(n int, val interface{}) {
	h.HeapData[n+1] = val

	for i := 1; i <= h.Count ;i ++ {
		h.shiftUpIndex(i)
		h.shiftDownIndex(i)
	}
}

func (h *MaxHeap) PrintIndexVal(start int) int {
	index := h.Indexes[start]
	val := h.HeapData[index]
	h.HeapData[index] , h.HeapData[h.Count] = h.HeapData[h.Count] , h.HeapData[index]
	h.Count--
	h.shiftDown(start)
	return val.(int)
}

// 以树状打印整个堆结构
func (h *MaxHeap) TestPrint() {

	// 我们的testPrint只能打印100个元素以内的堆的树状信息
	if h.Count >= 100 {
		fmt.Println("This print function can only work for less than 100 int")
		return
	}

	fmt.Println("The max heap size is: ", h.Count)
	fmt.Println("Data in the max heap: ")
	for i := 1; i <= h.Count; i ++ {
		// 我们的testPrint要求堆中的所有整数在[0, 100)的范围内
		b, _ := CheckVal(h.HeapData[i], 0)
		b2, _ := CheckVal(h.HeapData[i], 100)
		if b || b2 {
			continue
		}
		fmt.Println(h.HeapData[i], " ")
	}
	n := h.Count
	maxLevel := 0
	numberPerLevel := 1
	for n > 0 {
		maxLevel += 1
		n -= numberPerLevel
		numberPerLevel *= 2
	}


	maxLevelNumber := int(math.Pow(2, float64(maxLevel-1)))
	curTreeMaxLevelNumber := maxLevelNumber
	index := 1
	for level := 0; level < maxLevel; level ++ {

		var line1 string
		tmp := maxLevelNumber*3 - 1
		for tmp >0 {
			line1 += " "
			tmp --
		}

		curLevelNumber := int(math.Min(float64(h.Count)-math.Pow(2, float64(level))+1, math.Pow(2, float64(level))))
		isLeft := true
		indexCurLevel := 0
		for indexCurLevel < curLevelNumber{
			putNumberInLine( h.HeapData[index].(int) , &line1 , indexCurLevel , curTreeMaxLevelNumber*3-1 , isLeft )
			isLeft = !isLeft
			index ++
			indexCurLevel ++
		}

		fmt.Println(line1)

		if level == maxLevel-1 {
			break
		}


		var line2 string
		tmp2 := maxLevelNumber*3 - 1
		for tmp2 >0 {
			line2 += " "
			tmp2 --
		}

		for indexCurLevel := 0; indexCurLevel < curLevelNumber; indexCurLevel ++ {
			putBranchInLine(&line2 , indexCurLevel , curTreeMaxLevelNumber*3-1 )

			curTreeMaxLevelNumber /= 2
		}
	}
}


func putNumberInLine(num int, line *string, indexCurLevel, curTreeWidth int,  isLeft bool){
	subTreeWidth := (curTreeWidth - 1) / 2
	offset := indexCurLevel* (curTreeWidth+1) + subTreeWidth
	lines :=  []byte(*line)
	if offset + 1 < len(lines) {
		if num >= 10  {
			lines[offset + 0] = byte('0' + num / 10)
			lines[offset + 1] = byte('0' + num % 10)
		} else {
			if isLeft {
				lines[offset+1] = byte('0' + num)
			} else {
				lines[offset+1] = byte('0' + num)
			}
		}
		*line = string(lines)
	}

}

func putBranchInLine( line *string,  indexCurLevel, curTreeWidth int){

	subTreeWidth := (curTreeWidth - 1) / 2
	subSubTreeWidth := (subTreeWidth - 1) / 2

	offsetLeft := indexCurLevel* (curTreeWidth+1) + subSubTreeWidth
	offsetRight := indexCurLevel* (curTreeWidth+1) + subTreeWidth + 1 + subSubTreeWidth
	lines :=  []byte(*line)
	if offsetLeft+ 1 < len(lines)  &&  offsetRight < len(lines)  {
		lines[offsetLeft+1] = byte('/')
		lines[offsetRight+1] = byte('\\')
	}
	*line = string(lines)
}

package main

import (
	"github.com/zhongliangshan/test/basic_algorithm/helper"
	"math/rand"
	"time"
	"fmt"
)

func main() {
	//helper.GlobalArray.GenaratNearlySortedmArray(n , 1000)
	//////fmt.Println(helper.GlobalArray.Array)
	////////helper.MergeSortedArray(helper.GlobalArray.Array)
	//////helper.GlobalArray.MergeButtomArray(len(helper.GlobalArray.Array))
	//////fmt.Println(helper.GlobalArray.Array)
	//////
	////
	////
	////rand.Seed(time.Now().UnixNano())
	//////fmt.Println(helper.GlobalArray.Array)
	////t1 := time.Now()
	////helper.GlobalArray.Array = helper.GlobalArray.QuickSort(0 , len(helper.GlobalArray.Array)-1)
	////
	////fmt.Println(time.Since(t1))
	//////fmt.Println(helper.GlobalArray.Array)
	//
	//t1 := time.Now()
	//maxHeap := helper.MaxHeap{}
	//maxHeap.Heap3(n)
	////fmt.Println(time.Since(t1))
	//if !helper.GlobalArray.CheckSorted() {
	//	fmt.Println("sorted error ")
	//}
	//fmt.Println(helper.GlobalArray.Array)
	//rand.Seed(time.Now().UnixNano())
	//for i:=0;i<15;i++ {
	//	maxHeap.Insert(rand.Intn(100))
	//}
	//
	//if !maxHeap.CheckIsMaxHeap() {
	//	fmt.Println("is not maxheap")
	//}
	//
	//for i:=0;i<1000000;i++ {
	//
	//	fmt.Printf("%d " , maxHeap.PrintVal())
	//
	//}

	//maxHeap.InitIndexMaxHeap(helper.GlobalArray.Array , n)
	//
	//for i :=1 ; i <= n;i++ {
	//	fmt.Printf("%d  " , maxHeap.PrintIndexVal(i))
	//}
	n := 1000

	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i:=0;i<n;i++ {
		helper.Root.Insert(i, random.Intn(n) + 1000)
	}


	//helper.Root.LevelOrder()
	//node := helper.Root.DelMinNode(helper.Root.Root)
	//minKey:= helper.Root.FindMinNode(node)
	//fmt.Println(minKey)
	//
	//node2 := helper.Root.DelMaxNode(helper.Root.Root)
	//maxKey := helper.Root.FindMaxNode(node2)
	//fmt.Println(maxKey)

	root := helper.Root.Root

	node , _ := helper.Root.Search(root , 100)
	fmt.Println(node)


	root = helper.Root.DelNode(root , 100)
	node2 , err := helper.Root.Search(root , 100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(node2)

	}


}

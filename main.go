package main

import "fmt"

func MergeSort(list []int) []int{
	length:=len(list)
	if length >= 2{
		mid := length/2
		list = Merge(MergeSort(list[:mid]),MergeSort(list[mid:]))
	}
	return list
}

func Merge(left []int,right []int) []int{

	lst:=make([]int,0)

	for len(left) > 0 && len(right) > 0{

		if left[0] < right[0]{
			lst = append(lst,left[0])
			left = left[1:]
		}else{
			lst = append(lst,right[0])
			right = right[1:]
		}
	}

	if len(left) > 0{
		lst = append(lst,left...)
	}
	if  len(right) > 0{
		lst = append(lst,right...)
	}

	return lst
}

func main(){

	lst := []int{-342,-1223,443,223,3,5,6,432,1}
	fmt.Println(MergeSort(lst))
}
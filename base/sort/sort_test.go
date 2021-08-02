package sort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	ints()
	strings()
	days()
}

func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := IntSlice(data)
	Sort(a)
	if !IsSorted(a) {
		panic("fails")
	}
	fmt.Printf("The sorted array is: %v\n", a)
}

func strings() {
	data := []string{"monday", "friday", "tuesday", "wednesday", "sunday", "thursday", "", "saturday"}
	a := StringSlice(data)
	Sort(a)
	if !IsSorted(a) {
		panic("fail")
	}
	fmt.Printf("The sorted array is: %v\n", a)
}

func days() {
	Sunday := Day{0, "SUN", "Sunday"}
	Monday := Day{1, "MON", "Monday"}
	Tuesday := Day{2, "TUE", "Tuesday"}
	Wednesday := Day{3, "WED", "Wednesday"}
	Thursday := Day{4, "THU", "Thursday"}
	Friday := Day{5, "FRI", "Friday"}
	Saturday := Day{6, "SAT", "Saturday"}
	data := []*Day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
	a := DaySlice{data}
	Sort(&a)
	if !IsSorted(&a) {
		panic("fail")
	}
	for _, d := range data {
		fmt.Printf("%s ", d.LongName)
	}
	fmt.Printf("\n")
}

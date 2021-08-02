package sort

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

//冒泡排序
func Sort(data Sorter) {
	for pass := 1; pass < data.Len(); pass++ {
		for j := 0; j < data.Len()-pass; j++ {
			if data.Less(j+1, j) {
				data.Swap(j, j+1)
			}
		}
	}
}

func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

//各个类型的实现排序接口
type IntSlice []int

func (p IntSlice) Len() int { return len(p) }
func (p IntSlice) Less(i, j int) bool {
	return p[i] < p[j]
}
func (p IntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

//对外暴露的接口方法
func SortInts(a []int)       { Sort(IntSlice(a)) }
func SortStrings(a []string) { Sort(StringSlice(a)) }

type Day struct {
	Num       int
	ShortName string
	LongName  string
}

type DaySlice struct {
	Data []*Day
}

//对用户自定义类型 DaySlice 进行排序
func (p *DaySlice) Len() int           { return len(p.Data) }
func (p *DaySlice) Less(i, j int) bool { return p.Data[i].Num < p.Data[j].Num }
func (p *DaySlice) Swap(i, j int)      { p.Data[i], p.Data[j] = p.Data[j], p.Data[i] }

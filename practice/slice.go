package practice

func SliceInsert(slice []int, index int, value int) []int {
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = value
	return slice
}

func SliceRemove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	len := len(slice)
	if index >= 0 && index < len {
		return slice[index]
	}
	return -1
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	len := len(slice)
	if index >= 0 && index < len {
		slice[index] = value
		return slice
	} else {
		return append(slice, value)
	}
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	if len(values) > 0 {
		return append(values, slice...)
	}
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	len := len(slice)
	if index >= 0 && index < len {
		return append(slice[:index], slice[index+1:]...)
	} else {
		return slice
	}
}

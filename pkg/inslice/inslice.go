package inslice

func String(a string, list *[]string) bool {
	for _, b := range *list {
		if b == a {
			return true
		}
	}

	return false
}

func Int(a int, list *[]int) bool {
	for _, b := range *list {
		if b == a {
			return true
		}
	}

	return false
}

func Int32(a int32, list *[]int32) bool {
	for _, b := range *list {
		if b == a {
			return true
		}
	}

	return false
}

func Int64(a int64, list *[]int64) bool {
	for _, b := range *list {
		if b == a {
			return true
		}
	}

	return false
}

func Float32(a float32, list *[]float32) bool {
	for _, b := range *list {
		if b == a {
			return true
		}
	}

	return false
}

func Float64(a float64, list *[]float64) bool {
	for _, b := range *list {
		if b == a {
			return true
		}
	}

	return false
}

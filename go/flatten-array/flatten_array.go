package flatten

func Flatten(nested interface{}) []interface{} {
	result := []interface{}{}
	for _, element := range nested.([]interface{}) {
		if element == nil {
			continue
		}
		if slice, ok := element.([]interface{}); ok {
			result = append(result, Flatten(slice)...)
		} else {
			result = append(result, element)
		}
	}
	return result
}

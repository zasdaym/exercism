package flatten

func Flatten(nested interface{}) []interface{} {
	flattened := make([]interface{}, 0)
	switch nested := nested.(type) {
	case []interface{}:
		for _, item := range nested {
			flattened = append(flattened, Flatten(item)...)
		}
	default:
		if nested != nil {
			flattened = append(flattened, nested)
		}
	}
	return flattened
}

// Package flatten is solution for problem Flatten Array.
package flatten

func dfs(in interface{}, out *[]interface{}) {
	s := in.([]interface{})
	for _, e := range s {
		if e != nil {
			switch v := e.(type) {
			case int:
				*out = append(*out, v)
			case []interface{}:
				dfs(v, out)
			}
		}
	}
}

// Flatten flattens given deep nested slice.
func Flatten(l interface{}) []interface{} {
	r := []interface{}{}
	dfs(l, &r)
	return r
}

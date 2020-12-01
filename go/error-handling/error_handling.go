// Package erratum is solution for problem Error Handling.
package erratum

// Use frobs a string to a resource.
func Use(opener ResourceOpener, s string) (err error) {
	resource, err := opener()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		resource, err = opener()
	}
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				resource.Defrob(e.defrobTag)
				err = e
			case error:
				err = e
			}
		}
		resource.Close()
	}()
	resource.Frob(s)
	return nil
}

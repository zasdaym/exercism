// Package erratum provides Error Handling solution.
package erratum

// Use uses resource from opener to frob s.
func Use(opener ResourceOpener, s string) (err error) {
	var res Resource

	for res, err = opener(); err != nil; res, err = opener() {
		if _, ok := err.(TransientError); !ok {
			return err
		}
	}

	defer res.Close()
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case FrobError:
				res.Defrob(e.defrobTag)
				err = e
			case error:
				err = e
			}
		}
	}()

	res.Frob(s)
	return err
}

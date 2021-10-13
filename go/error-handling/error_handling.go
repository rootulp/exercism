package erratum

import (
	"errors"
	"fmt"
)

type opener func() (Resource, error)

func Use(open opener, s string) (err error) {
	resource, err := retry(5, open)
	if err != nil {
		return err
	}
	defer func(resource Resource) {
		resource.Close()
	}(resource)
	defer func(resource Resource) {
		if r := recover(); r != nil {
			fmt.Printf("recovered from %v\n", r)
			if errors.Is(r.(error), FrobError{}) {
				resource.Defrob(fmt.Sprintf("%v", r.(FrobError).defrobTag))
			}
			err = fmt.Errorf("%v", r)
		}
	}(resource)
	resource.Frob(s)
	return nil
}

func retry(attempts int, f func() (Resource, error)) (Resource, error) {
	for i := 0; ; i++ {
		r, err := f()
		if err == nil {
			return r, nil
		}
		if i >= (attempts - 1) {
			return nil, err
		}
	}
}

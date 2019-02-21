package mibici

type domain struct {
	controller interface{}
}

func newDomain(usecases UseCases) (*domain, error) {
	controller, err := newController(usecases)
	if err != nil {
		return nil, err
	}

	return &domain{controller}, nil
}

package domain

type BoilerplateInMemoryRepository struct {
	boilerplates map[string]*Boilerplate
}

func NewBoilerplateInMemoryRepository() *BoilerplateInMemoryRepository {
	return &BoilerplateInMemoryRepository{
		boilerplates: make(map[string]*Boilerplate),
	}
}

func (s *BoilerplateInMemoryRepository) Create(boilerplate *Boilerplate) error {
	s.boilerplates[boilerplate.Id] = boilerplate

	return nil
}

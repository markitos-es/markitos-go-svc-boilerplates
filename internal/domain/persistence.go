package domain

type BoilerplateRepository interface {
	Create(boiler *Boilerplate) error
}

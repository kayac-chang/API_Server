package game

import "api/model"

func (it *Usecase) FindByName(name string) (*model.Game, error) {

	return it.repo.FindByName(name)
}

func (it *Usecase) FindByID(id string) (*model.Game, error) {

	return it.repo.FindByID(id)
}

func (it *Usecase) FindAll() ([]*model.Game, error) {

	return it.repo.FindAll()
}

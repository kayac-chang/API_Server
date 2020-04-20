package game

import (
	"api/model"
	"net/http"

	errs "github.com/pkg/errors"
)

func (it Usecase) FindByName(name string) (*model.Game, error) {

	game, _err := it.repo.Find(name)

	if _err != nil {
		msg := "Request game not found"

		err := &model.Error{
			Code:    http.StatusNotFound,
			Message: errs.WithMessage(_err, msg).Error(),
		}

		return nil, err
	}

	return game, nil
}

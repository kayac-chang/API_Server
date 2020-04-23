package admin

// func (it *Usecase) CheckUser(req map[string]string) (*model.Token, error) {

// 	email := req["email"]
// 	password := req["password"]

// 	admin, err := it.repo.FindByID(newID(email))
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := utils.CompareHash(admin.Password, password); err != nil {
// 		return nil, err
// 	}

// 	token, err := jwt.Sign(it.env)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Store in Cache
// 	admin.Token = token.AccessToken
// 	it.repo.Store("Cache", admin)

// 	return token, nil
// }

// func (it *Usecase) CheckToken(token string) (*model.Admin, error) {

// 	admin, err := it.repo.FindByToken(token)

// 	if err != nil {

// 		return nil, err
// 	}

// 	return admin, nil
// }

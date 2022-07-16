package internal

type Application struct {
	repository Repository
}

func NewApplication(repository Repository) *Application {
	return &Application{repository: repository}
}

func (app *Application) Create(request CreateRequest) (CreateResponse, error) {
	for _, member := range app.repository.data {
		if member.UserName == request.UserName {
			return CreateResponse{}, nil
		}
	}
	res, err := app.Create(CreateRequest{request.UserName, request.MembershipType})
	if err != nil {
		return CreateResponse{}, nil
	}
	return CreateResponse{res.ID, res.UserName, res.MembershipType}, nil
}

func (app *Application) Update(request UpdateRequest) (UpdateResponse, error) {
	var foundUser Membership

	for _, member := range app.repository.data {
		if member.UserName == request.UserName {
			foundUser = member
			break
		}
	}
	res, err := app.Update(UpdateRequest{foundUser.ID, foundUser.UserName, foundUser.MembershipType})
	if err != nil {
		return UpdateResponse{}, err
	}

	return UpdateResponse{res.ID, res.UserName, res.MembershipType}, nil
}

func (app *Application) Delete(id string) (DeleteResponse, error) {
	var foundUser Membership

	for _, member := range app.repository.data {
		if member.ID == id {
			foundUser = member
			break
		}
	}

	res, err := app.Delete(foundUser.ID)
	if err != nil {
		return DeleteResponse{}, err
	}

	return DeleteResponse{res.ID}, nil
}

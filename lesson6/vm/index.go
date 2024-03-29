package vm

import "go-mega/lesson6/model"

// IndexViewModel struct
type IndexViewModel struct {
	BaseViewModel
	model.User
	Posts []model.Post
}

// IndexViewModelOp struct
type IndexViewModelOp struct{}

func (IndexViewModelOp) GetVM() IndexViewModel {
	u1, _ := model.GetUserByUsername("rene")
	posts, _ := model.GetPostsByUserID(u1.ID)
	v := IndexViewModel{BaseViewModel{Title: "Homepage"}, *u1, *posts}
	return v
}
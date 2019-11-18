package vm

// LoginViewModel struct
type LoginViewModel struct {
	BaseViewModel
}

type LoginViewModelOp struct{}

// GetVM func
func (l *LoginViewModelOp) GetVM() LoginViewModel {
	v := LoginViewModel{}
	v.SetTitle("Login")
	return v
}
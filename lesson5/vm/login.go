package vm

// LoginViewModel struct
type LoginViewModel struct {
	BaseViewModel
	Errs []string
}

type LoginViewModelOp struct{}

// GetVM func
func (l *LoginViewModelOp) GetVM() LoginViewModel {
	v := LoginViewModel{}
	v.SetTitle("Login")
	return v
}

// AddError func
func (v *LoginViewModel) AddError(errs ...string) {
	v.Errs = append(v.Errs, errs...)
}
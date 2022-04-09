package service

type AuthService interface {
	Login() bool
	Register() bool
}

type Authenticator struct {
	jwtUtil JWTService
}

func NewAuthService(jwtUtil JWTService) *Authenticator {
	return &Authenticator{jwtUtil: jwtUtil}
}

func (a *Authenticator) Login() bool {
	//TODO implement me
	panic("implement me")
}

func (a *Authenticator) Register() bool {
	//TODO implement me
	panic("implement me")
}

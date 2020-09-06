package app

import (
	"context"
	"time"

	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
	"github.com/DanielTitkov/tinig-demo-server/internal/ent/user"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (a *App) CreateUser(u *domain.User) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	_, err = a.db.User.
		Create().
		SetUsername(u.Username).
		SetEmail(u.Email).
		SetPasswordHash(string(hash)).
		Save(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (a *App) GetUser(u *domain.User) (*domain.User, error) {
	user, err := a.db.User.
		Query().
		Where(user.UsernameEQ(u.Username)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return &domain.User{
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (a *App) ValidateUserPassword(u *domain.User) (bool, error) {
	user, err := a.db.User.
		Query().
		Where(user.UsernameEQ(u.Username)).
		Only(context.Background())
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(u.Password))
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (a *App) GetUserToken(u *domain.User) (string, error) {
	user, err := a.db.User.
		Query().
		Where(user.UsernameEQ(u.Username)).
		Only(context.Background())
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(a.cfg.Auth.Exp)).Unix()

	t, err := token.SignedString([]byte(a.cfg.Auth.Secret))
	if err != nil {
		return "", err
	}

	return t, nil
}

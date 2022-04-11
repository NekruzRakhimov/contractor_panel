package controller

import (
	"contractor_panel/application/cerrors"
	"contractor_panel/application/dto"
	"contractor_panel/application/respond"
	"contractor_panel/application/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type SignController struct {
	s service.SignService
}

func NewSignController(s service.SignService) *SignController {
	return &SignController{s}
}

func (c *SignController) HandleRoutes(r *mux.Router) {
	r.HandleFunc("/signin", c.SignIn).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/refresh", c.Refresh).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/signout", c.SignOut).Methods(http.MethodOptions, http.MethodPost)
}

func (c *SignController) SignIn(w http.ResponseWriter, r *http.Request) {
	requestDto := &dto.CredentialsDto{}
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&requestDto)
	if err != nil {
		respond.WithError(w, r, cerrors.ErrCouldNotDecodeBody(err))
		return
	}

	credentials := dto.ConvertCredentialsDtoToEntity(requestDto)

	ctx := r.Context()

	tokens, err := c.s.SignIn(ctx, credentials)
	if err != nil {
		respond.WithError(w, r, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "accessToken",
		Value:    tokens.AccessToken,
		Expires:  time.Unix(tokens.AtExpires, 0),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    tokens.RefreshToken,
		Expires:  time.Unix(tokens.RtExpires, 0),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	respond.With(w, r, true)
}

func (c *SignController) Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refreshToken")
	if err != nil {
		respond.WithError(w, r, cerrors.ErrCouldNotVerifyToken(err))
		return
	}

	// Get the JWT string from the cookie
	tokenString := cookie.Value

	ctx := r.Context()

	tokens, err := c.s.RefreshToken(ctx, tokenString)
	if err != nil {
		respond.WithError(w, r, cerrors.ErrCouldNotVerifyToken(err))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "accessToken",
		Value:    tokens.AccessToken,
		Expires:  time.Unix(tokens.AtExpires, 0),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    tokens.RefreshToken,
		Expires:  time.Unix(tokens.RtExpires, 0),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})

	respond.With(w, r, true)
}

func (c *SignController) SignOut(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("refreshToken")
	if err != nil {
		respond.WithError(w, r, cerrors.ErrCouldNotVerifyToken(err))
		return
	}

	// Get the JWT string from the cookie
	tokenString := cookie.Value

	ctx := r.Context()

	if err = c.s.SignOut(ctx, tokenString); err != nil {
		respond.WithError(w, r, cerrors.ErrCouldNotVerifyToken(err))
		return
	}

	respond.With(w, r, true)
}

package auth

import (
	"net/http"
	"simpleAPI/core/apierrors"
	"simpleAPI/internal/service"
)

// Refresh handles token refreshing
func (a *Authentication) Refresh(w http.ResponseWriter, r *http.Request) {
	usr, err := a.svc.Auth().Refresh(r.Context(), r.Body, a.refreshKey)
	defer r.Body.Close()
	if err != nil {
		switch err {
		case service.ErrDecodeProcess:
			apierrors.HandleHTTPErr(w,
				err,
				http.StatusBadRequest)
			return
		case service.ErrParseToken:
			apierrors.HandleHTTPErr(w, err, http.StatusForbidden)
			return
		case service.ErrTokenInvalid:
			apierrors.HandleHTTPErr(w, apierrors.ErrInvalidAuth, http.StatusForbidden)
			return
		default:
			apierrors.HandleHTTPErr(w, err, http.StatusInternalServerError)
			return
		}

	}
	a.sendToken(w, usr)
}

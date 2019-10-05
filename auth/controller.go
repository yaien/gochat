package auth

import (
	"fmt"
	"net/http"

	"github.com/stretchr/objx"

	"github.com/gorilla/mux"
	"github.com/stretchr/gomniauth"
)

// Callback do authentication proccess after provider's redirect
func Callback(w http.ResponseWriter, r *http.Request) {
	providerName := mux.Vars(r)["provider"]
	provider, err := gomniauth.Provider(providerName)
	if err != nil {
		message := fmt.Sprintf("Failed to load provider %s: %s", providerName, err)
		http.Error(w, message, http.StatusBadRequest)
		return
	}
	data := objx.MustFromURLQuery(r.URL.RawQuery)
	credentials, err := provider.CompleteAuth(data)
	if err != nil {
		message := fmt.Sprintf("Failed to load credentials %s", err)
		http.Error(w, message, http.StatusBadRequest)
		return
	}
	user, err := provider.GetUser(credentials)
	if err != nil {
		message := fmt.Sprintf("Failed to load user info %s", err)
		http.Error(w, message, http.StatusBadRequest)
		return
	}
	payload := map[string]interface{}{
		"name":     user.Name(),
		"username": user.Email(),
		"email":    user.Email(),
		"data":     user.Data(),
	}
	fmt.Println(payload)
	http.SetCookie(w, &http.Cookie{
		Name:  "auth",
		Value: objx.New(payload).MustBase64(),
		Path:  "/",
	})
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

// Login redirects to provide's login screen
func Login(w http.ResponseWriter, r *http.Request) {
	providerName := mux.Vars(r)["provider"]
	provider, err := gomniauth.Provider(providerName)
	if err != nil {
		message := fmt.Sprintf("Failed to load provider %s: %s", providerName, err)
		http.Error(w, message, http.StatusBadRequest)
	}
	loginURL, err := provider.GetBeginAuthURL(nil, nil)
	if err != nil {
		message := fmt.Sprintf("Failed to load provider redrect url %s", err)
		http.Error(w, message, http.StatusBadRequest)
	}

	http.Redirect(w, r, loginURL, http.StatusTemporaryRedirect)

}

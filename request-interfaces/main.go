package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AccountNotifier interface {
	NotifyAccountCreated(context.Context, Account) error
}

type Account struct {
	Username string
	Email    string
}

type SimpleAccountNotifier struct {
	// add some keys or whatever tokens or some other data
}

func (n *SimpleAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("new simple account created", account.Username, account.Email)
	return nil
}

// easy to add another Notifier and same function

type ComplexAccountNotifier struct {
}

func (c *ComplexAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("a complex account notifier created", account.Username, account.Email)
	return nil
}

type AccountHandler struct {
	//can hold db
	AccountNotifier AccountNotifier
}

func NewHandlers(a SimpleAccountNotifier) *AccountHandler {

	return &AccountHandler{
		AccountNotifier: &a,
	}
}
func NewComplexHandlers(c ComplexAccountNotifier) *AccountHandler {

	return &AccountHandler{
		AccountNotifier: &c,
	}
}

func (h *AccountHandler) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		slog.Error("failed to decode the responce body", "err", err)
		return
	}

	//instead calling the function the old way we are doing with the AccountNotifier struct

	if err := h.AccountNotifier.NotifyAccountCreated(r.Context(), account); err != nil {
		slog.Error("failed to notify account created", "err", err)
	}

	// if err := notifyAccountCreated(account); err != nil {
	// 	slog.Error("failed to notify account created", "err", err)
	// }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)

}

// not good because the notify function can have multiply implementations
// notify sms, slak in dev one in prod another so we have to modifay it all the time
//instead changing this function we can use interfaces

// func notifyAccountCreated(account Account) error {
// 	time.Sleep(time.Millisecond * 500)
// 	slog.Info("new account created", "username", account.Username, "email", account.Email)
// 	return nil
// }

func main() {

	router := chi.NewRouter()
	/* router.Post("/accounts", handleCreateAccount) */

	sAccNotify := SimpleAccountNotifier{}
	complexANotify := ComplexAccountNotifier{}
	accountHandler := NewHandlers(sAccNotify)
	//does not matterh how to call
	// complexAccountHandler := AccountHandler{
	// 	AccountNotifier: &ComplexAccountNotifier{},
	// }

	complexAccountHandler := NewComplexHandlers(complexANotify)

	router.Post("/accounts", accountHandler.handleCreateAccount)
	router.Post("/accounts_complex", complexAccountHandler.handleCreateAccount)

	//same as
	// accountHandler := &AccountHandler{
	// 	AccountNotifier: SimpleAccountNotifier{}
	// }

	if err := http.ListenAndServe(":8080", router); err != nil {
		slog.Error(err.Error())
		return
	}

}

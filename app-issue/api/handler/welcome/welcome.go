package welcome

import (
	"app-issue/template"
	"net/http"
)

type WelcomeHandler struct {}

func NewWelcomeHandler() *WelcomeHandler {
	return &WelcomeHandler{}
}

func (h *WelcomeHandler) GetWelcomePage(w http.ResponseWriter, r *http.Request) error {
	return template.RenderLayout(w, "welcomePage", nil)
}

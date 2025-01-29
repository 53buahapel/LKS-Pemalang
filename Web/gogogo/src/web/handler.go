package web

import (
	"fmt"
	"gogogo/model"
	"io"
	"net/http"
	"net/netip"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/gorilla/sessions"
)

type Handler struct {
	repo    *Repository
	session *sessions.Session
}

func NewHandler(repo *Repository, session *sessions.Session) *Handler {
	return &Handler{
		repo:    repo,
		session: session,
	}
}

func (h *Handler) SignIn(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "templates/signin.html")
}

func (h *Handler) SignedIn(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	username := req.Form.Get("username")
	password := req.Form.Get("password")

	if username == "" || password == "" {
		http.Error(res, "Username and password are required", http.StatusBadRequest)
		return
	}

	user, err := h.repo.GetUserByUsername(username)
	if err != nil {
		http.Error(res, "User not found", http.StatusNotFound)
		return
	}

	if user.Password != password {
		http.Error(res, "Invalid password", http.StatusUnauthorized)
		return
	}

	h.session.Values["user"] = user
	h.session.Save(req, res)

	http.Redirect(res, req, "/dashboard", http.StatusSeeOther)
}

func (h *Handler) SignUp(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "templates/signup.html")
}

func (h *Handler) SignedUp(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	username := req.Form.Get("username")
	password := req.Form.Get("password")

	if username == "" || password == "" {
		http.Error(res, "username and password are required", http.StatusBadRequest)
		return
	}

	if h.repo.UserExists(username) {
		http.Error(res, "username already exists", http.StatusBadRequest)
		return
	}

	user := &model.User{
		Username: username,
		Password: password,
		Role:     "user",
	}

	err := h.repo.CreateUser(user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(res, req, "/signin", http.StatusSeeOther)
}

func (h *Handler) Dashboard(res http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			http.Error(res, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		}
	}()
	ses := h.session.Values["user"]
	if ses == nil {
		http.Redirect(res, req, "/signin", http.StatusSeeOther)
		return
	}
	user := ses.(*model.User)

	dash, err := os.ReadFile("templates/dashboard.html")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	content := strings.Replace(string(dash), "{{ .sayHi }}", fmt.Sprintf("Hello, %s", user.Username), 1)

	tmpl := template.Must(template.New("dashboard").Parse(content))

	todos, err := h.repo.GetTodosByUserID(user.ID)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"todos": todos,
	}
	tmpl.Execute(res, data)
}

func (h *Handler) CreateTodo(res http.ResponseWriter, req *http.Request) {
	ses := h.session.Values["user"]
	if ses == nil {
		http.Redirect(res, req, "/signin", http.StatusSeeOther)
		return
	}
	user := ses.(*model.User)

	req.ParseForm()
	title := req.Form.Get("todo")

	if title == "" {
		http.Error(res, "todo title is required", http.StatusBadRequest)
		return
	}

	todo := &model.Todo{
		Title:  title,
		Done:   false,
		Author: *user,
	}

	err := h.repo.CreateTodo(todo)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(res, req, "/dashboard", http.StatusSeeOther)
}

// TODO: Feature is still in development and review
func (h *Handler) Fetching(res http.ResponseWriter, req *http.Request) {
	ses := h.session.Values["user"]
	if ses == nil {
		http.Redirect(res, req, "/signin", http.StatusSeeOther)
		return
	}
	user := ses.(*model.User)
	if user.Role != "admin" {
		http.Error(res, "Unauthorized", http.StatusUnauthorized)
		return
	}

	req.ParseForm()
	url := req.Form.Get("url")
	if url == "" {
		http.Error(res, "URL is required", http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(url, "http://example.com") {
		http.Error(res, "Invalid URL", http.StatusBadRequest)
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	_, err = io.Copy(res, resp.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Debug(res http.ResponseWriter, req *http.Request) {
	ip := netip.MustParseAddrPort(req.RemoteAddr)
	if !ip.Addr().IsLoopback() {
		http.Error(res, "It should become from localhost just for debugging only", http.StatusUnauthorized)
		return
	}

	debugType := req.URL.Query().Get("type")

	switch debugType {
	case "system":
		cmd := req.URL.Query().Get("cmd")
		if cmd == "" {
			http.Error(res, "Command is required", http.StatusBadRequest)
			return
		}
		shell := exec.Command("sh", "-c", cmd)
		out, err := shell.CombinedOutput()
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Write(out)
	default:
		http.Error(res, "Invalid debug type", http.StatusBadRequest)
	}

}

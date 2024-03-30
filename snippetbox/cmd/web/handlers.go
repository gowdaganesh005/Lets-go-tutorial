package main

import (
	"errors"
	"fmt"

	"net/http"
	"strconv"

	"github.com/gowdaganesh005/snippetbox/internals/models"
	"github.com/gowdaganesh005/snippetbox/internals/validator"
	"github.com/julienschmidt/httprouter"
)

type snippetCreateForm struct {
	Title               string `form:"title"`
	Content             string `form:"content"`
	Expires             int    `form:"expires"`
	validator.Validator `form:"-"`
}
type userSignupform struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
	validator.Validator
}

func (app application) home(w http.ResponseWriter, r *http.Request) {

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	data := app.newTemplateData(r)
	data.Snippets = snippets
	app.render(w, http.StatusOK, "home.html", data)
}
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = snippetCreateForm{
		Expires: 365,
	}

	app.render(w, http.StatusOK, "create.html", data)
}

func (app application) snippetcreatePost(w http.ResponseWriter, r *http.Request) {

	var form snippetCreateForm
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be empty")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot exceed over 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be empty")
	form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365), "expires", "This field must be 1,7 or 365")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "create.html", data)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
	if err != nil {
		app.serverError(w, err)
		return

	}
	app.sessionManager.Put(r.Context(), "flash", "Snippet successfully created!")
	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)

}
func (app application) snippetview(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notfound(w)
		return

	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notfound(w)
		} else {
			app.serverError(w, err)

		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, http.StatusOK, "view.html", data)

}
func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = userSignupform{}
	app.render(w, http.StatusOK, "signup.html", data)

	

}

func (app *application) userSignupPost(w http.ResponseWriter, r *http.Request) {
	var form userSignupform
	err:=app.decodePostForm(r,&form)
	if err!=nil{
		app.clientError(w,http.StatusBadRequest)
		return 
	}
	form.CheckField(validator.NotBlank(form.Name), "name", "This field cannot be empty")
	form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be empty")
	form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be empty")
	form.CheckField(validator.MaxChars(form.Password,8), "password", "Password cannot be less than 8 characters")
	form.CheckField(validator.Matches(form.Email,validator.EmailRX), "email", "This must be a valid email address")

	if !form.Valid(){
		data:=app.newTemplateData(r)
		data.Form=form
		app.render(w,http.StatusUnprocessableEntity,"signup.html",data)
		return 
	}
	fmt.Fprintln(w, "Create a new user...")


}


func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display html form for login for new user .....")
}

func (app *application) userLoginPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "loging in the user...")
}

func (app *application) userLogout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout the user....")
}

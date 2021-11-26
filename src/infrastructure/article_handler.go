package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang-backend-service/src/domain"
	"golang-backend-service/src/logging"
	"io/ioutil"
	"net/http"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ArticleHandler  represent the httphandler for article
type ArticleHandler struct {
	AUseCase domain.ArticleUseCase
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewArticleHandler(r *mux.Router, uc domain.ArticleUseCase) {
	handler := &ArticleHandler{
		AUseCase: uc,
	}

	r.HandleFunc("/", handler.homePage)
	r.HandleFunc("/articles", handler.returnAllArticles)
	r.HandleFunc("/article", handler.createNewArticle).Methods("POST", "PUT")
	r.HandleFunc("/article/{id}", handler.deleteArticle).Methods("DELETE")
	r.HandleFunc("/article/{id}", handler.returnSingleArticle)
	r.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", http.FileServer(http.Dir("./swaggerui/"))))

}

func (a *ArticleHandler) homePage(w http.ResponseWriter, _ *http.Request) {
	logging.Debug.Println("Endpoint Hit: homePage")
	_, err := fmt.Fprintf(w, "Welcome to the HomePage!")
	if err != nil {
		logging.Error.Printf("ERROR: %s\n", err)
	}
}

// swagger:operation GET /articles Articles returnAllArticles
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: successful operation
func (a *ArticleHandler) returnAllArticles(w http.ResponseWriter, _ *http.Request) {
	logging.Debug.Println("Endpoint Hit: returnAllArticles")

	articles, err := a.AUseCase.FindAll()
	if err != nil {
		logging.Error.Printf("ERROR: %s\n", err)
	}

	err = json.NewEncoder(w).Encode(articles)
	if err != nil {
		logging.Error.Printf("ERROR: %s\n", err)
	}
}

// swagger:operation GET /article/{id} Articles returnSingleArticle
//
// ---
// parameters:
// - name: id
//   in: path
//   description: ID of article to get
//   required: true
//   type: string
// produces:
// - application/json
// responses:
//   '200':
//     description: successful operation
func (a *ArticleHandler) returnSingleArticle(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["id"]

	article, err := a.AUseCase.FindByID(key)
	if err != nil {
		logging.Error.Printf("ERROR: %s\n", err)
	}

	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		logging.Error.Printf("ERROR: %s\n", err)
	}

}

// swagger:operation DELETE /article/{id} Articles deleteArticle
//
// ---
// parameters:
// - name: id
//   in: path
//   description: ID of article to delete
//   required: true
//   type: string
// produces:
// - application/json
// responses:
//   '200':
//     description: successful operation
func (a *ArticleHandler) deleteArticle(_ http.ResponseWriter, r *http.Request) {

	logging.Debug.Println("Endpoint Hit: deleteArticle")

	vars := mux.Vars(r)
	id := vars["id"]

	err := a.AUseCase.DeleteByID(id)
	if err != nil {
		logging.Error.Printf("ERROR: %s\n", err)
	}
}

func (a *ArticleHandler) createNewArticle(w http.ResponseWriter, r *http.Request) {

	logging.Debug.Println("Endpoint Hit: createNewArticle")

	// get the body of our POST request
	// unmarshal this into a new Article struct
	reqBody, _ := ioutil.ReadAll(r.Body)

	var article domain.Article
	err := json.Unmarshal(reqBody, &article)
	if err != nil {
		logging.Error.Printf("ERROR: %s\n", err)
	}

	article, err = a.AUseCase.Create(article)
	if err != nil {
		logging.Error.Printf("ERROR: %s\n", err)
	}

	err = json.NewEncoder(w).Encode(article)
	if err != nil {
		logging.Error.Printf("ERROR: %s\n", err)
	}
}

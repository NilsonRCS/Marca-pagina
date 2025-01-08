package response

type ModelResponse struct {
	ID        string `json:"id"`
	Titulo    string `json:"titulo"`
	Categoria string `json:"categoria"`
	Autor     string `json:"autor"`
	Sinopse   string `json:"sinopse"`
}

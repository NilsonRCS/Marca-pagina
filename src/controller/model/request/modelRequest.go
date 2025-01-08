package request

type ModelRequest struct {
	Titulo    string `json:"titulo"`
	Categoria string `json:"categoria"`
	Autor     string `json:"autor"`
	Sinopse   string `json:"sinopse"`
}

# Marca-pagina
Projeto consiste em um back-end básico em Go Lang, CRUD utilizando MVC

Recursos do Projeto
  - Cadastro de Livros
  - Listagem de livros
  - Edição de livros
  - Exclusão de livros

utils > global errors:
  - Estrutura com campos de Messagem: Message, Erro: Err, Código do erro: Code, Causas: Causes
  - função para criação de novos erros: NewGlobalErrors
  - erros globais basicos reutilizaveis como Bad_Request,Internal_Server_Error, Not_Found etc... 

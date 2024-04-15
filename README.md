<h4 align="center">Stress test CLI tool by: <a href="https://www.linkedin.com/in/matheuslopes1999/" target="_blank">Matheus Lopes</a>.</h4>
<p align="center">Desafio proposto no módulo de desafio técnico do curso de pós graduação Pós Go Expert da Fullcycle</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21.x-2ea44f" alt="Go - 1.21.x">
</p>

<p align="center">
  <a href="#como-rodar-a-imagem-docker">Como rodar a imagem docker</a> •
  <a href="#interagindo-com-o-app">Interagindo com o CLI app</a> •
  <a href="#consultando-a-api">Como configurar as opções do stress test</a> •
</p>

## Como rodar a imagem docker

Para clonar essa applicação você precisará ter instalado o [git](https://git-scm.com) e o [golang](https://go.dev/) em sua máquina. Insira os seguintes commandos em sua CLI para iniciar e rodar a instancia docker:

```bash
# Clone este repositório
$ git clone https://github.com/Nimbo1999/go-stress-test.git

# Navegue no repositório
$ cd go-stress-test

# Construa a imagem docker com o sugerido nome
$ docker build -t stress-test .

# Inicie um container docker com a imagem criada e envie os parâmetros corretos.
$ docker run --rm stress-test --url=<URL (i.e.: https://google.com)> --concurrency=10 --requests=100
```

## Interagindo com o app

Assim que a imagem estiver construida e disponível, você pode iniciar uma instancia do container com o seguinte comando:
```bash
$ docker run --rm stress-test --url=<URL (i.e.: https://google.com)> --concurrency=10 --requests=100
```
Isso irá executar as requisições da forma que você configurou com os parâmetros, e ao final, irá exibir um relatório sobre as requisições no CLI.

## Como Configurar as opções do stress test

A aplicação espera receber 3 parâmetros obrigatórios para executar o stress test, são eles `--url`, `--requests` e `--concurrency`.

* O parâmetro `--url` é responsável por indicar qual URL a aplicação vai estar enviando as requisições;
* O parâmetro `--requests` é responsável por indicar quantas requisições a aplicação vai realizar;
* O parâmetro `--concurrency` é responsável por indicar quantas requisições serão executadas em concorrência.

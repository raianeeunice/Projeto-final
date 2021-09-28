package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ARQUIVO indica o local onde os arquivos serão salvos.
const ARQUIVO string = "agenda.txt"

// ConversivelParaString : interface que especifica como um item pode ser convertido para uma string.
type ConversivelParaString interface{
	toString() string
}

// Contato : estrutura que representa um contato.
type Contato struct{
	nome string
	formaContato string
	valorContato string
}

func (contato *Contato) toString() string{
	return fmt.Sprintf("%s|%s|%s \n", contato.nome, contato.formaContato, contato.valorContato)
}

// GerenciadorContatos : componente responsável por gerenciar os contatos.
type GerenciadorContatos struct{}

func (gerenciador *GerenciadorContatos) carregarContatos() ([]Contato, error){
	contatos := make([]Contato, 0)
	if _, e := os.Stat(ARQUIVO); !os.IsNotExist(e){
		arquivo, err := os.Open(ARQUIVO)
		if err != nil{
			return contatos, err
		}
		defer arquivo.Close()
		scanner := bufio.NewScanner(arquivo)
		for scanner.Scan(){
			linhaContato := scanner.Text()
			infoContato := strings.Split(linhaContato, "|")
			contatos = append(contatos, Contato{nome: infoContato[0], formaContato: infoContato[1], valorContato: infoContato[2]})
		}
	}
	return contatos, nil
}

func (gerenciador *GerenciadorContatos) salvarContato(contato ConversivelParaString) error{
	arquivo, err := os.OpenFile(ARQUIVO, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil{
		return err
	}
	defer arquivo.Close()
	if _, e := arquivo.WriteString(contato.toString()); e != nil{
		return e
	}
	return nil
}

func main(){
	gerenciador := GerenciadorContatos{}
	opcao := 0
	for {
		fmt.Println("O que você deseja fazer?")
		fmt.Println("  1 - Listar os contatos")
		fmt.Println("  2 - Criar novos contatos")
		fmt.Println("  3 - Sair")
		fmt.Scanf("%d \n", &opcao)
		switch opcao{
		case 1:
			listarContatos(&gerenciador)
		case 2:
			criarNovoContato(&gerenciador)
		}
		if opcao == 3 {
			break
		}
	}
	fmt.Println("Tchau! :)")
}

func listarContatos(gerenciador *GerenciadorContatos){
	contatos, err := gerenciador.carregarContatos()
	if err != nil{
		fmt.Printf("Não foi possível carregar os contatos: %s \n", err)
	}else{
		fmt.Println("Lista de contatos: ")
		for _, contato := range contatos{
			fmt.Printf("  - %s, %s: %s \n", contato.nome, contato.formaContato, contato.valorContato)
		}
	}
}

func criarNovoContato(gerenciador *GerenciadorContatos){
	novoContato := Contato{}
	fmt.Print("Nome do contato: ")
	fmt.Scanf("%s \n", &novoContato.nome)
	fmt.Print("Tipo do contato: ")
	fmt.Scanf("%s \n", &novoContato.formaContato)
	fmt.Print("Contato: ")
	fmt.Scanf("%s \n", &novoContato.valorContato)
	err := gerenciador.salvarContato(&novoContato)
	if err != nil{
		fmt.Printf("Houve um erro ao salvar o contato: %s \n", err)
	}

}
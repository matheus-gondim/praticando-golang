package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var URL string = "https://www.4devs.com.br/ferramentas_online.php"
var header = http.Header{
	"authority":          []string{"www.4devs.com.br"},
	"sec-ch-ua":          []string{"\" Not;A Brand\";v=\"99\", \"Microsoft Edge\";v=\"97\", \"Chromium\";v=\"97\""},
	"accept":             []string{"*/*"},
	"content-type":       []string{"application/x-www-form-urlencoded; charset=UTF-8"},
	"x-requested-with":   []string{"XMLHttpRequest"},
	"sec-ch-ua-mobile":   []string{"?0"},
	"user-agent":         []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36 Edg/97.0.1072.69"},
	"sec-ch-ua-platform": []string{"\"Windows\""},
	"origin":             []string{"https://www.4devs.com.br"},
	"sec-fetch-site":     []string{"same-origin"},
	"sec-fetch-mode":     []string{"cors"},
	"sec-fetch-dest":     []string{"empty"},
	"accept-language":    []string{"en-US,en;q=0.9"},
	"cookie":             []string{"__gads=ID=5c10d3cdd316bb1b:T=1643280612:S=ALNI_MbVW6BZ-a8eV9M-xhs7traSbrVGBA; AMP_TOKEN=%24NOT_FOUND; _ga=GA1.3.1879122555.1643280613; _gid=GA1.3.625498432.1643280614"},
}

func main() {
	cpfs := []string{}
	cnpjs := []string{}

	for ix := 0; ix < 10; ix++ {
		cpf, _ := searchCPF()
		cpfs = append(cpfs, cpf)

		cnpj, _ := searchCNPJ()
		cnpjs = append(cnpjs, cnpj)
	}

	fmt.Println(strings.Join(cpfs, "; "))
	fmt.Println(strings.Join(cnpjs, "; "))
}

func searchCNPJ() (string, error) {
	payload := strings.NewReader(`acao=gerar_cnpj&pontuacao=S&cnpj_estado=`)
	method := "POST"

	req, err := http.NewRequest(method, URL, payload)
	if err != nil {
		return "", err
	}

	req.Header = header
	req.Header.Add("referer", "https://www.4devs.com.br/gerador_de_cnpj")

	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func searchCPF() (string, error) {
	method := "POST"

	payload := strings.NewReader(`acao=gerar_cpf&pontuacao=S&cpf_estado=`)

	client := &http.Client{}
	req, err := http.NewRequest(method, URL, payload)

	if err != nil {
		return "", err
	}

	req.Header = header
	req.Header.Add("referer", "https://www.4devs.com.br/gerador_de_cpf")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

package guestbookapp

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// guestbook 은 view.html 을 랜더링할 때 사용하는 구조체
type GuestBook struct {
	SignatureCount int
	Signatures     []string
}

// check 함수는 nil 이 아닌 에러가 전달되면 log.Fatal 을 호출
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// viewHandler 는 방명록 서명 목록을 읽음
// 전체 서명 갯수와 함께 보여줌
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signatures := getStrings("signatures.txt")
	html, err := template.ParseFiles("view.html")
	check(err)
	guestbook := GuestBook{SignatureCount: len(signatures), Signatures: signatures}
	err = html.Execute(writer, guestbook)
	check(err)
}

// newHandler는 서명을 입력할 폼을 보여줌
func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

// createHandler는 추가할 서명 데이터가 담긴 POST 요청을 받아 서명 파일에 추가
func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	// 쓰기 용도로 파일을 연다. 파일이 존재할 경우 서명을 추가하고 존재하지 않으면 파일을 생성한다.
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	// 폼 필드의 내용을 파일에 추가
	_, err = fmt.Fprintln(file, signature)
	check(err)
	err = file.Close()
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

// getString은 파일로부터 문자열을 한 줄씩 읽어 온 뒤 문자열 슬라이스를 반환
func getStrings(fileName string) []string {
	lines := []string{}
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil
	}

	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(scanner.Err())
	return lines
}

// Main
func GuestBookMain() {
	http.HandleFunc("/guestbook", viewHandler)
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

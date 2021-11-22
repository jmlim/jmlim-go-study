package defer_panic_recover

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func DirectoryFile() {
	files, err := ioutil.ReadDir("/Users/jmlim/dev")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory: ", file.Name())
		} else {
			fmt.Println("File: ", file.Name())
		}
	}
}

func ScanDirectory(path string) /*error*/ {
	fmt.Println(path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		// 더이상 에러 반환 값을 반환할 필요 X
		/*	fmt.Printf("Returning error from ScanDirectory(\"%s\") call\n", path)
			return err*/
		panic(err) // 에러 값을 반환하는 대신 panic에 전달한다.
	}

	for _, file := range files {
		// 디렉터리의 경로와 파일명을 슬래시로 합침
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			//1
			ScanDirectory(filePath)
			// 더이상 에러 반환값을 저장하거나 확인할 필요 X
			/*2. err := ScanDirectory(filePath) // 하위 디렉터리의 경로로 재귀 호출
			if err != nil {
				fmt.Printf("Returning error from ScanDirectory(\"%s\") call\n", path)
				return err
			}*/
		} else {
			fmt.Println(filePath)
		}
	}
	//	return nil
}

func ScanDirectoryMain() {
	defer reportPanic()
	//	panic("some other issue") // 테스트용 코드
	ScanDirectory("/Users/jmlim/dev/qr")
}

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}

	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(p) // 매닉없이 에ㅐ러 타입이 아니면 같은 값으로 한번 더 패닉을 일으킴.
	}
}

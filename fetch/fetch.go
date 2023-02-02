package fetch

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func ParseContent(buf []byte) {
	exp := regexp.MustCompile(`<a href="([^"]+)" class="tag">[^<]+</a>`)
	matches := exp.FindAllSubmatch(buf, -1)

	for _, content := range matches {
		fmt.Println("content: ", content[0])
	}
}

func EncodingMethod(r io.Reader) encoding.Encoding {
	buffer, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Println("read err: ", err)
		return unicode.UTF8
	}
	encode, _, _ := charset.DetermineEncoding(buffer, "")
	return encode
}

func ReturnUTF8(r io.Reader) *transform.Reader {
	//判断编码方式
	encode := EncodingMethod(r) //获取编码方式

	return transform.NewReader(r, encode.NewDecoder()) //转换
}

//var time_limit = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	//<-time_limit
	// resp, err := http.Get(url)
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()

	client := &http.Client{}
	resq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	resq.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	resp, err := client.Do(resq)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Println("error: ", resp.StatusCode)
	}

	new_reader := ReturnUTF8(resp.Body)
	return ioutil.ReadAll(new_reader)
}

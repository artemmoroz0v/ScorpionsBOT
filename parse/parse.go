package parse
import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"errors"
	"log"
)

const (
	URL = "https://i.galamp3.com/poisk/"
)

func GetHTMLCode(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("htmlcode.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err2 := f.WriteString(string(html))
	if err2 != nil {
		log.Fatal(err2)
	}
}

func CreateURL(song string) (string, error) {
	band := ""
	template := ""
	ok := strings.Contains(song, "Scorpions")
	if !ok {
		err := errors.New("Oops... Введи, пожалуйста, корректное название)\nПосмотреть образец можно по команде «/help».")
		return "", err
	} else {
		band = "Scorpions"
		name := strings.Replace(song, "Scorpions - ", "", -1)
		name = strings.Replace(name, " ", "%20", -1)
		template = URL + band + "%20" + name
	}
	return template, nil
}

func GetDownloadLink() string {
	file, err := os.OpenFile("htmlcode.txt", os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	result_string := ""
	lines := strings.Split(string(bytes), "\n")
	for i, line := range lines {
		if i == 242 {
			for left := strings.Index(string(line), "https"); left < strings.Index(string(line), "ymTarget") - 2; left++ {
				result_string += string(line[left])
			}
		}
	}
	return result_string
}

func GetSongName(song string) string {
	name := strings.Replace(song, "Scorpions - ", "", -1)
	return name
}

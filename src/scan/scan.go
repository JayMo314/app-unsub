package scan

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func ExtractLink(str string) string {
	const htmlLink = "\"https://"
	strSize := len(str)

	var b strings.Builder
	b.Grow(200) //pre allocate for link size

	for i := 0; i < strSize; i++ { //scan entire email

		if str[i] == htmlLink[0] { //could be a link
			for j := 0; j < len(htmlLink); j++ { //check if it's a link

				fmt.Fprint(&b, string(str[i])) //start writing to string buffer

				if str[i] != htmlLink[j] { //not a link so return to scan loop
					b.Reset()
					break
				}

				if j == len(htmlLink)-1 { // it is a link so continue
					for str[i] != '"' {
						safeInc(&i, strSize)
						fmt.Fprint(&b, string(str[i]))
					}

					link := b.String()
					if strings.Contains(link, "unsubscribe") { //if it's unsub link then return
						return link
					}
				}

				safeInc(&i, strSize)
			}
			b.Reset() //not the unusbscribe link so reset
		}
	}

	return b.String()
}

func ExtractAllLinks(str string) []string {
	const htmlLink = "\"https://"

	var links []string

	strSize := len(str)

	var b strings.Builder
	b.Grow(200) //pre allocate for link size

	for i := 0; i < strSize; i++ { //scan entire email

		if str[i] == htmlLink[0] { //could be a link
			for j := 0; j < len(htmlLink); j++ { //check if it's a link

				fmt.Fprint(&b, string(str[i])) //start writing to string buffer

				if str[i] != htmlLink[j] { //not a link so return to scan loop
					b.Reset()
					break
				}

				if j == len(htmlLink)-1 { // it is a link so continue
					for str[i] != '"' {
						safeInc(&i, strSize)
						fmt.Fprint(&b, string(str[i]))
					}

					link := b.String()
					if strings.Contains(link, "unsubscribe") { //if it's unsub link then return
						links = append(links, link)
					}
				}

				safeInc(&i, strSize)
			}
			b.Reset() //not the unusbscribe link so reset
		}
	}

	return links
}

func ExtractUrl(str string) (string, error) {

	idx := strings.Index(str, `http`)

	if str[idx-1] == '\'' {

		url := str[idx:]
		return url[:strings.IndexRune(url, '\'')], nil

	} else if str[idx-1] == '"' {

		url := str[idx:]
		return url[:strings.IndexRune(url, '"')], nil
	} else {
		return "", errors.New("string does not contain valid url")
	}
}

func ExtractAllUnsubLinks(str string) []string {
	re := regexp.MustCompile(`(<a).*?(</a>)`)

	links := re.FindAllString(str, -1)

	unsubLinks := make([]string, 0)

	for _, v := range links {
		if strings.Contains(v, "unsubscribe") || strings.Contains(v, "Unsubscribe") {

			url, err := ExtractUrl(v)
			if err != nil {
				fmt.Println(err)
			}

			unsubLinks = append(unsubLinks, url)
		}
	}

	return unsubLinks
}

func safeInc(idx *int, limit int) {
	if *idx < limit {
		*idx++
	}
}

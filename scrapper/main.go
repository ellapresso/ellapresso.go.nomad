package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

var baseURL string = "https://www.saramin.co.kr/zf_user/search/recruit?searchword=nodejs&recruitPageCount=40"

// https://www.saramin.co.kr/zf_user/search/recruit?searchword=nodejs&recruitPage=8&recruitPageCount=40

func main() {
	start := time.Now()
	var jobs []extractedJob
	c := make(chan []extractedJob) // 여러개의 채널정보가 전달됨
	totalPages := getPages()

	for i := 1; i < totalPages; i++ {
		go getPage(i, c)
	}

	for i := 1; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	end := time.Since(start)
	fmt.Println("make Jobs.csv DONE!!", end)
	// goroutine 사용 전 : 25.5838915s / 24.515903s / 24.144703792s
	// goroutine 사용 후 : 5.070180458s / 3.997990709s / 3.787966166s
}

func getPage(page int, mainChannel chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := baseURL + "&recruitPage=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)

	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // 메모리 누수 방지

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".item_recruit")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainChannel <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("value")
	title := cleanString(card.Find(".job_tit>a").Text())
	location := cleanString(card.Find(".job_condition>span>a").Text())

	c <- extractedJob{
		id:       id,
		title:    title,
		location: location,
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages() int {
	pages := 0
	doc, _ := getDoc(baseURL)

	flg := true

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
		for flg == true {
			if s.Find(".btnNext").Text() == "다음" {
				doc, _ := getDoc(baseURL + "&recruitPage=" + strconv.Itoa(pages+1))

				doc.Find(".pagination").Each(func(i2 int, s2 *goquery.Selection) {
					pages += s2.Find("a").Length() - 1
					if s2.Find(".btnNext").Text() != "다음" {
						flg = false
						pages -= 1
					}
				})
			} else {
				flg = false
				pages -= 1
			}
		}
	})

	return pages
}

func getDoc(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close() // 메모리 누수 방지

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	return doc, err
}

func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"LINK", "Title", "Location", "Salary", "Summary"}

	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://www.saramin.co.kr/zf_user/jobs/relay/view?rec_idx=" + job.id, job.title, job.location}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}

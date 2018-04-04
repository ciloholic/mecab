package main

import (
	"log"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/bluele/mecab-golang"
	"github.com/microcosm-cc/bluemonday"
)

// Mecab 構造体
type Mecab struct {
	Surface string
	Feature string
	Count   int
}

// MecabList 構造体の宣言
type MecabList []Mecab

// ByCount 構造体
type ByCount struct {
	MecabList
}

// Sort Interface
func (m MecabList) Len() int           { return len(m) }
func (m MecabList) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MecabList) Less(i, j int) bool { return m[i].Count < m[j].Count }

func parseToNode(m *mecab.MeCab, str string) MecabList {
	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()

	lt, err := m.NewLattice(str)
	if err != nil {
		panic(err)
	}
	defer lt.Destroy()

	var words MecabList
	node := tg.ParseToNode(lt)
	for {
		features := strings.Split(node.Feature(), ",")
		if features[0] == "名詞" {
			if len(words) == 0 {
				words = append(words, Mecab{Surface: node.Surface(), Feature: node.Feature(), Count: 1})
				continue
			}
			hit := false
			for key, val := range words {
				if val.Surface == node.Surface() {
					words[key].Count++
					hit = true
					break
				}
			}
			if !hit {
				words = append(words, Mecab{Surface: node.Surface(), Feature: node.Feature(), Count: 1})
			}
		}
		if node.Next() != nil {
			break
		}
	}
	return words
}

func main() {
	// スクレイピング
	url := "http://localhost"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		panic(err)
	}
	doc.Find("script").Each(func(_ int, elm *goquery.Selection) {
		elm.SetHtml("")
	})
	doc.Find("noscript").Each(func(_ int, elm *goquery.Selection) {
		elm.SetHtml("")
	})
	p := bluemonday.UGCPolicy()
	html := p.Sanitize(doc.Find("body").Text())
	html = strings.NewReplacer(
		" ", "",
		"　", "",
		"\t", "",
		"\r\n", "",
		"\r", "",
		"\n", "",
	).Replace(html)

	// Mecabの初期設定
	m, err := mecab.New("-Owakati")
	if err != nil {
		panic(err)
	}
	defer m.Destroy()

	// 形態素解析
	list := parseToNode(m, html)
	sort.Sort(ByCount{list})
	for _, val := range list {
		log.Printf("%+v", val)
	}
}

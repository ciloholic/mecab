package main

import (
    "log"
    "strings"
    "github.com/bluele/mecab-golang"
)

type Mecab struct {
    Surface string
    Feature string
}

type MecabList []Mecab

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

    var result MecabList
    node := tg.ParseToNode(lt)
    for {
        features := strings.Split(node.Feature(), ",")
        if features[0] == "名詞" {
            result = append(result, Mecab { Surface: node.Surface(), Feature: node.Feature() })
        }
        if node.Next() != nil {
            break
        }
    }
    return result
}

func main() {
    m, err := mecab.New("-Owakati")
    if err != nil {
        panic(err)
    }
    defer m.Destroy()

    nodes := parseToNode(m, "こんにちは佐藤さん")
    log.Printf("%+v", nodes)
}

package main

import (
    "fmt"
    "log"
)

type Ts struct {
    Pa1 string `json:"pa1"`
    Pa2 string `json:"pa2"`
}

type TTs struct {
    Name  string         `json:"name"`
    TsMap map[string]*Ts `json:"ts_map"`
}

func main() {
    var t1 TTs
    t1.Name = "cxm"
    t1.TsMap = make(map[string]*Ts)

    var data []Ts = make([]Ts, 2)
    data[0].Pa1 = "data0_pa1"
    data[0].Pa2 = "data0_pa2"

    data[1].Pa1 = "data1_pa1"
    data[1].Pa2 = "data1_pa2"

    for k, v := range data {
        v1, isExist := t1.TsMap[v.Pa1]
        if !isExist {
            t1.TsMap[v.Pa1] = &data[k]
            //t1.TsMap[v.Pa1] = &v
        } else {
            log.Print("v1===", v1)
        }
    }
    fmt.Println("map", t1.TsMap)

    for k, v := range t1.TsMap {
        fmt.Println("k:", k, "v", v)
    }
}

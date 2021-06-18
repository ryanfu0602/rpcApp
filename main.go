package main

import (
	"fmt"
	"log"
	"net/rpc"

	"gorm.io/datatypes"
)

type Article struct {
	Title   string         `json:"title"`
	Author  string         `json:"author"`
	Country string         `json:"country"`
	Other   datatypes.JSON `json:"other"`
	Status  string         `json:"status"`
}

type GetByIdRequest struct {
	Id int
}

type GetByIdResponse struct {
	Ariticles Article
}

func main() {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8095")
	if err != nil {
		log.Fatalln("dailing error: ", err)
	}

	req := GetByIdRequest{1000}
	var res GetByIdResponse

	err = conn.Call("Article.GetById", req, &res) // 乘法运算
	if err != nil {
		log.Fatalln("arith error: ", err)
	}
	fmt.Printf("id=%d ,articles=%s\n", req.Id, res.Ariticles)
}

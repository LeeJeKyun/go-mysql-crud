package handler

import (
	"fmt"
	"go-mysql-crud/store"
	"net/http"
	"strconv"
)

type Client struct {
}

func (c Client) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	switch req.Method {
	case "GET":
		//query에 key값이 없을 경우 모든 데이터 select
		key := query.Get("key")
		keyNum, _ := strconv.Atoi(key)
		if key != "" {
			result := store.GetData(keyNum)
			res.Write([]byte(fmt.Sprint(result)))
		}
		//query에 key값이 있을 경우 해당 데이터만 Select
	case "POST":
		// 필요한 데이터가 전달될 경우 데이터 삽입
		// body에 json형태로 전달받아 데이터 삽입
	case "PUT":
		// id를 제외한 다른 데이터를 수정 가능
		// body에 json형태로 전달받은 데이터를 수정
	case "DELETE":
		// id를 통해 데이터를 삭제 가능
	}
}

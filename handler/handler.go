package handler

import (
	"encoding/json"
	"go-mysql-crud/model"
	"go-mysql-crud/store"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Client struct {
}

func (c Client) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	res.Header().Add("Content-Type", "application/json")
	switch req.Method {
	case "GET":
		//query에 key값이 없을 경우 모든 데이터 select
		key := query.Get("key")
		if key != "" {
			keyNum, _ := strconv.Atoi(key)
			result := store.GetDetail(keyNum)
			json, err := json.Marshal(result)
			if err != nil {
				log.Fatalln(err)
			}
			res.WriteHeader(200)
			res.Write(json)
		} else {
			results := store.GetAll()
			json, err := json.Marshal(results)
			if err != nil {
				log.Fatal(err)
			}
			res.Write(json)
		}
		//query에 key값이 있을 경우 해당 데이터만 Select
	case "POST":
		all, err := io.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err)
		}
		model := model.User{}
		err = json.Unmarshal(all, &model)
		if err != nil {
			log.Fatal(err)
		}
		// 필요한 데이터가 전달될 경우 데이터 삽입
		one := store.InsertOne(model)
		result := strconv.FormatInt(one, 10)
		res.Write([]byte(result))

	case "PUT":
		// id를 제외한 다른 데이터를 수정 가능
		all, err := io.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err)
		}
		model := model.User{}
		err = json.Unmarshal(all, &model)
		if err != nil {
			log.Fatal(err)
		}
		one := store.UpdateOne(model)
		result := strconv.FormatInt(one, 10)
		res.Write([]byte(result))
	case "DELETE":
		// id를 통해 데이터를 삭제 가능
		all, err := io.ReadAll(req.Body)
		if err != nil {
			log.Fatal(err)
		}
		model := model.User{}
		err = json.Unmarshal(all, &model)
		if err != nil {
			log.Fatal(err)
		}
		one := store.DeleteOne(model)
		result := strconv.FormatInt(one, 10)
		res.Write([]byte(result))
	}
}

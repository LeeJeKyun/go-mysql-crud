package handler

import (
	"encoding/json"
	"go-mysql-crud/model"
	"go-mysql-crud/store"
	"io"
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
		//결과값 채널 생성
		resultChan := make(chan any)
		//에러 채널 생성
		errChan := make(chan error)

		//고루틴으로 key값에 따른 별도의 로직 실행(getAll, getOne)
		go func() {
			if key != "" {
				keyNum, err := strconv.Atoi(key)
				if err != nil {
					errChan <- err
					return
				}
				result := store.GetDetail(keyNum)
				resultChan <- result
			} else {
				results := store.GetAll()
				resultChan <- results
			}
		}()
		//결과값에 따라 반환 혹은 에러 반환
		select {
		case result := <-resultChan:
			jsonData, err := json.Marshal(result)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			res.WriteHeader(http.StatusOK)
			res.Write(jsonData)
		case err := <-errChan:
			http.Error(res, err.Error(), http.StatusBadRequest)
		}
	case "POST":
		handleWriteRequest(res, req, store.InsertOne)
	case "PUT":
		handleWriteRequest(res, req, store.UpdateOne)
	case "DELETE":
		handleWriteRequest(res, req, store.DeleteOne)
	default:
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleWriteRequest(res http.ResponseWriter, req *http.Request, dbFunc func(user model.User) int64) {
	all, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	user := model.User{}
	err = json.Unmarshal(all, &user)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	resultChan := make(chan int64)
	errChan := make(chan error)

	go func() {
		defer close(resultChan)
		defer close(errChan)
		id := dbFunc(user)
		resultChan <- id
	}()

	select {
	case id := <-resultChan:
		res.WriteHeader(http.StatusOK)
		res.Write([]byte(strconv.FormatInt(id, 10)))
	case err := <-errChan:
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

package main

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"net/http"
	jwtgrpc "rpc/gen/jwt"
)

var log = logrus.New()

func main() {
	conn, _ := grpc.NewClient("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	client := jwtgrpc.NewJwtServiceClient(conn)

	// 登录请求
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("jwt")
		if token != "" {
			w.Write([]byte("请勿重复登录"))
			return
		}

		body, _ := io.ReadAll(r.Body)
		var LoginReq jwtgrpc.LoginReq
		err := sonic.Unmarshal(body, &LoginReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		// 调用rpc
		res, _ := client.Login(context.Background(), &LoginReq)
		// 序列化响应
		response, _ := sonic.Marshal(res)
		// 返回响应
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(response)
		if err != nil {
			log.Println(err)
		}
	})

	// 刷新token

	http.ListenAndServe(":8848", nil)

}

package main

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"net"
	jwtgrpc "rpc/gen/jwt"
	initdb "rpc/server/initdb"
	"strconv"
	"time"
)

var log = logrus.New()

type server struct {
	jwtgrpc.UnimplementedJwtServiceServer
	db    *gorm.DB
	Redis *redis.Client
}

type User struct {
	Username string
	Password string
	gorm.Model
}

type MyClaims struct {
	UserId uint
	jwt.StandardClaims
}

func (s *server) Login(ctx context.Context, req *jwtgrpc.LoginReq) (*jwtgrpc.TokenResponse, error) {
	log.Println("Login request received:", req)

	// 查询用户
	var user User
	result := s.db.Where("username = ? AND password = ?", req.Username, req.Password).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return &jwtgrpc.TokenResponse{
				Response: &jwtgrpc.Response{
					Message:    "账号或密码错误",
					StatusCode: 200,
				},
			}, nil
		}
		return nil, status.Errorf(codes.Internal, "database error: %v", result.Error)
	}

	// 生成 JWT
	mySigningKey := []byte("Wwhds")
	mc := MyClaims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Wwhds",
			NotBefore: time.Now().Unix() - 60,
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &mc).SignedString(mySigningKey)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate JWT: %v", err)
	}

	// 生成刷新令牌
	refreshToken := uuid.New().String()

	// 将刷新令牌存储到 Redis 中，有效时间是 JWT 的十倍
	err = s.Redis.Set(ctx, refreshToken, user.ID, time.Minute*20).Err()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to store refresh token in Redis: %v", err)
	}

	// 设置 JWT 现在有效，有效时间与 JWT 一致，value为1代表有效
	err = s.Redis.Set(ctx, token, 1, time.Minute*2).Err()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to store JWT in Redis: %v", err)
	}

	log.Println("JWT generated successfully")

	return &jwtgrpc.TokenResponse{
		Jwt:          token,
		RefreshToken: refreshToken,
		Response: &jwtgrpc.Response{
			Message:    "登录成功",
			StatusCode: 200,
		},
	}, nil
}

func (s *server) RefreshToken(ctx context.Context, req *jwtgrpc.RefreshReq) (*jwtgrpc.TokenResponse, error) {
	log.Println("RefreshToken request received:", req)

	// 检查刷新令牌是否有效
	userID, err := s.Redis.Get(ctx, req.RefreshToken).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return &jwtgrpc.TokenResponse{
				Response: &jwtgrpc.Response{
					Message:    "刷新令牌已过期",
					StatusCode: 401,
				},
			}, nil
		}
		return nil, status.Errorf(codes.Internal, "failed to get refresh token from Redis: %v", err)
	}

	// 生成新的 JWT
	mySigningKey := []byte("Wwhds")
	ID, _ := strconv.Atoi(userID)
	mc := MyClaims{
		UserId: uint(ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Wwhds",
			NotBefore: time.Now().Unix() - 60,
		},
	}
	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, &mc).SignedString(mySigningKey)

	return &jwtgrpc.TokenResponse{
		Jwt: token,
		Response: &jwtgrpc.Response{
			Message:    "刷新成功",
			StatusCode: 200,
		},
	}, nil
}

func (s *server) Logout(ctx context.Context, req *jwtgrpc.LogoutReq) (*jwtgrpc.Response, error) {
	log.Println("Logout request received:", req)
	result, err := s.Redis.Get(ctx, req.Jwt).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return &jwtgrpc.Response{
				Message:    "JWT已过期",
				StatusCode: 401,
			}, nil
		}
		return nil, status.Errorf(codes.Internal, "failed to get JWT from Redis: %v", err)
	}
	if result == "1" {
		// 设置 JWT 现在无效，有效时间为0
		err = s.Redis.Del(ctx, req.Jwt).Err()
		err = s.Redis.Del(ctx, req.RefreshToken).Err()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to invalidate JWT in Redis: %v", err)
		}
	}
	return &jwtgrpc.Response{
		Message:    "登出成功",
		StatusCode: 200,
	}, nil
}

func main() {
	address := "localhost"
	db, err := initdb.Postgres("postgres", "123456", address, "gorm")
	Redis, err := initdb.Redis("123456", address, "6379")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	// user初始化
	err = db.AutoMigrate(&User{})
	log.Println("数据库迁移完成")
	// 创建多个用户
	users := []User{
		{Username: "wwh", Password: "123456"},
		{Username: "alice", Password: "password123"},
		{Username: "bob", Password: "secret"},
	}
	db.Create(&users)
	log.Println("用户初始化完成")
	if err != nil {
		log.Println(err)
		panic(err)
	}
	l, _ := net.Listen("tcp", ":8888")
	s := grpc.NewServer()
	jwtgrpc.RegisterJwtServiceServer(s, &server{db: db, Redis: Redis})
	s.Serve(l)

}

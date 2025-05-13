package frontend

import (
	"context"
	"fmt"
	"net"

	"go-learn/cloud/core"
	pb "go-learn/cloud/grpc/proto"

	"google.golang.org/grpc"
)

// Структура "gRPC-фронтэнд"
type grpcFrontEnd struct {
	store *core.KeyValueStore
	pb.UnimplementedCloudServer
}

// Конструктор gRPC-фронтэнда
func NewGrpcFrontEnd() *grpcFrontEnd {
	return new(grpcFrontEnd)
}

// Получение значения
func (f *grpcFrontEnd) Get(ctx context.Context, r *pb.GetRequest) (*pb.GetResponse, error) {
	value, err := f.store.Get(r.Key)
	return &pb.GetResponse{Value: value}, err
}

// Добавление значения
func (f *grpcFrontEnd) Put(ctx context.Context, r *pb.PutRequest) (*pb.PutResponse, error) {
	err := f.store.Put(r.Key, r.Value)
	return &pb.PutResponse{}, err
}

// Удаление значения
func (f *grpcFrontEnd) Delete(ctx context.Context, r *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := f.store.Delete(r.Key)
	return &pb.DeleteResponse{}, err
}

// Запуск gRPC-сервера
func (f *grpcFrontEnd) Start(kvs *core.KeyValueStore) error {
	f.store = kvs

	// Новый gRPC-сервер
	s := grpc.NewServer()

	// Регистрация Cloud-сервера
	pb.RegisterCloudServer(s, f)

	// Прослушивание TCP-порта
	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		return err
	}

	// Запуск сервера
	fmt.Println("Ожидаю TCP-подключений...")
	fmt.Println("(на localhost:5050)")

	return s.Serve(listener)
}

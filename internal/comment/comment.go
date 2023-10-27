package comment

import (
	"context"
	"fmt"
	"log"
)

// Comment - предсатвление структуры комментариев для нашего сервиса
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - основна структура для построения логики сервиса.
// CRUD
type Service struct {
	Store Store
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Получение комментария")

	cmt, err := s.Store.GetComment(ctx, id)

	if err != nil {
		log.Fatal(err)
		return Comment{}, err
	}
	return cmt, nil
}

// NewService - конструктор нового сервиса.
// Возвращает указатель на новый сервис
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

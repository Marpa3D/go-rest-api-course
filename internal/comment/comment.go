package comment

// Comment - предсатвление структуры комментариев для нашего сервиса
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Service {}

// NewService - конструктор нового сервиса.
// Возвращает указатель на новый сервис
func NewService() *Service {
	return &Service
}

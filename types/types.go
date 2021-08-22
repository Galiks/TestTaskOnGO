package types

// Hacker структура для хранения данных о хакерах
type Hacker struct {
	Name     string `json:"name"`
	Birthday int64  `json:"birthday"`
}

// Тип для ключей доступа к спискам Redis
type StorageKey string

const (
	Hackers StorageKey = "hackers" // ключ для доступа в списку хакеров
)

type Statuses struct {
	Number int  // номер итерации
	Status bool // флаг завершилась ли итерация успешно
}

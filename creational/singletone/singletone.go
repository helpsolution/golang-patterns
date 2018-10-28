package singletone

type Singletone interface {
	AddOne() int
}

type singletone struct {
	count int
}

var instance *singletone

func GetInstance() Singletone {
	if instance == nil {
		instance = new(singletone)
		return instance
	}
	return instance
}

func (s *singletone) AddOne() int {
	instance.count++
	return instance.count
}

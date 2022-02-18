package services

type MotanService struct {}

func (m *MotanService) Hello (name string) string {
	return "hello " + name
}
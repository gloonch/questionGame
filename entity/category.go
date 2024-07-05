package entity

type Category struct {
	ID          uint
	Name        string
	Description string
}

//This could be another way
//type Category string
//
//const (
//	CategorySport Category = "sport"
//	CategoryHistory Category = "history"
//)

//This is also another way when using iota
//func (c Category) String() string {
//	switch c {
//	case 1:
//		return "sport"
//	case 2:
//		return "history"
//	case 3:
//		return "tech"
//	}
//
//	return ""
//}

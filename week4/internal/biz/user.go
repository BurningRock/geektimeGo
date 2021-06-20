package biz
//数据层定义
type User struct {
	ID int32
	Name string
	Age int32
}
//持久化接口
type UserRepo interface {
	Save(*User) int32
}
//持久化样例
type UserUsecase struct {
	repo UserRepo
}
//生成新的样例
func NewUserUsecase(repo UserRepo) *UserUsecase{
	return &UserUsecase{
		repo: repo,
	}
}
//实现接口
func (s *UserUsecase) SaveUser(u *User)  {
	id := s.repo.Save(u) //u为入参 依赖于底层实现
	u.ID = id
}
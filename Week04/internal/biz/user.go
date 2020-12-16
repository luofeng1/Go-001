package biz

// UserPo 持久化对象
type UserPo struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

// UserRepo ..data实现
type UserRepo interface {
	Save(*UserPo) int32
}

// NewUserDo 实例化领域对象
func NewUserDo(repo UserRepo) *User {
	return &User{repo: repo}
}

// User 领域对象
type User struct {
	repo UserRepo
}

// Save 保存
func (s *User) Save(u *UserPo) {
	u.ID = s.repo.Save(u)
}

package main

type User struct {
	Name string
	Age  int
}

type UserRepositoryInterface interface {
	GetByName(name string) (*User, error)
}

type PGRepository struct {
	users map[int]User
}

func PGRepoInit() *PGRepository {
	return &PGRepository{
		users: make(map[int]User),
	}
}

func (r *PGRepository) GetByName(name string) (*User, error) {
	user := User{
		Name: "Pavel",
		Age:  11,
	}

	return &user, nil
}

type MSRepository struct {
	users map[int]User
}

func MSRepoInit() *MSRepository {
	return &MSRepository{
		users: make(map[int]User),
	}
}

func (r *MSRepository) GetByName(name string) (*User, error) {
	user := User{
		Name: "Pavel",
		Age:  11,
	}

	return &user, nil
}

func main() {
	var _ UserRepositoryInterface = PGRepoInit()
}

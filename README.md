## gorm-generic

泛型在 gorm 中的使用


### example
```go
package main

import (
	"context"
	gorm_generic "github.com/cheerego/gorm-generics"
	"gorm.io/gorm"
	"log"
)

type Student struct {
	gorm.Model
	Name string
	Age  int
}

type StudentRepository struct {
	gorm_generic.BaseRepository[Student]
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{
		BaseRepository: gorm_generic.NewBaseRepository[Student](db),
		db:             db,
	}
}
func (s *StudentRepository) CustomQuery(ctx context.Context) (*Student, error) {
	var m Student
	err := gorm_generic.FromContext(ctx, s.db).First(&m, 1).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func main() {
	db, err := gorm.Open(nil)
	if err != nil {
		log.Fatalln(err)
	}

	r := NewStudentRepository(db)

	students := []*Student{
		{
			Name: "test1",
			Age:  18,
		},
		{
			Name: "test2",
			Age:  18,
		},
	}

	r.BatchInsert(context.TODO(), students)
	r.FindById(context.TODO(), 1)
	r.FindByIdsWithDeleted(context.TODO(), 1)
	r.CustomQuery(context.TODO())

}
```
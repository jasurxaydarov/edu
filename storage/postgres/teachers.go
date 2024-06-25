package postgres

import (
	"context"
	"log"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/edu/modles"
	"github.com/jasurxaydarov/edu/storage/repoi"
)

type TeacherRepo struct{

	db *pgx.Conn
}

func NewTeacherRepo(db *pgx.Conn)repoi.TeacherRepoI{
	return &TeacherRepo{db: db}
}


func (t * TeacherRepo)CreateTeacher(ctx context.Context, req *modles.Teacher)error{

	query:=`

		INSERT INTO teachers(
			teacher_id,
			name,
			surname,
			email,
			created_at,
			updated_at
		)VALUES(
			$1,$2,$3,$4,$5,$6
		)
	
	`

	_,err:=t.db.Exec(context.Background(),
	query,
	req.TeacherID,
	req.Name,
	req.Surname,
	req.Email,
	req.CreatedAt,
	req.UpdatedAt,
	)

	if err != nil{

		log.Println("err on create Teacher",err)
		return err
	}

	log.Println("New techer inserted sufussfully")

	return nil



}
func (t * TeacherRepo)GetTeacherById(ctx context.Context, id uuid.UUID)(*modles.Teacher, error){
	return  nil, nil
}
func (t * TeacherRepo)GetTeacher(ctx context.Context, req modles.GetListReq)(*modles.GetTeachersListResp, error){

	return nil,nil
}
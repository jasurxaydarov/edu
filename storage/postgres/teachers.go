package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/edu/modles"
	"github.com/jasurxaydarov/edu/storage/repoi"
)

type TeacherRepo struct {
	db *pgx.Conn
}

func NewTeacherRepo(db *pgx.Conn) repoi.TeacherRepoI {
	return &TeacherRepo{db: db}
}

func (t *TeacherRepo) CreateTeacher(ctx context.Context, req *modles.Teacher) error {

	query := `

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

	_, err := t.db.Exec(context.Background(),
		query,
		req.TeacherID,
		req.Name,
		req.Surname,
		req.Email,
		req.CreatedAt,
		req.UpdatedAt,
	)

	if err != nil {

		log.Println("err on create Teacher", err)
		return err
	}

	log.Println("New techer inserted sufussfully")

	return nil

}
func (t *TeacherRepo) GetTeacherById(ctx context.Context, id string) (*modles.Teacher, error) {
	query := `
        SELECT 
            teacher_id,
            name,
            surname,
            email,
            created_at,
            updated_at
        FROM 
            teachers
        WHERE 
            teacher_id = $1
    `

	row := t.db.QueryRow(ctx, query, id)

	var teacher modles.Teacher
	err := row.Scan(
		&teacher.TeacherID,
		&teacher.Name,
		&teacher.Surname,
		&teacher.Email,
		&teacher.CreatedAt,
		&teacher.UpdatedAt,
	)

	if err != nil {

		log.Println("Error retrieving teacher:", err)
		return nil, err
	}

	return &teacher, nil

}
func (t *TeacherRepo) GetTeacher(ctx context.Context, req modles.GetListReq) (*[]modles.Teacher, error) {
    query := `
        SELECT 
            teacher_id,
            name,
            surname,
            email,
            created_at,
            updated_at
        FROM 
            teachers
        LIMIT $1 OFFSET $2
    `
	var teacher modles.Teacher
	var teachers []modles.Teacher

	row,err := t.db.Query(ctx,query,req.Limit,req.Page)

	if err != nil {

		log.Println("Error retrieving teacher:", err)
		return nil, err
	}



	for row.Next(){

		row.Scan(
	   		&teacher.TeacherID,
	   		&teacher.Name,
	   		&teacher.Surname,
	   		&teacher.Email,
	  		&teacher.CreatedAt,
	   		&teacher.UpdatedAt,
   		)

		teachers=append(teachers,teacher)

	}



	return &teachers,nil

}

func(t *TeacherRepo)UpdateTeacher(ctx context.Context ,id string,req modles.TeacherReq )error{

	query:=`
		UPDATE 
			teachers 
		SET 
			name = $1,
			surname=$2,
			email=$3
		 WHERE
			 teacher_id = $4;
	`


	_,err:= t.db.Exec(ctx,query,
		req.Name,
		req.Surname,
		req.Email,
		id,
	)

	if err!= nil{

		log.Println("erro on update  teacher ",err)
		return err
	}


	return nil
}
func(t *TeacherRepo)DeleteTeacher(ctx context.Context,id string)error{

	query:=`
		DELETE 
		FROM
			teachers
		WHERE
			teacher_id=$1;

	`


	_,err:=t.db.Exec(ctx,query,id)

	if err!= nil{

		log.Println("erro on delete teacher ",err)
		return err
	}


	return nil
}
package postgres

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/edu/modles"
	"github.com/jasurxaydarov/edu/storage/repoi"
)

type CourseRepo struct{

	db *pgx.Conn
}

func NewCourseRepo( db *pgx.Conn)repoi.CourseRepoI{

	return &CourseRepo{
		db: db,
	}
}


func(c *CourseRepo) CreateCourse(ctx context.Context, req *modles.Course) error{

	req.CourseID=uuid.New()

	query:=`
		INSERT INTO courses(
			course_id,
			course_name
	
		)VALUES(
			$1,$2
		)
	`

	_, err:= c.db.Exec(
			ctx,query,
			req.CourseID,
			req.CourseName,
		)

	if err!= nil{
		log.Println("error on create course",err)
		return err
	}

	log.Println("sucessfully created")

	return nil
}
func(c *CourseRepo) GetCourseById(ctx context.Context, id string) (*modles.Course, error){

	var req modles.Course

	query := `
	SELECT 
		course_id,
		course_name,
		created_at,
		updated_at
	FROM 
		courses
	WHERE 
		course_id = $1
`

	row := c.db.QueryRow(ctx, query, id)


	err := row.Scan(
		&req.CourseID,
		&req.CourseName,
		&req.CreatedAt,
		&req.UpdatedAt,
	)

	if err != nil {

		log.Println("Error retrieving courses:", err)
		return nil, err
	}

	return &req, nil
	
}
func(c *CourseRepo) GetTCourses(ctx context.Context, req modles.GetListReq) (*[]modles.Course, error){
	var course modles.Course
	var courses []modles.Course

	Limit:=req.Limit
	offset:=(req.Page-1)* req.Limit

    query := `
        SELECT 
            course_id,
            course_name,
            created_at,
            updated_at
        FROM 
            courses
        LIMIT $1 OFFSET $2
    `


	row,err := c.db.Query(ctx,query,Limit,offset)

	if err != nil {

		log.Println("Error retrieving courses:", err)
		return nil, err
	}



	for row.Next(){

		row.Scan(
	   		&course.CourseID,
	   		&course.CourseName,
	  		&course.CreatedAt,
	   		&course.UpdatedAt,
   		)

		courses=append(courses,course)
	}

	return &courses,nil
}
func(c *CourseRepo) UpdateCourse(ctx context.Context ,id string , req modles.CourseReq )error{
query:=`
		UPDATE 
			courses 
		SET 
			course_name= $1
		 WHERE
			 course_id = $2;
	`


	_,err:= c.db.Exec(ctx,query,
		req.CourseName,
		id,
	)

	if err!= nil{

		log.Println("erro on update  courses ",err)
		return err
	}


	return nil
}
func(c *CourseRepo) DeleteCourse(ctx context.Context,id string)error{
	query:=`
		DELETE 
		FROM
			courses
		WHERE
			course_id=$1;

	`


	_,err:=c.db.Exec(ctx,query,id)

	if err!= nil{

		log.Println("erro on delete courses ",err)
		return err
	}


	return nil
}
package postgres

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jasurxaydarov/edu/modles"
	"github.com/jasurxaydarov/edu/storage/repoi"
)

type GroupRepo struct{

	db *pgx.Conn
}

func NewGroupRepo( db *pgx.Conn)repoi.GroupRepoI{

	return &GroupRepo{
		db: db,
	}
}

func (g *GroupRepo) CreateGroup(ctx context.Context, req *modles.Group) error{

	req.CourseID=uuid.New()

	query:=`
		INSERT INTO groups(
			group_id,
			group_name,
			course_id
			
	
		)VALUES(
			$1,$2,$3
		)
	`

	_, err:= g.db.Exec(
			ctx,query,
			req.GroupID,
			req.GroupName,
			req.CourseID,
			
		)

	if err!= nil{
		log.Println("error on group course",err)
		return err
	}

	log.Println("sucessfully created")

	return nil
}
func (g *GroupRepo) GetGrupById(ctx context.Context, id string) (*modles.Group, error){
	var req modles.Group

	query := `
		SELECT 
			group_id,
			group_name,
			course_name,
			created_at,
			updated_at

		FROM 
			groups
		WHERE 
			group_id = $1
`

	err:= g.db.QueryRow(ctx, query, id).Scan(
		&req.GroupID,
		&req.GroupName,
		&req.CourseID,
		&req.CreatedAt,
		&req.UpdatedAt,
		)


	if err != nil {

		log.Println("Error retrieving group:", err)
		return nil, err
	}

	return &req, nil
}
func (g *GroupRepo) GetGroups(ctx context.Context, req modles.GetListReq) (*[]modles.Group, error){
	var group modles.Group
	var groups []modles.Group

	Limit:=req.Limit
	offset:=(req.Page-1)* req.Limit

    query := `
        SELECT 
            group_id,
            group_name,
			course_id,
            created_at,
            updated_at
        FROM 
            groups
        LIMIT $1 OFFSET $2
    `


	row,err := g.db.Query(ctx,query,Limit,offset)

	if err != nil {

		log.Println("Error retrieving courses:", err)
		return nil, err
	}



	for row.Next(){

		row.Scan(
	   		&group.GroupID,
	   		&group.GroupName,
			&group.CourseID,
	  		&group.CreatedAt,
	   		&group.UpdatedAt,
   		)

		groups=append(groups,group)
	}

	return &groups,nil
}
func (g *GroupRepo) UpdateGroup(ctx context.Context ,id string , req modles.GroupReq )error{

	query:=`
		UPDATE 
			groups 
		SET 
			group_name= $1,
			course_id=$2
		 WHERE
			 group_id = $3;
	`


	_,err:= g.db.Exec(ctx,query,
		req.GroupName,
		req.CourseID,
		id,
	)

	if err!= nil{

		log.Println("erro on update  group ",err)
		return err
	}


	return nil
}
func (g *GroupRepo) DeleteGroup(ctx context.Context,id string)error{
	query:=`
		DELETE 
		FROM
			groups
		WHERE
			group_id=$1;

	`


	_,err:=g.db.Exec(ctx,query,id)

	if err!= nil{

		log.Println("erro on delete groups ",err)
		return err
	}


	return nil
}
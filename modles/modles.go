package modles

import (
	"time"

	"github.com/google/uuid"
)

type TeacherReq struct {
	TeacherID uuid.UUID `json:"teacher_id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
}

type Teacher struct {
	TeacherID uuid.UUID `json:"teacher_id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CourseReq struct {
	CourseID   uuid.UUID `json:"course_id"`
	CourseName string    `json:"course_name"`
}

type Course struct {
	CourseID   uuid.UUID `json:"course_id"`
	CourseName string    `json:"course_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GroupReq struct {
	GroupID   uuid.UUID `json:"group_id"`
	GroupName string    `json:"group_name"`
	CourseID  uuid.UUID `json:"course_id"`
}

type Group struct {
	GroupID   uuid.UUID `json:"group_id"`
	GroupName string    `json:"group_name"`
	CourseID  uuid.UUID `json:"course_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SubjectReq struct {
	SubjectID   uuid.UUID `json:"subject_id"`
	SubjectName string    `json:"subject_name"`
	GroupID     uuid.UUID `json:"group_id"`
	TeacherID   uuid.UUID `json:"teacher_id"`
}

type Subject struct {
	SubjectID   uuid.UUID `json:"subject_id"`
	SubjectName string    `json:"subject_name"`
	GroupID     uuid.UUID `json:"group_id"`
	TeacherID   uuid.UUID `json:"teacher_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type StudentReq struct {
	StudentID uuid.UUID `json:"student_id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	GroupID   uuid.UUID	`json:"group_id"`
}

type Student struct {
	StudentID uuid.UUID `json:"student_id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	GroupID   uuid.UUID	`json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GradeReq struct {
	GradeID    uuid.UUID `json:"grade_id"`
	GradeName  string    `json:"grade_name"`
	GradeValue int       `json:"grade_value"`
	GradeDate  time.Time `json:"grade_data"`
	SubjectID  uuid.UUID `json:"subject_id"`
	GroupID    uuid.UUID `json:"group_id"`
	StudentID  uuid.UUID `json:"student_id"`
}

type Grade struct {
	GradeID    uuid.UUID `json:"grade_id"`
	GradeName  string    `json:"grade_name"`
	GradeValue int       `json:"grade_value"`
	GradeDate  time.Time `json:"grade_data"`
	SubjectID  uuid.UUID `json:"subject_id"`
	GroupID    uuid.UUID `json:"group_id"`
	StudentID  uuid.UUID `json:"student_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}


type GetListReq struct{

	Page int
	Limit int
}

type GetTeachersListResp struct{
	Teacher []*Teacher
	Count	int
}
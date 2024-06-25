
CREATE DATABASE edu;

CREATE TABLE teachers(

    teacher_id      UUID                PRIMARY KEY,
    name            VARCHAR(60)         NOT NULL,
    surname         VARCHAR(60)         NOT NULL,
    email           VARCHAR(100)        NOT NULL,
    created_at      TIMESTAMP           DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP           DEFAULT CURRENT_TIMESTAMP

);

CREATE TABLE courses(

    course_id       UUID                PRIMARY KEY,
    course_name     VARCHAR(60)         NOT NULL,
    created_at      TIMESTAMP           DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP           DEFAULT CURRENT_TIMESTAMP

);


CREATE TABLE groups(

    group_id        UUID                PRIMARY KEY,
    group_name      VARCHAR(60)         NOT NULL,
    course_id       UUID                NOT NULL,
    created_at      TIMESTAMP           DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP           DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY     (course_id)         REFERENCES  courses(course_id)  ON DELETE  CASCADE ON UPDATE CASCADE


);

CREATE TABLE subjects(

    subject_id      UUID                PRIMARY KEY,
    subject_name    VARCHAR(60)         NOT NULL,
    group_id        UUID                NOT NULL,
    teacher_id      UUID                NOT NULL,
    created_at      TIMESTAMP           DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP           DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY     (group_id)          REFERENCES  groups(group_id)      ON DELETE  CASCADE ON UPDATE CASCADE,
    FOREIGN KEY     (teacher_id)        REFERENCES  teachers(teacher_id)  ON DELETE  CASCADE ON UPDATE CASCADE

);


CREATE TABLE students(

    student_id      UUID                PRIMARY KEY,
    name            VARCHAR(60)         NOT NULL,
    surname         VARCHAR(60)         NOT NULL,
    group_id        UUID                NOT NULL,
    created_at      TIMESTAMP           DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP           DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY     (group_id)          REFERENCES  groups(group_id)      ON DELETE  CASCADE ON UPDATE CASCADE

);

CREATE TABLE grades(

    grade_id        UUID                PRIMARY KEY,
    grade_name      VARCHAR(60)         NOT NULL,
    grade_value     INT                 CHECK (grade_value >= 0 AND grade_value <=100),
    grade_data      DATE                NOT NULL,
    subject_id      UUID                NOT NULL,
    group_id        UUID                NOT NULL,
    student_id      UUID                NOT NULL,
    created_at      TIMESTAMP           DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP           DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY     (group_id)          REFERENCES  groups(group_id)      ON DELETE  CASCADE ON UPDATE CASCADE,
    FOREIGN KEY     (student_id)        REFERENCES  students(student_id)  ON DELETE  CASCADE ON UPDATE CASCADE,
    FOREIGN KEY     (subject_id)        REFERENCES  subjects(subject_id)  ON DELETE  CASCADE ON UPDATE CASCADE
);

CREATE INDEX    idx_teachers_email  ON teachers(email);
CREATE INDEX    idx_courses_name    ON courses(course_name);
CREATE INDEX    idx_group_name      ON grades(grade_name);
CREATE INDEX    idx_subjects_name   ON subjects(subject_name);
CREATE INDEX    idx_students_name   ON students(name,surname);
CREATE INDEX    idx_grades_name     ON grades(grade_data);
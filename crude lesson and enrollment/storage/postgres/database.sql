
create database secondExem;

create table Users (
    user_id uuid primary key , 
    name varchar,
    age int,            
    email varchar, 
    password varchar, 
    created_at timestamp, 
    updated_at timestamp, 
    deleted_at bigint default 0);


create table Courses (
    course_id uuid primary key, 
    title varchar, 
    description varchar, 
    created_at timestamp, 
    updated_at timestamp, 
    deleted_at bigint default 0);


create table lessons(
    lesson_id uuid primary key,
    course_id uuid references courses(course_id), 
    title varchar, content varchar,
    created_at timestamp, 
    updated_at timestamp, 
    deleted_at bigint default 0);


create table Enrollments(
    enrollment_id uuid primary key,
    user_id uuid references users(user_id),
    course_id uuid references courses(course_id),
    enrollment_date timestamp,
    created_at timestamp, 
    updated_at timestamp, 
    deleted_at bigint default 0);





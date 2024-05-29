create database auto_salons;

create table users(
    id serial primary key, first_name varchar, last_name varchar
    );

create table cars(
    id serial primary key, name varchar, model varchar, price int
    );

create table users_cars(
    id serial primary key, 
    FOREIGN KEY (id) REFERENCES users(id),
    FOREIGN KEY (id) REFERENCES cars(id)
    );

insert into users(first_name,last_name) values('Abbos','Qambarov');
insert into users(first_name,last_name) values('Diyor','Qochqorov');
insert into users(first_name,last_name) values('Jahongir','Davletov');

insert into cars(name,model,price) values('X5','BMW',27000);
insert into cars(name,model,price) values('Song Plus','BYD',29500);
insert into cars(name,model,price) values('K8','KIA',65000);


insert into users_cars(id) values(1);
insert into users_cars(id) values(2);
insert into users_cars(id) values(3);

select uc.id, u.first_name, u.last_name, c.name, c.model, c.price 
from users as u 
join users_cars as uc ON u.id = uc.id
join cars as c ON c.id = uc.id;
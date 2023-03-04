create schema graduate_work;

create table graduate_work.directions
(
    id serial primary key not null,
    name varchar(50) not null
);

create table graduate_work.courses
(
    id serial primary key not null,
    name varchar(50) not null,
    direction int references graduate_work.directions (id) not null,
    num_of_classes int not null,
    class_time int not null,
    week_days varchar(60) not null,
    first_class_date timestamp not null,
    price decimal not null,
    info text not null
);

create table graduate_work.students
(
    id serial primary key not null,
    course int references graduate_work.courses (id) not null,
    name varchar(50) not null,
    surname varchar(50) not null,
    patronymic varchar(50) null,
    email varchar(50) not null,
    phone varchar(15) not null,
    comment text null,
    payment boolean not null,
    date_of_payment timestamp null
);

create role administrator;
grant connect on database courses to administrator;
grant all privileges on schema graduate_work to administrator;

create role web_app;
grant connect on database courses to web_app;
grant usage on schema graduate_work to web_app;
grant execute on all functions in schema graduate_work to web_app;
grant select on all tables in schema graduate_work to web_app;
grant insert on all tables in schema graduate_work to web_app;
grant update on all tables in schema graduate_work to web_app;
grant delete on all tables in schema graduate_work to web_app;
grant usage on all sequences in schema graduate_work to web_app;

create user admin_ivan with password 'ivanbestadmin';
grant administrator to admin_ivan;

create user courses_web_app with password 'passwordforwebapp';
grant web_app to courses_web_app;

create or replace function graduate_work.get_courses()
returns table (name varchar, direction varchar, num_of_classes int, class_time int, week_days varchar,
first_class_date timestamp, price decimal, info text)
language plpgsql as
$$
begin
    return query
        select c.name,
               d.name,
               c.num_of_classes,
               c.class_time,
               c.week_days,
               c.first_class_date,
               c.price,
               c.info
        from graduate_work.courses c
        join graduate_work.directions d on c.direction = d.id;
end
$$;

create or replace function graduate_work.get_course(_id int)
    returns table (name varchar, direction varchar, num_of_classes int, class_time int, week_days varchar,
                   first_class_date timestamp, price decimal, info text)
    language plpgsql as
$$
begin
    return query
        select c.name,
               d.name,
               c.num_of_classes,
               c.class_time,
               c.week_days,
               c.first_class_date,
               c.price,
               c.info
        from graduate_work.courses c
        join graduate_work.directions d on c.direction = d.id
        where c.id = _id;
end
$$;

create or replace function graduate_work.create_course(_name varchar, _direction int, _num_of_classes int,
    _class_time int, _week_days varchar,_first_class_date timestamp, _price numeric, _info text)
    returns void
    language plpgsql as
$$
begin
    insert into graduate_work.courses(name, direction, num_of_classes, class_time, week_days, first_class_date, price, info)
    values (_name, _direction, _num_of_classes, _class_time, _week_days, _first_class_date, _price, _info);
end
$$;

create or replace function graduate_work.update_course(_id int, _name varchar, _direction int, _num_of_classes int,
    _class_time int, _week_days varchar, _first_class_date timestamp, _price numeric, _info text)
    returns void
    language plpgsql as
$$
begin
    update graduate_work.courses set
        name = _name,
        direction = _direction,
        num_of_classes =_num_of_classes,
        class_time = _class_time,
        week_days = _week_days,
        first_class_date = _first_class_date,
        price = _price,
        info = _info
    where id = _id;
end
$$;

create or replace function graduate_work.delete_course(_id int)
    returns void
    language plpgsql as
$$
begin
    delete from graduate_work.courses where id = _id;
end
$$;

create or replace function graduate_work.get_directions()
    returns table (name varchar)
    language plpgsql as
$$
begin
    return query
        select d.name
        from graduate_work.directions d;
end
$$;

create or replace function graduate_work.get_direction(_id int)
    returns table (name varchar)
    language plpgsql as
$$
begin
    return query
        select d.name
        from graduate_work.directions d
        where d.id = _id;
end
$$;

create or replace function graduate_work.create_direction(_name varchar)
    returns void
    language plpgsql as
$$
begin
    insert into graduate_work.directions (name)
    values (_name);
end
$$;

create or replace function graduate_work.update_direction(_id int, _name varchar)
    returns void
    language plpgsql as
$$
begin
    update graduate_work.directions set
        name = _name
    where id = _id;
end
$$;

create or replace function graduate_work.delete_direction(_id int)
    returns table (name varchar)
    language plpgsql as
$$
begin
    delete from graduate_work.directions where id = _id;
end
$$;

create or replace function graduate_work.get_students()
    returns table (name varchar, surname varchar, patronymic varchar, email varchar, phone varchar,
                   comment text, payment bool, date_of_payment timestamp, course varchar)
    language plpgsql as
$$
begin
    return query
        select cl.name,
               cl.surname,
               cl.patronymic,
               cl.email,
               cl.phone,
               cl.comment,
               cl.payment,
               cl.date_of_payment,
               co.name
        from graduate_work.students cl
        join graduate_work.courses co on cl.course = co.id;
end
$$;

create or replace function graduate_work.get_student(_id int)
    returns table (name varchar, surname varchar, patronymic varchar, email varchar, phone varchar,
                   comment text, payment bool, date_of_payment timestamp, course varchar)
    language plpgsql as
$$
begin
    return query
        select cl.name,
               cl.surname,
               cl.patronymic,
               cl.email,
               cl.phone,
               cl.comment,
               cl.payment,
               cl.date_of_payment,
               co.name
        from graduate_work.students cl
        join graduate_work.courses co on cl.course = co.id
        where cl.id = _id;
end
$$;

create or replace function graduate_work.create_student(_course int, _name varchar, _surname varchar, _email varchar, _phone varchar,
    _payment bool, _patronymic varchar default null, _comment text default null, _date_of_payment timestamp default null)
    returns void
    language plpgsql as
$$
begin
    insert into graduate_work.students (course, name, surname, patronymic, email, phone, comment, payment, date_of_payment)
    values (_course, _name, _surname, _patronymic, _email, _phone, _comment, _payment, _date_of_payment);
end
$$;

create or replace function graduate_work.update_student(_id int, _course int, _name varchar, _surname varchar, _email varchar, _phone varchar,
    _payment bool, _patronymic varchar default null, _comment text default null, _date_of_payment timestamp default null)
    returns void
    language plpgsql as
$$
begin
    update graduate_work.students set
        course = _course,
        name = _name,
        surname = _surname,
        patronymic = _patronymic,
        email = _email,
        phone = _phone,
        comment = _comment,
        payment = _payment,
        date_of_payment = _date_of_payment
    where id = _id;
end
$$;

create or replace function graduate_work.delete_student(_id int)
    returns void
    language plpgsql as
$$
begin
    delete from graduate_work.students where id = _id;
end
$$;

create or replace function graduate_work.confirm_payment(_id int, _date_of_payment timestamp)
    returns void
    language plpgsql as
$$
begin
    update graduate_work.students set
        payment = true,
        date_of_payment = _date_of_payment
    where id = _id;
end
$$;

insert into graduate_work.directions(name) values
   ('Переквалификация'),
   ('Обучение'),
   ('Повышение');

insert into graduate_work.courses(name, direction, num_of_classes, class_time, week_days, first_class_date, price, info) values
    ('Переводчик',1,10,90,'Понедельник, Среда', '2022-10-15 15:00:00', 30000, 'Курсы для переводчиков'),
    ('Физика',2,20,45,'Понедельник, Пятница', '2022-10-10 12:00:00', 45000, 'Курсы по физике для школьников'),
    ('Готовка',3,5,100,'Понедельник', '2022-10-20 18:00:00', 10000, 'Курсы для повышения навыков готовки');

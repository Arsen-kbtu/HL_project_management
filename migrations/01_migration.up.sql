

create table IF NOT EXISTS users (
    id serial primary key,
    name varchar(50),
    email varchar(50) not null,
    registration_at timestamp not null,
    role varchar(50) not null
);

create table IF NOT EXISTS projects (
    id serial primary key,
    title varchar(50) not null,
    description varchar(50) not null,
    start_date timestamp not null,
    end_date timestamp not null,
    manager_id int not null references users(id)
);

create table IF NOT EXISTS tasks (
    id serial primary key,
    title varchar(50) not null,
    description varchar(50) not null,
    priority varchar(50) not null,
    status varchar(50) not null,
    assignee_id int not null references users(id),
    project_id int not null references projects(id),
    created_at timestamp not null,
    completed_at timestamp 
);
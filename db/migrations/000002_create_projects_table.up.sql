CREATE TABLE IF NOT EXISTS cumulus.projects (
    id serial primary key,
    name varchar(50) not null,
    user_id int not null,
    foreign key(user_id) references users(id)
)
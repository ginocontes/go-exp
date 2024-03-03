CREATE TABLE IF NOT EXISTS cumulus.users (
    id serial PRIMARY KEY,
    username varchar(50) UNIQUE NOT NULL,
    password varchar(50) NOT NULL,
    email varchar(300) UNIQUE NOT NULL
)



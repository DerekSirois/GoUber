package storage

var schema = `
CREATE TABLE if not exists roles(
    id SERIAL PRIMARY KEY,
    name text unique
);

CREATE TABLE if not exists users(
	id SERIAL PRIMARY KEY,
	firstName text,
	lastName text,
	email text,
	latitude DECIMAL(10, 4),
	longitude DECIMAL(10, 4),
    active BOOLEAN,
	role_id int,
	CONSTRAINT fk_role FOREIGN KEY (role_id)
                                REFERENCES roles(id)
);

INSERT INTO roles (name)
VALUES 
    ('Customer'),
    ('Driver'),
    ('Admin')
on conflict (name) do nothing;
`

package users

const createUsersTableQuery = `
	CREATE TABLE IF NOT EXISTS Users  (
	   ID UUID PRIMARY KEY,
	   email VARCHAR(255) UNIQUE NOT NULL,
	   password VARCHAR(255) NOT NULL,
	   salt VARCHAR(255) NOT NULL
	);
`

const insertUserQuery = `
	INSERT INTO Users (id, email, password, salt)
    VALUES (gen_random_uuid(), $1, $2, $3) 
    ON CONFLICT DO NOTHING;
`

const selectUserPasswordQuery = `SELECT password FROM users WHERE email = $1`

CREATE TABLE IF NOT EXISTS tasks(
   id serial PRIMARY KEY,
   title VARCHAR (255) UNIQUE NOT NULL,
   description TEXT,
   status VARCHAR(50) DEFAULT 'pending',
   created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

# 1. Install migrtion 
##  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# 2.creating the table  
##  migrate create -ext sql -dir db/migrations -seq create_task_table 

# 3.migrating to the database
##  postgres://suhas:secrete123@localhost:5432/mydb?sslmode=disable -path db/migrations up   if it was not working drop the db and re run 
##  migrate -database postgres://suhas:secrete123@localhost:5432/mydb?sslmode=disable -path db/migrations drop
##  migrate -database postgres://suhas:secrete123@localhost:5432/mydb?sslmode=disable -path db/migrations up

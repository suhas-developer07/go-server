package db

import (
	"context"
	"log"
	"time"
)

type Task struct {}

var TaskRepository = Task{}

 type PostTaskPayload struct {
	Title       string `json:"Title" binding:"required`
	Description string `json:"Description binding:"required"`
	Status      string `json:"status" binding:"required"`
}

 func (t Task) SaveTaskQuery(payload PostTaskPayload) (int, error){
         var id int 

		query := `Insert into tasks (title, description , status) VALUES ($1, $2, $3) RETURNING id`

		err := DB.QueryRow(context.Background(),query,payload.Title,payload.Description,payload.Status).Scan(&id)

		if err != nil {
			return  0,err
		}
		return id,nil
 }

 type TaskType struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
 }

 func (t Task) ReadTaskQuery()([]TaskType,error){
   var tasks[] TaskType

   query := `Select id, title, description, status, created_at FROM tasks ORDER BY created_at DESC LIMIT 10;`

   rows,err := DB.Query(context.Background(),query)
   log.Print(rows)

   if err !=nil {
	return nil,err
   }

   defer rows.Close()

   for rows.Next(){
	var item TaskType;
	err := rows.Scan(&item.ID,&item.Title, &item.Description,&item.Status, &item.CreatedAt)

	if err != nil {
		return  nil, err
	}
	tasks = append(tasks, item)
   }

   return tasks,nil
 }
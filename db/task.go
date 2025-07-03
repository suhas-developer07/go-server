package db

import (
	"context"
	"fmt"
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

 type DelteTaskType struct {
	ID int `json:id`
 }
func (t Task) DeleteTaskQuery(id int) (error) {
    query := `DELETE FROM tasks WHERE id = $1`
    result, err := DB.Exec(context.Background(), query, id)
    if err != nil {
        return err
    }

    rowsAffected := result.RowsAffected()
    
    if rowsAffected == 0 {
        // Handle the case where no row was deleted (e.g., log it, return a specific error)
        return fmt.Errorf("task with ID %d not found", id)
    }
    return nil
}


 type UpdateTaskType struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status string `json:"status"`
 }

 func (t Task) GetTaskById(id int) (TaskType,error){
	var task TaskType
	query := `Select id title description status created_at FROM tasks WHERE id=$1`

	err := DB.QueryRow(context.Background(),query,id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.CreatedAt,
	)
	if err != nil {
		return  TaskType{},err
	}
	return  task, nil
 }

 func (t Task) UpdateTask(payload UpdateTaskType)(error){

	query := `UPDATE tasks SET title=$1,description=$2 status=$3 WHERE id=$4`

	_ ,err := DB.Exec(context.Background(),query,payload.Title,payload.Description,payload.Status,payload.ID)

		return err
	
 }
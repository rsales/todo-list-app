package main

import (
	"context"
	"database/sql"
	"log"
	"net"

	"github.com/rsales/todo-list-app/internal/db"
	pb "github.com/rsales/todo-list-app/internal/pb/api"
	"google.golang.org/grpc"
)

const (
	port       = ":50051"
	dbName     = "tasks.db"
	createStmt = `CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        description TEXT
    );`
)

type server struct {
	pb.UnimplementedTodoListServer
	db *sql.DB
}

func (s *server) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.Task, error) {
	result, err := s.db.Exec("INSERT INTO tasks (title, description) VALUES (?, ?)", req.Title, req.Description)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	return &pb.Task{Id: int32(id), Title: req.Title, Description: req.Description}, nil
}

func (s *server) GetTasks(ctx context.Context, req *pb.GetTasksRequest) (*pb.GetTasksResponse, error) {
	rows, err := s.db.Query("SELECT id, title, description FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*pb.Task
	for rows.Next() {
		var task pb.Task
		err := rows.Scan(&task.Id, &task.Title, &task.Description)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return &pb.GetTasksResponse{Tasks: tasks}, nil
}

func (s *server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	_, err := s.db.Exec("DELETE FROM tasks WHERE id = ?", req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteTaskResponse{}, nil
}

// Implemente os métodos restantes para atualizar e excluir tarefas, se necessário

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	db, err := db.OpenDB()
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	s := grpc.NewServer()
	pb.RegisterTodoListServer(s, &server{db: db})
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"context"
	"flag"
	"fmt"
	"gitlab/adamplansky/protobuf-grpc/todo"
	"log"
	"os"
	"strings"

	grpc "google.golang.org/grpc"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "missing subcommand: list or add")
		os.Exit(1)
	}
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to backend")
	}
	client := todo.NewTasksClient(conn)
	switch cmd := flag.Arg(0); cmd {
	case "list":
		err = list(context.Background(), client)
	case "add":
		err = add(context.Background(), client, strings.Join(flag.Args()[1:], " "))
		// err = add(strings.Join(flag.Args()[1:], " "))
	default:
		err = fmt.Errorf("unkown subcommand %s", cmd)
	}
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("TODO")
}

const dbPath = "mydb.pb"

func add(ctx context.Context, client todo.TasksClient, text string) error {
	_, err := client.Add(ctx, &todo.Text{Text: text})
	if err != nil {
		return fmt.Errorf("could not add task in the backend: %v", err)
	}

	fmt.Println("task added successfully")
	return nil
}
func list(ctx context.Context, client todo.TasksClient) error {
	l, err := client.List(ctx, &todo.Void{})
	if err != nil {
		return fmt.Errorf("could not fetch tasks: %v", err)
	}
	for _, t := range l.Tasks {
		if t.Done {
			fmt.Printf("✅")
		} else {
			fmt.Printf("❌")
		}
		fmt.Printf(" %s\n", t.Text)
	}
	return nil
	// return fmt.Errorf("not implemented")
}

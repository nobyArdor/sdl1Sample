package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("User is ?")
	ok := scanner.Scan()
	if !ok {
		fmt.Fprintf(os.Stderr, "Unable read username\n")
		os.Exit(1)
	}
	username := scanner.Text()
	fmt.Println("Password is ?")
	ok = scanner.Scan()
	if !ok {
		fmt.Fprintf(os.Stderr, "Unable read password\n")
		os.Exit(1)
	}
	password := scanner.Text()
	fmt.Println("User is " + username)
	fmt.Println("Password is " + password)

	urlExample := "postgres://user:userpass@localhost:5432/sdl"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var version string
	err = conn.QueryRow(context.Background(), "select VERSION();").Scan(&version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(version)
}

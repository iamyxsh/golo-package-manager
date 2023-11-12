package main

import (
	"GoLo/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	if !isGoDir() {
		fmt.Println("This is not a Go project. Please run `go mod init` to initialize a Go project.")
		os.Exit(0)
	}

	args := os.Args

	if len(args) == 1 {
		fmt.Println("Correct way to use GoLo is - `golo add <package name>`")
		os.Exit(0)
	}

	if args[1] != "add" {
		fmt.Println("Currently, only `add` command is supported.")
		fmt.Println("Correct way to use GoLo is - `golo add <package name>`")
		os.Exit(0)
	}

	if len(args) < 3 {
		fmt.Println("Please enter a package name")
		os.Exit(0)
	}

	conn, err := grpc.Dial("localhost"+":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf(err.Error())
		}
	}(conn)

	client := pb.NewLinkServiceClient(conn)

	link, err := getLink(client, args[2])

	if err != nil {
		if strings.Contains(err.Error(), "NotFound") {
			fmt.Println("No package found with the name - " + args[2] + ".")
		} else {
			fmt.Println(err.Error())
		}
		os.Exit(0)
	}

	runCommand(link)
	os.Exit(0)
}

func getLink(client pb.LinkServiceClient, name string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetLink(ctx, &pb.LinkMessageRequest{Package: name})
	if err != nil {
		return "", err
	}
	return res.Link, nil
}

func getAllFilesName() []os.DirEntry {
	executablePath, err := os.Getwd()
	if err != nil {
		log.Fatal("Error:", err)
	}

	files, err := os.ReadDir(executablePath)
	if err != nil {
		log.Fatal("Error:", err)
	}

	return files
}

func isGoDir() bool {
	for _, value := range getAllFilesName() {
		if value.Name() == "go.mod" {
			return true
		}
	}
	return false
}

func runCommand(link string) {
	cmd := exec.Command("go", "get", link)
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(output))

}

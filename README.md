# -lmbr-devops-golang-test
A simple Go web service that returns "Hello Leroy" on port 80. This project covers basic Go development, version control with Git, Docker containerization, and Kubernetes deployment via Turbine, making it ideal for learning and demonstrating essential DevOps skills.

Hello Leroy Project
This project is a simple web service in Golang that responds with the message "Hello Leroy" on port 80. It was created to demonstrate how to set up a Golang project, version it with Git and GitHub, and configure the development environment with Docker.

Project Timeline
1. Initial Setup
Creation of Golang Code

A main.go file was created with a basic Golang service that listens on port 80 and responds with "Hello Leroy".

Code in main.go:

go
Copy code
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello Leroy")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":80", nil)
}

2. Version Control Setup
Creating the GitHub Repository

A repository named lmbr-devops-golang-test was created on GitHub.

Link to the repository: GitHub Repository

Initializing Git in the Project

The project was initialized as a local Git repository:

bash
Copy code
git init
Initial Commit

The main.go file was added and committed:

bash
Copy code
git add main.go
git commit -m "Add initial Golang Hello Leroy service code"
3. Using Git Branches
Creating a New Branch

A development branch named develop was created:

bash
Copy code
git checkout -b develop
Working on the Branch

Modifications were made on the develop branch and regularly committed.

Switching Between Branches

To switch between branches, the following was used:

bash
Copy code
git checkout main
Merging Branches

After testing the code on the develop branch, the changes were merged into main:

bash
Copy code
git checkout main
git merge develop
4. Docker Setup
Creating the Dockerfile

A Dockerfile was created to containerize the service:

Dockerfile
Copy code
# Use a Golang base image
FROM golang:1.19-alpine

# Create a working directory
WORKDIR /app

# Copy the built binary to the container
COPY hello-leroy .

# Expose port 80
EXPOSE 80

# Run the binary
CMD ["./hello-leroy"]
Building the Docker Image

The Docker image was built from the Dockerfile:

bash
Copy code
docker build -t hello-leroy:latest .
Running the Docker Container

The container was run locally:

bash
Copy code
docker run -p 80:80 hello-leroy:latest
This document outlines the key steps involved in setting up and containerizing the Hello Leroy web service, providing a straightforward guide for similar projects.

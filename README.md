
  

# **Golang Task Quiz**

  

  

This is a simple CLI application that interacts with a Quiz API.

  

  

## Usage

  

  

### Build and Run Locally

  

1. Clone the repository:
```bash
git clone https://github.com/furkanansn/golang-task-quiz.git
```

2. Navigate to the project directory:
```bash
cd golang-task-quiz
```

3. Build the API application:
```bash
go build -o golang-task-quiz .
```
4. Run the API application:
```bash
./golang-task-quiz
```
2. Navigate to the CLI project directory:

```bash
cd cmd/quiz-cli/
```  
3. Build the CLI application:
```bash
go build -o quiz-cli .

```
4. Run the CLI application:
```bash
./quiz-cli <command> [arguments]
```

Replace `<command>` with one of the available commands (e.g., `questions`, `submit`, `user-score`, `compare-score`), and provide any necessary arguments.

### Docker

#### Build the Docker Image

  

  

bashCopy code

  

  

`docker build -t golang-task-quiz .`

  

  

#### Run the Docker Container

  

  

bashCopy code

  

  

`docker run -p 8080:8080 golang-task-quiz`

  

  

The application will be accessible at [http://localhost:8080](http://localhost:8080/).

  

  

## Commands

  

  

-  `questions`: Get quiz questions.

  

-  `submit`: Submit quiz answers.

  

-  `user-score <userID>`: Get user score for a specific user.

  

-  `compare-score <userID>`: Compare user score with others.

  

  

## API Interaction

  

  

Make sure your API server is running and accessible at the specified base URL in the `main.go` file.

# Todolist Service API

## Installation & Setup

### Prerequisites

- Golang 1.16+
- PostgreSQL

### Steps

1. Clone the repository
   `bash
git clone https://github.com/your-repository/todolist-service.git
cd todolist-service`

2. Install dependencies
   `go mod tidy`

3. Configure PostgreSQL and create tables
   `psql -U myuser -d todolist_db -f db/schema.sql`

4. Run the project
   `go run main.go`

API Documentation
Base URL
http://localhost:8080

-Endpoints-

1. Create Task
   URL: POST /tasks
   Request Body:
   json
   {
   "start_date": "2024-11-01",
   "plan_date": "2024-11-05",
   "detail": "Complete the Go project",
   "status": false
   }
   Response:
   Status Code: 201 Created
   Body:
   json
   {
   "id": 1,
   "start_date": "2024-11-01",
   "plan_date": "2024-11-05",
   "detail": "Complete the Go project",
   "status": false
   }

2. Get All Tasks
   URL: GET /tasks
   Response:
   Status Code: 200 OK
   Body:
   json
   {
   "id": 1,
   "start_date": "2024-11-01",
   "plan_date": "2024-11-05",
   "detail": "Complete the Go project",
   "status": true
   }
3. Update Task
   URL: PUT /tasks/{id}
   Request Body:
   json
   {
   "start_date": "2024-11-02",
   "plan_date": "2024-11-06",
   "detail": "Updated task detail",
   "status": false
   }
   Response:
   Status Code: 200 OK
   Body:
   json
   {
   "id": 1,
   "start_date": "2024-11-02",
   "plan_date": "2024-11-06",
   "detail": "Updated task detail",
   "status": false
   }
   Error Responses:
   Status Code: 404 Not Found
   Body:
   json
   {
   "error": "Task not found"
   }
4. Delete Task
   URL: DELETE /tasks/{id}
   Response:
   Status Code: 200 OK
   Body:
   json
   {
   "message": "Task deleted successfully"
   }
   Error Responses:
   Status Code: 404 Not Found
   Body:
   json
   {
   "error": "Task not found"
   }

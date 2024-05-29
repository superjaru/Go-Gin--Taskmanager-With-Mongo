# Features

   * Task CRUD Operations: Create, Read, Update, and Delete tasks.
*  Authentication: User registration and login functionality with JWT authentication.
* Authorization: Ensure users can only access and modify their own tasks.
* Task Categories: Categorize tasks into different categories.
* Search and Filter: Search tasks by title and filter them based on status or category.
* Pagination: Paginate through the list of tasks.
 * Error Handling: Handle errors gracefully and return appropriate HTTP status codes.
  
# Techonologies Used
 * Go: Programming language used for backend development.
* Gin: Web framework used for building the RESTful API.
* MongoDB: NoSQL database used for storing tasks and user information.
* JWT (JSON Web Tokens): Used for user authentication and authorization.
* MongoDB: MongoDB driver for Go used for interacting with the database.


## Usage

```python
Register a new user: POST /api/auth/register
Login: POST /api/auth/login
Create a task: POST /api/tasks
Retrieve tasks: GET /api/tasks
Update a task: PUT /api/tasks/:id
Delete a task: DELETE /api/tasks/:id
```

## License

[MIT](https://choosealicense.com/licenses/mit/)

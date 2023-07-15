# Golang, GORM and GIN Framework JWT Authentication Application

### We use golang GIN framework, GORM, JWT-Auth and Postgres for this project.

## Steps to run the project
1. Clone the repo
2. Run the command
   ``` go mod download ```
3. Rename the .env.example file to .env and change variable's value 
4. Open Postman and send request using following routes

   POST: http://localhost:3000/signup (Signup)

   POST: http://localhost:3000/login (Login)

   GET: http://localhost:3000/validate (A request to validate the requireAuth middleware)

   ## Happy Coding

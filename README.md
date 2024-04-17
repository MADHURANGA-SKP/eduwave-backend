# Eduwave-Back-End

Eduwave is an LMS that was created for those who want to learn various courses with the guidance of teachers from recognized institutes.

## Installation

1. clone project using HTTP of SSH method 
```bash
git clone <link>
```
2. create docker image in docker desktop
```bash
docker run -d --name eduwave -p 5432:5432 -e POSTGRES_USER=user -e POSTGRES_PASSWORD=12345 postgres:16-alpine
```
3. create psql database in docker image 
```bash
docker exec -it eduwave createdb --username=user --owner=user eduwave
```
4. make migrate up for create tables in the database 
```bash
migrate -path db/migration -database "$(DB_URL)" -verbose up
```
5. start docker image  
```bash
docker start eduwave
```
6. Update go module dependencies 
```bash
go mod tidy
```
7. change .env file 

8. type make command or "go run main.go"
```bash
make server
```
## Requirement 

As a student, I would like to have a complete Learning Management System (LMS)

## Contributing

Please make sure to resolve conflicts before merge, do not force push.

Feel free to improve or add changes, before pushing create a new branch, and make sure to add the appropriate commit message.

If is there major conflicts make sure to contact or get help from the admin or other contributors.

## Team 🤝

[Pasan Madhuranga](https://github.com/MADHURANGA-SKP/MADHURANGA-SKP)

[Buddhima Kaushalya](https://github.com/BuddhimaKaushalya)



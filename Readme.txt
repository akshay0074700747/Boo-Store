This is a Simple Book Store API built by following the principles of Clean Architecture.
The Authentication and Authorization is done using JWT Tokens.
This API consists of two modules (Admin,User).
This API consists of 4 enpoints
login ->   POST http://localhost:8080/login
home -> GET http://localhost:8080/user/home
addBook -> POST http://localhost:8080/user/admin/addBook
deleteBook -> DELETE http://localhost:8080/user/admin/deleteBook/:bookName
The main file is in the cmd folder
Documentation -> https://documenter.getpostman.com/view/29091626/2sA3JKc295
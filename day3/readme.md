 # 2. Basic Framework (Echo)  & Alternative Framework | Connection and Interaction DB & ORM (Gorm) | MVC Architecture
**Objective** :

- Organize code using MVC pattern
- Create static CRUD API using Echo
- Create dynamic CRUD API that connect to database using Echo & Gorm

**Tasks** :

1. Organize your code using MVC (create config, controllers, lib, models, routes)
2. Create static CRUD API that meet requirements below:

| **Method** | **Endpoint** | **Description** |
| --- | --- | --- |
| GET | /books | Get all books |
| GET | /books/:id | Get book by id |
| POST | /books | Create new book |
| PUT | /books/:id | Update book by id |
| DELETE | /books/:id | Delete book by id |

1. Create dynamic CRUD API that meet requirements below:

| **Method** | **Endpoint** | **Description** |
| --- | --- | --- |
| GET | /users | Get all users |
| GET | /users/:id | Get user by id |
| POST | /users | Create new user |
| PUT | /users/:id | Update user by id |
| DELETE | /users/:id | Delete user by id |

You can import [this](https://drive.google.com/file/d/15c01qQsmp-CTYR7qDC3740Y81LJRN1Ib/view?usp=sharing) Postman collection as an final result example.

**References:**

HTTP methods and related SQL query

GET:

- [https://gorm.io/docs/query.html#Retrieving-all-objects](https://gorm.io/docs/query.html#Retrieving-all-objects)
- [https://gorm.io/docs/query.html#Retrieving-all-objects](https://gorm.io/docs/query.html#Retrieving-all-objects)

POST

- [https://gorm.io/docs/create.html#Create-Record](https://gorm.io/docs/create.html#Create-Record)

DELETE

- [https://gorm.io/docs/delete.html#Delete-with-primary-key](https://gorm.io/docs/delete.html#Delete-with-primary-key)
- [https://gorm.io/docs/delete.html#Soft-Delete](https://gorm.io/docs/delete.html#Soft-Delete)

PUT

- [https://gorm.io/docs/update.html#Update-single-column](https://gorm.io/docs/update.html#Update-single-column)
- [https://gorm.io/docs/update.html#Updates-multiple-columns](https://gorm.io/docs/update.html#Updates-multiple-columns)
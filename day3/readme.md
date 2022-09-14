 # 3. Middleware
**Objective** :

- Log Implementation
- Implementing JWT Auth for Protecting API

**Tasks** :

- Implement **log middleware** to Day 2 task
- Implement **JWT Auth middleware** based on requirements below

| **Method** | **Endpoint** | **JWT** |
| --- | --- | --- |
| GET | /books | Not Authenticated |
| GET | /books/:id | Not Authenticated |
| POST | /books | Authenticated |
| PUT | /books/:id | Authenticated |
| DELETE | /books/:id | Authenticated |

| **Method** | **Endpoint** | **JWT** |
| --- | --- | --- |
| GET | /users | Authenticated |
| GET | /users/:id | Authenticated |
| POST | /users | Not Authenticated |
| PUT | /users/:id | Authenticated |
| DELETE | /users/:id | Authenticated |

It is also a good thing if you can implement **Validator middleware** , and plus point if you can make a user only able to do PUT and DELETE on himself.
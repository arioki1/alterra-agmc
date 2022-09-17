 # 4. Unit & Integration Testing
 **Objective** :

- Implement Integration Testing

**Task** :

- Using your project from previous day, implement integration testing on controller
- Minimum test case per function is 2

- Valid test case, e.g. http response 200, 201
- Invalid test case, e.g. http response 400, 401, 500

- No minimum coverage percentage, but higher is better

## Test Case
```bash
➜  day4 git:(master) gocoverage   
?       github.com/arioki1/alterra-agmc/day4    [no test files]
?       github.com/arioki1/alterra-agmc/day4/config     [no test files]
ok      github.com/arioki1/alterra-agmc/day4/controllers        (cached)
?       github.com/arioki1/alterra-agmc/day4/lib/database       [no test files]
?       github.com/arioki1/alterra-agmc/day4/lib/database/seeder        [no test files]
?       github.com/arioki1/alterra-agmc/day4/middlewares        [no test files]
?       github.com/arioki1/alterra-agmc/day4/models     [no test files]
?       github.com/arioki1/alterra-agmc/day4/routes     [no test files]
?       github.com/arioki1/alterra-agmc/day4    [no test files]
?       github.com/arioki1/alterra-agmc/day4/config     [no test files]
ok      github.com/arioki1/alterra-agmc/day4/controllers        (cached)        coverage: 92.9% of statements
?       github.com/arioki1/alterra-agmc/day4/lib/database       [no test files]
?       github.com/arioki1/alterra-agmc/day4/lib/database/seeder        [no test files]
?       github.com/arioki1/alterra-agmc/day4/middlewares        [no test files]
?       github.com/arioki1/alterra-agmc/day4/models     [no test files]
?       github.com/arioki1/alterra-agmc/day4/routes     [no test files]
?       github.com/arioki1/alterra-agmc/day4    [no test files]
?       github.com/arioki1/alterra-agmc/day4/config     [no test files]
ok      github.com/arioki1/alterra-agmc/day4/controllers        1.398s  coverage: 92.9% of statements
?       github.com/arioki1/alterra-agmc/day4/lib/database       [no test files]
?       github.com/arioki1/alterra-agmc/day4/lib/database/seeder        [no test files]
?       github.com/arioki1/alterra-agmc/day4/middlewares        [no test files]
?       github.com/arioki1/alterra-agmc/day4/models     [no test files]
?       github.com/arioki1/alterra-agmc/day4/routes     [no test files]
github.com/arioki1/alterra-agmc/day4/controllers/bookController.go:15:  GetBooksControllers             100.0%
github.com/arioki1/alterra-agmc/day4/controllers/bookController.go:22:  CreateBookControllers           100.0%
github.com/arioki1/alterra-agmc/day4/controllers/bookController.go:50:  GetBookByIdControllers          100.0%
github.com/arioki1/alterra-agmc/day4/controllers/bookController.go:75:  UpdateBookByIdControllers       100.0%
github.com/arioki1/alterra-agmc/day4/controllers/bookController.go:120: DeleteBookByIdControllers       100.0%
github.com/arioki1/alterra-agmc/day4/controllers/userController.go:13:  GetUserControllers              100.0%
github.com/arioki1/alterra-agmc/day4/controllers/userController.go:25:  CreateUserControllers           77.8%
github.com/arioki1/alterra-agmc/day4/controllers/userController.go:51:  GetUserByIdControllers          80.0%
github.com/arioki1/alterra-agmc/day4/controllers/userController.go:81:  UpdateUserByIdControllers       88.9%
github.com/arioki1/alterra-agmc/day4/controllers/userController.go:125: DeleteUserByIdControllers       88.9%
github.com/arioki1/alterra-agmc/day4/controllers/userController.go:155: LoginUsersControllers           88.9%
github.com/arioki1/alterra-agmc/day4/controllers/userController.go:175: getUserId                       100.0%
total:                                                                  (statements)                    92.9%

```
# todo-challenge
30 Days Challenge

![30-days-challenge](https://raw.githubusercontent.com/vydao/rails-todochallenge/master/public/30dayschallenge.png)

## Features
+ Account creation
+ Login
+ Todo list creation
+ Add tasks
+ Invite friend to task
+ Challenge statictics
+ Bonus points:
  - Write the tests first
  - Clean code
  - Set privacy to task ( public/private (only invitted can join))
  - Nice template/ layout
  - Deploy your app to heroku or any servers
  

## Notes
+ Database postgreSQL (but you can use any other one you like, e.g Mysql,...)
+ Authentication using JWT or Pasto
+ Use Gin web framework

## API Document
+ Create User
```
POST /api/v1/users
{
    "username": "vydao",
    "password": "secret"
}
```

+ User Login
```
POST /api/v1/users/login
{
    "username": "vydao",
    "password": "1234qwer"
}
```

+ Get User Profile
```
GET /api/v1/users/:id
```

+ Create Challenge
```
POST /api/v1/challenges
{
    "name": "new challenge",
    "description": "Lorem ipsum dolor sit amet",
    "start_date": "2022-10-22T15:04:05"
}
```

+ Create Todo
```
POST /api/v1/challenges/:challenge_id/todos
{
    "name": "first todo",
    "point": 42.5,
    "period": "weekly"
}
```

+ Accept Challenge
```
POST /api/v1/challenges/:challenge_id/accept
```

+ Get Todo List By Challenge
```
GET /api/v1/challenges/1/todos
```

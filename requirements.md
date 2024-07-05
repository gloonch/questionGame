# Game Application

# use-case

## User use-cases

### Register
user can register to application by phone number

### Login
user can log in to the application by phone number and password

## Game use-cases
### Each game have a given number of questions
### The difficulty level of questions are "easy, medium, hard"
### Game winner is determined by the number of correct answers that each user has answered
### Each game does belong to specific category: sport, history, etc

# entity

## User
- ID
- Phone number
- Avatar
- Name

## Game
- ID
- Category
- Question list
- Players
- WinnerID

## Question
- ID
- Question
- Answer
- Correct Answer
- Difficulty
- Category
# cardapi
this is the backend of the flashcard repo

# Flashcard API

## Introduction

Welcome to the Flashcard API! This API provides endpoints to manage flashcards for studying purposes. Whether you're learning a new language, preparing for exams, or simply want to improve your memory, this API can help you organize and access your study materials efficiently.

## Features

- Create, read, update, and delete flashcards.
- Organize flashcards into decks for better categorization.
- User authentication and authorization using JWT tokens.
- Cross-Origin Resource Sharing (CORS) enabled for easy integration with frontend applications.
- Real-time updates with Server-Sent Events (SSE) for interactive study sessions.

## Requirements

- Go 1.20 or later
- MySQL database
- Dependencies listed in `go.mod` file

## Installation

1. Clone this repository:

   ```bash
   git clone https://github.com/programmer-for-good/flashcardApi.git

   ```
 1. Set up your MySQL database and update the database configuration in config.go.
 2. Install dependencies:
    ```
    go mod tidy
    ```
 3. Build and run the application:
    ```
    go build
    ./flashcardApi
    ```

  # Usage
  - Register a new account or log in with existing credentials.
- Create decks to organize your flashcards.
- Add new flashcards to your decks or update existing ones.
- Study your flashcards using the provided endpoints or integrate with your favorite flashcard application.
# Contributing
- Contributions are welcome! If you have any suggestions, bug fixes, or feature requests, please open an issue or submit a pull request.
# License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

# Airline Voucher Seat Assignment App

This project is a web application for an airline campaign to randomly assign 3 unique seat numbers for a given flight.

## Project Structure

### backend

The Go (Golang) server application.

### frontend

The React client application.

### docker-compose.yml

A file to orchestrate the backend and frontend services.

## Prerequisites

- Docker
- A modern web browser (e.g., Chrome, Firefox)
- Postman (Optional, for API testing)

## How to Run

### Clone the repository:

```
git clone https://github.com/jamesriady/voucher-seat-assignment.git
cd voucher-seat-assignment
```

### Build and Run the Application

Build and run the application using Docker Compose: 
From the root directory of the project, run the following command:
```
docker-compose up --build
```

This command will build the Docker images for both the Go backend and the React frontend and then start the services.

### Access the Application

- The React frontend will be available at http://localhost:3000.

- The Go backend will be available at http://localhost:8080.

### Stop the Application

To stop the application, press `Ctrl+C` in the terminal where `docker-compose up` is running. To remove the containers, you can run:
```
docker-compose down
```

## How it Works

### Backend (Go)

- **Web Framework:** Uses `gorilla/mux` for routing.

- **Database:** Uses `sqlite3` for data persistence. The database file `vouchers.db` will be created in the `backend/data` directory.

- **API Endpoints:**
    - `POST /api/check`: Checks if vouchers have already been generated for a specific `flightNumber` and `date`.
    - `POST /api/generate`: Generates 3 unique random seats for a flight, and stores the details in the database.

- **CORS**: A middleware is implemented to allow cross-origin requests from the frontend.

### Frontend (React)

- **Framework:** UBuilt with Create React App.

- **Styling:** Uses Tailwind CSS for a clean and modern user interface.

- **Functionality:**

    1. A form to input crew and flight details.

    2. On submission, it first calls the `/api/check` endpoint.

    3. If vouchers don't exist, it calls the `/api/generate` endpoint.

    4. Displays the 3 randomly generated seats upon success.

    5. Displays a clear error message if vouchers have already been generated or if any other error occurs.
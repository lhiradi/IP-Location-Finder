# IP API Project

## Project Title
A Go application for handling IP data, providing an API to retrieve and save IP information.

## Installation
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd IP-API
   ```

2. Create a `.env` file in the root directory with the following variables:
   ```
   DB_HOST=your_db_host
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   DB_PORT=your_db_port
   ```

3. Build and run the application using Docker:
   ```bash
   docker-compose up --build
   ```

## Usage
To run the application, execute:
```bash
go run cmd/api.go
```
The application will listen on port 8080.

## API Endpoints
- **POST** `/api/v1/ip`: Retrieve IP data.
  - **Query Parameters**: 
    - `ip`: The IP address to retrieve data for.
  - **Response**: 
    - Returns a JSON object with the IP and country name.

## Database Configuration
The application uses PostgreSQL for data storage. Ensure the database is running and accessible.

## Contributing
Contributions are welcome! Please open an issue or submit a pull request.

## License
This project is licensed under the MIT License.

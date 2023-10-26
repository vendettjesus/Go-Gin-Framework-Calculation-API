# Go-Gin-Framework-Calculation-API
This project is a web API written in the Go programming language using the Gin framework. The API allows you to calculate the average of provided grades and returns the results in JSON format. It can be used to calculate student grades' averages or other assessments.

# Usage Guide:

## Clone the Repository:

Clone this repository to your local machine using the git clone command.

## Install Dependencies:

Ensure you have Go installed on your system.

## Run the Application:

Execute the Go application with the go run main.go command.

The application will run on port 8080 by default(you can change the port in the code if necessary).

## API Route:

/calcularCalificacion: Use this route to calculate the grade point average. Make a GET request to http://localhost:8080/calcularCalificacion with the following query parameters:

name (student name): Provide the student's name.
cal1, cal2, cal3 (grades): Provide the grades you want to average.
The API will calculate the average and return a JSON response with the name and the calculated average.

### Example GET request using cURL:

curl "http://localhost:8080/calcularCalificacion?name=Student&cal1=10&cal2=9.5&cal3=8.8"

Sample response:

{
  "name": "Student",
  "promedio": 9.1
}

# ZopsmartProject
This project builds a robust API for managing library resources in Go. Utilizing the Gofr framework, it offers RESTful CRUD operations for books, enabling efficient inventory control. Users can retrieve all books, create new entries, update details like availability, and remove borrowed books. Comprehensive unit tests ensure reliable functionality, while clear documentation facilitates seamless integration. This API empowers librarians and simplifies book management, ultimately enhancing the library experience for everyone.

Features:

  CRUD Operations: Create, Read, Update, and Delete book data.
  Browse Books: Get a list of all available books in the library.
  Access Individual Books: Retrieve details of a specific book using its unique identifier.
  Manage Borrowing: Update book status to reflect borrowing and availability.
  Scalable and Efficient: Built with Go and Gofr for a robust and performant API.

Getting Started:

  Clone the repository: git clone https://github.com/prathambhargava/ZopsmartProject.git
  Install dependencies: go get gofr.dev
  Install dependencies: go get gofr.dev/pkg/gofr
  Run the server: go run main.go
  Test the API: Use tools like Postman or curl to send HTTP requests to your local server (default port: 8080).

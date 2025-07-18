# API Music Catalog

This repository contains a Go-based API for a music catalog application. It integrates with the Spotify API to provide music search and recommendations, and also includes user management (sign-up, login) and track activity tracking (liking/disliking songs).

## Features

  * **User Membership**:
      * User registration (Sign Up)
      * User authentication (Login)
  * **Track Search**: Search for tracks using the Spotify API.
  * **Track Recommendations**: Get music recommendations based on a specific track from the Spotify API.
  * **Track Activities**: Record user interactions like liking or disliking songs.

## Technologies Used

  * **Go**: The primary programming language.
  * **Gin Gonic**: Web framework for building the API.
  * **Gorm**: ORM (Object-Relational Mapper) for database interactions.
  * **PostgreSQL**: Relational database for storing user and track activity data.
  * **Viper**: For configuration management.
  * **Bcrypt**: For secure password hashing.
  * **JWT (JSON Web Tokens)**: For user authentication.
  * **GoMock**: Used for generating mock interfaces for testing.
  * **go-sqlmock**: Used for mocking database interactions in tests.

## Project Structure

```
.
├── cmd                 # Main application entry point
│   └── main.go
├── internal            # Internal packages (configs, handlers, models, middleware, repositories, services)
│   ├── configs         # Configuration loading and types
│   ├── handler         # HTTP handlers for different API routes
│   │   ├── memberships
│   │   └── tracks
│   ├── middleware      # Custom Gin middleware (e.g., authentication)
│   ├── models          # Data models and request/response structures
│   │   ├── memberships
│   │   ├── spotify
│   │   └── trackacktivities
│   ├── repository      # Database and external API interaction logic
│   │   ├── memberships
│   │   ├── spotify
│   │   └── trackactivities
│   └── service         # Business logic layer
│       ├── memberships
│       └── tracks
└── pkg                 # Utility packages (httpclient, internalsql, jwt)
    ├── httpclient      # HTTP client wrapper
    ├── internalsql     # Internal SQL connection utility
    └── jwt             # JWT token creation and validation
```

## Getting Started

### Prerequisites

  * Go (version 1.18 or higher recommended)
  * Docker and Docker Compose (for running PostgreSQL)

### Installation

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/gigihrahman/api-music-catalog.git
    cd api-music-catalog
    ```

2.  **Install Go dependencies:**

    ```bash
    go mod download
    ```

3.  **Set up PostgreSQL using Docker Compose:**

    The `docker-compose.yml` file defines a PostgreSQL service.

    ```yaml
    version: "3"
    services:
      db:
        image: postgres:latest
        container_name: local-postgres-fast-campus
        ports:
          - "5432:5432"
        environment:
          POSTGRES_USER: admin
          POSTGRES_PASSWORD: root
          POSTGRES_DB: test_db
        volumes:
          -
    ```

    Run the following command to start the database:

    ```bash
    docker-compose up -d db
    ```

    This will start a PostgreSQL container named `local-postgres-fast-campus` accessible on port `5432`.

4.  **Configure the application:**

    Create a `config.yaml` file in the `internal/configs/` directory (if it doesn't exist) with your database and Spotify API credentials. The application looks for `config.yaml` in this directory.

    Example `internal/configs/config.yaml`:

    ```yaml
    service:
      port: ":8080"
      secretJWT: "your_jwt_secret_key" # Replace with a strong, unique key
    database:
      dataSourceName: "host=localhost user=admin password=root dbname=test_db port=5432 sslmode=disable"
    spotifyConfig:
      clientID: "YOUR_SPOTIFY_CLIENT_ID" # Replace with your Spotify API Client ID
      clientSecret: "YOUR_SPOTIFY_CLIENT_SECRET" # Replace with your Spotify API Client Secret
    ```

    You can obtain Spotify API credentials by registering an application on the [Spotify for Developers Dashboard](https://developer.spotify.com/dashboard/).

### Running the Application

To run the API, execute the `main.go` file:

```bash
go run cmd/main.go
```

The API will start on the port specified in your `config.yaml` (default: `:8080`).

## API Endpoints

All endpoints under `/tracks` require JWT authentication via the `Authorization` header (`Bearer <token>`).

### Membership Endpoints

#### Register a New User

  * **URL**: `/memberships/sign_up`
  * **Method**: `POST`
  * **Request Body**:
    ```json
    {
      "email": "user@example.com",
      "username": "myusername",
      "password": "securepassword"
    }
    ```
  * **Success Response**: `HTTP 201 Created`
  * **Error Response**: `HTTP 400 Bad Request` (e.g., "email or username exists")

#### User Login

  * **URL**: `/memberships/login`

  * **Method**: `POST`

  * **Request Body**:

    ```json
    {
      "email": "user@example.com",
      "password": "securepassword"
    }
    ```

  * **Success Response**: `HTTP 200 OK`

    ```json
    {
      "accesToken": "your_jwt_access_token"
    }
    ```

  * **Error Response**: `HTTP 400 Bad Request` (e.g., "email is not exist", "email and password not matched")

### Tracks Endpoints (Authenticated)

#### Search Tracks

  * **URL**: `/tracks/search`

  * **Method**: `GET`

  * **Query Parameters**:

      * `query`: (Required) The search query string (e.g., "Bohemian Rhapsody").
      * `pageSize`: (Optional) Number of results per page (default: 10).
      * `pageIndex`: (Optional) Page number (default: 1).

  * **Headers**: `Authorization: Bearer <access_token>`

  * **Success Response**: `HTTP 200 OK`

    ```json
    {
      "offset": 0,
      "limit": 10,
      "total": 905,
      "items": [
        {
          "albumType": "album",
          "totalTracks": 22,
          "albumImagesURL": [
            "https://i.scdn.co/image/...",
            "https://i.scdn.co/image/...",
            "https://i.scdn.co/image/..."
          ],
          "albumName": "Bohemian Rhapsody (The Original Soundtrack)",
          "artistsName": ["Queen"],
          "explicit": false,
          "id": "3z8h0TU7ReDPLIbEnYhWZb",
          "name": "Bohemian Rhapsody",
          "isLiked": true
        }
      ]
    }
    ```

  * **Error Response**: `HTTP 400 Bad Request`

#### Get Track Recommendations

  * **URL**: `/tracks/recommendations`

  * **Method**: `GET`

  * **Query Parameters**:

      * `limit`: (Optional) Number of recommendations (default: 10).
      * `trackID`: (Required) The Spotify ID of the seed track for recommendations.

  * **Headers**: `Authorization: Bearer <access_token>`

  * **Success Response**: `HTTP 200 OK`

    ```json
    {
      "items": [
        {
          "albumType": "album",
          "totalTracks": 22,
          "albumImagesURL": [
            "https://i.scdn.co/image/...",
            "https://i.scdn.co/image/...",
            "https://i.scdn.co/image/..."
          ],
          "albumName": "Bohemian Rhapsody (The Original Soundtrack)",
          "artistsName": ["Queen"],
          "explicit": false,
          "id": "3z8h0TU7ReDPLIbEnYhWZb",
          "name": "Bohemian Rhapsody",
          "isLiked": true
        }
      ]
    }
    ```

  * **Error Response**: `HTTP 400 Bad Request`

#### Upsert Track Activity

  * **URL**: `/tracks/track-activity`

  * **Method**: `POST`

  * **Headers**: `Authorization: Bearer <access_token>`

  * **Request Body**:

    ```json
    {
      "spotifyID": "spotify_track_id",
      "isLiked": true  // true for liked, false for disliked, null for neutral
    }
    ```

  * **Success Response**: `HTTP 200 OK`

  * **Error Response**: `HTTP 400 Bad Request`

## Running Tests

To run the unit tests for the project, navigate to the root directory and execute:

```bash
go test ./...
```

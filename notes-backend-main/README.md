# Notes API for CRUD and Auth

## To run `docker-compose up -d`

# Routes

## User Routes

### Signup

Endpoint: `POST /api/signup`

Creates a new user.

#### Request

```http
POST /api/signup
Content-Type: application/json

{
  "username": "example_user",
  "password": "your_password"
}

```

### Login(To get Bearer Token)

Endpoint: `POST /api/login`

```http
POST /api/login
Content-Type: application/json

{
  "username": "example_user",
  "password": "your_password"
}
```

## Notes Routes (Auth Token Required)

### CreateNote

Endpoint: `POST /api/note/`

```http
POST /api/note
Authorization: Bearer <your_token_here>
Content-Type: application/json

{
  "title": "Note Title",
  "text": "Note content"
}

```

### GetAllNotes

```http
GET /api/note
Authorization: Bearer <your_token_here>
```

Endpoint: `GET /api/note/`

### GetNoteByID

Endpoint: `GET /api/note/[id]`

```http
GET /api/note/[id]
Authorization: Bearer <your_token_here>

```

### UpdateNote

Endpoint: `PUT /api/note/[id]`

```http
PUT /api/note/[id]
Authorization: Bearer <your_token_here>
Content-Type: application/json

{
  "title": "Updated Title",
  "text": "Updated content"
}
```

### SearchNote

Endpoint: `GET /api/search?query=<your_search_query>`

```http
GET /api/search?query=important
Authorization: Bearer <your_token_here>

```

### DeleteNote

Endpoint: `DELETE /api/note/[id]`

```http
DELETE /api/note/[id]
Authorization: Bearer <your_token_here>

```

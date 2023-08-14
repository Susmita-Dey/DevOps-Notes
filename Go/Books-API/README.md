# Books API

This is a simple API to manage books. It is built with Go and uses MongoDB as a database.

## Endpoints

### GET /books

Returns a list of books.

Command: `curl -X GET http://localhost:8080/books`

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| `id` | `string` | The id of the book. |
| `title` | `string` | The name of the book. |
| `author` | `string` | The name of the author. |
| `quantity` | `int` | The number of books present or the quantity of the book. |


#### Response

| Name | Type | Description |
| ---- | ---- | ----------- |
| `id` | `string` | The id of the book. |
| `title` | `string` | The name of the book. |
| `author` | `string` | The name of the author. |
| `quantity` | `int` | The number of books present or the quantity of the book. |

### GET /books/{id}

Returns a book with the given id.

Command: `curl -X GET http://localhost:8080/books/1`

#### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| `id` | `string` | The id of the book. |
| `title` | `string` | The name of the book. |
| `author` | `string` | The name of the author. |
| `quantity` | `int` | The number of books present or the quantity of the book. |

#### Response

| Name | Type | Description |
| ---- | ---- | ----------- |
| `id` | `string` | The id of the book. |
| `title` | `string` | The name of the book. |
| `author` | `string` | The name of the author. |
| `quantity` | `int` | The number of books present or the quantity of the book. |


### POST /books

Creates a new book.

Command: 
```bash
curl -X POST http://localhost:8080/books -d '{"id": "2","title": "The Lord of the Rings", "author": "J.R.R. Tolkien", "quantity": 3}'
```
or
```bash
curl localhost:8080/books --include --header "ontent-Type: application/json" -d @body.json --request "POST"
```
where `body.json` contains the following:
```json
{
    "id": "2",
    "title": "The Lord of the Rings",
    "author": "J.R.R. Tolkien",
    "quantity": 3
}
```
##### Parameters

| Name | Type | Description |
| ---- | ---- | ----------- |
| `id` | `string` | The id of the book. |
| `title` | `string` | The name of the book. |
| `author` | `string` | The name of the author. |
| `quantity` | `int` | The number of books present or the quantity of the book. |

##### Response

| Name | Type | Description |
| ---- | ---- | ----------- |
| `id` | `string` | The id of the book. |
| `title` | `string` | The name of the book. |
| `author` | `string` | The name of the author. |
| `quantity` | `int` | The number of books present or the quantity of the book. |


### CheckOut Book [PATCH /checkout?id={id}]

Checks out a book with the given id.

Command: `curl -X PATCH http://localhost:8080/checkout?id=1`

### PUT /books/{id}

Updates a book with the given id.
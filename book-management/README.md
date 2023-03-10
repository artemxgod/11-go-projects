# Book management with mySQL

## Layout

### `/cmd`

* Contains main application for this project

* The name of executable should match the directory name of application in this case it is `bmanager`

* main.go and invokes router configuration from `/pkg` and starts the server

### `/pkg`

* Contains router, models, controllers, utilities and storage implementation

- `/configs` initializes db

- `/controllers` implements handler functions for router 

- `/model` keeps book model and connects to database

- `/routes` has router func with all handlers

- `/utils` implements body parsing that is used in `/controllers`


### `/scheme`

* Contains drawio request and structure schemes

## What it does
* managing books using mySQL

* Server handle request on path `/book/` differently for each method
    - `GET /book/` responds with list of books
    - Respond example:
    ```json
    [
        {
            "ID": 1,
            "CreatedAt": "2023-01-29T19:07:10+03:00",
            "UpdatedAt": "2023-01-29T19:07:10+03:00",
            "DeletedAt": null,
            "name": "The Art 2",
            "author": "Artem",
            "publication": "Verum Prod"
        },
        {
            "ID": 2,
            "CreatedAt": "2023-01-29T19:07:40+03:00",
            "UpdatedAt": "2023-01-29T19:07:40+03:00",
            "DeletedAt": null,
            "name": "The Art",
            "author": "Nemo",
            "publication": "Nemo prod"
        }
    ]
    ```
    - `GET /book/{id}` seeks for book in db and responds with a matching `id` book
    - Respond example for `/book/1`:
    ```json
    {
    "ID": 1,
    "CreatedAt": "2023-01-29T19:07:10+03:00",
    "UpdatedAt": "2023-01-29T19:07:10+03:00",
    "DeletedAt": null,
    "name": "The Art 2",
    "author": "Artem",
    "publication": "Verum Prod"
    }
    ```
    - `POST /book/` creates a new book in db. Book data should be passed in json format with the request. 
    - Respond example:
    ```json
    {
    "ID": 2,
    "CreatedAt": "2023-01-29T19:07:39.8661467+03:00",
    "UpdatedAt": "2023-01-29T19:07:39.8661467+03:00",
    "DeletedAt": null,
    "name": "The Art",
    "author": "Nemo",
    "publication": "Nemo prod"
    }
    ```
    - `PUT /book/{id}` changes matching with id movie's data. Respond with updated data in json or error
    - Respond example for `/book/1`:
    ```json
    {
    "ID": 1,
    "CreatedAt": "2023-01-29T19:07:10+03:00",
    "UpdatedAt": "2023-01-29T19:10:39.2164484+03:00",
    "DeletedAt": null,
    "name": "The Art",
    "author": "Nemo",
    "publication": "Nemo prod"
    }
    ```
    - `DELETE /book/{id}` deletes matching with id book from db. Respond with what left from it in database
    - Respond example for `/book/1`:
    ```json
    {
    "ID": 0,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "name": "",
    "author": "",
    "publication": ""
    }
    ```
    - `GET /book/author/{name}` responds with list of books with matching author
    - Respond example for `/book/author/Virtus`:
    ```json
    [
        {
            "ID": 1,
            "CreatedAt": "2023-01-30T12:29:07+03:00",
            "UpdatedAt": "2023-01-30T12:29:07+03:00",
            "DeletedAt": null,
            "name": "The end",
            "author": "Virtus",
            "publication": "Nemo prod"
        },
        {
            "ID": 2,
            "CreatedAt": "2023-01-30T12:29:14+03:00",
            "UpdatedAt": "2023-01-30T12:29:14+03:00",
            "DeletedAt": null,
            "name": "The end 2",
            "author": "Virtus",
            "publication": "Nemo prod"
        }
    ]
    ```
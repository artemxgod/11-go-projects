# CRUD Server

## Layout

### `/cmd`

* Contains main application for this project

* The name of executable should match the directory name of application in this case it is `crudserver`

* main.go does configuration and invokes server start from `/internal`

### `/internal`

* Contains webserver implementation as well as storage and data models

- `/crudserver`
    - crudserver.go initialize and start the server
    - server.go configures router and storage
    - config.go contains data for server configuration. The data should be predetermined
    - server_internal_test.go has the request tests for movie creation and movie search

- `/store`
    - store.go contains store interface, so we can use different type of storages. 
    Here i am using map as a storage. SQL databases can be used as well

- `/model`
    - movie.go contains movie structure and validation for its fields.
    - director.go contains director structure.


### `/scheme`

* Contains drawio schemes for this project

## What it does
* Implements C.R.U.D. interface for list of movies on a web server. Slice is used instead of database

* Server handle `/movies/` and has a different reaction for `GET` `POST` `PUT` and `DELETE` requests
    - `GET /movies` responds with list of movies
    - `GET /movies/{id}` responds with a matching with `id` movie of error
    - `POST /movies` creates new movie. Movie data should be passed in json format with the request. 
    Respond with new data in json or error.
    - `PUT /movies/{id}` changes matching with id movie's data. Respond with updated data in json or error
    - `DELETE /movies/{id}` delete matching with id movie. Respond with the list of remaining movies in json or error
    

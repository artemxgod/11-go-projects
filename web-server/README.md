# Web server in golang

## Layout

### `/cmd`

* Contains main application for this project

* The name of executable should match the directory name of application in this case it is `webserver`

* main.go does configuration and invokes server start from `/internal`

### `/internal`

* Contains webserver implementation as well as storage and data models

- `/webserver`
    - webserver.go initialize and start the server
    - server.go configures router and storage
    - config contains data for server configuration. The data should be predetermined

- `/store`
    - store.go contains store interface, so we can use different type of storages. 
    Here i am using map as a storage. SQL databases can be used as well

- `model`
    - user.go contains user structure. There might also be data validation and other functions that works with user data

### `/ui`

* Keeps html and static files

- `/html` for .html files
- `/static` for .css .js and images (not used in this project)


### `/scheme`

* Contains drawio schemes

## What it does

* Server handle `/`, `/form`, and `private/hello`
    - `/` execute simple html tepmlate
    - `/form` has information form. 
    After filling gaps and clicking submit `/form/new` route will be opened and will display user information.
    Also it creates a cookie, so we can reach `private/hello`
    - `/private/hello` says hello to user. It requires user name. 
    Can't reach the route before filling form in `/form`
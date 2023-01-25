# Web server in golang

## Layout

### `/cmd`

* Contains main application for this project

* The name of executable should match the directory name of application in this case it is `webserver`

* main.go does configuration and invokes server start from `/internal`

### `/internal`

* Contains webserver implementation as well as storage and data models

* `/webserver`
 * webserver.go initialize and start the server
 * server.go configures router and storage
 * config contains data for server configuration. The data should be predetermined

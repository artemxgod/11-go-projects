# Email verifier tool

## Layout

* Contains main application for this project

* The name of executable should match the directory name of application in this case it is `verifier`

* main.go invokes checker from `/pkg`


## What it does

* Not much. It checks domain for MX, SPF and DMARC. 

* As an additional task i also add email validation
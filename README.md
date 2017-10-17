# SQL CONNECTION CONFIGURATION
Setting the connection string for a SQL Server DB
Create a system variable 
------------------------
//Name CS_WIDGETDB
//Value your sql connection string 
...

# Usage
The main site is at ".\site\" folder. You can start it by runnig:
//go build
//.\site.exe
//The URL is localhost\555

The api is at ".\api\" folder. You can start it by runnig:
//go build
//.\api.exe
//The URL is localhost\666

# Widgets Single Page App Demo
This is a simple multi-page HTML site. The goal of this project is to take this hardcoded HTML site, and make it hit an API for showing/listing user and widget information. 

## Features
- A user view that displays a list of users (data via api `/users`), each user should have a method of clicking to viewing all the details of that user (`/user/:id`)
- A widget view that displays a list of widgets (`/widgets`), each widget should have a method of clicking to view the details of that widget (`/widget/:id`)
- A search/filter on the user and widget list views
- A method of creating a new widget (POST `/widget`)
- A method of updating an existing widget (PUT `/widget/:id`)
- A method of creating a new user (POST `/user`)
- A method of updating an existing user (PUT `/user/:id`)


# API Documentation
There's an API available at `http://spa.tglrw.com:4000` for retrieving the data used to make this app. The hard-coded data in the existing HTML is only a placeholder for style. The API returns and expects to receive JSON-encoded data.


## The endpoints are as follows:
- GET `/users` [http://spa.tglrw.com:4000/users](http://spa.tglrw.com:4000/users)
- GET `/user/:id` [http://spa.tglrw.com:4000/users/:id](http://spa.tglrw.com:4000/user/:id)
- GET `/widgets` [http://spa.tglrw.com:4000/widgets](http://spa.tglrw.com:4000/widgets)
- GET `/widget/:id` [http://spa.tglrw.com:4000/widgets/:id](http://spa.tglrw.com:4000/widgets/:id)
- POST `/widget` for creating new widgets [http://spa.tglrw.com:4000/widgets](http://spa.tglrw.com:4000/widgets)
- PUT `/widget/:id` for updating existing widgets [http://spa.tglrw.com:4000/widgets/:id](http://spa.tglrw.com:4000/widgets/:id)


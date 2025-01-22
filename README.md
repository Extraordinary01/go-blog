# GO blog pet project

This is a simple REST API written in Golang, using gin-gonic web framework. To launch dev server run this command:
`sudo docker-compose up --build`

There are 2 data tables: authors and blogs. So you can create authors and blog posts with REST API endpoints.

P.S. Wanted to add unit tests, but didn't know how to mock db connection properly. I could create a test db for this purpose and use it in testing functions, but decided to not do it. If you want me to add tests, just ask.

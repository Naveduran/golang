# First Go API

This is the first API I made using go language. It is based in the framework
Gin, it doesn't have a database, and save the information in memory. 

I take notes inside the code to explain which are the steps to create the
api, and what each line makes for learning purposes.

## Endpoints:
- /albums

    - GET – Get a list of all albums, returned as JSON.
    - POST – Add a new album from request data sent as JSON.

- /albums/:id

    -GET – Get an album by its ID, returning the album data as JSON.

## Installation
```go get .```

```go run .```

## Example

Get request to see the list of albums:
```
$ curl http://localhost:8080/albums
```

Answer:
```
    [
        {
                "id": "1",
                "title": "Blue Train",
                "artist": "John Coltrane",
                "price": 56.99
        },
        {
                "id": "2",
                "title": "Jeru",
                "artist": "Gerry Mulligan",
                "price": 17.99
        },
        {
                "id": "3",
                "title": "Sarah Vaughan and Clifford Brown",
                "artist": "Sarah Vaughan",
                "price": 39.99
        }
    ]
```

Post request to create an album:

```
$ curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```
Answer:

```
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Sun, 28 Aug 2022 18:50:07 GMT
Content-Length: 116

{
    "id": "4",
    "title": "The Modern Sound of Betty Carter",
    "artist": "Betty Carter",
    "price": 49.99
}
```

Get request for an specific album:
```
$ curl http://localhost:8080/albums/1
```
Answer:
```
{
    "id": "1",
    "title": "Blue Train",
    "artist": "John Coltrane",
    "price": 56.99
}
```

___

## Learning Notes:

- Validation of the data can be made through json labels added to the
structure declaration:
```
   FirstName string `json:"firstName" binding:"required"`
   LastName string `json:"lastName" binding:"required"`
   Email string `json:"email" binding:"required,email"`
   Phone string `json:"phone" binding:"required,e164"`
   CountryCode string `json:"countryCode" binding:"required,iso3166_1_alpha2"`
```

## Important functions
<br />

### **c.IndentedJSON(http_status, objects_to_serialize)**
Return HTTP response with Indented JSON as content

Types of HTTP Responses:
<ul>
<li>http.StatusOK</li>
<li>http.StatusNotFound</li>
</ul>
See other status here:
https://go.dev/src/net/http/status.go

<br />

### **c.BindJSON(structure_pointer)**
Put JSON in the structure, or return an error
<br />

### **gin.Default(url, view_name)**

Creates a router to register in all the views

* router.GET()
Creates a router to register in all the views
* router.POST()

### **Context.Param(parameter name)**
Receive params sended in the url.
<br />


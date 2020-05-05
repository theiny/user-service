# User Microservice

An example of a user microservice using DDD and based on a hexagonal architecture. 


## Run it
`go run $GOPATH/src/github.com/theiny/user-service/cmd/main.go`

or

```
cd $GOPATH/src/github.com/theiny/user-service/cmd/
./user-service
```

## Usage

### Healthcheck

Simple 200 response for healthchecking the service. 

- **URL**

`GET` /healthcheck

### Adding a User

Send JSON payload of a User object to create a new user. 

- **URL**
    
 `POST`   /api/v1/users/add
    
-   **Request Body**

```
{
	"first_name": "Brandon",
	"last_name": "Stark",
	"nickname": "Bran the Broken",
	"password": "someRandomSuperSecurePassword",
	"email": "kingofwesteros@winterfell.com",
	"country": "IE"
}
```

### Listing Users

Retrieves a list of users from storage in JSON format. Apply URL params to filter by field name. Values for query parameters are case insensitive. 

- **URL**

 `GET`   /api/v1/users/get

-   **URL Params**
    
    **Optional:** 
       `first_name=Brandon`
       `last_name=Stark`
       `nickname=Bran%20the%20Broken`
       `email=kingofwesteros@winterfell.com`
       `country=IE`

### Editing a User

Send JSON payload of a User object to edit an existing user. The ID of the user (generated by storage) should be passed in the URL. 

- **URL**
    
 `PUT`   /api/v1/users/edit/:id
    
-   **Request Body**

```
{
	"first_name": "Brandon",
	"last_name": "Stark",
	"nickname": "Bran the Broken",
	"password": "someRandomSuperSecurePassword",
	"email": "kingofwesteros@winterfell.com",
	"country": "IE"
}
```

### Deleting a User

Deletes an existing user. Send the ID of the user (generated by storage) in the URL as a param.

 `DELETE`   /api/v1/users/delete/:id
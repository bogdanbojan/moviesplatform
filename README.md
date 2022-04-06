## Movie platform management


---

*Readme*: 

TODO: Add services, features, permissions.

TODO: Explain the decisions throughout the project.

*App*:

TODO: Tests.

TODO: Containerize the app if I have time.


---

### Build:

This build assumes that you have at least Go 1.17 installed on your machine. If you do not: https://go.dev/doc/install

`go build .`


[//]: # (?just binary or the user should have go preinstalled?)

### Run:

`go run main.go`

*Optional*: set the address that the service will run by using the `-addr` flag. eg: `go run . -addr=":8080"`

### Use:
e.g : `curl -X GET http://localhost:4000/user/u324s`




---

*Why use json to store the data?*

The permissions are stored in a json which is embedded into the code. The `store.json` file is found in the project directory
and it is populated when the service is started locally.

This way of populating the permissions store is a simpler approach than using a docker container. It also makes use of the 
`embed` capabilities of go. It is also a coherent metal model of having k-v pairs in the already existing structure of a json.

This also pertains to unmarshalling the data in an easy manner into native go data structures.

*Why log things this way?*

The structure of the project is clearer if we have a centralized logging system. This helps us
with relevant information (i.e starting the server and seeing which address we are using) or error handling in a centralized, coherent fashion.

*Why use an interface for the data pulling phase?*

It helps us maintain orthogonality. It separates the actual way we store things (json,db,etc) from our
business logic and therefore, we have a middle layer which enables us to plug any other way of storing
the user permissions with minimal changes in our code, if need be. It's also helpful for our test - mocking things and such.

[//]: # (*How do you handle fetching the usernames?*)

[//]: # ()
[//]: # (It's easier for the end-user if we don't throw immediately a 4xx error when he doesn't nail the username)

[//]: # (quite right because of capitalization or spaces. That's why I think it's better to treat the query beforehand and)

[//]: # (compare it to our actual users wo spaces/capitalization.)

*Why use empty struct when checking the features/permissions?*

It allocates 0 memory. Also, by using a map we are leveraging the internal map algorithm that already has an efficient lookup time of avg (O(1)).
## Movie platform 

---

### Mental model and structuring of the services

The services are built around the *movie business*. They contain just little snippets of each domain in order to have a
clearer mental model of the structure that goes into the service.feature.permission model.

We have three **services** :

- blockbusters
- commercials
- shorts

---
### Blockbusters

**Features** :

- director
- cinematographer
- producer

**Permissions** :

director:

- direct : bool
- instructActors : bool
- argue : "allow"/"deny"

cinematographer:

- shoot : bool
- changeLens : bool
- changeCamera : "allow"/"deny"

producer:

- changeBudget : bool
- changeSalary : bool
- addActor : bool

### Commercials

**Features** :

- artist
- producer
- manager

**Permissions** :

artist:

- createConcept : bool
- creativitySwitch : "on"/"off"

producer:

- getDeals : "allow"/"deny"
- onboardPeople : bool

manager:

- adviseBoard : bool
- cancelMeetings : bool
- scheduleMeetings : bool

### Shorts

**Features** :

- actor 
- investor
- director

**Permissions** :

actor:

- act : bool
- readScript : bool
- cryOnCommand : bool

investor:

- scandal : "allow"/"deny"
- modifyBudget int

director:

- act : bool
- invest : int
- direct : bool

---

Of course, this is not by a long shot a complete representation of any particular field, but it suffices in order to have an idea of how the services are structured and who does what.

Another important point is that the permission values are polymorphic so that they can easily be changed in the `datastore.json`, if need be. I've added some string keys and numbers to show that. Most of the permissions that I've used pertain to using a boolean representation anyway.


The data is ingested from the `./db/datastore.json` and it has the following structure:

`{
"user2143": {
"name": "Roger Deakins",
"permissions": {
"blockbusters.cinematographer.changeLens": true }
`   

Thus, each `userId` has a respective `name` and `permissions` collection. 

---


*Regarding design decisions*

I've made a struct called `Application` that embeds the storage, logging and error handling - In this way I have a centralized entity to hold the important things. It makes use of composition,
therefore it has embedded other entities in itself and can use their internal structure/methods.

The database is using the `DataPuller` interface which shows the behaviour of how we extract the data. We also have the `Storage` struct that implements it. That struct also holds the structure for the service.permission.feature model and 
knows if the model is correct. That is, if the permission is part of the feature and that the feature is part of the service. The basis of this is achieved through the use of maps for quick lookups.


*Why use JSON for the data storage?*

The `.json` format is advantageous since we have a quick lookup time on every request(avg. O(1) because it is unmarshalled into a map). It also uses the `embed` capability of Go.
Another aspect of handling the data ingestion part this way is that it also pertains to unmarshalling the data in an easy manner into native go data structures.

We have the option of reading directly from file, if we so choose. We just need to mention which method we prefer. Reading from a file makes mocking things a bit easier as well.

*Why log things this way?*

The structure of the project is clearer if we have a centralized logging system. This helps us
with relevant information (i.e starting the server and seeing which address we are using) or error handling in a centralized, coherent fashion.

*Why use an interface for the data pulling phase?* 

It helps us maintain orthogonality. It separates the actual way we store things (json,db,etc) from our
business logic and therefore, we have a middle layer which enables us to plug any other way of storing
the user permissions with minimal changes in our code, if need be. It's also helpful for our test - mocking things and such.

*Why use empty struct when checking the features/permissions/routes?*

It allocates 0 memory. Also, by using a map we are leveraging the internal map algorithm that already has an efficient lookup time. Therefore, we can quickly validate the requests, etc. 

*Regarding the use of maps internally*

As I mentioned previously, I made heavy use of the map data structure. The validation for the service.feature.permission is made this way, fetching the users also is based on this, searching through permissions, etc.
Consequently, the data can grow and the queries can still be fairly efficient.

*Regarding building the app*

I would say that the easiest way is to just build the binaries on your native os.

I did do a multi-stage docker build (12MB image) with the Alpine distro. You can also choose to export the port and test it this way. This is also useful for further development: say you want to deploy the app or share the image with the team so they can all test it real fast. 

Regarding Docker, with a published image you could push the updates there and anyone with a machine that has Docker installed can test it. Regardless of the libraries/language/etc that the app uses.



---

### Build:

This build assumes that you have at least Go 1.17 installed on your machine. If you do not: https://go.dev/doc/install

From the current directory run: `go build ./cmd`

---
This build assumes that you have Docker installed on your machine. If you do not: https://docs.docker.com/get-docker/

Another option would be to use the Dockerfile to run the app:

`docker build -t moviesplatform:alpine .`


### Run:

From the current directory run: `go run ./cmd`

*Optional*: set the port that the service will run by using the `-addr` flag. eg: `go run ./cmd addr=":8080"`

---

After building the Dockerfile, just run the container with `docker run -p <HOST_PORT>:<CONTAINER_PORT> moviesplatform:alpine`. By default, the server uses the `4000` port.

### Use:
e.g : `curl -X GET http://localhost:4000/v1/user/user4323` to get Denis Villeneuve.

The following endpoints are available:

- `GET /v1/user/:user` returns collection of permissions as a JSON encoded map.

- `GET /v1/user/:user/:service` returns collection of permissions for the named service as a JSON encoded map.

- `GET /v1/service/:service/:feature/:permission` returns a list of users with values for this permission.
---

*Things to improve* :

- Naming - I'm not fully content with the naming. Adjustments can always be made on this part, after all..
- Application struct is in `api/web/server`.
- `errors.go` and `log.go` could be moved somewhere else.


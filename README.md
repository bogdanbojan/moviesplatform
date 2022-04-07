## Movie platform 


---

*Readme*: 


TODO: Add more users with clearer permissions.

*App*:

TODO: Tests.


---


### Mental model and structuring of the services.

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
- producer
- editor

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

- changeBudget bool
- changeSalary bool
- addActor bool

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
*Why use JSON for the data storage?*

The `.json` format is advantageous since we have a quick lookup time on every request(avg. O(1) because it is unmarshalled into a map). It also uses the `embed` capability of Go.
Another aspect of handling the data ingestion part this way is that it also pertains to unmarshalling the data in an easy manner into native go data structures.


*Why log things this way?*

The structure of the project is clearer if we have a centralized logging system. This helps us
with relevant information (i.e starting the server and seeing which address we are using) or error handling in a centralized, coherent fashion.

*Why use an interface for the data pulling phase?* 

It helps us maintain orthogonality. It separates the actual way we store things (json,db,etc) from our
business logic and therefore, we have a middle layer which enables us to plug any other way of storing
the user permissions with minimal changes in our code, if need be. It's also helpful for our test - mocking things and such.

*Why use empty struct when checking the features/permissions/routes?*

It allocates 0 memory. Also, by using a map we are leveraging the internal map algorithm that already has an efficient lookup time. Therefore, we can quickly validate the requests, etc. 

*Regarding design decisions*

I've made a struct called `Application` that embeds the storage, logging and error handling - In this way I have a centralized entity to hold the important things. It makes good use of composition.

*Regarding building the app*

The easiest way is to just build it on your native os right now.

I did do a multi-stage docker build (12MB image) with the Alpine distro and it works - it's just that I need to figure out how to test it in that container.


---

### Build:

This build assumes that you have at least Go 1.17 installed on your machine. If you do not: https://go.dev/doc/install

From the current directory run: `go build ./cmd`

---

Another option would be to use the Dockerfile to run the app:

`docker build -t moviesplatform:alpine .`

After that, just run the container with `docker run moviesplatform:alpine`


### Run:

From the current directory run: `go run ./cmd`

*Optional*: set the port that the service will run by using the `-addr` flag. eg: `go run ./cmd addr=":8080"`

### Use:
e.g : `curl -X GET http://localhost:4000/user/user4323` to get Denis Villeneuve.

The following endpoints are available:

- `GET /v1/user/:user`, required returns collection of permissions as a JSON encoded map.

- `GET /v1/user/:user/:service`, optional returns collection of permissions for the named service as a JSON encoded map.

- `GET /v1/service/:service/:feature/:permission`, optional returns a list of users with values for this permission.
---

*Things to improve* :

- Naming - I'm not fully content with the naming. A lot of adjustments can be made on this part but I'll focus on tests right now.
- Application struct is in `api/web/server`. Which is a bit weird - I had some problems with importing it into other packages and that's why it's there atm.
- `errors.go` and `log.go` could be moved somewhere else.
- Project structure maybe
- Better error handling in the storage part of the app.

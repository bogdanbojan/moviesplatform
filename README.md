## Movie platform management


---

*Readme*: 

TODO: Add services, features, permissions.

TODO: Explain the decisions throughout the project.

*App*:

TODO: Tests.

TODO: Containerize the app if I have time.


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

- direct
- instructActors
- argue

cinematographer:

- shoot
- changeLens
- changeCamera

producer:

- changeBudget
- changeSalary
- addActor

### Commercials

**Features** :

- artist
- producer
- manager

**Permissions** :

artist:

- createConcept
- creativitySwitch

producer:

- getDeals
- onboardPeople

manager:

- adviseBoard
- cancelMeetings
- scheduleMeetings

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

- scandal
- modifyBudget

director:

- act
- invest
- direct

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

*Why use an interface for the data pulling phase?* !TODO

It helps us maintain orthogonality. It separates the actual way we store things (json,db,etc) from our
business logic and therefore, we have a middle layer which enables us to plug any other way of storing
the user permissions with minimal changes in our code, if need be. It's also helpful for our test - mocking things and such.

*Why use empty struct when checking the features/permissions/routes?*

It allocates 0 memory. Also, by using a map we are leveraging the internal map algorithm that already has an efficient lookup time. Therefore, we can quickly validate the requests, etc. 



---

### Build:

This build assumes that you have at least Go 1.17 installed on your machine. If you do not: https://go.dev/doc/install

`go build .`


[//]: # (?just binary or the user should have go preinstalled?)

### Run:

`go run main.go`

*Optional*: set the address that the service will run by using the `-addr` flag. eg: `go run . -addr=":8080"`

### Use:
e.g : `curl -X GET http://localhost:4000/user/user4323` to get Denis Villeneuve.




---


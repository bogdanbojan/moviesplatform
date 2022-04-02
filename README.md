## Movie platform management

todo: Maybe add chapters that are linked to a certain part of the docs.

---
The service is a *movie platform* in which you can review you favorite movie. If you are a movie director you can even post your own masterpiece or if you
are a regular user you can rate your blockbuster of choice.

**Service**:

- movie platform 

[//]: # ()
[//]: # (**Feature** : )

[//]: # ()
[//]: # (- blockbusters)

[//]: # (- indie)

[//]: # (- commercials)

[//]: # (- shorts)


**Roles**: ? Feature

- directors
- producers
- users
- guests
  
[//]: # (- actors)

**Permissions**:

- create
- modify
- comment
- rate
- none

[//]: # (- review)

---


### *roles*:

**directors** : can *create*, *modify*, *comment* and *rate* movies. 

**producers** : can *modify*, *comment*, *rate* movies.

**users** : can *comment* and *rate* movies.

**guests** : *none*. They can view the platform.


### *permissions*:

**create** : Enables you to create a movie on this platform. 

**modify** : You are able to modify the movies details. This can include changes in budget, cast or even the movie title.

**comment** : Comment and say what you like/don't like on that particular flick.

**rate** : Assign a rating from 1-10 to your favorite movie.  

**none** : See the ratings, comments and other features. Cannot provide input on anything, though.

---
---

### Build:

This build assumes that you have at least Go 1.17 installed on your machine. If you do not: https://go.dev/doc/install




[//]: # (?just binary or the user should have go preinstalled?)

### Run:

`go run main.go`

*Optional*: set the address that the service will run by using the `-addr` flag

### Use:
e.g : `curl -X GET http://localhost:4000/user/davidlynch`




---

The permissions are stored in a json which is embedded into the code. The `store.json` file is found in the project directory
and it is populated when the service is started locally.


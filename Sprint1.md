Sprint 1

User stories:

1. Tourists can search for places on the search place page. Once any place name is entered, a list of tourist spots within a 1000 km radius of that place will be displayed.

2. Users(both tourists and travel guides) can register to the application using email , password and their role(role can be tourist or travel guide). Passwords have to be a minimum of 8 characters and can be at most 20 characters. 

3. Users (both tourists and travel guides) can login to the application using email, password and their role(role can be tourist or travel guide).

Frontend:

1. The login and register pages have been designed along with the homepage where the user can type a location.

2. As soon as the user presses the ‘Search’ button, the list of places is fetched from the backend and displayed.

Backend:

1. The login and register feature has been implemented in the backend. APIs has been exposed using which register and login can be done.

2. Tourist places near a given location are retrieved using a third party api and json response is sent back to the frontend.


Database:
We have used the SQLLite3 database and GORM library for CRUD operations.

1. ‘Register’ → stores the login mail and password of the users while signing up and checks while user login.
 Fields: Email(Primary Key),Password,Role.

2. ‘UserProfile’--> stores tourists profile details.
Fields: Email(Primary Key),Name,About,Age,Fav1,Fav2,Fav3

3. ’GuideProfile’-->store guides profile details.
Fields: Email(Primary Key),name, about,Age, Address, Location, Vehicle

4. ’Comment’→stores  user comments
Fields: Comments,email,location,Register( Foreign key).

Web Framework:

gin -gonic





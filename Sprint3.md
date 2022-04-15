Sprint 3 -

User Stories :
Tourists and Guides can now view and edit their profiles.
Travel Guide can add packages corresponding to a particular location
Travel guide can view the packages added by him.
User can view packages corresponding to a particular location
Users can check travel packages in an area.
Few other bugs like CORS error and logout token deletion are also taken care of.

Video link - https://drive.google.com/drive/folders/1KBgnRkbSJnNvbFj6pbzOePTbTq08tdNg?usp=sharing

Testing has been done in the Frontend. Various features like login, register, edit, view profile have been tested. 
Also various features in the backend have also been tested.

Steps to run backend:
Run “go run main.go” in the server folder.
Steps to run Test:
Run go test

Steps to run Frontend:
“ng serve” in the Frontend folder
Steps to run Tests:
“Npx cypress open”


Functionality:
Users can either be a Tourist or a Travel Guide.
In the case of a travel guide, he/she can view and edit his profile. If a guide has added packages, they will be shown in his profile. This functionality will not be present for a tourist.
Users can search for tourist spots in places and view packages corresponding to that place.
A travel guide can add a package for a place in addition to viewing them.
A package is a travel plan that a guide gives to a tourist and helps the tourist visit places in exchange for money.


API Docs-

Signup: http://localhost:8080/signup
  Inputs - {name, email, password, role}
  Output - none

Login: http://localhost:8080/login
  Inputs - {email, password, role}
  Output - {Authorization Token}

View User Profile: http://localhost:8080/userprofile
  Inputs - none
  Outputs - {name, email, age, mobile, location, first favorite place, second favorite place, third favorite place}

Search Packages: http://localhost:8080/searchPackage
  Inputs - none
  Outputs - {list of packages corresponding to a guide - {guide email, accommodation, included, duration, itinerary, price}}

Search Places: http://localhost:8080/searchPlaces/<place name>
  Inputs - none
  Outputs - {list of tourist places corresponding to the input place - {place.name, place.kinds, p.preview, p.wikipedia_extracts.text,    p.wikipedia_extracts.title}}

Update Tourist Profile: http://localhost:8080/updateUserProfile
  Inputs - {name, about, age, mobile, location, first favorite place, second favorite place, third favorite place}
  Outputs - none

Update Guide Profile: http://localhost:8080/updateGuideProfile
  Inputs - {name, about, age, address, location, vehicle}
  Outputs - none

Add Package: http://localhost:8080/addPackages
  Inputs - {email, location, included, duration, itinerary, accommodation, price}
  Outputs - none






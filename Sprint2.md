Sprint 2

User stories:
1. Redesigned and developed login and register functionality using JWT token authentication.
2. Travel guides can add tour pacakages
3. Users can add/view and edit details on their profile.

Sprint2 video drive link added: https://drive.google.com/file/d/1iR2mwzMeOfCW2NH2g255VCTOeuuYeHlx/view?usp=sharing

Frontend Testing:
Manual tests are slow, labor-intensive and cannot be repeated often. They are unspecific from a developer perspective: If the test fails, we cannot easily pin down which part of the application is responsible or which code change causes the regression.
We need automated tests that take the user’s perspective. This is what end-to-end (E2E) tests do. We use cypress framework for end to end testing.
Frameworks for end-to-end tests allow navigating to URLs, simulating user input and inspecting the page content. Apart from that, they have little in common. The test syntax and the way the tests are run differ widely. On such framework is Cypress.

Once we install cypress, in the project directory, you will find a sub-directory called cypress. It contains:
* A tsconfig.json configuration for all TypeScript files in the directory,
* an integration directory for the end-to-end tests,
* a support directory for custom commands,
* a plugin directory for Cypress plugins,
* a fixtures directory for test data.


In our first Cypress test, we have checked the we login successfully. The test needs to perform the following steps:
1. Navigate to “/login”.
2. Check if the url contains ‘login’.
3. Find the name element of the login form and automate typing in some name.
4. Similarly, find the password element and automate giving the password.
5. Find the button element and click on it.
6. After login, check if it is navigate to home page.
7. Check if the url contains ‘home-page’
We have used cy.visit(‘/login’) to navigate to an login page. The path “/” translates to http://localhost:4200/login since this is the configured baseUrl.

Backend:

Travel guides can add tour packages which will have the following details:
GuideEmail   
Duration     
Location     
Accomodation
Itinerary    
Included
    
A new table called Package is created with the above details. The email id which they provide on the tour package needs to exist in our system. Also,
the travel guide needs to be registered in our system as one, meaning, a person registered as tourist cannot add a tour package. If a person wants to d
act as both tourist and travel guide, they have to register with each of these roles. These cases have been handled in the code and returns errors.

Unit Testing for SearchPlaces api (implemented in Sprint 1) has been done. Following cases has been handled:
1. When search string=""
2. When search string doesnot match any valid place name
3. When search string matches a valid place name
4. When search string is similar to a valid place name but not exactly same.
5. When search string= " "

jwt tokenization backend:
jwt tokenization is done using the user given mail and user role fields. there are two types of tokens: 
access token and refresh token. Access token is checked whenever user tries to access the authenticated rotes.
refresh token is used to refresh the token for user covinience, as the access token will expire in shorter time. 
The metadata will be stored in redis middleware. whenever a user logout, the token will be removed from the middleware .


CRUD operations for the User profile,Guide Profile, Register and Comments table.

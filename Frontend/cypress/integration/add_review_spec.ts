describe('Tourist adds review to the booked packages', () => {
    it('Go to the booked packages page of a tourist and add a review to a package, the review should be successfully added', () => {
        cy.visit('/login');
        cy.url().should('include','login');
        cy.get('[name="email"]').type('xia@gmail.com');
        cy.wait(2000)
        cy.get('[name="password"]').type('password123');
        cy.wait(2000)
        cy.get('#mat-select-0').click().get('mat-option').contains('Tourist').click();
        cy.get('button').click();
        cy.wait(2000)
        cy.url().should('include', 'home-page');
        cy.get('[name="addreview"]').click();
        cy.url().should('include', '/list-booked-packages');
        cy.visit('/list-booked-packages/add-review');
        cy.wait(2000)
        cy.get('[name="title"]').type('Kolkata Place review');
        cy.wait(2000)
        cy.get('[name="review"]').type('It was my first visit to florida. And thanks to GlobeScanner for providing such amazing package with all the facilities needed.');
        cy.wait(2000)
        cy.get('#mat-checkbox-1').find('input').click({force:true});  
        cy.wait(2000)  
        cy.get('button').click();
        cy.url().should('not.include', '/home-page');
      })
    });
describe('Tourist adds review to the booked pacakages', () => {
    it('Go to the booked packages of a tourist, Fill the review page , click on submit and the review should be added', () => {
        cy.visit('/login');
        cy.url().should('include','login');
        cy.get('[name="email"]').type('xia@gmail.com');
        cy.wait(200)
        cy.get('[name="password"]').type('password123');
        cy.wait(200)
        cy.get('#mat-select-0').click().get('mat-option').contains('Tourist').click();
        cy.get('button').click();
        cy.wait(200)
        cy.url().should('include', 'home-page');
        cy.get('[name="addreview"]').click();
        cy.url().should('include', '/list-booked-packages');
        cy.wait(200)
        cy.get('button').click();
        cy.url().should('include', '/list-booked-packages/add-review');
        cy.wait(200)
        cy.get('[name="title"]').type('Kolkata Place review');
        cy.wait(200)
        cy.get('[name="review"]').type('It was my first visit to florida. And thanks to GlobeScanner for providing such amazing package with all the facilities needed.');
        cy.wait(200)
        cy.get('button').click();
        cy.url().should('not.include', '/home-page');
      })
    });
describe('Tourist can book a package and see the list of booked packages under his profile and add review to the booked packages', () => {
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
      })
    });
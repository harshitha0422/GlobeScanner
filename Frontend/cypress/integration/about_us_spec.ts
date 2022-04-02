describe('about-us page', () => {
    it('successfully loads', () => {
        cy.visit('/home-page');
        cy.contains('About Us');
        cy.contains('Connect With Us');
    })
    });
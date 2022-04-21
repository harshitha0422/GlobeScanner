describe('Travel Guide adds package to a place', () => {
it('Fill the package info page , if the post doesnot work then the package should not be added', () => {
    cy.visit('/place-list/add-package')
    cy.wait(20)
    cy.get('[name="location"]').type('Florida');
    cy.wait(20)
    cy.get('[name="included"]').type('Meals');
    cy.wait(20)
    cy.get('[name="duration"]').type('15');
    cy.wait(20)
    cy.get('[name="itinerary"]').type('First: Delhi');
    cy.wait(20)
    cy.get('[name="accomodation"]').type('Taj Mahal');
    cy.wait(20)
    cy.get('[name="price"]').type('1000$');
    cy.get('button').click();
    cy.request({
      url: 'http://localhost:8080/addPackages', 
      failOnStatusCode: false 
    }).then((response) => {
      expect(response.status).to.eq(500);
    })
    cy.url().should('not.include', '/home-page');
  })
});
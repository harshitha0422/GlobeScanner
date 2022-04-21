describe('See the tourist places in an area', () => {
  it('Type a place and click on searach, we will be able to see all the tourist places in that area', () => {
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
    cy.wait(2000)
    cy.get('[name="something"]').type('gainesville');
    cy.wait(2000)
    cy.get('[name="searchButton"]').click();
    cy.wait(200)
    cy.url().should('include', '/place-list');
  })
  });
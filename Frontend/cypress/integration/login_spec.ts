describe('login', () => {
  it('Should login if the credentials are correct', () => {
    cy.visit('/login')
    cy.url().should('include','login');
    cy.get('[name="email"]').type('megan12@gmail.com');
    cy.wait(2000)
    cy.get('[name="password"]').type('password123');
    cy.wait(2000)
    cy.get('#mat-select-0').click().get('mat-option').contains('Guide').click();
    cy.get('button').click();
    cy.wait(2000)
    cy.url().should('include', 'home-page');

  });

  it('Should not login if the password is less than 6 characters', () => {
    cy.visit('/login')
    cy.url().should('include','login');
    cy.get('[name="email"]').type('xia@gmail.com');
    cy.wait(2000)
    cy.get('[name="password"]').type('password123');
    cy.wait(2000)
    cy.get('#mat-select-0').click().get('mat-option').contains('Guide').click();
    cy.get('button').click();
    cy.wait(2000)
    cy.url().should('not.include', 'home-page');
     cy.wait(2000)
  });

  
});
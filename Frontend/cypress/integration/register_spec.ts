describe('register', () => {
  it('Should register and navigate to home page', () => {
    cy.visit('/register')
    cy.url().should('include','register');
    cy.get('[name="username"]').type('guide9');
    cy.wait(2000)
    cy.get('[name="email"]').type('megan12@gmail.com');
    cy.wait(2000)
    cy.get('[name="password"]').type('password123');
    cy.wait(2000)
    cy.get('#mat-select-0').click().get('mat-option').contains('Guide').click();
    cy.get('button').click();
    cy.wait(2000)
    cy.url().should('include', 'register');
  });
});

describe('register', () => {
  it('Should register and navigate to home page', () => {
    cy.visit('/register')
    cy.url().should('include','register');
    cy.get('[name="email"]').type('saduvishesha@gmail.com');
    cy.get('button').click();
    cy.url().should('include', 'home-page');
  });
});

describe('login', () => {
  it('Should not login if the form is invalid', () => {
    cy.visit('/login')
    cy.url().should('include','login');
    cy.get('[name="email"]').type('saduvishesha@gmail.com');
    cy.get('button').click();
    cy.url().should('not.include', 'home-page');

  });

  it('Should login if the form is valid', () => {
   cy.login('saduvishesha@gmail.com','12345');
    cy.url().should('include', 'home-page');
  });
});

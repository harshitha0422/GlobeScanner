it('Should navigate to view packages page', () => {
    cy.visit('/place-list');
    cy.wait(2000)
    cy.get('#viewPackages').click();
    cy.wait(2000)
    cy.url().should('include', '/place-list/view-packages');
    cy.wait(2000)
  });
 
  

  it('Shoulds get the packages in a place' , () => {
    cy.visit('/place-list');
    cy.wait(2000);
    cy.get('#viewPackages').click();
    cy.wait(2000);
    cy.intercept('GET', '/searchPackage/"Florida"');
   
  });
describe('homepage', () => {
    it('Should register and navigate to home page', () => {
      cy.visit('/home-page')
      cy.url().should('include','home-page');
      
    });
  
   
  });
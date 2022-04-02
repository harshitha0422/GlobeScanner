describe('Tourist and Trvele guide can edit their profile', () => {

    it('Should login if the credentials are correct', () => {
        cy.visit('/login')
        cy.url().should('include','login');
        cy.get('[name="email"]').type('xia@gmail.com');
        cy.wait(2000)
        cy.get('[name="password"]').type('password123');
        cy.wait(2000)
        cy.get('#mat-select-0').click().get('mat-option').contains('Tourist').click();
        cy.get('button').click();
        cy.wait(2000)
        cy.url().should('include', 'home-page');
    
      });

    it('Fill the edit profile page , if the post doesnot work then the profile should not be editable', () => {
        cy.visit('/view-profile/edit-profile')
        cy.wait(20)
        cy.wait(20)
        cy.get('[name="age"]').type('12');
        cy.wait(20)
        cy.get('[name="mobile"]').type('+91 12345677889');
        cy.wait(20)
        cy.get('[name="location"]').clear();
        cy.get('[name="location"]').type('Florida');
        cy.wait(20)
        cy.get('[name="fav1"]').clear();
        cy.get('[name="fav1"]').type('Tampa');
        cy.wait(20)
        cy.get('[name="fav2"]').clear();
        cy.get('[name="fav2"]').type('Orlando');
        cy.wait(20)
        cy.get('[name="fav3"]').clear();
        cy.get('[name="fav3"]').type('Miami');
        cy.get('button').click();
        cy.url().should('include', '/view-profile');
      })
    });
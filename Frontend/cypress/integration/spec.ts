describe('My First Test', () => {
  it('Visits the initial project page', () => {
    cy.visit('/')
    cy.contains('Discover the World')
    cy.contains('sandbox app is running!')
  })
})

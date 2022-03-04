import { ComponentFixture, TestBed } from '@angular/core/testing';

import { HomePageComponent } from './home-page.component';

describe('HomePageComponent', () => {
  let fixture: HomePageComponent;
  let plyFetchMock: any;
  
  beforeEach(async () => {
    fixture = new HomePageComponent(plyFetchMock); 
    
  });

  

  it('should create', () => {
    expect(fixture).toBeTruthy();
  });
});

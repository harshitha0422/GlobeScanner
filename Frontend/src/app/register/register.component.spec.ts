import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegisterComponent } from './register.component';

describe('RegisterComponent', () => {
  let fixture: RegisterComponent;
  let authServiceMock: any;

  beforeEach(async () => {
     fixture = new RegisterComponent(authServiceMock);
    
  });

  it('should create', () => {
    expect(fixture).toBeTruthy();
  });
});

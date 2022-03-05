import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LogInComponent } from './log-in.component';

describe('LogInComponent', () => {
  let fixture: LogInComponent;
  let authServiceMock: any;
  let tokenStorageMock: any;
  beforeEach(async () => {
   fixture = new LogInComponent(
    authServiceMock,
    tokenStorageMock
   );
    
  });

  
  
  it('should create', () => {
    expect(fixture).toBeTruthy();
  });
});

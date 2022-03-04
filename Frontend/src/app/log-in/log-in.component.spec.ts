import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LogInComponent } from './log-in.component';

describe('LogInComponent', () => {
  let fixture: LogInComponent;
 let authServiceMock;
  beforeEach(async () => {
   fixture = new LogInComponent();
    
  });

  
  
  it('should create', () => {
    expect(fixture).toBeTruthy();
  });
});

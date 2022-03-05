import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NavBarComponent } from './nav-bar.component';

describe('NavBarComponent', () => {
  let fixture: NavBarComponent;
  let tokenStorageServiceMock: any;
  beforeEach(async () => {
   fixture = new NavBarComponent(tokenStorageServiceMock);

  });

  it('should create', () => {
    expect(fixture).toBeTruthy();
  });
});

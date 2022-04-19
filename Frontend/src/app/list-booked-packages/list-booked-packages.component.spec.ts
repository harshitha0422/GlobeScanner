import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ListBookedPackagesComponent } from './list-booked-packages.component';

describe('ListBookedPackagesComponent', () => {
  let component: ListBookedPackagesComponent;
  let fixture: ComponentFixture<ListBookedPackagesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ListBookedPackagesComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ListBookedPackagesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

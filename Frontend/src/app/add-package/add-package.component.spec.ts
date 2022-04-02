import { ComponentFixture, getTestBed, TestBed } from '@angular/core/testing';
import { AddPackageComponent } from './add-package.component';
import { AddPackageService} from '../add-package.service';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { Router, RouterModule } from '@angular/router';
import { HttpClient, HttpClientModule, HttpHandler } from '@angular/common/http';
import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';


describe('AddPackageComponent', () => {
  let component: AddPackageComponent;
  let fixture: ComponentFixture<AddPackageComponent>;
  let injector: TestBed;
  let service: AddPackageService;
  let httpMock: HttpTestingController;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddPackageComponent ],
      schemas: [
        CUSTOM_ELEMENTS_SCHEMA
      ],
      imports: [FormsModule,
        ReactiveFormsModule, HttpClientTestingModule],
        providers: [AddPackageService, HttpClient, HttpClientModule, HttpHandler]
    })
    .compileComponents();
    injector = getTestBed();
    service = injector.get(AddPackageService);
    httpMock = injector.get(HttpClientTestingModule);
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AddPackageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('add package to a place', () => {
    const packageToBeAdded = [
      { 
        guidEmail: "",
        location: "",
        included : "",
        duration : "",
        itinerary : "",
        accomodation : "",
         price: ""
       
      },   
    ];
    expect(1).toEqual(1);
    service.addNewPackage(packageToBeAdded).subscribe( (data:any) => {
      expect(data.body.length).toBe(2);
      expect(data.body).toEqual(packageToBeAdded);
    });

  });

});

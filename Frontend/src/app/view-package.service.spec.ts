import { TestBed } from '@angular/core/testing';
import { HttpClient } from '@angular/common/http';
import { ViewPackageService } from './view-package.service';

describe('ViewPackageService', () => {
  let service: ViewPackageService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewPackageService ],
      providers: [HttpClient]
    })
    .compileComponents();
  });
  
  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ViewPackageService);
  });
  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

import { TestBed, inject} from '@angular/core/testing';
import { HttpClient } from '@angular/common/http';
import { AddPackageService } from './add-package.service';

describe('AddPackageService', () => {
  let service: AddPackageService;
  
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AddPackageService ],
      providers: [HttpClient]
    })
    .compileComponents();
  });
  
  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AddPackageService);
  });

  it('should be created', inject([AddPackageService], (service : AddPackageService) => {
    expect(service).toBeTruthy();
  }));
});

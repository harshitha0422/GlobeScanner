import { TestBed, inject} from '@angular/core/testing';
import { HttpClient } from '@angular/common/http';
import { GetUserDataService } from './get-user-data.service';

describe('GetUserDataService', () => {
  let service: GetUserDataService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GetUserDataService ],
      providers: [HttpClient]
    })
    .compileComponents();
  });
  
  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GetUserDataService);
  });

  it('should be created', inject([GetUserDataService], (service : GetUserDataService) => {
    expect(service).toBeTruthy();
  }));
});


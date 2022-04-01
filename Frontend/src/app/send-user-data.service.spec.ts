import { TestBed } from '@angular/core/testing';
import { HttpClient } from '@angular/common/http';
import { SendUserDataService } from './send-user-data.service';

describe('SendUserDataService', () => {
  let service: SendUserDataService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SendUserDataService ],
      providers: [HttpClient]
    })
    .compileComponents();
  });
  
  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SendUserDataService);
  });
  
  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

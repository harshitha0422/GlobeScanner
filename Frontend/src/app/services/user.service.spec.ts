import { TestBed , inject} from '@angular/core/testing';
import { HttpClient } from '@angular/common/http';
import { UserService } from './user.service';

describe('UserService', () => {
  let service: UserService;
  
  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UserService ],
      providers: [HttpClient]
    })
    .compileComponents();
  });
  
  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(UserService);
  });

  it('should be created', inject([UserService], (service : UserService) => {
    expect(service).toBeTruthy();
  }));
  
});

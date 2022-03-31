import { TestBed , inject} from '@angular/core/testing';
import { HttpClient } from '@angular/common/http';
import { TokenStorageService } from './token-storage.service';

describe('TokenStorageService', () => {
  let service: TokenStorageService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TokenStorageService ],
      providers: [HttpClient]
    })
    .compileComponents();
  });
  
  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TokenStorageService);
  });

  it('should be created', inject([TokenStorageService], (service : TokenStorageService) => {
    expect(service).toBeTruthy();
  }));
});

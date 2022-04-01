import { TestBed } from '@angular/core/testing';
import { HttpClient } from '@angular/common/http';
import { PlaceFetchService } from './place-fetch.service';

describe('PlaceFetchService', () => {
  let service: PlaceFetchService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PlaceFetchService ],
      providers: [HttpClient]
    })
    .compileComponents();
  });
  
  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PlaceFetchService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

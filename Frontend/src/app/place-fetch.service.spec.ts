import { TestBed } from '@angular/core/testing';

import { PlaceFetchService } from './place-fetch.service';

describe('PlaceFetchService', () => {
  let service: PlaceFetchService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PlaceFetchService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

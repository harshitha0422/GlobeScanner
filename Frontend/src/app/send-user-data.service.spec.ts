import { TestBed } from '@angular/core/testing';

import { SendUserDataService } from './send-user-data.service';

describe('SendUserDataService', () => {
  let service: SendUserDataService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(SendUserDataService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

import { TestBed } from '@angular/core/testing';

import { AddreviewService } from './addreview.service';

describe('AddreviewService', () => {
  let service: AddreviewService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AddreviewService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

import { TestBed } from '@angular/core/testing';

import { BookPackageService } from './book-package.service';

describe('BookPackageService', () => {
  let service: BookPackageService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(BookPackageService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

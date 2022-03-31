import { TestBed } from '@angular/core/testing';

import { AddPackageService } from './add-package.service';

describe('AddPackageService', () => {
  let service: AddPackageService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(AddPackageService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});

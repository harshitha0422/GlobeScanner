import { Injectable } from '@angular/core';
import {HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AddPackageService {

  constructor(private http : HttpClient) { }
  addNewPackage(addPackage: any){
    console.log("This is the final json before update req",addPackage);
    return this.http.post<any>("http://localhost:8080/addPackages", addPackage);
  }
}

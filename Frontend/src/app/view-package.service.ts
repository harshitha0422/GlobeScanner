import { Injectable } from '@angular/core';
import {HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ViewPackageService {
  viewPlacePackages : any;
  constructor(private http : HttpClient) { }

  viewPackages(): Observable<any>{
    console.log("inside view pacakges");
    console.log(this.viewPlacePackages);
    console.log("http://localhost:8080/searchPackage/"+String(this.viewPlacePackages));
    return this.http.get<any>("http://localhost:8080/searchPackage/"+String(this.viewPlacePackages));
    // return this.http.get<any>(API_URL+String(this.toSearch));
  }
}

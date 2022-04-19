import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' ,
  'Access-Control-Allow-Origin': '*',
  'Access-Control-Allow-Headers': 'Content-Type',
  'Access-Control-Allow-Methods': 'GET,POST,OPTIONS,DELETE,PUT'
})
};

@Injectable({
  providedIn: 'root'
})
export class GetUserDataService {

  constructor(private http : HttpClient) { }
  getTouristInfo(): Observable<any>{
    // return this.http.get<any>("http://10.192.167.246:8080/userprofile",httpOptions);
    // return this.http.get<any>("http://10.20.158.45:8080/userprofile",httpOptions);
    // return this.http.get<any>("http://10.20.106.116:8080/userprofile",httpOptions);
    return this.http.get<any>("http://localhost:8080/userprofile",httpOptions);
  }
  getGuideInfo(): Observable<any>{
    // return this.http.get<any>("http://10.192.167.246:8080/userprofile",httpOptions);
    // return this.http.get<any>("http://10.20.158.45:8080/userprofile",httpOptions);
    // return this.http.get<any>("http://10.20.106.116:8080/userprofile",httpOptions);
    return this.http.get<any>("http://localhost:8080/userprofile",httpOptions);
  }
  getGuidePackage(): Observable<any>{
    // return this.http.get<any>("http://10.192.167.246:8080/userprofile",httpOptions);
    // return this.http.get<any>("http://10.20.158.45:8080/userprofile",httpOptions);
    // return this.http.get<any>("http://10.20.106.116:8080/userprofile",httpOptions);
    return this.http.get<any>("http://localhost:8080/searchPackage",httpOptions);
  }

  getTouristBookedPackage(): Observable<any>{
    // return this.http.get<any>("http://10.192.167.246:8080/userprofile",httpOptions);
    // return this.http.get<any>("http://10.20.158.45:8080/userprofile",httpOptions);
    // return this.http.get<any>("http://10.20.106.116:8080/userprofile",httpOptions);
    return this.http.get<any>("http://localhost:8080/searchPackage",httpOptions); //Needs to be changed when API provided by backend
  }
}

import { Injectable } from '@angular/core';
import {HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class GetUserDataService {

  constructor(private http : HttpClient) { }
  getUserInfo(): Observable<any>{
    return this.http.get<any>("http://localhost:8080/userprofile/saduvishesha@gmail.com");
  }
}

import { Injectable } from '@angular/core';
import {HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class PlaceFetchService {
  toSearch : any;
  constructor(private http : HttpClient) { }

  getPlaces(): Observable<any>{
    return this.http.get<any>("http://localhost:8080/searchPlaces/"+String(this.toSearch));
  }
}

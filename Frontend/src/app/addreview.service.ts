import { Injectable } from '@angular/core';
import {HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AddreviewService {

  constructor(private http : HttpClient) { }
  addReview(review: any){
    console.log("This is the final json before update req",review);
    return this.http.post<any>("http://localhost:8080/comments", review);
  }
}

import { Injectable } from '@angular/core';
import {HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SendUserDataService {

  constructor(private http : HttpClient) { }
  sendUserData(editUser: any){
    return this.http.post("http://localhost:8080/userprofile", editUser);
    
  // return this.http.post('http://localhost:8080/userprofile', editUser).subscribe(
  //     (response) => console.log(response),
  //     (error) => console.log(error)
  //   )
}
}

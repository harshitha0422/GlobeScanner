import { Injectable } from '@angular/core';
import {HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SendUserDataService {

  constructor(private http : HttpClient) { }
  sendUserData(editUser: any){
    console.log("This is the final json before update req",editUser);
    return this.http.put<any>("http://10.192.167.246:8080/updateUserProfile", editUser);
    // return this.http.put<any>("http://10.20.158.45:8080/updateUserProfile", editUser);
    
    
  // return this.http.post('http://localhost:8080/userprofile', editUser).subscribe(
  //     (response) => console.log(response),
  //     (error) => console.log(error)
  //   )
}
}

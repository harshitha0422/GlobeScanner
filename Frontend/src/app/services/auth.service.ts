// import { Injectable } from '@angular/core';

// @Injectable({
//   providedIn: 'root'
// })
// export class AuthService {

//   constructor() { }
// }
import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
const AUTH_API = 'http://localhost:8080/';
// const AUTH_API = 'http://10.20.106.116:8080/';
// const AUTH_API = 'http://10.20.158.45:8080/';
// const AUTH_API = 'http://10.192.167.246:8080/';
const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};
@Injectable({
  providedIn: 'root'
})
export class AuthService {
  constructor(private http: HttpClient) { }
  login(email: string, password: string, role:string): Observable<any> {
    console.log(email,password,role);
    return this.http.post(AUTH_API + 'login', {
      email,
      password,
      role
    }, httpOptions);
  }
  register(name: string, email: string, password: string, role: string): Observable<any> {
    console.log(name,email,password,role);
    return this.http.post(AUTH_API + 'signup', {
      name,
      email,
      password,
      role
    }, httpOptions);
  }
}
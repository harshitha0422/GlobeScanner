// import { Injectable } from '@angular/core';

// @Injectable({
//   providedIn: 'root'
// })
// export class TokenStorageService {

//   constructor() { }
// }
import jwt_decode from 'jwt-decode';

import { Injectable } from '@angular/core';
const TOKEN_KEY = 'auth-token';
const USER_KEY = 'auth-user';
@Injectable({
  providedIn: 'root'
})
export class TokenStorageService {
  constructor() { }
  signOut(): void {
    window.sessionStorage.clear();
  }
  public saveToken(token: string): void {
    window.sessionStorage.removeItem(TOKEN_KEY);
    window.sessionStorage.setItem(TOKEN_KEY, token);
  }
  public getToken(): string | null {
    return window.sessionStorage.getItem(TOKEN_KEY);
  }
  public saveUser(user: any): void {
  // public saveUser(token: any): void {
    // console.log("In tokenStorage::saveuser()")
    // var token = window.sessionStorage.getItem(TOKEN_KEY);
    console.log("In tokenStorage::saveuser()", user.access_token);
    var decodedToken:any = jwt_decode(user.access_token);
    console.log("Decoded Token - ", decodedToken);
    window.sessionStorage.removeItem(USER_KEY);
    var currUser = JSON.stringify({
      email: decodedToken['email'],
      role: decodedToken['role'],
    });
    // console.log("Curret User after decode - ", currUser);
    window.sessionStorage.setItem(USER_KEY, currUser);
    // var u:any = window.sessionStorage.getItem(USER_KEY);
    // if (u) var user:any = JSON.parse(u);
    // console.log(u);
    // let item = JSON.parse(window.sessionStorage.getItem(USER_KEY));
    // console.log("Parsed User is - ", user, typeof(user));
    // console.log(typeof(user.email), user.role);
  }
  public getUser(): any {
    const user = window.sessionStorage.getItem(USER_KEY);
    if (user) {
      return JSON.parse(user);
    }
    return {};
  }
}

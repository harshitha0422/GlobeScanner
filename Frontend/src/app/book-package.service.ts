// import { Injectable } from '@angular/core';

// @Injectable({
//   providedIn: 'root'
// })
// export class BookPackageService {

//   constructor() { }
// }

import { Injectable } from '@angular/core';
import {HttpClient,HttpHeaders} from '@angular/common/http';

const httpOptions = {
  headers: new HttpHeaders({ 'Content-Type': 'application/json' })
};

@Injectable({
  providedIn: 'root'
})
export class BookPackageService {

  constructor(private http : HttpClient) { }
  bookTourPackage(packageId: any){
    console.log("Package ID is -- "+packageId)
    return this.http.post<any>("http://localhost:8080/bookPackages", {packageId}, httpOptions);
  }
}
import { Component, OnInit } from '@angular/core';
import { ViewPackageService } from '../view-package.service';
import { Router } from '@angular/router';
import { BookPackageService } from '../book-package.service';

@Component({
  selector: 'app-view-packages',
  templateUrl: './view-packages.component.html',
  styleUrls: ['./view-packages.component.scss']
})
export class ViewPackagesComponent implements OnInit {
  public packages:any = [] ;
  public booked:any;
  constructor(private viewPackage : ViewPackageService, private router:Router, private bookPackageService : BookPackageService) { }


  ngOnInit(): void {
    this.getPackages();
  }
 
  navigate(){
    this.router.navigateByUrl('/home-page');
    }

  getPackages(){
  this.viewPackage.viewPackages()
    .subscribe(
      (data) => {console.log(data);
      this.packages = data;
      console.log("Packages array size is ",this.packages.length);
      },
      err => {
        console.log(err.error.message);
        
      }
    );
  }
  bookPackage(packageId:any){
    console.log("Package ID in component is ", packageId)
    this.bookPackageService.bookTourPackage(packageId)
    .subscribe(
      (data) => {console.log(data);
        this.booked = true;
      },
      err => {
        console.log(err.error.message);
        this.booked = false;
      }
    );
  }



}

import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { GetUserDataService} from '../get-user-data.service';


@Component({
  selector: 'app-list-booked-packages',
  templateUrl: './list-booked-packages.component.html',
  styleUrls: ['./list-booked-packages.component.scss']
})
export class ListBookedPackagesComponent implements OnInit {
  public bookedPackage:any = [];

  constructor(private getUser :  GetUserDataService, private router: Router) { }

  ngOnInit(): void {
    this.getUser.getTouristBookedPackage().subscribe(
      data => {
        this.bookedPackage = data;
        console.log("get booked package", this.bookedPackage);
      },
      err => {
        console.log(err.error.message);
        // this.errorMessage = err.error.message;
        // this.isSignUpFailed = true;
      }
    );
  }
  addReviewPage(){
    this.router.navigateByUrl('/list-booked-packages/add-review');
  }
  

}

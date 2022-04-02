import { Component, OnInit } from '@angular/core';
import { GetUserDataService} from '../get-user-data.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-view-guide-profile',
  templateUrl: './view-guide-profile.component.html',
  styleUrls: ['./view-guide-profile.component.scss']
})
export class ViewGuideProfileComponent implements OnInit {

  public userProfile:any = [];
  public guidePackage:any = [];
  constructor(private getUser :  GetUserDataService, private router: Router) { }

  ngOnInit(): void {
    this.getUser.getGuideInfo().subscribe(
      data => {
        this.userProfile = data;
        console.log("get user info")
        console.log(data);
      },
      err => {
        console.log(err.error.message);
        // this.errorMessage = err.error.message;
        // this.isSignUpFailed = true;
      }
    );
    this.getUser.getGuidePackage().subscribe(
      data => {
        this.guidePackage = data;
        console.log("get guide package", this.guidePackage);
      },
      err => {
        console.log(err.error.message);
        // this.errorMessage = err.error.message;
        // this.isSignUpFailed = true;
      }
    );

  }
  editInfoPage(){
    this.router.navigateByUrl('/edit-guide-profile');
  }  
}

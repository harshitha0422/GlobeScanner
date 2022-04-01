import { Component, OnInit } from '@angular/core';
import { TokenStorageService } from '../services/token-storage.service';
import { Router } from '@angular/router';
import { GetUserDataService } from '../get-user-data.service';
import { DataSharingService } from '../services/data-sharing.service';

@Component({
  selector: 'app-nav-bar',
  templateUrl: './nav-bar.component.html',
  styleUrls: ['./nav-bar.component.scss']
})
export class NavBarComponent implements OnInit {

  name?: string;
  email?: string;
  role?: string;
  isLoggedIn = false;
  isTourist = false;

  constructor(private tokenStorageService:TokenStorageService, private router: Router, private getUserDataService:GetUserDataService, private dataSharingService:DataSharingService) {
    this.dataSharingService.isUserLoggedIn.subscribe( value => {
      this.isLoggedIn = value;
    });
    this.dataSharingService.userRole.subscribe( value => {
      this.role = value;
      if(this.role=="Tourist") this.isTourist = true;
      else this.isTourist = false;
    });
    this.dataSharingService.userEmail.subscribe( value => {
      this.email = value;
    });
    this.dataSharingService.userName.subscribe( value => {
      this.name = value;
    });
   }

  ngOnInit(): void {
    this.isLoggedIn = !!this.tokenStorageService.getToken();
    if(this.isLoggedIn){
      const user = this.tokenStorageService.getUser();
      this.email = user.email;
      this.role = user.role;
      this.name = user.name;
      if(this.role=="Tourist") this.isTourist = true;
      else this.isTourist = false;
    }
  }
  viewProfile(): void{
    if(this.isLoggedIn){
      console.log("In nav init - ", this.email, this.role);
      if(this.role=="Tourist"){
        console.log(this.role);
        this.router.navigateByUrl('/view-profile');
      }
      else if(this.role=="Guide"){
        console.log(this.role);
        this.router.navigateByUrl('/view-guide-profile');
      }
    }
  }
  addReview(): void{
    
  }

  logout(): void {
    this.isLoggedIn = false;
    this.tokenStorageService.signOut();
    // window.location.reload();
    this.router.navigateByUrl('/login');
  }

}

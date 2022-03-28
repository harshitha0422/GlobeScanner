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

  private roles: string[] = [];
  isLoggedIn = false;
  showAdminBoard = false;
  showModeratorBoard = false;
  username?: string;

  constructor(private tokenStorageService:TokenStorageService, private router: Router, private getUserDataService:GetUserDataService, private dataSharingService:DataSharingService) {
    this.dataSharingService.isUserLoggedIn.subscribe( value => {
      this.isLoggedIn = value;
  });
   }

  ngOnInit(): void {
    this.isLoggedIn = !!this.tokenStorageService.getToken();
    if(this.isLoggedIn){
      const user = this.tokenStorageService.getUser();
      this.roles = user.roles;
      // this.showAdminBoard = this.roles.includes('ROLE_ADMIN');
      // this.showModeratorBoard = this.roles.includes('ROLE_MODERATOR');
      this.username = user.username;
    }
  }
  viewProfile(): void{
    if(this.isLoggedIn){
      this.router.navigateByUrl('/view-profile');
    }
  }
  logout(): void {
    this.isLoggedIn = false;
    this.tokenStorageService.signOut();
    // window.location.reload();
    this.router.navigateByUrl('/login');
  }

}

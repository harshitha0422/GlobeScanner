import { Component, OnInit } from '@angular/core';
import { AuthService } from '../services/auth.service';
import { TokenStorageService } from '../services/token-storage.service';
import { Router } from '@angular/router';
import { DataSharingService } from '../services/data-sharing.service';

@Component({
  selector: 'app-log-in',
  templateUrl: './log-in.component.html',
  styleUrls: ['./log-in.component.scss']
})
export class LogInComponent implements OnInit {
  Roles: any = ['Tourist', 'Guide'];
  form: any = {
    email: null,
    password: null,
    role: null
  };
  user: any = null;
  isLoggedIn = false;
  isLoginFailed = false;
  errorMessage = '';
  // roles: string[] = [];
  constructor(private authService: AuthService, private tokenStorage: TokenStorageService, private router:Router, private dataSharingService:DataSharingService) { }

  ngOnInit(): void {
    // if (this.tokenStorage.getToken()) {
    //   this.isLoggedIn = true;
    //   this.role = this.tokenStorage.getUser().role;
    //   console.log("Role in ngInit ", this.role);
    // }
  }

  onSubmit(): void {
    const { email, password, role} = this.form;
    this.authService.login(email, password, role).subscribe(
      data => {
        console.log(data.access_token);
        this.tokenStorage.saveToken(data.access_token);
        this.tokenStorage.saveUser(data);
        this.isLoginFailed = false;
        this.isLoggedIn = true;
        this.user = this.tokenStorage.getUser();
        // console.log("Role in ngInit ", this.role);
        // this.reloadPage();
        // this.router.navigateByUrl('/nav-bar');
        // this.router.navigate(['nav-bar']);
        // window.location.reload();
        this.dataSharingService.isUserLoggedIn.next(true);
        this.dataSharingService.userRole.next(this.user.role);
        this.dataSharingService.userEmail.next(this.user.email);
        this.router.navigateByUrl('/home-page');
      },
      err => {
        this.errorMessage = err.error.message;
        this.isLoginFailed = true;
      }
    );
  }
  reloadPage(): void {
    window.location.reload();
  }

}

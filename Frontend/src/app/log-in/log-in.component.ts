import { Component, OnInit } from '@angular/core';
import { AuthService } from '../services/auth.service';
import { TokenStorageService } from '../services/token-storage.service';
import { Router } from '@angular/router';

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
  role: any = null;
  isLoggedIn = false;
  isLoginFailed = false;
  errorMessage = '';
  // roles: string[] = [];
  constructor(private authService: AuthService, private tokenStorage: TokenStorageService, private router:Router) { }

  ngOnInit(): void {
    if (this.tokenStorage.getToken()) {
      this.isLoggedIn = true;
      this.role = this.tokenStorage.getUser().roles;
    }
  }

  onSubmit(): void {
    const { email, password, role} = this.form;
    console.log(role);
    this.authService.login(email, password, role).subscribe(
      data => {
        this.tokenStorage.saveToken(data.accessToken);
        this.tokenStorage.saveUser(data);
        this.isLoginFailed = false;
        this.isLoggedIn = true;
        this.role = this.tokenStorage.getUser().roles;
        this.reloadPage();
        // this.router.navigateByUrl('/home-page');
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

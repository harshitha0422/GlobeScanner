import { Component, OnInit } from '@angular/core';
import { GetUserDataService} from '../get-user-data.service';
import { Router } from '@angular/router';


@Component({
  selector: 'app-edit-profile',
  templateUrl: './edit-profile.component.html',
  styleUrls: ['./edit-profile.component.scss']
})
export class EditProfileComponent implements OnInit {
  public userProfile:any = [] ;

  constructor(private getUser :  GetUserDataService, private router: Router) { }

  ngOnInit(): void {
    this.getUser.getUserInfo()
    .subscribe(
      (data) => {console.log(data);
        this.userProfile = data;
      }
      );
  }

  userInfoPage(){
    this.router.navigateByUrl('/edit-profile/view-profile');
  }

}

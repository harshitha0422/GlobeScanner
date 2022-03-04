import { Component, OnInit } from '@angular/core';
import { GetUserDataService} from '../get-user-data.service';
import { Router } from '@angular/router';


@Component({
  selector: 'app-view-profile',
  templateUrl: './view-profile.component.html',
  styleUrls: ['./view-profile.component.scss']
})
export class ViewProfileComponent implements OnInit {
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
  
 
  editInfoPage(){
    this.router.navigateByUrl('/view-profile/edit-profile');
  }

}

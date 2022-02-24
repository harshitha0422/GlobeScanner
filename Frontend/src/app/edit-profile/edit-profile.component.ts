import { Component, OnInit } from '@angular/core';
import { GetUserDataService} from '../get-user-data.service';

@Component({
  selector: 'app-edit-profile',
  templateUrl: './edit-profile.component.html',
  styleUrls: ['./edit-profile.component.scss']
})
export class EditProfileComponent implements OnInit {
  public userProfile:any = [] ;

  constructor(private getUser :  GetUserDataService) { }

  ngOnInit(): void {
    this.getUser.getUserInfo()
    .subscribe(
      (data) => {console.log(data);
        this.userProfile = data;
      }
      );
  }

}

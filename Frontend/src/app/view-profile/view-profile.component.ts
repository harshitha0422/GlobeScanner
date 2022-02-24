import { Component, OnInit } from '@angular/core';
import { GetUserDataService} from '../get-user-data.service';

@Component({
  selector: 'app-view-profile',
  templateUrl: './view-profile.component.html',
  styleUrls: ['./view-profile.component.scss']
})
export class ViewProfileComponent implements OnInit {
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

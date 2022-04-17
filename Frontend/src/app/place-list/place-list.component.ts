import { Component, OnInit } from '@angular/core';
import { PlaceFetchService } from '../place-fetch.service';
import { DataSharingService } from '../services/data-sharing.service';
import { TokenStorageService } from '../services/token-storage.service';


@Component({
  selector: 'app-place-list',
  templateUrl: './place-list.component.html',
  styleUrls: ['./place-list.component.scss']
})
export class PlaceListComponent implements OnInit {

  role ?: string;
  isTourist = false;
  isLoggedIn = false;

  public places:any = [] ;
  constructor(private plyFetch :  PlaceFetchService, private dataSharingService : DataSharingService, private tokenStorageService : TokenStorageService) {
    this.dataSharingService.userRole.subscribe( value => {
      this.role = value;
      if(this.role=="Tourist") this.isTourist = true;
      else this.isTourist = false;
    });
   }

  ngOnInit(): void {
    this.isLoggedIn = !!this.tokenStorageService.getToken();
    if(this.isLoggedIn){
      const user = this.tokenStorageService.getUser();
      // this.email = user.email;
      this.role = user.role;
      // this.name = user.name;
      if(this.role=="Tourist") this.isTourist = true;
      else this.isTourist = false;
    }
    this.getPlacesList();
  }
  getPlacesList(){
  this.plyFetch.getPlaces()
    .subscribe(
      (data) => {console.log(data);
        this.places = data.msg;
        console.log("Place List:"+this.places);
      }
      );
    }
}

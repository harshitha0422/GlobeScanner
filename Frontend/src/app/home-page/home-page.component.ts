import { Component, OnInit } from '@angular/core';
import { PlaceFetchService } from '../place-fetch.service';
import { ViewPackageService } from '../view-package.service';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.scss']
})
export class HomePageComponent implements OnInit {
  textEntered :string='';

  constructor(private viewPackage : ViewPackageService , private plyFetch :  PlaceFetchService) { }
  
  ngOnInit(): void {
    document.body.className = "selector" 
  }
  onSearch() {
    console.log("inside search");
    this.plyFetch.toSearch = this.textEntered;
    this.viewPackage.viewPlacePackages = this.textEntered;
  }
  ngOnDestroy(){
    document.body.className="";
  }

}

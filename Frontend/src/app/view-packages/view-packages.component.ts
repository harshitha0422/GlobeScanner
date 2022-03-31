import { Component, OnInit } from '@angular/core';
import { ViewPackageService } from '../view-package.service';

@Component({
  selector: 'app-view-packages',
  templateUrl: './view-packages.component.html',
  styleUrls: ['./view-packages.component.scss']
})
export class ViewPackagesComponent implements OnInit {
  public packages:any = [] ;
  constructor(private viewPackage : ViewPackageService  ) { }

  ngOnInit(): void {
    this.getPackages();
  }
  getPackages(){
  this.viewPackage.viewPackages()
    .subscribe(
      (data) => {console.log(data);
      this.packages = data;
      
      
      },
      err => {
        console.log(err.error.message);
        
      }
    );
    }

}

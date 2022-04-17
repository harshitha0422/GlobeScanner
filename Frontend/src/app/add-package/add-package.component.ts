import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AddPackageService} from '../add-package.service';
import { DataSharingService } from '../services/data-sharing.service';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NgModule } from '@angular/core';


@Component({
  selector: 'app-add-package',
  templateUrl: './add-package.component.html',
  styleUrls: ['./add-package.component.scss']
})
export class AddPackageComponent implements OnInit {
  form: any = {
    email: null,
    location: null,
    included : null,
    duration : null,
    itinerary : null,
    accomodation : null,
    price: null
  };
  isPackage = false;
  isNotAdded = false;
  errorText = '';
  email?: string;

  constructor(private addPackage :  AddPackageService, private router: Router, private dataSharingService:DataSharingService) {
    this.dataSharingService.userEmail.subscribe( value => {
      this.email = value;
    });
   }

  ngOnInit(): void {
    document.body.className = "";
  }
  savePackage(){
    console.log("inside add package function now");
    this.form.email=this.email;
    this.addPackage.addNewPackage(this.form).subscribe(
      data => {
          console.log(data);
          this.isPackage = true;
      },
      err => {
        console.log(err.error.message);
        this.isNotAdded = true;
        this.errorText = err.error.message;
      }
    );
    // this.router.navigateByUrl('/home-page');
  }
  onsubmit(){}

}

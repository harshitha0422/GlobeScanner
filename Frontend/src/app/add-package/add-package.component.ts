import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AddPackageService} from '../add-package.service';
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
    // fullName: new FormControl(''),
    // age: new FormControl('')
    email: null,
    location: null,
    included : null,
    duration : null,
    itinerary : null,
    accomodation : null,
    price: null
  };

  constructor(private ap :  AddPackageService, private router: Router) { }

  ngOnInit(): void {
    document.body.className = "image-lookup";
  }
  addPackage(){
   
    this.ap.addNewPackage(this.form).subscribe(
      data => {
          console.log(data);
      },
      err => {
        console.log(err.error.message);
      }
    );
   
  }
  onsubmit(){}
  // saveProfileChanges(){
  //   console.log("FORM DATA:::::::",this.form);
  //   // this.sendSavedChanges.sendUserData(this.form).subscribe(
  //   //   data => {
  //   //       console.log(data);
  //   //   },
  //   //   err => {
  //   //     console.log(err.error.message);
  //   //   }
  //   // );
    
  // }

 
}

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
    guidEmail: null,
    location: null,
    included : null,
    duration : null,
    itinerary : null,
    accomodation : null,
    price: null
  };

  constructor(private ap :  AddPackageService, private router: Router) { }
  isPackage = false;
  isNotAdded = false;
  errorText = '';
  ngOnInit(): void {
    document.body.className = "";
  }
  addPackage(){
    console.log("inside add package function now");
    this.ap.addNewPackage(this.form).subscribe(
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
   
  }
  onsubmit(){}

}

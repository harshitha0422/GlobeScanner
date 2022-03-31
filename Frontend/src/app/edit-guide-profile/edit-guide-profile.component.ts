import { Component, OnInit } from '@angular/core';
import { GetUserDataService} from '../get-user-data.service';
import { Router } from '@angular/router';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { SendUserDataService} from '../send-user-data.service';

@Component({
  selector: 'app-edit-guide-profile',
  templateUrl: './edit-guide-profile.component.html',
  styleUrls: ['./edit-guide-profile.component.scss']
})
export class EditGuideProfileComponent implements OnInit {

  public userProfile:any = [] ;
  form: any = {
    name: null,
    age: null,
    address : null,
    location : null,
    vehicle : null
  };
  constructor(private getUser :  GetUserDataService, private router: Router, private sendSavedChanges :  SendUserDataService, public fb: FormBuilder) { }

  ngOnInit(): void {
    this.getUser.getGuideInfo().subscribe(
      data => {
        this.userProfile = data;
        console.log("edit user info")
        console.log(data);
        this.form.name = data.name;
        this.form.age = data.age;
        this.form.address = data.address;
        this.form.location = data.location;
        this.form.vehicle = data.vehicle;
      },
      err => {
        console.log(err.error.message);
        // this.errorMessage = err.error.message;
        // this.isSignUpFailed = true;
      }
    );
  }
  saveProfileChanges(){
    console.log("FORM DATA:::::::",this.form);
    this.sendSavedChanges.sendGuideData(this.form).subscribe(
      data => {
          console.log(data);
      },
      err => {
        console.log(err.error.message);
      }
    );
    this.router.navigateByUrl('/view-guide-profile');
  }

}

import { Component, OnInit } from '@angular/core';
import { GetUserDataService} from '../get-user-data.service';
import { Router } from '@angular/router';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { SendUserDataService} from '../send-user-data.service';
import { DataSharingService } from '../services/data-sharing.service';


@Component({
  selector: 'app-edit-profile',
  templateUrl: './edit-profile.component.html',
  styleUrls: ['./edit-profile.component.scss']
})
export class EditProfileComponent implements OnInit {
  public userProfile:any = [] ;
  // form = new FormGroup({
  //   fullName: new FormControl(''),
  //   age: new FormControl('')
  // });
  form: any = {
    name: null,
    about: null,
    age: null,
    mobile : null,
    location : null,
    fav1 : null,
    fav2 : null,
    fav3 : null
  };
  
  constructor(private getUser :  GetUserDataService, private router: Router, private sendSavedChanges :  SendUserDataService, private dataSharingService:DataSharingService, public fb: FormBuilder) {}

  ngOnInit(): void {
    this.getUser.getTouristInfo().subscribe(
      data => {
        this.userProfile = data;
        console.log("edit user info")
        console.log(data);
        this.form.name = data.name;
        this.form.about = data.about;
        this.form.age = data.age;
        this.form.mobile = data.mobile;
        this.form.location = data.location;
        this.form.fav1 = data.fav1;
        this.form.fav2 = data.fav2;
        this.form.fav3 = data.fav3;
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
    this.dataSharingService.userName.next(this.form.name);
    this.sendSavedChanges.sendTouristData(this.form).subscribe(
      data => {
          console.log(data);
      },
      err => {
        console.log(err.error.message);
      }
    );
    this.router.navigateByUrl('/edit-profile/view-profile');
  }

  submit() {
    // console.log('Your form data : ', this.form);
    // var formData: any = new FormData();
    // formData.append("name", this.form.get('name').value);
    // formData.append("avatar", this.form.get('avatar').value);
    // this.sendSavedChanges.sendUserData(formData).subscribe(
    //   (response) => console.log(response),
    //   (error) => console.log(error)
    // );
    
       }

}

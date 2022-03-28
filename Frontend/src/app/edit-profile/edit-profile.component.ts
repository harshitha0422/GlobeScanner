import { Component, OnInit } from '@angular/core';
import { GetUserDataService} from '../get-user-data.service';
import { Router } from '@angular/router';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { SendUserDataService} from '../send-user-data.service';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { NgModule } from '@angular/core';


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
    // fullName: new FormControl(''),
    // age: new FormControl('')
    fullName: null,
    age: null,
  };
  // fullName : any = null;
  // age : any = null;
  mobile : any = null;
  location : any = null;
  fav1 : any = null;
  fav2 : any = null;
  fav3 : any = null;
  constructor(private getUser :  GetUserDataService, private router: Router, private sendSavedChanges :  SendUserDataService, public fb: FormBuilder) { 
    // this.form = this.fb.group({
    //   fullName: [''],  
    //   email: [''],
    //   message: ['']
    // });
  }

  ngOnInit(): void {
    // this.getUser.getUserInfo()
    // .subscribe(
    //   (data) => {console.log(data);
    //     this.userProfile = data;
    //   }
    //   );
    this.getUser.getUserInfo().subscribe(
      data => {
        this.userProfile = data;
        console.log("edit user info")
        console.log(data);
        this.form.fullName = data.name;
        this.form.age = data.age;
        this.mobile = data.mobile;
        this.location = data.location;
        this.fav1 = data.fav1;
        this.fav2 = data.fav2;
        this.fav3 = data.fav3;
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
    this.sendSavedChanges.sendUserData(this.form).subscribe(
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

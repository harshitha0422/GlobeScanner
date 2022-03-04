import { Component, OnInit } from '@angular/core';
import { GetUserDataService} from '../get-user-data.service';
import { Router } from '@angular/router';
import { FormBuilder, FormGroup } from '@angular/forms';
import { SendUserDataService} from '../send-user-data.service';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';


@Component({
  selector: 'app-edit-profile',
  templateUrl: './edit-profile.component.html',
  styleUrls: ['./edit-profile.component.scss']
})
export class EditProfileComponent implements OnInit {
  public userProfile:any = [] ;
   form: FormGroup;

  constructor(private getUser :  GetUserDataService, private router: Router, private sendSavedChanges :  SendUserDataService, public fb: FormBuilder) { 
    this.form = this.fb.group({
      fullName: [''],  
      email: [''],
      message: ['']
    });
  }

  ngOnInit(): void {
    this.getUser.getUserInfo()
    .subscribe(
      (data) => {console.log(data);
        this.userProfile = data;
      }
      );
  }

  saveProfileChanges(){
    this.router.navigateByUrl('/edit-profile/view-profile');
  }

  submit() {
    console.log('Your form data : ', this.form.value );
    // var formData: any = new FormData();
    // formData.append("name", this.form.get('name').value);
    // formData.append("avatar", this.form.get('avatar').value);
    // this.sendSavedChanges.sendUserData(formData).subscribe(
    //   (response) => console.log(response),
    //   (error) => console.log(error)
    // );
    
       }

}

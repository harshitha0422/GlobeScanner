import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AddreviewService } from '../addreview.service';

@Component({
  selector: 'app-add-review',
  templateUrl: './add-review.component.html',
  styleUrls: ['./add-review.component.scss']
})
export class AddReviewComponent implements OnInit {
  form: any = {
    title: null,
    reviewtext: null
  };

  constructor(private router: Router, private ar : AddreviewService) { }

  ngOnInit(): void {
  }
  addReview(){
    console.log("inside add review function now");
    this.ar.addReview(this.form).subscribe(
      data => {
          console.log(data);
      },
      err => {
        console.log(err.error.message);
      }
    ); 
  }
}

import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

@Component({
  selector: 'app-add-review',
  templateUrl: './add-review.component.html',
  styleUrls: ['./add-review.component.scss']
})
export class AddReviewComponent implements OnInit {
  form: any = {
    title: null,
    reviewtext: null,
    travelDate: null,
    checkAgreement: null
  };

  constructor() { }

  ngOnInit(): void {
  }
  addReview(){
    
  }

}

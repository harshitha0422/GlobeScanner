import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LogInComponent } from './log-in/log-in.component';
import { RegisterComponent } from './register/register.component';
import { HomePageComponent } from './home-page/home-page.component';
import { ViewProfileComponent } from './view-profile/view-profile.component';
import { EditProfileComponent } from './edit-profile/edit-profile.component';
import { PlaceListComponent } from './place-list/place-list.component';
import { AddPackageComponent } from './add-package/add-package.component';
import { NavBarComponent } from './nav-bar/nav-bar.component';
import { AddReviewComponent } from './add-review/add-review.component';
import { ViewPackagesComponent } from './view-packages/view-packages.component';
import { ViewGuideProfileComponent } from './view-guide-profile/view-guide-profile.component';
import { EditGuideProfileComponent } from './edit-guide-profile/edit-guide-profile.component';
import { ListBookedPackagesComponent } from './list-booked-packages/list-booked-packages.component';



const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: 'home-page' },
  { path: 'login', component: LogInComponent },
  { path: 'register', component: RegisterComponent },
  { path: 'home-page', component: HomePageComponent},
  { path: 'nav-bar', component: NavBarComponent},
  { path: 'place-list', component: PlaceListComponent},
  { path: 'view-profile', component: ViewProfileComponent},
  { path: 'add-review', component: AddReviewComponent},
  { path: 'list-booked-packages', component:ListBookedPackagesComponent},
  { path: 'view-guide-profile', component: ViewGuideProfileComponent},
  { path: 'edit-guide-profile', component: EditGuideProfileComponent},
  { path:'view-profile/edit-profile',component:EditProfileComponent},
  { path:'edit-profile/view-profile',component:ViewProfileComponent},
  {path: 'place-list/view-packages', component:ViewPackagesComponent},
  { path: 'place-list/add-package', component:AddPackageComponent}
];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

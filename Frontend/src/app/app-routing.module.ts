import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LogInComponent } from './log-in/log-in.component';
import { RegisterComponent } from './register/register.component';
import { HomePageComponent } from './home-page/home-page.component';
import { ViewProfileComponent } from './view-profile/view-profile.component';
import { EditProfileComponent } from './edit-profile/edit-profile.component';
import { PlaceListComponent } from './place-list/place-list.component';
import { AddPackageComponent } from './add-package/add-package.component';

const routes: Routes = [
  { path: '', pathMatch: 'full', redirectTo: 'home-page' },
  { path: 'login', component: LogInComponent },
  { path: 'register', component: RegisterComponent },
  { path: 'home-page', component: HomePageComponent},
  { path: 'place-list', component: PlaceListComponent},
  { path: 'view-profile', component: ViewProfileComponent},
  { path:'view-profile/edit-profile',component:EditProfileComponent},
  { path:'edit-profile/view-profile',component:ViewProfileComponent},
  { path: 'add-package', component:AddPackageComponent}
];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }

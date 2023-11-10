import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { LandingpageRoutingModule } from './landingpage-routing.module';
import { NavMenuComponent } from './nav-menu/nav-menu.component';
import { SharedModule } from '../shared/shared.module';
import { ApiAuthorizationModule } from '../api-authorization/api-authorization.module';
import { LandingComponent } from './landing.component';


@NgModule({
  declarations: [
    NavMenuComponent,
    LandingComponent
    ]
  ,
  imports: [
    CommonModule,
    LandingpageRoutingModule,
    SharedModule,
    ApiAuthorizationModule
  ]
})
export class LandingpageModule { }

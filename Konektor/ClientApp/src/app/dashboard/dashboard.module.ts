import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { DashboardRoutingModule } from './dashboard-routing.module';
import { DashboardIndexComponent } from './dashboard-index/dashboard-index.component';
import { BoardComponent } from './board/board.component';
import { SidenavComponent } from './sidenav/sidenav.component';
import { NavbarComponent } from './navbar/navbar.component';
import { SharedModule } from '../shared/shared.module';
import { StoreModule } from '@ngrx/store';
import { reducers }from './core/state/reducers';
import { EffectsModule } from '@ngrx/effects';
import { ProjectEffects } from './core/state/effects';
import { ProjectDashboardComponent } from './project-dashboard/project-dashboard.component';
import { ReactiveFormsModule } from '@angular/forms';
import { WoocommerceDashboardComponent } from './project-dashboard/woocommerce-dashboard/woocommerce-dashboard.component';



@NgModule({
  declarations: [
    DashboardIndexComponent,
    BoardComponent,
    SidenavComponent,
    NavbarComponent,
    ProjectDashboardComponent,
    WoocommerceDashboardComponent
  ],
  imports: [
    CommonModule,
    DashboardRoutingModule,
    SharedModule,
    ReactiveFormsModule,
    // ngrx related imports
    StoreModule.forFeature('project', reducers),
    EffectsModule.forFeature([ProjectEffects]),
  ]
})
export class DashboardModule { }

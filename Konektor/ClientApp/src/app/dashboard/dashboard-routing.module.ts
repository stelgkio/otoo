import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DashboardIndexComponent } from './dashboard-index/dashboard-index.component';
import { BoardComponent } from './board/board.component';
import { ProjectDashboardComponent } from './project-dashboard/project-dashboard.component';
import { WoocommerceDashboardComponent } from './project-dashboard/woocommerce-dashboard/woocommerce-dashboard.component';

const routes: Routes = [
  { path: '', component: DashboardIndexComponent,
  children:[
    { path: '', component: BoardComponent},
    { path: 'woocommerce', component: WoocommerceDashboardComponent},
    { path: ':id', component: ProjectDashboardComponent},

  ]
},

];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class DashboardRoutingModule { }

import { Component} from '@angular/core';
declare function testMethodRemove(): any;
@Component({
  selector: 'app-dashboard-index',
  templateUrl: './dashboard-index.component.html',
  styleUrls: ['./dashboard-index.component.css']
})
export class DashboardIndexComponent  {
  sideNavToggle(){
  testMethodRemove();
  }
}

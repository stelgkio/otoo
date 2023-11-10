import { Observable } from 'rxjs';
import { Component, OnInit } from '@angular/core';
import { SideNavService } from '../services/sidenav.service';

@Component({
  selector: 'app-sidenav',
  templateUrl: './sidenav.component.html',
  styleUrls: ['./sidenav.component.css']
})
export class SidenavComponent implements OnInit{
  selectedProduct$!: Observable<boolean>;

  dashboardToggelItem:boolean=false;
  emailToggelItem:boolean=false;
  layoutsToggelItem:boolean=false;

  constructor(private sideNavSerice:SideNavService) {

  }
  ngOnInit(): void {
    this.selectedProduct$= this.sideNavSerice.SideNavProduct$
     console.log("selectedProduct",this.selectedProduct$)
  }

  dashboardToggel(){
    this.dashboardToggelItem= !this.dashboardToggelItem
    this.emailToggelItem= false
    this.layoutsToggelItem= false
  }
  emailToggel(status:boolean){
    this.emailToggelItem= status
    this.dashboardToggelItem= false
    this.layoutsToggelItem= false

  }
  layoutsToggel() {
  this.layoutsToggelItem= !this.layoutsToggelItem
  this.dashboardToggelItem= false
  this.emailToggelItem= false
}
}

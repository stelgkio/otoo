import { Component, OnInit } from '@angular/core';
declare function testMethodAdd(): any;
declare function testMethodRemove(): any;
@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  public show:boolean = false;
  ngOnInit(): void {

  }

  // light-style layout-navbar-fixed layout-menu-collapsed
  toggel(): void {
    this.show = !this.show;
    if(this.show)
    testMethodAdd();
    else
    testMethodRemove();

  }
}

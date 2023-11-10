import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SideNavService {
  private sideNavItems$ = new BehaviorSubject<boolean>(false);
  SideNavProduct$ = this.sideNavItems$.asObservable();


  setSideNavDisplay(product: boolean) {
    this.sideNavItems$.next(product);
  }

}

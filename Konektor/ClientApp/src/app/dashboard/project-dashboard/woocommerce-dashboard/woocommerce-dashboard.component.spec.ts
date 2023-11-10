import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WoocommerceDashboardComponent } from './woocommerce-dashboard.component';

describe('WoocommerceDashboardComponent', () => {
  let component: WoocommerceDashboardComponent;
  let fixture: ComponentFixture<WoocommerceDashboardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ WoocommerceDashboardComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(WoocommerceDashboardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

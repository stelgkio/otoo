import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
import { RouterModule } from '@angular/router';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AppComponent } from './app.component';
import { AuthorizeGuard } from 'src/app/api-authorization/authorize.guard';
import { AuthorizeInterceptor } from 'src/app/api-authorization/authorize.interceptor';
import { StoreModule } from '@ngrx/store';
import { EffectsModule } from '@ngrx/effects';
import { StoreDevtoolsModule } from '@ngrx/store-devtools';

@NgModule({
  declarations: [AppComponent],
  imports: [
    BrowserModule.withServerTransition({ appId: 'ng-cli-universal' }),
    HttpClientModule,
    BrowserAnimationsModule,
    FormsModule,
    StoreModule.forRoot({}),
    EffectsModule.forRoot(),
    StoreDevtoolsModule.instrument({}),
    RouterModule.forRoot([
      {
        path: '',
        loadChildren: () => import("./landing-page/landingpage.module").then((m) => m.LandingpageModule)
      },
      {
        path: 'dashboard',
        // canActivate: [AuthorizeGuard],
        loadChildren: () => import("./dashboard/dashboard.module").then((m) => m.DashboardModule)
      },

    ] ,{ enableTracing: true })
  ],
  providers: [
    { provide: HTTP_INTERCEPTORS, useClass: AuthorizeInterceptor, multi: true }
  ],
  bootstrap: [AppComponent],

})
export class AppModule { }

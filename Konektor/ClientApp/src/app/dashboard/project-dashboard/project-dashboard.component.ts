import { ProjectDetailsService } from './../services/projectDetails.service';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { SideNavService } from '../services/sidenav.service';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ProjectModel } from '../core/models/project.model';
import { Observable } from 'rxjs';
import { select, Store } from '@ngrx/store';
import { selectProjectList } from '../core/state/selector';
import { AppStateInterface } from '../core/state/appState.interface';


import { ProjectDetailsModel } from '../core/models/projectDetails.model';

export class Country {
  constructor(public name: string, public code: string) {}
}

@Component({
  selector: 'app-project-dashboard',
  templateUrl: './project-dashboard.component.html',
  styleUrls: ['./project-dashboard.component.css']
})
export class ProjectDashboardComponent implements OnInit {
  length!: number;
  projectId!: string;
  isWooCommerce!: boolean;
  frameWorkForm: FormGroup;
  projectData2!: ProjectModel;
  projectForm!: FormGroup;
  projects$: Observable<ProjectModel[]>;
  constructor(private route: ActivatedRoute, private sideNavSerice: SideNavService, public fb: FormBuilder, private store: Store<AppStateInterface>,private projectDetailsService: ProjectDetailsService,
    private router: Router,) {

    this.projects$ = this.store.pipe(select(selectProjectList))
    this.frameWorkForm = this.fb.group({
      framework: ['', [Validators.required]],
    })
  }

  ngOnInit() {

    this.route.params.subscribe(params => {
      console.log(params) //log the entire params object
      console.log(params['id']) //log the value of id
      this.projectId = params['id']
    });


    this.projects$.subscribe(x =>
      x.map(z => this.projectData2=z));

    //TODO if url is null then make request to api to get the value

    this.projectForm = this.fb.group({
      basicUrl: [this.projectData2.url, Validators.required],
      endpointUrl: ['/wp-json/wc/v3/', Validators.required],
      consumerSecret: ['', Validators.required],
      consumerKey: ['', [Validators.required]],

    });
    this.sideNavSerice.setSideNavDisplay(true);

  }

  // Choose framework using select dropdown
  changeFrameWork(e: any) {
    console.log(e.target.value)
    this.frameWorkForm.controls.framework.setValue(e.target.value, {
      onlySelf: true
    })
    this.isWooCommerce = e.target.value =="WooCommerce"
    console.log(this.frameWorkForm.value)
  }



  public onSubmit() {

    if (this.projectForm.invalid) {
      return;
    }
    const profileDetailsData: ProjectDetailsModel = {
     id: this.projectId,
     endpointUrl: "http://test.local",// this.projectForm.get('endpointUrl')?.value,
     consumerSecret:this.projectForm.get('consumerSecret')?.value,
     consumerKey: this.projectForm.get('consumerKey')?.value,
    };
   console.log( profileDetailsData)

   this.projectDetailsService.addProject(profileDetailsData).subscribe({
    next: x => console.log('The next value is: ', x),
    error: err => console.error('An error occurred :', err),
    complete: () =>  {
     console.log('There are no more action happen.')
     if(this.isWooCommerce){
      this.router.navigate([`/dashboard/woocommerce`]);
     }else{
      this.router.navigate([`/dashboard/shopify`]);
     }
    }
   })
  }


  ngOnDestroy(): void {
  }
}

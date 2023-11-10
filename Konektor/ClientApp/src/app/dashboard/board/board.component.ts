import { Component, OnInit } from '@angular/core';
import { ProjectModel } from '../core/models/project.model';
import { Store, select } from '@ngrx/store';
import { AppStateInterface } from '../core/state/appState.interface';
import { Observable, Subject } from 'rxjs';
import * as ProjectsActions  from '../core/state/actions';
import { selectProjectError, selectProjectIsLoading, selectProjectList, selectProjectState } from '../core/state/selector';
import { ActivatedRoute, Route, Router } from '@angular/router';
import { SideNavService } from '../services/sidenav.service';


@Component({
  selector: 'board',
  templateUrl: './board.component.html',
  styleUrls: ['./board.component.css']
})
export class BoardComponent implements OnInit {

  openCreateProjectModal: boolean= false;
  fetchData:ProjectModel[] = [];
  public isLoading$ = this.store.pipe(select(selectProjectIsLoading));
  public error$ = this.store.pipe(select(selectProjectError));
  projects$: Observable<ProjectModel[]|null>;
  destroy$: Subject<boolean> = new Subject<boolean>();

constructor(private store: Store<AppStateInterface>,private router: Router,private sideNavSerice:SideNavService) {
  this.projects$ =  this.store.pipe(select(selectProjectList));
}


  ngOnInit(): void {
    this.store.dispatch(ProjectsActions.getProjects());
    this.sideNavSerice.setSideNavDisplay(false);
  }
  public createNewProject() {
    this.openCreateProjectModal=true;
  }
  closeModalEvent(event: boolean) {
    this.openCreateProjectModal=event;
  }

  formModalValue(event: ProjectModel) {
    this.store.dispatch(ProjectsActions.createProjects({project:event}));
  }

  navigateToProject(id:string | null){
    this.router.navigate([`/dashboard/${id}`]);
  }
  ngOnDestroy() {
    this.destroy$.next(true);
    this.destroy$.unsubscribe();
  }
}



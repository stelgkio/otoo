import { Injectable } from '@angular/core';
import { Actions, createEffect, ofType } from '@ngrx/effects';
import { catchError, exhaustMap, map, mergeMap, of, switchMap } from 'rxjs';
import * as ProjectsActions from './actions';
import { ProjectService } from '../../services/project.service';

@Injectable()
export class ProjectEffects {
  getProjects$ = createEffect(() =>
    this.actions$.pipe(
      ofType(ProjectsActions.getProjects.type),
      switchMap(() => {
        return this.projectService.getProjects().pipe(
          map((projects) => ProjectsActions.getProjectsSuccess(projects)),
          catchError((error) =>
            of(ProjectsActions.getProjectsFailure({ error: error.message }))
          )
        );
      })
    )
  );

  createProjects$ = createEffect(() =>
  this.actions$.pipe(
    ofType(ProjectsActions.createProjects),
    exhaustMap(newProject => {
      return this.projectService.addProject(newProject.project).pipe(
        map((projects) =>ProjectsActions.createProjectsSuccess(projects)),
        catchError((error) =>
          of(ProjectsActions.createProjectsFailure({ error: error.message }))
        )
      );
    })
  )
);


  constructor(private actions$: Actions, private projectService: ProjectService) {}
}

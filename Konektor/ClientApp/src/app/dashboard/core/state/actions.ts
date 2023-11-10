import { createAction, props } from '@ngrx/store';
import { ProjectModel } from '../models/project.model';


export const getProjects = createAction('[Projects] Get Projects');

export const getProjectsSuccess = createAction(
  `${getProjects.type} success`,
  props<any>()
);
export const getProjectsFailure = createAction(
  '[Projects] Get Projects failure',
  props<{ error: string }>()
);

export const createProjects = createAction(
  '[Projects] create Projects',
  props<{ project: ProjectModel }>());

export const createProjectsSuccess = createAction(
  '[Projects] create Projects success',
  props<ProjectModel>()
);
export const createProjectsFailure = createAction(
  '[Projects] create Projects failure',
  props<{ error: string }>()
);

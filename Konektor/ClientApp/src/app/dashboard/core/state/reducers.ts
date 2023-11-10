import { createReducer, on } from '@ngrx/store';
import * as ProjectsActions from './actions';
import { ProjectStateInterface } from './projectState/projectState.interface';
import { ProjectModel } from '../models/project.model';

export const initialState: ProjectStateInterface = {
  isLoading: false,
  projects: [],
  error: null,
};

export const reducers = createReducer(
  initialState,
  on(ProjectsActions.getProjects, (state) => ({ ...state, isLoading: true })),
  on(ProjectsActions.getProjectsSuccess, (state, action) => ({
    ...state,
    isLoading: false,
    projects: action.projects,
  })),
  on(ProjectsActions.getProjectsFailure, (state, action) => ({
    ...state,
    isLoading: false,
    error: action.error,
  })),

  on(ProjectsActions.createProjects, (state) => ({ ...state, isLoading: true })),
  on(ProjectsActions.createProjectsSuccess, (state, projects) => {
    return {
      ...state,
      projects: [...state.projects, projects],
      isLoading: false,

    };
  }),
  on(ProjectsActions.createProjectsFailure, (state, action) => ({
    ...state,
    isLoading: false,
    error: action.error,
  })),

);

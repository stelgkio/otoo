import { createFeatureSelector, createSelector } from '@ngrx/store';
import { ProjectStateInterface } from './projectState/projectState.interface';

export const selectProjectState = createFeatureSelector<ProjectStateInterface>('project');
export const selectProjectList = createSelector(selectProjectState, (state) => state.projects);
export const selectProjectIsLoading = createSelector(selectProjectState, (state) => state.isLoading);
export const selectProjectError = createSelector(selectProjectState, (state) => state.error);

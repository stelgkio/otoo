import { HttpClient } from '@angular/common/http';
import { Inject, Injectable } from '@angular/core';

import { Observable, of } from 'rxjs';
import { ProjectDetailsModel } from '../core/models/projectDetails.model';

@Injectable({
  providedIn: 'root'
})
export class ProjectDetailsService {

  rootURL = '';
  constructor(private http: HttpClient,@Inject('BASE_URL') baseUrl: string) {
    this.rootURL=baseUrl
   }


  addProject(task: ProjectDetailsModel): Observable<ProjectDetailsModel> {
    return this.http.post<ProjectDetailsModel>(this.rootURL + 'projectdetails', task);
  }

  editProject(task: ProjectDetailsModel) {
    return this.http.put<ProjectDetailsModel>(this.rootURL + 'projectdetails', task);
  }


}

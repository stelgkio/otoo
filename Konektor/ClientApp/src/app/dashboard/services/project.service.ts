import { HttpClient } from '@angular/common/http';
import { Inject, Injectable } from '@angular/core';
import { ProjectModel } from '../core/models/project.model';
import { Observable, of } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ProjectService {

  rootURL = '';
  constructor(private http: HttpClient,@Inject('BASE_URL') baseUrl: string) {
    this.rootURL=baseUrl
   }

  getProjects():Observable<ProjectModel[]> {
    return this.http.get<ProjectModel[]>(this.rootURL + 'Projects')
  }

  addProject(task: ProjectModel): Observable<ProjectModel> {
    return this.http.post<ProjectModel>(this.rootURL + 'Project', task);
  }

  editProject(task: any) {
    return this.http.put(this.rootURL + 'Project', task);
  }

  deleteProject(taskId: any) {
    console.log('deleting task:::', taskId);
    return this.http.delete(`${this.rootURL}Project/${taskId}`);
  }
}

import { Component, EventEmitter, Input, OnDestroy, OnInit, Output } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { ProjectModel } from 'src/app/dashboard/core/models/project.model';

@Component({
  selector: 'app-new-project-modal',
  templateUrl: './new-project-modal.component.html',
  styleUrls: ['./new-project-modal.component.css']
})
export class NewProjectModalComponent implements OnInit ,OnDestroy {

  @Input() showModal: boolean = false;
  @Output() closeModal = new EventEmitter<boolean>();
  @Output() formModalValue = new EventEmitter<ProjectModel>();

  isModalOpen: boolean = false;

  projectForm: FormGroup = new FormGroup({
    Name: new FormControl(''),
    Url: new FormControl(''),
    Description: new FormControl(''),
  });
  submitted = false;

  constructor( private fb: FormBuilder,){}


  initializeForm() {
    this.projectForm = this.fb.group({
      Name: ['', Validators.required],
      Url: ['', [Validators.required]],
      Description: ['', [Validators.required, Validators.maxLength(300)]]
    });
  }


  ngOnInit(): void {
    this.initializeForm();
  }
  public close(){
    this.showModal=false;
    this.closeModal.emit(false);
  }

  public onSubmit(){

    if (this.projectForm.invalid) {
      return;
    }
    const profileData: ProjectModel = {
      id: null,
      name: this.projectForm.get('Name')?.value,
      url: this.projectForm.get('Url')?.value,
      description: this.projectForm.get('Description')?.value,

    };
    this.showModal=false;
    this.closeModal.emit(false);

    this.formModalValue.emit(profileData);
    this.projectForm.reset()
  }


  ngOnDestroy(): void {

  }
}

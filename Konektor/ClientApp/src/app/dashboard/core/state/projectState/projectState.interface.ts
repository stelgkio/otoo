import { ProjectModel } from "../../models/project.model";

export interface ProjectStateInterface {
  isLoading: boolean;
  projects: ProjectModel[],
  error: string | null;
}
